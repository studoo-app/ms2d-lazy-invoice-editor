package models

import "time"

type Former struct {
    ID          string    `json:"id" validate:"required"`
    LastName    string    `json:"nom" validate:"required"`
    FirstName   string    `json:"prenom" validate:"required"`
    Email       string    `json:"email" validate:"required,email"`
    Address     string    `json:"address" validate:"required"`
    PostalCode  string    `json:"code_postal" validate:"required"`
    Phone       string    `json:"tel" validate:"required"`
    SiretNumber string    `json:"siret" validate:"required,len=14"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}