package food

import (
	"encoding/json"
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

	data := make([]SearchResponseDto, len(d))
	for i := range d {
		data[i] = TOSearchResponse(d[i])
	}

	response.Success(w, data)
}

func (h *FoodHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body createDto

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, appErr.New(appErr.BadRequest, "invalid body"))
		return
	}

	if err := h.validate.Struct(body); err != nil {
		response.Error(w, appErr.Wrap(appErr.BadRequest, "validation error", err))
		return
	}

	result, err := h.service.Create(r.Context(), *body.toDomain())
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}
