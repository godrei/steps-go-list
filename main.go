package main

import (
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"github.com/godrei/steps-go-list/envman"
	"github.com/godrei/steps-go-list/gotool"
)

// Config ...
type Config struct {
	Exclude string `env:"exclude"`
}

func failf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {
	var cfg Config
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Error: %s\n", err)
	}
	stepconf.Print(cfg)

	dir, err := os.Getwd()
	if err != nil {
		failf("Failed to get working directory: %s", err)
	}

	excludes := strings.Split(cfg.Exclude, ",")

	packages, err := gotool.ListPackages(dir, excludes...)
	if err != nil {
		failf("Failed to list packages: %s", err)
	}

	log.Infof("\nList of packages:")
	for _, p := range packages {
		log.Printf("- %s", p)
	}

	if err := envman.Export("BITRISE_GO_PACKAGES", strings.Join(packages, ",")); err != nil {
		failf("Failed to export packages, error: %s", err)
	}
}
