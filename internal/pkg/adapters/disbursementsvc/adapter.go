package disbursementsvc

import "context"

type Adapter interface {
	Disburse(ctx context.Context, req DisburseReq) error
}

type DisburseReq struct {
	AccountIdFrom int
	AccountIdTo   int
	Amount        float64
	Description   string
}
