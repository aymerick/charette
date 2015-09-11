package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/harvester"
)

const (
	version = "0.0.1"

	defaultRegions = "France,Europe,World,USA,Japan"
	defaultGarbage = "_GARBAGE_"
)

var (
	// flags
	fDir     string
	fGarbage string
	fNoop    bool

	fRegions      string
	fLeaveMeAlone bool
	fMame         bool
	fSane         bool
	fUnzip        bool

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
	// get current directory
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// flags
	flag.StringVar(&fDir, "dir", curDir, "Roms absolute directory")
	flag.StringVar(&fGarbage, "garbage", "", "Garbage absolute directory (default is '<dir>/_GARBAGE_'")
	flag.BoolVar(&fNoop, "noop", false, "Noop mode: do nothing, usefull for debugging")

	flag.StringVar(&fRegions, "regions", defaultRegions, "Preferred regions")
	flag.BoolVar(&fLeaveMeAlone, "leave-me-alone", false, "Skip games that are not in preferred regions")
	flag.BoolVar(&fMame, "mame", false, "MAME roms")
	flag.BoolVar(&fSane, "sane", false, "Activates flags: -no-proto -no-beta -no-sample -no-demo -no-pirate -no-promo")
	flag.BoolVar(&fUnzip, "unzip", false, "Unzip roms")

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
		fmt.Println(version)
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

	if fGarbage == "" {
		fGarbage = path.Join(fDir, defaultGarbage)
	}

	if fDebug {
		log.Printf("charette v%s", version)
		log.Printf("   dir: %s", fDir)
		log.Printf("   mame: %v", fMame)
	}

	// computes options
	options := harvester.NewOptions()

	options.Regions = core.ExtractRegions(fRegions)

	options.LeaveMeAlone = fLeaveMeAlone

	options.NoProto = fNoProto
	options.NoBeta = fNoBeta
	options.NoSample = fNoSample
	options.NoDemo = fNoDemo
	options.NoPirate = fNoPirate
	options.NoPromo = fNoPromo

	options.Mame = fMame
	options.Verbose = fVerbose
	options.Debug = fDebug
	options.Noop = fNoop
	options.Unzip = fUnzip

	// run harvester
	h := harvester.New(fDir, fGarbage, options)
	if err := h.Run(); err != nil {
		panic(err)
	}
}
