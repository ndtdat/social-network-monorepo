package base

type Identity struct {
	Metadata  map[string]string `json:"metadata"`
	APIKey    map[string]string `json:"apiKey"`
	ID        string            `json:"id,required"` //nolint:staticcheck
	SessionID string            `json:"sid"`
	Domain    string            `json:"dom"`
	DeviceID  string            `json:"did"`
	Roles     []string          `json:"roles,required"` //nolint:staticcheck
	Exp       int               `json:"exp"`
}
