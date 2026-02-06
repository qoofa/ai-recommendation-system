package food

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qoofa/AI-Recommendation-System/internal/domain/food"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/response"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type FoodHandler struct {
	validate *validator.Validate
	service  food.Service
}

func NewFoodHandler(s food.Service) *FoodHandler {
	return &FoodHandler{
		validate: validator.New(),
		service:  s,
	}
}

func (h *FoodHandler) Search(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()

	query := q.Get("query")

	if query == "" {
		response.Error(w, appErr.New(appErr.BadRequest, "query is required"))
		return
	}

	d, err := h.service.Search(r.Context(), query)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, d)
}
