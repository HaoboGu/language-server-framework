package server

// LanguageServer defines the behaviors of a real language server
type LanguageServer interface {
	// TODO: Add behaviors
	Start() error
	Init() error
}
