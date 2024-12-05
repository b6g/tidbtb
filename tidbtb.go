package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/zncoder/check"
	"github.com/zncoder/mygo"
)

func main() {
	mygo.ParseFlag("[sql]")

	var sql string
	if flag.NArg() == 0 {
		b := check.V(io.ReadAll(os.Stdin)).F("read from stdin")
		sql = string(b)
	} else {
		sql = flag.Arg(0)
	}

	normSQL, digest := parser.NormalizeDigest(sql)
	fmt.Printf("digest: %s\nnormalized: %s", digest, normSQL)
}
