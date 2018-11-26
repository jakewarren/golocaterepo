package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
)

func main() {

	verbose := flag.BoolP("verbose", "v", false, "Verbose error output")
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(0)
	}
	switch flag.Args()[0] {
	// easier than defining four flags and checking for the string "help" also:
	case "-h", "-help", "--help", "-?", "help":
		flag.Usage()
		os.Exit(0)
	}

	exec := flag.Arg(0)

	// Determine the location of `exec`
	path, err := getExecPath(exec)
	if err != nil {
		log.Println(exec + ": command not found")
		if *verbose {
			log.Println(errors.WithStack(err))
		}
		return
	}

	// Extract the source path from the binary
	src, err := getMainPath(path)
	if err != nil {
		log.Println("No source path in", path, "-", exec, "is perhaps not a Go binary")
		if *verbose {
			log.Println(errors.WithStack(err))
		}
		return
	}

	fmt.Println(src)
}

func init() {
	flag.Usage = func() {
		fmt.Printf("Description: %s\n\n", "prints the source repository found in the strings of a Go binary.")
		fmt.Printf("Usage: %s [binary in your PATH]\n\n", os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}
}

// getExecPath receives the name of an executable and determines its path
// based on $PATH, $GOPATH, or the current directory (in this order).
func getExecPath(name string) (string, error) {

	// Try $PATH first.
	s, err := exec.LookPath(name)
	if err == nil {
		return s, nil
	}

	// Next, try $GOPATH/bin
	paths := gopath()
	for i := 0; s == "" && i < len(paths); i++ {
		s, err = exec.LookPath(filepath.Join(paths[i], name))
	}
	if err == nil {
		return s, nil
	}

	// Finally, try the current directory.
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "Unable to determine current directory")
	}
	s, err = exec.LookPath(filepath.Join(wd, name))
	if err != nil {
		return "", errors.New(name + " not found in any of " + os.Getenv("PATH") + ":" + strings.Join(paths, ":") + ":" + wd)
	}

	return s, nil
}

// gopath returns a list of paths as defined by the GOPATH environment
// variable, or the default gopath if $GOPATH is empty.
func gopath() []string {

	gp := os.Getenv("GOPATH")
	if gp == "" {
		return []string{build.Default.GOPATH}
	}

	return strings.Split(gp, pathssep())
}

// pathssep returns the separator between the paths of $PATH or %PATH%.
func pathssep() string {

	sep := ":"
	if runtime.GOOS == "windows" {
		sep = ";"
	}

	return sep
}
