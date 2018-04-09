package envman

import (
	"strings"

	"github.com/bitrise-io/go-utils/command"
)

// Export ...
func Export(keyStr, valueStr string) error {
	cmd := command.New("envman", "add", "--key", keyStr)
	cmd.SetStdin(strings.NewReader(valueStr))
	return cmd.Run()
}
