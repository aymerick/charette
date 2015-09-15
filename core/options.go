package core

// Options holds the settings for Harvester
type Options struct {
	Input  string
	Output string
	Tmp    string

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
	Unzip   bool
}

// NewOptions instanciates a new Options
func NewOptions() *Options {
	return &Options{}
}
