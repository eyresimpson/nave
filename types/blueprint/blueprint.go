package blueprint

// BluePrint 蓝图，其中包含了流水线的具体结构
type BluePrint struct {
	Label    string     `json:"label"`
	Version  string     `json:"version"`
	Request  Request    `json:"request"`
	FlowType string     `json:"flowType"`
	Port     string     `json:"port"`
	Path     string     `json:"path"`
	Route    []Route    `json:"route"`
	Mods     []string   `json:"mods"`
	Vars     []Variable `json:"vars"`
	Steps    []Worker   `json:"steps"`
}
