package cmd

import "flag"

// The naming has been taken from "pflag" i.e. used in "cobra". URL: https://pkg.go.dev/github.com/spf13/pflag#BoolVarP
func flagBoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	flag.BoolVar(p, name, value, usage)
	flag.BoolVar(p, shorthand, value, usage)
}
