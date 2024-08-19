package brapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://brapi.dev/api/quote/"

// GetQuote faz uma requisição à API Brapi e retorna a resposta deserializada.
func GetCotacaoBrapi(symbol, token string) (*BrapiCotacaoResponse, error) {
	url := fmt.Sprintf("%s%s?token=%s", baseURL, symbol, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cotacaoResponse BrapiCotacaoResponse
	err = json.Unmarshal(body, &cotacaoResponse)
	if err != nil {
		return nil, err
	}

	return &cotacaoResponse, nil
}

// FormatResponse formata a resposta da API em um JSON indentado.
func FormatResponse(cotacaoResponse *BrapiCotacaoResponse) string {
	output, err := json.MarshalIndent(cotacaoResponse, "", "    ")
	if err != nil {
		return fmt.Sprintf("Failed to marshal JSON: %v", err)
	}

	return string(output)
}
