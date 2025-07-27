package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ovitorvalente/aurcli/src/models"
)

func BuscarPacotes(termo string) ([]models.Package, error) {
	url := fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s", termo)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pacotes: %w", err)
	}
	defer res.Body.Close()

	var response models.AurResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	return response.Results, nil
}