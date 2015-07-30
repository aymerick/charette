package harvester

// Options holds the settings for Harvester
type Options struct {
	Regions []string

	NoProto  bool
	NoBeta   bool
	NoSample bool
	NoDemo   bool
	NoPirate bool
	NoPromo  bool

	Mame    bool
	Verbose bool
	Debug   bool
	Noop    bool
}

// NewOptions instanciates a new Options
func NewOptions() *Options {
	return &Options{}
}
