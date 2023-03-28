package main

import (
	"github.com/bwmarrin/snowflake"
	"testing"
)

func BenchGenerate(b *testing.B) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node.Generate()
	}
}
