package setup

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-getter"
	"github.com/sparkymat/oxm/path"
)

// Init initializes the oxm installation
func Init() error {
	var err error

	if err = createHomeFolder(); err != nil {
		return err
	}

	if err = checkCompiler(); err != nil {
		return err
	}

	if err = downloadGCC(); err != nil {
		return err
	}

	if err = compileGCC(); err != nil {
		return err
	}

	return nil
}

func createHomeFolder() error {
	var err error

	if _, err = os.Stat(path.Home()); !os.IsNotExist(err) {
		return fmt.Errorf("%v already exists", path.Home())
	}

	if err = os.MkdirAll(path.Src(), 0755); err != nil {
		return err
	}

	return nil
}

func checkCompiler() error {
	typeGccCommand := exec.Command("type", "gcc")
	err := typeGccCommand.Run()
	if err != nil {
		return err
	}

	return nil
}

func downloadGCC() error {
	destPath := path.Src()
	gccLink := "https://www.artfiles.org/gnu.org/gcc/gcc-8.3.0/gcc-8.3.0.tar.xz"
	track := progressTracker{}
	return getter.GetAny(destPath, gccLink, getter.WithProgress(&track))
}

func compileGCC() error {
	return errors.New("Not implemented")
}
