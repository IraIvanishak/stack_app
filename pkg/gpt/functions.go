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

type GPTModelProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

const (
	GPTPropertyTypeString  = "string"
	GPTPropertyTypeInteger = "integer"
)
