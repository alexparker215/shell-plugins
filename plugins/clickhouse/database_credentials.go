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
		Name:    credname.DatabaseCredentials,
		DocsURL: sdk.URL("https://clickhouse.com/docs/operations/access-rights"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Host,
				MarkdownDescription: "The hostname of the ClickHouse server to connect to. Can either be a hostname or an IPv4 or IPv6 address.",
			},
			{
				Name:                fieldname.Port,
				MarkdownDescription: "The port the server is accepting connections on. The default ports are 9440 (TLS) and 9000 (no TLS).",
				Optional:            true,
			},
			{
				Name:                fieldname.User,
				MarkdownDescription: "The database user to connect as.",
			},
			{
				Name:                fieldname.Password,
				MarkdownDescription: "The password of the database user.",
				Secret:              true,
			},
			{
				Name:                fieldname.Secure,
				MarkdownDescription: "Whether to use TLS. Enabled automatically when connecting to port 9440 (the default secure port) or ClickHouse Cloud.",
				Optional:            true,
			},
		},
		DefaultProvisioner: newProvisioner(),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TryClickHouseConfigFile(),
		),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"CLICKHOUSE_HOST":     fieldname.Host,
	"CLICKHOUSE_USER":     fieldname.User,
	"CLICKHOUSE_PASSWORD": fieldname.Password,
}
