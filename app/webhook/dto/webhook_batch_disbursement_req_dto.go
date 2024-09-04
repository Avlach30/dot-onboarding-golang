package dto

type WebhookBatchDisbursementReq struct {
	ID                       string  `json:"id"`
	TotalExecutedBatch       int     `json:"total_disbursed_count"`
	TotalExecutedBatchAmount float64 `json:"total_disbursed_amount"`
	Reference                string  `json:"reference"`
	TotalRequestBatch        int     `json:"total_uploaded_count"`
	TotalRequestBatchAmount  float64 `json:"total_uploaded_amount"`
	TotalErrorBatch          int     `json:"total_error_count"`
	TotalErrorBatchAmount    float64 `json:"total_error_amount"`
	Status                   string  `json:"status"`
}
