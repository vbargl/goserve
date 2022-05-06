package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vbargl/goserve/pkg/http"
	gs_os "github.com/vbargl/goserve/pkg/os"
)

var address = flag.String("address", "", "specific address to use (instead of port)")
var port = flag.Int("port", http.GetPort(49152, 65535), "specific port to use (instead of address)")
var public = flag.Bool("public", false, "use 0.0.0.0 instead of localhost")

func main() {
	defer catchErr()
	flag.Parse()

	var addr string
	switch {
	case *address != "":
		addr = *address
	case *public:
		addr = fmt.Sprintf("0.0.0.0:%d", *port)
	default:
		addr = fmt.Sprintf("localhost:%d", *port)
	}

	http.Serve(addr, getPath(flag.Args()))
}

func catchErr() {
	if r := recover(); r != nil {
		log.Fatal(r)
	}
}

func getPath(args []string) string {
	switch len(args) {
	case 0:
		return gs_os.GetWorkingDirectory()
	case 1:
		return gs_os.GetDirectory(args[0])
	default:
		panic("expected single optional parameter path")
	}
}
