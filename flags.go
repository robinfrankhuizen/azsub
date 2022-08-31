package main

import "flag"

type Flags struct {
	Verbose bool
}

func parseFlags() Flags {
	v := flag.Bool("v", false, "Show subscription ids")
	flag.Parse()

	return Flags{
		Verbose: *v,
	}
}
