package blueprint

type Route struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Type      string   `json:"type"`
	GetParams []string `json:"getParams"`
}
