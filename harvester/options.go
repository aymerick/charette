package harvester

// Options holds the settings for Harvester
type Options struct {
	Regions []string
	Mame    bool
	Debug   bool
}

// NewOptions instanciates a new Options
func NewOptions() *Options {
	return &Options{}
}
