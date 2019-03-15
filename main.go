package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	switch args := os.Args[1:]; len(args) {
	case 1:
		if name := strings.Replace(args[0], "-", "", -1); name == "h" || name == "help" {
			printhelp()
			os.Exit(0)
		}
		os.Exit(1)
	case 2:
		if err := runcommit(strings.Replace(args[0], "-", "", -1), args[1]); err != nil {
			panic(err)
		}
		os.Exit(0)
	default:
		printhelp()
		os.Exit(1)
	}
}

func runcommit(name, val string) error {
	var prefix string
	switch name {
	case "f", "feat":
		prefix = "Feat: "
	case "x", "fix":
		prefix = "Fix: "
	case "d", "docs":
		prefix = "Docs: "
	case "s", "style":
		prefix = "Style: "
	case "r", "refactor":
		prefix = "Refactor: "
	case "p", "perf":
		prefix = "Perf: "
	case "t", "test":
		prefix = "Test: "
	case "c", "chore":
		prefix = "Chore: "
	default:
		printhelp()
		os.Exit(1)
	}
	return exec.Command("git", "commit", "-m", prefix+val).Run()
}

func printhelp() {
	str := `Useage:
  git ng -flag message.
Example:
  git ng -f new feat.
Flags:
  must be one (and the only one) of the following:
  -f --feat: A new feature
  -x --fix: A bug fix
  -d --docs: Documentation only changes
  -s --style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  -r --refactor: A code change that neither fixes a bug or adds a feature
  -p --perf: A code change that improves performance
  -t --test: Adding missing tests
  -c --chore: Changes to the build process or auxiliary tools and libraries such as documentation generation

If you read the source code, you'll know I'm teasing you.`
	fmt.Println(str)
}
