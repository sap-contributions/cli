package v7

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/resources"
)

type CreateAppCommand struct {
	BaseCommand

	RequiredArgs    flag.AppName `positional-args:"yes"`
	AppType         flag.AppType `long:"app-type" choice:"buildpack" choice:"docker" choice:"cnb" description:"App lifecycle type to stage and run the app" default:"buildpack"`
	usage           interface{}  `usage:"CF_NAME create-app APP_NAME [--app-type (buildpack | docker | cnb)]"`
	relatedCommands interface{}  `related_commands:"app, apps, push"`
}

func (cmd CreateAppCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Actor.GetCurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Creating app {{.AppName}} in org {{.CurrentOrg}} / space {{.CurrentSpace}} as {{.CurrentUser}}...", map[string]interface{}{
		"AppName":      cmd.RequiredArgs.AppName,
		"CurrentSpace": cmd.Config.TargetedSpace().Name,
		"CurrentOrg":   cmd.Config.TargetedOrganization().Name,
		"CurrentUser":  user.Name,
	})

	cmd.UI.DisplayText(fmt.Sprintf("Using app type %q", constant.AppLifecycleType(cmd.AppType)))

	_, warnings, err := cmd.Actor.CreateApplicationInSpace(
		resources.Application{
			Name:          cmd.RequiredArgs.AppName,
			LifecycleType: constant.AppLifecycleType(cmd.AppType),
		},
		cmd.Config.TargetedSpace().GUID,
	)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		switch err.(type) {
		case ccerror.NameNotUniqueInSpaceError:
			cmd.UI.DisplayText(err.Error())
		default:
			return err
		}
	}

	cmd.UI.DisplayOK()

	return nil
}
