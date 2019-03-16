package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var m commitmessage
	m.Type = "feat"
	m.Scope = "scope"
	m.Subject = "subject"
	m.Body = "body"
	m.Footer = "footer"
	s, err := parse(m)
	if err != nil {
		fmt.Printf("parse failed, err: %v", err)
		t.Fail()
		return
	}

	fmt.Printf("%v", s)
}
