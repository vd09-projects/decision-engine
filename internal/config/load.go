package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (Config, error) {
	var cfg Config
	b, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return cfg, err
	}
	applyDefaults(&cfg)
	if err := validate(cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func applyDefaults(cfg *Config) {
	// Output defaults
	if cfg.DecisionEngine.Output.ScoreMax == 0 {
		cfg.DecisionEngine.Output.ScoreMax = 100
	}
	if cfg.DecisionEngine.Output.MaxReasons == 0 {
		cfg.DecisionEngine.Output.MaxReasons = 6
	}
	if cfg.DecisionEngine.Output.MaxConcerns == 0 {
		cfg.DecisionEngine.Output.MaxConcerns = 6
	}

	// Per-usecase defaults
	for k, uc := range cfg.DecisionEngine.UseCases {
		if uc.LLM.TimeoutMs == 0 {
			uc.LLM.TimeoutMs = 20000
		}
		if uc.LLM.NumPredict == 0 {
			uc.LLM.NumPredict = 512
		}
		cfg.DecisionEngine.UseCases[k] = uc
	}
}

func validate(cfg Config) error {
	if len(cfg.DecisionEngine.UseCases) == 0 {
		return fmt.Errorf("decision_engine.use_cases is empty")
	}
	if cfg.DecisionEngine.Output.ScoreMin > cfg.DecisionEngine.Output.ScoreMax {
		return fmt.Errorf("output.score_min > output.score_max")
	}
	return nil
}
