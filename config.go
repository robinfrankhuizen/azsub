package main

type Config struct {
	Verbose   bool
	NameWidth int
}

func createConfig(subs []Subscription, flags Flags) Config {
	return Config{
		Verbose:   flags.Verbose,
		NameWidth: maxSubName(subs),
	}
}
