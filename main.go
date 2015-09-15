package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/harvester"
)

const (
	version = "0.0.1"

	defaultRegions = "France,Europe,World,USA,Japan"
	defaultOutput  = "roms"
	defaultTmpDir  = ".~charette"
)

var (
	// flags
	fInput  string
	fOutput string
	fTmpDir string

	fRegions string
	fStrict  bool
	fInsane  bool
	fUnzip   bool

	fKeepProto  bool
	fKeepBeta   bool
	fKeepSample bool
	fKeepDemo   bool
	fKeepPirate bool
	fKeepPromo  bool

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
	flag.StringVar(&fInput, "input", curDir, "Path to no-intro archives directory, or path to a single no-intro archive file")
	flag.StringVar(&fOutput, "output", path.Join(curDir, defaultOutput), "Path to output directory")
	flag.StringVar(&fTmpDir, "tmp", path.Join(curDir, defaultTmpDir), "Path to temporary working directory")

	flag.StringVar(&fRegions, "regions", defaultRegions, "Preferred regions")
	flag.BoolVar(&fStrict, "strict", false, "Skip games that are not in preferred regions")
	flag.BoolVar(&fInsane, "insane", false, "Activates flags: -keep-proto -keep-beta -keep-sample -keep-demo -keep-pirate -keep-promo")
	flag.BoolVar(&fUnzip, "unzip", false, "Unzip roms")

	flag.BoolVar(&fKeepProto, "keep-proto", false, "Keep roms tagged with 'Promo'")
	flag.BoolVar(&fKeepBeta, "keep-beta", false, "Keep roms tagged with 'Beta'")
	flag.BoolVar(&fKeepSample, "keep-sample", false, "Keep roms tagged with 'Sample'")
	flag.BoolVar(&fKeepDemo, "keep-demo", false, "Keep roms tagged with 'Demo'")
	flag.BoolVar(&fKeepPirate, "keep-pirate", false, "Keep roms tagged with 'Pirate'")
	flag.BoolVar(&fKeepPromo, "keep-promo", false, "Keep roms tagged with 'Promo'")

	flag.BoolVar(&fVerbose, "verbose", false, "Activate verbose output")
	flag.BoolVar(&fDebug, "debug", false, "Activate debug logs")
	flag.BoolVar(&fVersion, "version", false, "Display charette version")
}

func main() {
	flag.Parse()

	// check flags
	if fVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if fInsane {
		fKeepProto = true
		fKeepBeta = true
		fKeepSample = true
		fKeepDemo = true
		fKeepPirate = true
		fKeepPromo = true
	}

	if fInput == "" {
		fInput = curDir()
	}

	fInput = path.Clean(fInput)

	if fOutput == "" {
		fOutput = path.Join(curDir(), defaultOutput)
	}

	fOutput = path.Clean(fOutput)

	if fTmpDir == "" {
		fTmpDir = path.Join(curDir(), defaultTmpDir)
	}

	fTmpDir = path.Clean(fTmpDir)

	if fVerbose {
		fmt.Printf("charette v%s\n", version)
		fmt.Printf("   input: %s\n", fInput)
		fmt.Printf("   output: %s\n", fOutput)
		fmt.Printf("   tmp: %s\n", fTmpDir)
	}

	if (fInput == fOutput) || (fInput == fTmpDir) {
		panic("Output and tmp directories can't be the same as input directory")
	}

	// computes options
	options := core.NewOptions()

	options.Input = fInput
	options.Output = fOutput
	options.Tmp = fTmpDir

	options.Regions = core.ExtractRegions(fRegions)

	options.Strict = fStrict

	options.KeepProto = fKeepProto
	options.KeepBeta = fKeepBeta
	options.KeepSample = fKeepSample
	options.KeepDemo = fKeepDemo
	options.KeepPirate = fKeepPirate
	options.KeepPromo = fKeepPromo

	options.Verbose = fVerbose
	options.Debug = fDebug
	options.Unzip = fUnzip

	// run harvester
	h := harvester.New(options)
	if err := h.Run(); err != nil {
		panic(err)
	}
}

// curDir returns current directory
func curDir() string {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return curDir
}
