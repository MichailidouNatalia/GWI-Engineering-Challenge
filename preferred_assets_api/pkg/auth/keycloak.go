package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/config"
	"github.com/golang-jwt/jwt/v4"
)

type KeycloakClient struct {
	config *config.KeycloakConfig
	client *http.Client
}

type TokenIntrospection struct {
	Active    bool     `json:"active"`
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
	ExpiresAt int64    `json:"exp"`
	IssuedAt  int64    `json:"iat"`
	Subject   string   `json:"sub"`
}

type CustomClaims struct {
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
	ResourceAccess map[string]struct {
		Roles []string `json:"roles"`
	} `json:"resource_access"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	PreferredName string `json:"preferred_username"`
	jwt.StandardClaims
}

func NewKeycloakClient(cfg *config.KeycloakConfig) *KeycloakClient {
	return &KeycloakClient{
		config: cfg,
		client: &http.Client{Timeout: cfg.Timeout},
	}
}

// VerifyToken verifies the JWT token and returns claims
func (kc *KeycloakClient) VerifyToken(tokenString string) (*CustomClaims, error) {
	// Get Keycloak realm public key
	publicKey, err := kc.getPublicKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %v", err)
	}

	// Parse and verify token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// Check if token is expired
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// IntrospectToken uses Keycloak's token introspection endpoint
func (kc *KeycloakClient) IntrospectToken(token string) (*TokenIntrospection, error) {
	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token/introspect",
		kc.config.URL, kc.config.Realm)

	formData := strings.NewReader(fmt.Sprintf(
		"token=%s&client_id=%s&client_secret=%s",
		token, kc.config.ClientID, kc.config.ClientSecret,
	))

	req, err := http.NewRequest("POST", url, formData)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := kc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var introspection TokenIntrospection
	if err := json.Unmarshal(body, &introspection); err != nil {
		return nil, err
	}

	return &introspection, nil
}

func (kc *KeycloakClient) getPublicKey() (interface{}, error) {
	// In production, you should cache this public key
	certURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/certs",
		kc.config.URL, kc.config.Realm)

	resp, err := kc.client.Get(certURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var certs struct {
		Keys []struct {
			Kid string `json:"kid"`
			Kty string `json:"kty"`
			Alg string `json:"alg"`
			Use string `json:"use"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&certs); err != nil {
		return nil, err
	}

	if len(certs.Keys) == 0 {
		return nil, fmt.Errorf("no public keys found")
	}

	// Use the first key (you might want to match by kid in production)
	key := certs.Keys[0]
	return jwt.ParseRSAPublicKeyFromPEM([]byte(fmt.Sprintf(
		"-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----",
		key.N,
	)))
}

// GetUserRoles returns roles from claims
func (kc *KeycloakClient) GetUserRoles(claims *CustomClaims) []string {
	var roles []string

	// Add realm roles
	roles = append(roles, claims.RealmAccess.Roles...)

	// Add client-specific roles
	if clientRoles, exists := claims.ResourceAccess[kc.config.ClientID]; exists {
		roles = append(roles, clientRoles.Roles...)
	}

	return roles
}
