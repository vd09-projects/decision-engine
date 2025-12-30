package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeConfigFile(t *testing.T, contents string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}

func TestLoadAppliesDefaults(t *testing.T) {
	path := writeConfigFile(t, `
decision_engine:
  output:
    score_min: 10
  use_cases:
    example:
      llm:
        provider: "ollama"
        base_url: "http://localhost:11434"
        model: "test-model"
        temperature: 0.2
        top_p: 0.8
      policy:
        thresholds:
          yes_min_score: 60
          no_max_score: 40
      prompt:
        template_path: "./prompts/example.tmpl"
`)

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	output := cfg.DecisionEngine.Output
	if output.ScoreMax != 100 {
		t.Fatalf("ScoreMax = %d, want 100", output.ScoreMax)
	}
	if output.MaxReasons != 6 {
		t.Fatalf("MaxReasons = %d, want 6", output.MaxReasons)
	}
	if output.MaxConcerns != 6 {
		t.Fatalf("MaxConcerns = %d, want 6", output.MaxConcerns)
	}

	llm := cfg.DecisionEngine.UseCases["example"].LLM
	if llm.NumPredict != 512 {
		t.Fatalf("NumPredict = %d, want 512", llm.NumPredict)
	}
	if llm.TimeoutMs != 20000 {
		t.Fatalf("TimeoutMs = %d, want 20000", llm.TimeoutMs)
	}
}

func TestLoadErrorsWhenUseCasesEmpty(t *testing.T) {
	path := writeConfigFile(t, `
decision_engine:
  output:
    score_min: 0
    score_max: 100
  use_cases: {}
`)

	if _, err := Load(path); err == nil {
		t.Fatalf("Load() error = nil, want non-nil")
	}
}

func TestLoadErrorsWhenScoreMinGreaterThanMax(t *testing.T) {
	path := writeConfigFile(t, `
decision_engine:
  output:
    score_min: 100
    score_max: 50
  use_cases:
    example:
      llm:
        provider: "ollama"
        base_url: "http://localhost"
        model: "test-model"
        temperature: 0.1
        top_p: 0.8
      policy:
        thresholds:
          yes_min_score: 70
          no_max_score: 30
      prompt:
        template_path: "./prompts/example.tmpl"
`)

	if _, err := Load(path); err == nil {
		t.Fatalf("Load() error = nil, want non-nil")
	}
}
