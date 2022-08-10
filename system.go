package GoConsole

import (
	"errors"
	"os/user"
	"runtime"

	"github.com/zcalusic/sysinfo"
)

var ErrUnableToGetUser error = errors.New("unable to get current user")
var ErrInputValueUnexpected error = errors.New("unexpected input value")

func IsRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, ErrUnableToGetUser
	}
	return currentUser.Username == "root", nil
}

func GetSystemInfo() sysinfo.SysInfo {
	var si sysinfo.SysInfo
	if runtime.GOOS == "linux" {
		si.GetSysInfo()
	} else {
		si.OS = sysinfo.OS{
			Vendor: "darwin",
		}
	}
	return si
}
