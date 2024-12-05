package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/tikv/client-go/v2/oracle"
	"github.com/zncoder/check"
	"github.com/zncoder/mygo"
)

type OpList struct{}

func (OpList) Digest() {
	digest()
}

func digest() {
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

func (OpList) TSO() {
	tso()
}

func tso() {
	mygo.ParseFlag("[tso]")

	if flag.NArg() == 0 {
		tso := oracle.GoTimeToTS(time.Now())
		fmt.Println(tso)
	} else {
		tso := check.V(strconv.ParseUint(flag.Arg(0), 10, 64)).F("parse tso")
		t := oracle.GetTimeFromTS(tso)
		fmt.Printf("%d.%d\n%v\n%v\n", t.Unix(), t.Nanosecond()/1e3, t.UTC(), t.Local())
	}
}

func main() {
	mygo.RunOpMapCmd[OpList]()
}
