package main

import (
	"flag"
	"fmt"
	"github.com/velour/bred"
	"os"
)

var (
	listenAddr = flag.String("addr", "localhost:8001", "HTTP listen address")
	tplDir = flag.String("tpl", "./tpl", "template directory path")
)

func main() {
	flag.Parse()
	fmt.Fprintf(os.Stderr, "Listening on http://%s\n", *listenAddr)
	bred.New(*listenAddr, *tplDir).Start()
}
