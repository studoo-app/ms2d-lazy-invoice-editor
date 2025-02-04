package handlers

import (
    "net/http"
)

type InvoiceHandler struct {
    // Add dependencies here later (storage, pdf generator etc)
}

func NewInvoiceHandler() *InvoiceHandler {
    return &InvoiceHandler{}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
    // Handler for POST /invoices
}

func (h *InvoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
    // Handler for GET /invoices/{id}
}

func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
    // Handler for PUT /invoices/{id}
}

func (h *InvoiceHandler) ListInvoices(w http.ResponseWriter, r *http.Request) {
    // Handler for GET /invoices
}

func (h *InvoiceHandler) GeneratePDF(w http.ResponseWriter, r *http.Request) {
    // Handler for POST /invoices/{id}/pdf
}

func (h *InvoiceHandler) DownloadPDF(w http.ResponseWriter, r *http.Request) {
    // Handler for GET /invoices/{id}/pdf
}