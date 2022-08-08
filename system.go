package GoConsole

import (
	"errors"
	"os/user"

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
	si.GetSysInfo()
	return si
}
