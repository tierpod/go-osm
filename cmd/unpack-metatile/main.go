// Unpack tiles data from metatile(s) to directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/tierpod/go-osm/metatile"
	"github.com/tierpod/go-osm/point"
)

const defaultExt = ".png"
const usage = `Unpack tiles data from metatiles to directory.

Usage: unpack-metatiles -dir DIR [-ext EXT] /path/to/file1 [path/to/file2]
`

var version string

func main() {
	// Command line flags
	var (
		flagDir     string
		flagExt     string
		flagVersion bool
	)

	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}

	flag.StringVar(&flagDir, "dir", "", "output directory")
	flag.StringVar(&flagExt, "ext", defaultExt, "append extension")
	flag.BoolVar(&flagVersion, "v", false, "Show version and exit")
	flag.Parse()
	files := flag.Args()

	if flagVersion {
		fmt.Printf("Version: %v\n", version)
		os.Exit(0)
	}

	if flagDir == "" {
		fmt.Println("-dir is not set")
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("files(s) is not set")
		os.Exit(1)
	}

	for _, f := range files {
		fmt.Printf("unpack %v to %v\n", f, flagDir)

		f, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		decoder, err := metatile.NewDecoder(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := decoder.TilesMap()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = writeToFile(data, flagDir, flagExt)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("----------")
	}
}

func writeToFile(data map[point.ZXY][]byte, dir, ext string) error {
	for p, d := range data {
		subDir := path.Join(dir, strconv.Itoa(p.Z), strconv.Itoa(p.X))
		err := os.MkdirAll(subDir, 0777)
		if err != nil {
			return nil
		}
		fName := path.Join(subDir, strconv.Itoa(p.Y)+ext)

		err = ioutil.WriteFile(fName, d, 0666)
		if err != nil {
			return err
		}
	}

	return nil
}
