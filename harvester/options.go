package harvester

type Options struct {
	Mame  bool
	Debug bool
}

func NewOptions() *Options {
	return &Options{}
}
