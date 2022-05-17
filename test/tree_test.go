package main_test

import (
	"testing"

	"github.com/weisrc/nsink"
)

func TestNewTreeIP(t *testing.T) {
	tree := nsink.NewTree("1.2.3.4")
	if tree.Ip != "1.2.3.4" {
		t.Errorf(`expected "1.2.3.4", got %q`, tree.Ip)
	}
}

func TestNullTreeIP(t *testing.T) {
	tree := nsink.NullTree()
	if tree.Ip != "" {
		t.Errorf(`expected "", got %q`, tree.Ip)
	}
}

func TestTreeInsert(t *testing.T) {
	tree := nsink.NullTree()
	tree.Insert("example.org.", "0.0.0.0")
	org, ok := tree.Nodes["org"]
	if !ok {
		t.Errorf(`expected tree["org"] to exist`)
	}
	example, ok1 := org.Nodes["example"]
	if !ok1 {
		t.Errorf(`expected tree["example"] to exist`)
	}
	if example.Ip != "0.0.0.0" {
		t.Errorf(`expected tree["org"]["example"].Ip to be "0.0.0.0", got %q`, example.Ip)
	}
}
