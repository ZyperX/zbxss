package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"
)


type Options struct {
	Method, Sql string
}

var o *Options

func init() {
	o = &Options{}

	flag.StringVar(&o.Method, "X", "GET", "")
	flag.StringVar(&o.Method, "method", "GET", "")

	flag.StringVar(&o.Method, "d", "", "")
	flag.StringVar(&o.Method, "data", "", "")

	flag.StringVar(&o.Sql, "sql", "", "")

	flag.Usage = func() {
		showBanner()
		h := []string{
			"",
			"Usage:" + usage,
			"",
			"Options:" + options,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}

	flag.Parse()
}

func ParseOptions() *Options {

	return o
}
