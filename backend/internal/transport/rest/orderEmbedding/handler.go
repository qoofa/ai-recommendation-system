package orderembedding

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	orderembeddings "github.com/qoofa/AI-Recommendation-System/internal/domain/orderEmbeddings"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/response"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type OrderEmbeddingHanlder struct {
	validate *validator.Validate
	service  orderembeddings.Service
}

func New(s orderembeddings.Service) *OrderEmbeddingHanlder {
	return &OrderEmbeddingHanlder{
		validate: validator.New(),
		service:  s,
	}
}

func (h *OrderEmbeddingHanlder) Train(w http.ResponseWriter, r *http.Request) {
	var body TrainDto

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.Error(w, appErr.New(appErr.BadRequest, "invalid body"))
		return
	}

	if err := h.validate.Struct(body); err != nil {
		response.Error(w, appErr.Wrap(appErr.BadRequest, "validation error", err))
		return
	}

	result, err := h.service.Train(r.Context(), body.Items)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}
