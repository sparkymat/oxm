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

// Usr returns the /usr folder for oxm
func Usr() string {
	return path.Join(Home(), "usr")
}

// Src returns the /usr/src folder for oxm
func Src() string {
	return path.Join(Home(), "usr", "src")
}
