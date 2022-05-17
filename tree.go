package nsink

import (
	"strings"
)

type Tree struct {
	Nodes map[string]Tree
	Ip    string
}

func NewTree(ip string) Tree {
	t := Tree{}
	t.Ip = ip
	t.Nodes = make(map[string]Tree)
	return t
}

func NullTree() Tree {
	return NewTree("")
}

func (t *Tree) traverse(name []string, i int, stop int, write bool) *Tree {
	if i < stop {
		return t
	}
	if len(t.Nodes) == 0 && !write {
		return t
	}
	part := name[i]
	node, ok := t.Nodes[part]
	if !ok {
		if write {
			node = NullTree()
			t.Nodes[part] = node
		} else {
			return nil
		}
	}
	return node.traverse(name, i-1, stop, write)
}

func (t *Tree) Traverse(name string, write bool) *Tree {
	split := strings.Split(name, ".")
	return t.traverse(split, len(split)-2, 0, write)
}

func (t *Tree) Find(name string) *Tree {
	return t.Traverse(name, false)
}

func (t *Tree) FindIP(name string) string {
	found := t.Find(name)
	if found != nil {
		return found.Ip
	}
	return ""
}

func (t *Tree) Insert(name string, ip string) {
	split := strings.Split(name, ".")
	t.traverse(split, len(split)-2, 1, true).Nodes[split[0]] = NewTree(ip)
}

func (t *Tree) Delete(name string) {
	split := strings.Split(name, ".")
	delete(t.traverse(split, len(split)-2, 1, false).Nodes, split[0])
}

func (t *Tree) String(name string, tab int) string {
	lines := []string{strings.Repeat("  ", tab) + "- " + name + "@" + t.Ip}
	for sub, node := range t.Nodes {
		lines = append(lines, node.String(sub, tab+1))
	}
	return strings.Join(lines, "\n")
}
