package initializer

import (
	"flag"
)

var (
	flConfigPath *string
	FlVersion    *bool
)

func init() {
	flConfigPath = flag.String("c", "", "app conf file")
	FlVersion = flag.Bool("v", false, "show version")
}
