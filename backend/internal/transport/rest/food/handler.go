package food

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/response"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type FoodHandler struct {
	validate *validator.Validate
}

func NewFoodHandler() *FoodHandler {
	return &FoodHandler{
		validate: validator.New(),
	}
}

func (h *FoodHandler) Search(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()

	query := q.Get("query")

	if query == "" {
		response.Error(w, appErr.New(appErr.BadRequest, "query is required"))
		return
	}

	response.Success(w, nil)
}
