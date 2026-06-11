package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func DatabaseCredentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.DatabaseCredentials,
		DocsURL:       sdk.URL("https://clickhouse.com/docs/integrations/sql-clients/cli"),
		ManagementURL: sdk.URL("https://console.clickhouse.cloud"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Host,
				MarkdownDescription: "Host of the ClickHouse service to connect to, e.g. abc123.us-east-1.aws.clickhouse.cloud. For any host ending in .clickhouse.cloud the client automatically enables TLS and connects on port 9440.",
			},
			{
				Name:                fieldname.Username,
				MarkdownDescription: "ClickHouse user to authenticate as. For ClickHouse Cloud services this is `default` unless you created another user.",
			},
			{
				Name:                fieldname.Password,
				MarkdownDescription: "Password used to authenticate to ClickHouse.",
				Secret:              true,
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer:           importer.TryEnvVarPair(defaultEnvVarMapping),
	}
}

// Only these three environment variables are read by the clickhouse client. CLICKHOUSE_PORT
// and CLICKHOUSE_DATABASE are intentionally omitted: the client ignores them, and ClickHouse
// Cloud hosts already imply port 9440 + TLS from the .clickhouse.cloud suffix.
var defaultEnvVarMapping = map[string]sdk.FieldName{
	"CLICKHOUSE_HOST":     fieldname.Host,
	"CLICKHOUSE_USER":     fieldname.Username,
	"CLICKHOUSE_PASSWORD": fieldname.Password,
}
