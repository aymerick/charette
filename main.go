package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aymerick/charette/harvester"
)

const (
	VERSION = "0.0.1"
)

var (
	// allowed file extensions
	extensions []string

	// flags
	fDir     string
	fMame    bool // @todo Handle that
	fDebug   bool
	fVersion bool
)

func init() {
	// no-intro file extensions
	extensions = []string{".zip", ".7z"}

	// flags
	flag.StringVar(&fDir, "dir", "", "Roms absolute directory (default is current working dir)")
	flag.BoolVar(&fMame, "mame", false, "MAME roms")
	flag.BoolVar(&fDebug, "debug", false, "Activate debug logs")
	flag.BoolVar(&fVersion, "version", false, "Display charette version")
}

func main() {
	var err error

	flag.Parse()

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
	options.Mame = fMame
	options.Debug = fDebug

	// run harvester
	h := harvester.New(fDir, options)
	if err := h.Run(); err != nil {
		panic(err)
	}
}
