package embeddings

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type Response struct {
	status    bool
	embedding []float64
}

type pythonProvider struct {
	url    string
	client *http.Client
}

func NewPythonProvider(url string) *pythonProvider {
	return &pythonProvider{
		url:    strings.TrimRight(url, "/") + "/",
		client: &http.Client{},
	}
}

func (p *pythonProvider) GetEmbedding(ctx context.Context, text string) ([]float64, error) {

	resp, err := p.client.Get(p.url)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, appErr.Wrap(appErr.Internal, "internal error", errors.New("PYTHON EMBEDDING SERVER INVALID STATUS CODE"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	return response.embedding, nil
}
