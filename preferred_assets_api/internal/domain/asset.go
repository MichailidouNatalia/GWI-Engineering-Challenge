package domain

import (
	"errors"
	"time"
)

type AssetBase struct {
	ID          string
	Type        AssetType
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Asset interface {
	GetID() string
	GetType() AssetType
	GetTitle() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	Validate() error
}

// Common getter methods
func (a AssetBase) GetID() string           { return a.ID }
func (a AssetBase) GetType() AssetType      { return a.Type }
func (a AssetBase) GetTitle() string        { return a.Title }
func (a AssetBase) GetDescription() string  { return a.Description }
func (a AssetBase) GetCreatedAt() time.Time { return a.CreatedAt }
func (a AssetBase) GetUpdatedAt() time.Time { return a.UpdatedAt }

// Domain validation
func (a AssetBase) Validate() error {
	if a.Title == "" {
		return errors.New("asset title is required")
	}
	return nil
}
