package models

import "time"

type Company struct {
    ID          string    `json:"id" validate:"required"`
    Name        string    `json:"name" validate:"required"`
    Email       string    `json:"email" validate:"required,email"`
    Address     string    `json:"address" validate:"required"`
    PostalCode  string    `json:"postal_code" validate:"required"`
    City        string    `json:"city" validate:"required"`
    Country     string    `json:"country" validate:"required"`
    Phone       string    `json:"phone" validate:"required"`
    SiretNumber string    `json:"siret" validate:"required,len=14"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}