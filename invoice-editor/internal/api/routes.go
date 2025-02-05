package api

import (
    "github.com/gorilla/mux"
    "invoice-editor/internal/api/handlers"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    
    invoiceHandler := handlers.NewInvoiceHandler()
    
    // Invoice routes
    router.HandleFunc("/api/invoices", invoiceHandler.CreateInvoice).Methods("POST")
    router.HandleFunc("/api/invoices", invoiceHandler.ListInvoices).Methods("GET")
    router.HandleFunc("/api/invoices/{id}", invoiceHandler.GetInvoice).Methods("GET")
    router.HandleFunc("/api/invoices/{id}", invoiceHandler.UpdateInvoice).Methods("PUT")
    router.HandleFunc("/api/invoices/{id}/pdf", invoiceHandler.GeneratePDF).Methods("POST")
    router.HandleFunc("/api/invoices/{id}/pdf", invoiceHandler.DownloadPDF).Methods("GET")

    return router
}