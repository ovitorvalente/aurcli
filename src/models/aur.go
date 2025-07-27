package models

type Package struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type AurResponse struct {
	Results []Package `json:"results"`
}
