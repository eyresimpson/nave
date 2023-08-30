package blueprint

type Variable struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default any    `json:"default"`
}
