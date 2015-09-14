package core

// Options holds the settings for Harvester
type Options struct {
	Regions      []string
	LeaveMeAlone bool

	KeepProto  bool
	KeepBeta   bool
	KeepSample bool
	KeepDemo   bool
	KeepPirate bool
	KeepPromo  bool

	Verbose bool
	Debug   bool
	Noop    bool
	Unzip   bool
	Tmp     string
}

// NewOptions instanciates a new Options
func NewOptions() *Options {
	return &Options{}
}
