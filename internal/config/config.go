package config

type Config struct {
	DecisionEngine DecisionEngine `yaml:"decision_engine"`
}

type DecisionEngine struct {
	Output   Output                 `yaml:"output"`
	UseCases map[string]UseCaseConf `yaml:"use_cases"`
}

type UseCaseConf struct {
	LLM    LLM    `yaml:"llm"`
	Policy Policy `yaml:"policy"`
	Prompt Prompt `yaml:"prompt"`
}

type LLM struct {
	Provider    string  `yaml:"provider"` // "ollama"
	BaseURL     string  `yaml:"base_url"`
	Model       string  `yaml:"model"`
	Temperature float64 `yaml:"temperature"`
	TopP        float64 `yaml:"top_p"`
	NumPredict  int     `yaml:"num_predict"`
	TimeoutMs   int     `yaml:"timeout_ms"`
}

type Output struct {
	Format         string   `yaml:"format"` // "tool_json"
	AllowDecisions []string `yaml:"allow_decisions"`
	ScoreMin       int      `yaml:"score_min"`
	ScoreMax       int      `yaml:"score_max"`
	MaxReasons     int      `yaml:"max_reasons"`
	MaxConcerns    int      `yaml:"max_concerns"`
}

type Policy struct {
	Thresholds Thresholds `yaml:"thresholds"`
}

type Thresholds struct {
	YesMinScore int `yaml:"yes_min_score"`
	NoMaxScore  int `yaml:"no_max_score"`
}

type Prompt struct {
	TemplatePath string             `yaml:"template_path"`
	Bindings     map[string]Binding `yaml:"bindings"`
}

type Binding struct {
	Source   string `yaml:"source"`  // "payload"|"context"
	Path     string `yaml:"path"`    // gjson path
	AsJSON   bool   `yaml:"as_json"` // emit JSON text instead of string
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
	MaxLen   int    `yaml:"max_len"` // truncate long strings/JSON
}
