// convert-path is the small tool for converting between tile path and metatile path.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tierpod/go-osm/metatile"
	"github.com/tierpod/go-osm/tile"
)

// Default flags values.
const (
	defaultPrefix = "/var/lib/mod_tile/"
	defaultExt    = ".png"
)
const usage = `Convert between tile and metatile file path.

Usage: convert-path [-prefix PREFIX] [-ext EXT] /path/to/file1.png [path/to/file2.meta]
`

var version string

func main() {
	// Command line flags
	var (
		flagPrefix  string
		flagExt     string
		flagVersion bool
	)

	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}

	flag.StringVar(&flagPrefix, "prefix", defaultPrefix, "Output string `prefix`")
	flag.StringVar(&flagExt, "ext", defaultExt, "Output `extension` for tile (metatile always has 'meta' ext)")
	flag.BoolVar(&flagVersion, "v", false, "Show version and exit")
	flag.Parse()
	paths := flag.Args()

	if flagVersion {
		fmt.Printf("Version: %v\n", version)
		os.Exit(0)
	}

	if len(paths) == 0 {
		fmt.Println("path(s) is not set")
		os.Exit(1)
	}

	for _, p := range paths {
		if strings.HasSuffix(p, ".meta") {
			mt, err := metatile.NewFromURL(p)
			if err != nil {
				fmt.Printf("[ERROR] skip %v: %v\n", p, err)
				continue
			}
			t := tile.New(mt.Zoom, mt.X, mt.Y, flagExt, mt.Style)
			fmt.Println(t.Filepath(flagPrefix))
			continue
		}

		t, err := tile.NewFromURL(p)
		if err != nil {
			fmt.Printf("[ERROR] skip %v: %v\n", p, err)
			continue
		}
		mt := metatile.NewFromTile(t)
		fmt.Println(mt.Filepath(flagPrefix))
	}
}
