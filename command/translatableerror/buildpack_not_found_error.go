package translatableerror

import "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"

type BuildpackNotFoundError struct {
	BuildpackName string
	StackName     string
	Lifecycle     constant.AppLifecycleType
}

func (e BuildpackNotFoundError) Error() string {
	if len(e.StackName) == 0 {
		return "Buildpack '{{.BuildpackName}}' not found"
	}
	return "Buildpack '{{.BuildpackName}}' with stack '{{.StackName}}' and lifecycle '{{.Lifecycle}}' not found"
}

func (e BuildpackNotFoundError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"BuildpackName": e.BuildpackName,
		"StackName":     e.StackName,
		"Lifecycle":     e.Lifecycle,
	})
}
