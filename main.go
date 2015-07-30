package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aymerick/charette/harvester"
	"github.com/aymerick/charette/utils"
)

const (
	VERSION = "0.0.1"

	DEFAULT_REGIONS = "France,Europe,World,USA,Japan"
)

var (
	// allowed file extensions
	extensions []string

	// flags
	fDir     string
	fRegions string
	fMame    bool // @todo Handle that
	fDebug   bool
	fVersion bool
)

func init() {
	// no-intro file extensions
	extensions = []string{".zip", ".7z"}

	// flags
	flag.StringVar(&fDir, "dir", "", "Roms absolute directory (default is current working dir)")
	flag.StringVar(&fRegions, "regions", DEFAULT_REGIONS, fmt.Sprintf("Preferred regions (default: %s)", DEFAULT_REGIONS))
	flag.BoolVar(&fMame, "mame", false, "MAME roms")
	flag.BoolVar(&fDebug, "debug", false, "Activate debug logs")
	flag.BoolVar(&fVersion, "version", false, "Display charette version")
}

func main() {
	var err error

	flag.Parse()

	// check flags
	if fVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if fDir == "" {
		fDir, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	if fDebug {
		log.Printf("charette v%s", VERSION)
		log.Printf("   dir: %s", fDir)
		log.Printf("   mame: %v", fMame)
	}

	// computes options
	options := harvester.NewOptions()

	options.Regions = utils.ExtractRegions(fRegions)
	options.Mame = fMame
	options.Debug = fDebug

	// run harvester
	h := harvester.New(fDir, options)
	if err := h.Run(); err != nil {
		panic(err)
	}
}
