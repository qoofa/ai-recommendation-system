package embeddings

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	appErr "github.com/qoofa/AI-Recommendation-System/internal/core/errors"
)

type Response struct {
	Status    bool      `json:"status"`
	Embedding []float64 `json:"embedding"`
}

type Request struct {
	Text string `json:"text"`
}

type pythonProvider struct {
	url    string
	client *http.Client
}

func NewPythonProvider(url string) *pythonProvider {
	return &pythonProvider{
		url:    strings.TrimRight(url, "/"),
		client: &http.Client{},
	}
}

func (p *pythonProvider) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	reqBody, err := json.Marshal(Request{Text: text})
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.url+"/embed", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, appErr.Wrap(appErr.Internal, "internal error", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
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

	return response.Embedding, nil
}
