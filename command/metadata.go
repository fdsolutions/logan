package command

type Metadata struct {
	Intent  string
	Target  string
	Context string
}

func New() *Metadata {
	return new(Metadata)
}

func FromJSON(name string) *Metadata {
	return newMetadata("", "", "", nil, nil)
}
