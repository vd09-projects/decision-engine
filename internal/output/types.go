package output

// Stable TOOL/JSON output across use-cases.
type DecisionOutput struct {
	Decision   string   `json:"decision"`   // "YES" | "NO" | "MAYBE"
	Score      int      `json:"score"`      // 0..100
	Reasons    []string `json:"reasons"`
	Concerns   []string `json:"concerns"`
	Confidence string   `json:"confidence"` // "low"|"medium"|"high"
	Notes      []string `json:"notes,omitempty"`
}