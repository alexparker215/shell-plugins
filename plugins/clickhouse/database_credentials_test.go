package clickhouse

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestDatabaseCredentialsImporter(t *testing.T) {
	plugintest.TestImporter(t, DatabaseCredentials().Importer, map[string]plugintest.ImportCase{
		"default": {
			Environment: map[string]string{
				"CLICKHOUSE_HOST":     "abc123.us-east-1.aws.clickhouse.cloud",
				"CLICKHOUSE_USER":     "default",
				"CLICKHOUSE_PASSWORD": "hunter2",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Host:     "abc123.us-east-1.aws.clickhouse.cloud",
						fieldname.Username: "default",
						fieldname.Password: "hunter2",
					},
				},
			},
		},
	})
}

func TestDatabaseCredentialsProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, DatabaseCredentials().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Host:     "abc123.us-east-1.aws.clickhouse.cloud",
				fieldname.Username: "default",
				fieldname.Password: "hunter2",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"CLICKHOUSE_HOST":     "abc123.us-east-1.aws.clickhouse.cloud",
					"CLICKHOUSE_USER":     "default",
					"CLICKHOUSE_PASSWORD": "hunter2",
				},
			},
		},
		"with port": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Host:     "abc123.us-east-1.aws.clickhouse.cloud",
				fieldname.Port:     "9440",
				fieldname.Username: "default",
				fieldname.Password: "hunter2",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"CLICKHOUSE_HOST":     "abc123.us-east-1.aws.clickhouse.cloud",
					"CLICKHOUSE_USER":     "default",
					"CLICKHOUSE_PASSWORD": "hunter2",
				},
				CommandLine: []string{"--port", "9440"},
			},
		},
	})
}
