package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "clickhouse",
		Platform: schema.PlatformInfo{
			Name:     "ClickHouse",
			Homepage: sdk.URL("https://clickhouse.com"),
		},
		Credentials: []schema.CredentialType{
			DatabaseCredentials(),
		},
		Executables: []schema.Executable{
			ClickHouseClient(),
		},
	}
}
