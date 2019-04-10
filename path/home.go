package path

import (
	"os/user"
	"path"
)

// Home returns the path to the oxm folder
func Home() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return path.Join(user.HomeDir, ".oxm")
}
