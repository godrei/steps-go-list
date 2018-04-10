package main

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-go-list/gotool"
	"github.com/bitrise-tools/go-steputils/tools"
	glob "github.com/ryanuber/go-glob"
)

func failf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {
	exclude := os.Getenv("exclude")

	log.Infof("Configs:")
	log.Printf("- exclude: %s", exclude)

	if exclude == "" {
		failf("Required input not defined: exclude")
	}

	excludes := strings.Split(exclude, "\n")

	packages, err := gotool.ListPackages()
	if err != nil {
		failf("Failed to list packages: %s", err)
	}

	log.Infof("\nList of packages:")
	var filteredPackages []string

packageLoop:
	for _, p := range packages {
		for _, e := range excludes {
			if glob.Glob(e, p) {
				log.Printf("- %s", p)
				continue packageLoop
			}
		}
		log.Donef("✓ %s", p)
		filteredPackages = append(filteredPackages, p)
	}

	if err := tools.ExportEnvironmentWithEnvman("BITRISE_GO_PACKAGES", strings.Join(filteredPackages, "\n")); err != nil {
		failf("Failed to export packages, error: %s", err)
	}
}
