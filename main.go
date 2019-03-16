package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func main() {
	switch args := os.Args[1:]; len(args) {
	case 1:
		if name := strings.Replace(args[0], "-", "", -1); name == "h" || name == "help" {
			printhelp()
			os.Exit(0)
		}
		os.Exit(1)
	case 2, 4, 6, 8:
		if err := runcommit(args); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	default:
		printhelp()
		os.Exit(1)
	}
}

func runcommit(args []string) error {
	m := new(commitmessage)
	for i := 0; i < len(args); i += 2 {
		if err := m.apply(args[i], args[i+1]); err != nil {
			return err
		}
	}
	s, err := parse(m)
	if err != nil {
		return err
	}
	return exec.Command("git", "commit", "-m", s).Run()
}

func printhelp() {
	str := `Useage:
  git ng --type-flag subject [-o scope] [-b body] [-e footer]
Example:
  git ng -f new feat
Flags:
  type flag must be (one and only) one of the following:
  -f --feat: A new feature
  -x --fix: A bug fix
  -d --docs: Documentation only changes
  -s --style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  -r --refactor: A code change that neither fixes a bug or adds a feature
  -p --perf: A code change that improves performance
  -t --test: Adding missing tests
  -c --chore: Changes to the build process or auxiliary tools and libraries such as documentation generation

  optional flags:
  -o scope: Scope can be anything specifying place of the commit change
  -b body: Motivation for the change and contrasts with previous behavior
  -e footer: Breaking changes, referencing issues, etc.

If you read the source code, you'll know I'm teasing you.`
	fmt.Println(str)
}

type commitmessage struct {
	Subject string
	Type    string
	Scope   string
	Body    string
	Footer  string
}

func (m *commitmessage) apply(flag, value string) error {
	switch strings.Replace(flag, "-", "", -1) {
	case "f", "feat":
		return m.typeandsubject("feat", value)
	case "x", "fix":
		return m.typeandsubject("fix", value)
	case "d", "docs":
		return m.typeandsubject("docs", value)
	case "s", "style":
		return m.typeandsubject("style", value)
	case "r", "refactor":
		return m.typeandsubject("refactor", value)
	case "p", "perf":
		return m.typeandsubject("perf", value)
	case "t", "test":
		return m.typeandsubject("test", value)
	case "c", "chore":
		return m.typeandsubject("chore", value)
	case "o", "scope":
		m.Scope = value
	case "b", "body":
		m.Body = value
	case "e", "footer":
		m.Footer = value
	default:
		return fmt.Errorf("unknown flag %s, use -h for help", flag)
	}
	return nil
}

func (m *commitmessage) typeandsubject(t, s string) error {
	if m.Type != "" {
		return fmt.Errorf("more than one type flag, %s, %s", m.Type, t)
	}
	m.Type = t
	m.Subject = s
	return nil
}

const format = `{{.Type}}{{if .Scope}}({{.Scope}}){{end}}: {{.Subject}}{{if .Body}}

{{.Body}}{{end}}{{if .Footer}}

{{.Footer}}{{end}}
`

func parse(m *commitmessage) (string, error) {
	t := template.Must(template.New("").Parse(format))
	var b strings.Builder
	if err := t.Execute(&b, m); err != nil {
		return "", err
	}
	return b.String(), nil
}
