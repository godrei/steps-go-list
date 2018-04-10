package gotool

import (
	"fmt"
	"strings"

	"github.com/bitrise-io/go-utils/command"
)

func parsePackages(out string) (list []string) {
	for _, l := range strings.Split(string(out), "\n") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		list = append(list, l)
	}
	return list
}

// ListPackages ...
func ListPackages() ([]string, error) {
	cmd := command.New("go", "list", "./...")
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %s", cmd.PrintableCommandArgs(), out)
	}
	return parsePackages(out), nil
}
