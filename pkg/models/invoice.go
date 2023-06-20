package models

import (
    "time"
)

type InvoiceData struct {
    Sender      string `json:"from"`
    Receiver    string `json:"to"`
    AmountDue   int64 `json:"total"`
    DueDate     time.Time `json:"dueDate"` 
    Amounts     []int64 `json:"amounts"`
    TaxRate     int64 `json:"tax"`
    TaxID       int64 `json:"TIN"`
}

type Invoice struct {
    UID             int64 `json:"id"`
    Creator         string `json:"creator"`
    Deliverable     string `json:"deliverables"`
    IsPaid          bool `json:"status"`
    InvoiceData     InvoiceData
}

func (inv *Invoice) GetDeliverables() string {
    if inv.IsPaid {
        return inv.Deliverable
    }
    return ""
}
