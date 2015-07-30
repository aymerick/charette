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
	fSane    bool

	fNoProto  bool
	fNoBeta   bool
	fNoSample bool
	fNoDemo   bool
	fNoPirate bool
	fNoPromo  bool

	fVerbose bool
	fDebug   bool
	fVersion bool
)

func init() {
	// no-intro file extensions
	extensions = []string{".zip", ".7z"}

	// flags
	flag.StringVar(&fDir, "dir", "", "Roms absolute directory (default is current working dir)")
	flag.StringVar(&fRegions, "regions", DEFAULT_REGIONS, "Preferred regions")
	flag.BoolVar(&fMame, "mame", false, "MAME roms")
	flag.BoolVar(&fSane, "sane", false, "Activates flags: -no-proto, -no-beta, -no-sample, -no-demo, -no-pirate, -no-promo")

	flag.BoolVar(&fNoProto, "no-proto", false, "Skip roms tagged with 'Promo'")
	flag.BoolVar(&fNoBeta, "no-beta", false, "Skip roms tagged with 'Beta'")
	flag.BoolVar(&fNoSample, "no-sample", false, "Skip roms tagged with 'Sample'")
	flag.BoolVar(&fNoDemo, "no-demo", false, "Skip roms tagged with 'Demo'")
	flag.BoolVar(&fNoPirate, "no-pirate", false, "Skip roms tagged with 'Pirate'")
	flag.BoolVar(&fNoPromo, "no-promo", false, "Skip roms tagged with 'Promo'")

	flag.BoolVar(&fVerbose, "verbose", false, "Activate verbose output")
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

	if fSane {
		fNoProto = true
		fNoBeta = true
		fNoSample = true
		fNoDemo = true
		fNoPirate = true
		fNoPromo = true
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

	options.NoProto = fNoProto
	options.NoBeta = fNoBeta
	options.NoSample = fNoSample
	options.NoDemo = fNoDemo
	options.NoPirate = fNoPirate
	options.NoPromo = fNoPromo

	options.Mame = fMame
	options.Verbose = fVerbose
	options.Debug = fDebug

	// run harvester
	h := harvester.New(fDir, options)
	if err := h.Run(); err != nil {
		panic(err)
	}
}
