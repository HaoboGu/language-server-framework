package protocol

// We extend the raw LSP types here

// ExtendedCompletionParam adds fileContent to completion params
type ExtendedCompletionParam struct {
	CompletionParams
	FileContent string `json:"fileContent,omitempty"`
}
