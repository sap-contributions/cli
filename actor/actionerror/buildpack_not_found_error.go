package actionerror

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
)

// BuildpackNotFoundError is returned when a requested buildpack is not found.
type BuildpackNotFoundError struct {
	BuildpackName string
	StackName     string
	Lifecycle     constant.AppLifecycleType
}

func (e BuildpackNotFoundError) Error() string {
	return fmt.Sprintf("Buildpack not found - Name: '%s'; Stack: '%s'; Lifecycle: '%s'", e.BuildpackName, e.StackName, e.Lifecycle)
}
