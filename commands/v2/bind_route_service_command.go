package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type BindRouteServiceCommand struct {
	RequiredArgs     flags.RouteServiceArgs `positional-args:"yes"`
	ParametersAsJSON string                 `short:"c" description:"Valid JSON object containing service-specific configuration parameters, provided inline or in a file. For a list of supported configuration parameters, see documentation for the particular service offering."`
	Hostname         string                 `long:"hostname" short:"n" description:"Hostname used in combination with DOMAIN to specify the route to bind"`
	Path             string                 `long:"path" description:"Path for the HTTP route"`
	usage            interface{}            `usage:"CF_NAME bind-route-service DOMAIN SERVICE_INSTANCE [--hostname HOSTNAME] [--path PATH] [-c PARAMETERS_AS_JSON]\n\nEXAMPLES:\n   CF_NAME bind-route-service example.com myratelimiter --hostname myapp --path foo\n   CF_NAME bind-route-service example.com myratelimiter -c file.json\n   CF_NAME bind-route-service example.com myratelimiter -c '{\"valid\":\"json\"}'\n\n   In Windows PowerShell use double-quoted, escaped JSON: \"{\\\"valid\\\":\\\"json\\\"}\"\n   In Windows Command Line use single-quoted, escaped JSON: '{\\\"valid\\\":\\\"json\\\"}'"`
	relatedCommands  interface{}            `related_commands:"routes, services"`
}

func (_ BindRouteServiceCommand) Setup(config commands.Config) error {
	return nil
}

func (_ BindRouteServiceCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
