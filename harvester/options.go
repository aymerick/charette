package harvester

type Options struct {
	Regions []string
	Mame    bool
	Debug   bool
}

func NewOptions() *Options {
	return &Options{}
}
