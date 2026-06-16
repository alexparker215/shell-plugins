package clickhouse

import (
	"context"
	"strings"

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

	if isTruthy(in.ItemFields[fieldname.Secure]) {
		out.AddArgs("--secure")
	}
}

func (p clickHouseProvisioner) Deprovision(ctx context.Context, in sdk.DeprovisionInput, out *sdk.DeprovisionOutput) {
	p.envVars.Deprovision(ctx, in, out)
}

func (p clickHouseProvisioner) Description() string {
	return "Provision the ClickHouse host, user, and password as environment variables, and the port and TLS settings as command-line flags when set."
}

func isTruthy(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "true", "1", "yes", "on":
		return true
	}
	return false
}
