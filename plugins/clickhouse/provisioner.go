package clickhouse

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

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
