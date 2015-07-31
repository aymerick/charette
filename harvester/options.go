package harvester

// Options holds the settings for Harvester
type Options struct {
	Regions      []string
	LeaveMeAlone bool

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
	Unzip   bool
	Scrap   bool
	Sus     bool
}

// NewOptions instanciates a new Options
func NewOptions() *Options {
	return &Options{}
}
