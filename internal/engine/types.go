package engine

import "encoding/json"

// Stable envelope input across use-cases.
type DecisionRequest struct {
	UseCase string          `json:"use_case"`
	Payload json.RawMessage `json:"payload"`
	Context json.RawMessage `json:"context,omitempty"`
}
