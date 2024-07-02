/*
 * Copyright 2022 Holoinsight Project Authors. Licensed under Apache-2.0.
 */

package listfiles

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	commonpb "github.com/traas-stack/holoinsight-agent/pkg/server/pb"
	"github.com/traas-stack/holoinsight-agent/pkg/server/registry/pb"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDir(t *testing.T) {
	fmt.Println(filepath.Clean("/saas/g/"))
	fmt.Println(filepath.Dir("/saas/g/"))
}

func TestSplit(t *testing.T) {
	fmt.Println(strings.Split("/a/b/c/", string(os.PathSeparator)))
	resp := &pb.ListFilesResponse{}
	ListFiles(&pb.ListFilesRequest{
		Name:           "/Users/xzchaoo/logs",
		MaxDepth:       3,
		IncludeParents: true,
	}, resp)
	originalNodes := resp.Nodes

	nodes, err := appendPrefixNodes(originalNodes, "/a//b/////")
	if err != nil {
		log.Fatalln(err)
	}
	print("", nodes[0])

	assert.Equal(t, "a", nodes[0].Name)
	assert.Equal(t, "b", nodes[0].Children[0].Name)
	assert.Equal(t, originalNodes[0].Name, nodes[0].Children[0].Children[0].Name)

	nodes, err = Rebase(originalNodes, "/Users/xzchaoo/logs", "/home/admin/logs")
	if err != nil {
		log.Fatalln(err)
	}
	print("", nodes[0])
}

func print(ident string, node *commonpb.FileNode) {
	fmt.Println(ident + node.Name)
	for _, child := range node.Children {
		print(ident+"  ", child)
	}
}

func TestJoin(t *testing.T) {
	fmt.Println(filepath.Clean("/a/b/../../.."))
}
