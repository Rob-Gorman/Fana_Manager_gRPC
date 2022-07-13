package models

type Flag struct {
	Name      string   `json:"name"`
	SDKKey    string   `json:"sdkKey"`
	Status    bool     `json:"status,omitempty"`
	Audiences []string `json:"audiences"`
}
