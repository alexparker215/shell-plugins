package clickhouse

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

// clickHouseProvisioner provisions the host, user, and password as environment variables
// (the only connection settings clickhouse-client reads from the environment) and, when a
// port is stored on the item, passes it as the --port flag. ClickHouse has no port
// environment variable, so the port can only be applied as a command-line argument.
type clickHouseProvisioner struct {
	envVars sdk.Provisioner
}

func newProvisioner() sdk.Provisioner {
	return clickHouseProvisioner{
		envVars: provision.EnvVars(defaultEnvVarMapping),
	}
}

func (p clickHouseProvisioner) Provision(ctx context.Context, in sdk.ProvisionInput, out *sdk.ProvisionOutput) {
	p.envVars.Provision(ctx, in, out)

	if port := in.ItemFields[fieldname.Port]; port != "" {
		out.AddArgs("--port", port)
	}
}

func (p clickHouseProvisioner) Deprovision(ctx context.Context, in sdk.DeprovisionInput, out *sdk.DeprovisionOutput) {
	p.envVars.Deprovision(ctx, in, out)
}

func (p clickHouseProvisioner) Description() string {
	return "Provision the ClickHouse host, user, and password as environment variables, and the port as a --port flag when set."
}
