package infrastructure

import kingpin "gopkg.in/alecthomas/kingpin.v2"

// Options represents option parameters.
type Options struct {
	Port *int
}

// Opts represents option parameters.
var Opts = parse()

func parse() *Options {
	o := &Options{
		Port: kingpin.Flag("port", "Server port").Short('p').Default("1234").Int(),
	}
	kingpin.Parse()
	return o
}
