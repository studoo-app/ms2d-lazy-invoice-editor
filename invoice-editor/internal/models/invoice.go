package models

import "time"

type InvoiceLine struct {
    ID        string    `json:"id" validate:"required"`
    Module    string    `json:"module" validate:"required"`
    Date      time.Time `json:"date" validate:"required"`
    Hours     float64   `json:"hours" validate:"required,gt=0"`
    Class     string    `json:"class" validate:"required"`
    RatePerHr float64   `json:"rate_per_hr" validate:"required,gt=0"`
}

type Invoice struct {
    ID          string        `json:"id" validate:"required"`
    TeacherName string        `json:"teacher_name" validate:"required"`
    Month       time.Time     `json:"month" validate:"required"`
    Lines       []InvoiceLine `json:"lines" validate:"required,dive"`
    TotalAmount float64       `json:"total_amount" validate:"required,gte=0"`
    CreatedAt   time.Time     `json:"created_at" validate:"required"`
    UpdatedAt   time.Time     `json:"updated_at" validate:"required"`
    PDFPath     string        `json:"pdf_path" validate:"required"`
}