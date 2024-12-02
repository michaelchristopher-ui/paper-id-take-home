package disbursement

import "github.com/labstack/echo"

type Adapter interface {
	Disbursement(c echo.Context) error
}

type DisbursementRequest struct {
	AccountIdFrom int     `json:"account_id_from"`
	AccountIdTo   int     `json:"account_id_to"`
	Amount        float64 `json:"amount"`
	Description   string  `json:"description"`
}

type DisbursementResponse struct{}
