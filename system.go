package GoConsole

import (
	"errors"
	"os/user"
	"runtime"
	"strings"

	"github.com/matishsiao/goInfo"
)

var ErrUnableToGetUser error = errors.New("unable to get current user")
var ErrInputValueUnexpected error = errors.New("unexpected input value")

type OS struct {
	Type         string
	Name         string // PRETTY_NAME
	Distro       string // ID
	Version      string // VERSION
	VersionID    string // VERSION_ID
	CodeName     string // VERSION_CODENAME
	Architecture string
}
type System struct {
	OS OS
}

func IsRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, ErrUnableToGetUser
	}
	return currentUser.Username == "root", nil
}

func SystemInfo() (System, error) {

	var err error
	osType := strings.ToLower(runtime.GOOS)
	sys := System{OS: OS{Type: osType}}

	if runtime.GOOS == "linux" {
		getLinuxInfo(&sys)
	} else if runtime.GOOS == "windows" {
		oi, _ := goInfo.GetInfo()
		sys.OS.Distro = oi.OS
	}

	return sys, err
}
