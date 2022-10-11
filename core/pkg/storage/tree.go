package storage

import (
	"github.com/obarbier/custom-app/core/pkg/models"
	"strings"
)

type node struct {
	children    map[string]*node
	isEndOfPath bool
	acls        models.PolicyAnon
}

func newNode() *node {
	return &node{
		children:    make(map[string]*node),
		isEndOfPath: false,
		acls:        models.PolicyAnon{},
	}
}

func (n *node) insert(path string, acl models.PolicyAnon) {
	curr := n
	for _, r := range strings.Split(path, "/") {
		_, ok := curr.children[string(r)]
		if !ok {
			curr.children[r] = newNode()
		}
		curr = curr.children[r]
	}
	curr.isEndOfPath = true
	curr.acls = acl
}

func (n *node) isInTree(path string) bool {
	// TODO(obarbier): how to check against wildcard path
	curr := n
	for _, r := range strings.Split(path, "/") {
		node, ok := curr.children[r]
		if !ok {
			return false
		}
		curr = node
	}
	return curr.isEndOfPath
}

func (n *node) getPolicy(path string) *node {
	curr := n
	for _, r := range strings.Split(path, "/") {
		node, ok := curr.children[r]
		if !ok {
			return nil
		}
		curr = node
	}
	if curr.isEndOfPath {
		return curr
	}
	return nil
}
