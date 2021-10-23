package clockin

import "io/fs"

var (
	ValidConfigPaths []string = []string{"./config.yml", "~/.clockin/config.yml"}
)

const (
	DEFAULT_FILE_MODE fs.FileMode = 0644
)
