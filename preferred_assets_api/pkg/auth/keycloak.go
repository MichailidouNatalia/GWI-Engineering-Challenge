package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/cmd/api/config"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt"
)

type KeycloakClient struct {
	config   *config.KeycloakConfig
	client   *http.Client
	verifier *oidc.IDTokenVerifier
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

func NewKeycloakClient(cfg *config.KeycloakConfig) (*KeycloakClient, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, cfg.URL+"/realms/"+cfg.Realm)
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID:          cfg.ClientID,
		SkipClientIDCheck: true,
	})

	return &KeycloakClient{
		config:   cfg,
		client:   &http.Client{Timeout: cfg.Timeout},
		verifier: verifier,
	}, nil
}

func (kc *KeycloakClient) VerifyToken(tokenString string) (*CustomClaims, error) {
	ctx := context.Background()

	idToken, err := kc.verifier.Verify(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	var claims CustomClaims
	if err := idToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}

	return &claims, nil
}

func (kc *KeycloakClient) GetUserRoles(claims *CustomClaims) []string {
	var roles []string
	roles = append(roles, claims.RealmAccess.Roles...)
	if clientRoles, exists := claims.ResourceAccess[kc.config.ClientID]; exists {
		roles = append(roles, clientRoles.Roles...)
	}
	return roles
}
