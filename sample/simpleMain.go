package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"fortio.org/version"
)

var shortVersion, longVersion, fullVersion = version.FromBuildInfo()

var programName = "fortio.org/version sample main"

func usage(msg string, args ...interface{}) {
	os := flag.CommandLine.Output()
	fmt.Fprintf(os, "%s %s usage:\n\t%s [flags]\nflags:\n",
		programName, shortVersion, flag.CommandLine.Name())
	flag.PrintDefaults()
	if msg != "" {
		fmt.Fprintf(os, msg, args...)
		fmt.Fprintln(os)
	}
}

func main() { os.Exit(Main()) }

func Main() int {
	buildInfo := flag.Bool("buildinfo", false, "Show full build info and exit.")
	flag.Usage = func() { usage("") }
	flag.Parse()
	if *buildInfo {
		fmt.Print(fullVersion)
		return 0
	}
	numArgs := len(flag.Args())
	if numArgs != 0 {
		usage("No argument expected, got %d.", numArgs)
		return 1
	}
	log.Printf("%s started %s\n", programName, longVersion)
	// ... do something ...
	return 0
}
