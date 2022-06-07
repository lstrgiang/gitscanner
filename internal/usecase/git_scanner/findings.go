package git_scanner

type LineType struct {
	Begin int `json:"begin"`
}
type FindingLocation struct {
	Path     string   `json:"path"`
	Position LineType `json:"position"`
}
type ScanFinding struct {
	Type     string          `json:"type"`
	RuleID   string          `json:"rule_id"`
	Location FindingLocation `json:"location"`
	Meta     FindingMeta     `json:"meta"`
}

type FindingMeta struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

func NewSecretKeyFinding(line int, fileName string) ScanFinding {
	return ScanFinding{
		Type:   "secret_key",
		RuleID: "G2202",
		Location: FindingLocation{
			Path: fileName,
			Position: LineType{
				Begin: line,
			},
		},
		Meta: FindingMeta{
			Description: "Secret key should not be exposed publicly",
			Severity:    "HIGH",
		},
	}
}
