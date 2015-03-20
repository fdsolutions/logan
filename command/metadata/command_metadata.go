package metadata

// Metadata holds command specification
type Metadata struct {
	Intent  string
	Target  string
	Context string
}

// ForName returns a new metadata object from the given command name
func FromCmdName(name string) *Metadata {

}

// New create a metadata object from its properties (intent, target, context).
// Make sure parameters are in the right order.
func New(intent string, target string, context string) *Metadata {
	return &Metadata{intent, target, context}
}
