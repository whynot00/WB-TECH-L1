package flags

import "flag"

func ParseFlag() int {

	flags := *flag.Int("wc", 5, "set workers count")

	flag.Parse()
	return flags
}
