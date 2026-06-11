package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

// ClickHouseClient covers the `clickhouse-client` symlink (a single binary invocation).
func ClickHouseClient() schema.Executable {
	return schema.Executable{
		Name:      "ClickHouse client",
		Runs:      []string{"clickhouse-client"},
		DocsURL:   sdk.URL("https://clickhouse.com/docs/integrations/sql-clients/cli"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.DatabaseCredentials,
			},
		},
	}
}
