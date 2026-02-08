package orderembedding

type TrainDto struct {
	Items []string `json:"items" validate:"required,min=1"`
}
