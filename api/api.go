package api

import (
	"github.com/project-flogo/cli/util"
)

var verbose = false

func SetVerbose(enable bool) {
	verbose = enable
	util.SetVerbose(enable)
}

func Verbose() bool {
	return verbose
}
