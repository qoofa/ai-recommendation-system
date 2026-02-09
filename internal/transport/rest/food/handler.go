package food

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/qoofa/AI-Recommendation-System/internal/core"
	"github.com/qoofa/AI-Recommendation-System/internal/transport/rest/response"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type FoodHandler struct {
	validate *validator.Validate
	service  core.FoodService
}

func New(s core.FoodService) *FoodHandler {
	return &FoodHandler{
		validate: validator.New(),
		service:  s,
	}
}

// Find godoc
//
//	@Summary		List food items
//	@Description	Get all food items from the database
//	@Tags			food
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.APIResponse{data=[]SearchResponseDto}
//	@Failure		500	{object}	response.APIResponse{error=response.APIError}
//	@Router			/food [get]
func (h *FoodHandler) Find(w http.ResponseWriter, r *http.Request) {
	d, err := h.service.Find(r.Context())
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

// Search godoc
//
//	@Summary		Search food items
//	@Description	Search food items by name or description using vector search
//	@Tags			food
//	@Accept			json
//	@Produce		json
//	@Param			query	query		string	true	"Search query"
//	@Success		200		{object}	response.APIResponse{data=[]SearchResponseDto}
//	@Failure		400		{object}	response.APIResponse{error=response.APIError}
//	@Failure		500		{object}	response.APIResponse{error=response.APIError}
//	@Router			/food/search [get]
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

// Create godoc
//
//	@Summary		Create food item
//	@Description	Create a new food item and generate its embedding
//	@Tags			food
//	@Accept			json
//	@Produce		json
//	@Param			food	body		createDto	true	"Food item to create"
//	@Success		200		{object}	response.APIResponse{data=string}
//	@Failure		400		{object}	response.APIResponse{error=response.APIError}
//	@Failure		500		{object}	response.APIResponse{error=response.APIError}
//	@Router			/food [post]
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

// Recommend godoc
//
//	@Summary		Recommend food item
//	@Description	Recommends food items related to the food item selected
//	@Tags			food
//	@Accept			json
//	@Produce		json
//	@Param			itemId	query		string	true	"Item Id Selected"
//	@Success		200		{object}	response.APIResponse{data=string}
//	@Failure		400		{object}	response.APIResponse{error=response.APIError}
//	@Failure		500		{object}	response.APIResponse{error=response.APIError}
//	@Router			/food/recommend [get]
func (h *FoodHandler) Recommend(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	itemId := q.Get("itemId")

	if itemId == "" {
		response.Error(w, appErr.New(appErr.BadRequest, "itemId is required"))
		return
	}

	result, err := h.service.Recommend(r.Context(), itemId)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, result)
}
