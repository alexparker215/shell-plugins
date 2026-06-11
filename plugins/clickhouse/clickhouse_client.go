package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func ClickHouseClient() schema.Executable {
	return schema.Executable{
		Name:      "ClickHouse Client",
		Runs:      []string{"clickhouse-client"},
		DocsURL:   sdk.URL("https://clickhouse.com/docs/interfaces/client"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.DatabaseCredentials,
			},
		},
	}
}
