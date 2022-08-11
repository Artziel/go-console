package GoConsole

import (
	"os"
	"regexp"
	"strings"
)

func getLinuxOSReleaseValue(source string, name string) string {
	result := ""

	r := regexp.MustCompile(`(\A|\r|\n|\r\n|\s.?)` + name + `=(.*)\n`)
	match := r.FindStringSubmatch(source)
	if len(match) > 0 {
		result = strings.ReplaceAll(match[2], "\"", "")
	}

	return result
}

func getLinuxInfo(sys *System) {
	if _, err := os.Stat("/lib64/ld-linux-x86-64.so.2"); err == nil {
		sys.OS.Architecture = "amd64"
	} else if _, err := os.Stat("/lib/ld-linux.so.2"); err == nil {
		sys.OS.Architecture = "i386"
	}

	content := ""
	var err error

	b, e := os.ReadFile("/etc/os-release")
	if e == nil {
		content = string(b)
	}

	if err == nil {
		sys.OS.Name = getLinuxOSReleaseValue(content, "PRETTY_NAME")
		sys.OS.Distro = getLinuxOSReleaseValue(content, "ID")
		sys.OS.Version = getLinuxOSReleaseValue(content, "VERSION")
		sys.OS.VersionID = getLinuxOSReleaseValue(content, "VERSION_ID")
		sys.OS.CodeName = getLinuxOSReleaseValue(content, "VERSION_CODENAME")
	}

}
