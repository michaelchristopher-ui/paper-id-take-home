package disbursementhandler

import (
	"net/http"
	da "paperid-entry-task/handler/http/adapters/disbursement"
	"paperid-entry-task/handler/http/common/constants"
	"paperid-entry-task/handler/http/common/models"
	"paperid-entry-task/internal/pkg/adapters/disbursementsvc"

	"github.com/labstack/echo"
)

var _ da.Adapter = &DisbursementAPIPIntegrator{}

type DisbursementAPIPIntegrator struct {
	DisbursementService disbursementsvc.Adapter
}

func NewAPIIntegrator(disbursementService disbursementsvc.Adapter) da.Adapter {
	return &DisbursementAPIPIntegrator{DisbursementService: disbursementService}
}

func API(e *echo.Echo, disbursementService disbursementsvc.Adapter) {
	api := e.Group("/disbursement")
	integrator := NewAPIIntegrator(disbursementService)

	api.POST("/disburse", integrator.Disbursement)
}

func (h *DisbursementAPIPIntegrator) Disbursement(c echo.Context) error {
	u := new(da.DisbursementRequest)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorJson{
			Error: constants.ErrorInternalServer,
		})
	}

	err := h.DisbursementService.Disburse(c.Request().Context(), disbursementsvc.DisburseReq{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorJson{
			Error: constants.ErrorInternalServer,
		})
	}

	return c.JSON(http.StatusOK, da.DisbursementResponse{})

}
