package fixtures

import (
	"path/filepath"
)

var (
	UnexistingPath      string = filepath.Join("..", "fixtures", "data", "nofile.metas")
	ExistingPath        string = filepath.Join("..", "fixtures", "data", "command_examples.metas")
	EmptyFilePath       string = filepath.Join("..", "fixtures", "data", "empty.metas")
	UnsupportedFilePath string = filepath.Join("..", "fixtures", "data", "unsupported_yaml.metas")
)
