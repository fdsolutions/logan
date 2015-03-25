package command

// Metadata holds command specification
type Metadata struct {
	Intent  string
	Target  string
	Context string
}

// NewMetadataFromName instanciate a new metadata from command name
func NewMetadataFromName(name string) Metadata {
	intent, target, ctx := GetCommandNameParts(name)
	return NewMetadata(intent, target, ctx)
}

// New create a metadata object from its properties (intent, target, context).
// Make sure parameters are in the right order.
func NewMetadata(intent string, target string, context string) Metadata {
	return Metadata{intent, target, context}
}
