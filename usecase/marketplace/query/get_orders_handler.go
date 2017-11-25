package query

import (
	"fmt"

	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/spec"
)

// OrderBySpecStorage fetches reports by the given criteria
type OrderBySpecStorage interface {
	BySpecWithLimit(spec intf.Specification, limit uint64) (report.GetOrdersReport, error)
}

// GetOrdersHandler returns Orders by the given Specification.
type GetOrdersHandler struct {
	s OrderBySpecStorage
}

// NewGetOrdersHandler creates a new instance of GetOrdersHandler.
func NewGetOrdersHandler(s OrderBySpecStorage) *GetOrdersHandler {
	return &GetOrdersHandler{s: s}
}

// Handle handles the given query and returns result.
// Retrieves Orders by the given Spec.
func (h *GetOrdersHandler) Handle(req intf.Query, result interface{}) error {

	q, ok := req.(GetOrders)
	if !ok {
		return fmt.Errorf("invalid query %v given", req)
	}

	r, ok := result.(*report.GetOrdersReport)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	s := spec.GetOrdersSpec(report.OrderType(q.OrderType), report.Slot{
		BuyerRating:    q.Slot.BuyerRating,
		SupplierRating: q.Slot.SupplierRating,
	})

	orders, err := h.s.BySpecWithLimit(s, q.Limit)
	if err != nil {
		return err
	}

	*r = orders

	return err
}