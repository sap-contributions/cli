package flag

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	flags "github.com/jessevdk/go-flags"
)

type AppType struct {
	Value constant.AppLifecycleType
}

func (AppType) Complete(prefix string) []flags.Completion {
	return completions([]string{string(constant.AppLifecycleTypeBuildpack), string(constant.AppLifecycleTypeCNB), string(constant.AppLifecycleTypeDocker)}, prefix, false)
}
