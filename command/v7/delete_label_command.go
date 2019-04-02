package v7

import (
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
)

type DeleteLabelCommand struct {
	RequiredArgs flag.DeleteLabelArgs `positional-args:"yes"`
	usage        interface{}          `usage:"CF_NAME delete-label RESOURCE RESOURCE_NAME KEY\n\n EXAMPLES:\n   cf delete-label app dora ci_signature_sha2\n\nRESOURCES:\n   APP\n\nSEE ALSO:\n   set-label, labels"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
}

func (cmd DeleteLabelCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}
	return nil
}
