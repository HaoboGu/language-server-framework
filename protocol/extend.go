package protocol

// We extend the raw LSP types here
type ExtendedCompletionParam struct {
	CompletionParams
	FileContent string
}
