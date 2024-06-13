package gpt

type GPTFunction struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Parameters  GPTFunctionParameter `json:"parameters"`
}
type GPTFunctionParameter struct {
	Type       string                 `json:"type"`
	Properties GPTFunctionObjectModel `json:"properties"`
}

type GPTFunctionObjectModel interface{}

// TODO expand
type GPTModelProperty struct {
	Type        string            `json:"type"`
	Description string            `json:"description"`
	Items       *GPTModelProperty `json:"items,omitempty"`
}

// TODO expand
const (
	GPTPropertyTypeObject  = "object"
	GPTPropertyTypeString  = "string"
	GPTPropertyTypeInteger = "integer"
	GPTPropertyTypeArray   = "array"
)
