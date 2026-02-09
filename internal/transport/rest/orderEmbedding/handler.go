package orderembedding

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qoofa/AI-Recommendation-System/internal/core"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/response"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type OrderEmbeddingHanlder struct {
	validate *validator.Validate
	service  core.OrderEmbeddingService
}

func New(s core.OrderEmbeddingService) *OrderEmbeddingHanlder {
	return &OrderEmbeddingHanlder{
		validate: validator.New(),
		service:  s,
	}
}

// Train godoc
//
//	@Summary		Train order embedding
//	@Description	Calculate order embedding based on a list of food item IDs
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			order	body		TrainDto	true	"Order items to train"
//	@Success		200		{object}	response.APIResponse{data=string}
//	@Failure		400		{object}	response.APIResponse{error=response.APIError}
//	@Failure		500		{object}	response.APIResponse{error=response.APIError}
//	@Router			/order/train [post]
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
