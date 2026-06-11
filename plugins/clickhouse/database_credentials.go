package clickhouse

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func DatabaseCredentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.DatabaseCredentials,
		DocsURL:       sdk.URL("https://clickhouse.com/docs/operations/access-rights"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Host,
				MarkdownDescription: "Host of the ClickHouse server to connect to.",
			},
			{
				Name:                fieldname.Port,
				MarkdownDescription: "Port of the ClickHouse server to connect to.",
				Optional:            true,
			},
			{
				Name:                fieldname.Username,
				MarkdownDescription: "ClickHouse user to authenticate as.",
			},
			{
				Name:                fieldname.Password,
				MarkdownDescription: "Password used to authenticate to ClickHouse.",
				Secret:              true,
			},
		},
		DefaultProvisioner: newProvisioner(),
		Importer:           importer.TryEnvVarPair(defaultEnvVarMapping),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"CLICKHOUSE_HOST":     fieldname.Host,
	"CLICKHOUSE_USER":     fieldname.Username,
	"CLICKHOUSE_PASSWORD": fieldname.Password,
}
