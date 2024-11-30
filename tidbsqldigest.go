package main

import (
	"flag"
	"fmt"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/zncoder/mygo"
)

func main() {
	mygo.ParseFlag("sql")
	sql := flag.Arg(0)

	normSQL, digest := parser.NormalizeDigest(sql)
	fmt.Printf("digest: %s\nnormalized: %s", digest, normSQL)
}
