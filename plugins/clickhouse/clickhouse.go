package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

// ClickHouse covers the `clickhouse client` subcommand form of the all-in-one binary.
// Matching is on the full command string, so only `clickhouse client ...` is intercepted,
// not other subcommands such as `clickhouse local` or `clickhouse server`.
func ClickHouse() schema.Executable {
	return schema.Executable{
		Name:      "ClickHouse",
		Runs:      []string{"clickhouse", "client"},
		DocsURL:   sdk.URL("https://clickhouse.com/docs/integrations/sql-clients/cli"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.DatabaseCredentials,
			},
		},
	}
}
