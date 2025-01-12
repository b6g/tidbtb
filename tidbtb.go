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

func isDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

const timestampFormat = "2006-01-02 15:04:05 -0700 MST"

func tso() {
	mygo.ParseFlag("[tso|timestamp...]")

	if flag.NArg() == 0 {
		tso := oracle.GoTimeToTS(time.Now())
		fmt.Println(tso)
	} else {
		for i, s := range flag.Args() {
			if isDigit(s) {
				tso := check.V(strconv.ParseUint(s, 10, 64)).F("parse tso")
				t := oracle.GetTimeFromTS(tso)
				if i != 0 {
					fmt.Println("----")
				}
				fmt.Printf("%d\n%d.%d\n%v\n%v\n", tso, t.Unix(), t.Nanosecond()/1e3, t.UTC(), t.Local())
			} else {
				t := check.V(time.Parse(timestampFormat, s)).F("parse time", "format", timestampFormat, "value", s)
				tso := oracle.GoTimeToTS(t)
				fmt.Println(tso)
			}
		}
	}
}

func main() {
	mygo.RunOpMapCmd[OpList]()
}
