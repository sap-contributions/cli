package flag

import (
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	flags "github.com/jessevdk/go-flags"
)

type Lifecycle struct {
	Value constant.AppLifecycleType
}

func (Lifecycle) Complete(prefix string) []flags.Completion {
	return completions([]string{string(constant.AppLifecycleTypeBuildpack), string(constant.AppLifecycleTypeCNB), string(constant.AppLifecycleTypeDocker)}, prefix, false)
}
