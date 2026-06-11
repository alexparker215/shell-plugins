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
		"XML config file": {
			Files: map[string]string{
				"~/.clickhouse-client/config.xml": `<config>
    <host>abc123.us-east-1.aws.clickhouse.cloud</host>
    <port>9440</port>
    <user>default</user>
    <password>hunter2</password>
    <secure>true</secure>
</config>`,
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Host:     "abc123.us-east-1.aws.clickhouse.cloud",
						fieldname.Port:     "9440",
						fieldname.Username: "default",
						fieldname.Password: "hunter2",
					},
				},
			},
		},
		"YAML config file": {
			Files: map[string]string{
				"~/.config/clickhouse/config.yaml": `host: abc123.us-east-1.aws.clickhouse.cloud
port: 9440
user: default
password: 'hunter2'
secure: true`,
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Host:     "abc123.us-east-1.aws.clickhouse.cloud",
						fieldname.Port:     "9440",
						fieldname.Username: "default",
						fieldname.Password: "hunter2",
					},
				},
			},
		},
		"config file without a password is skipped": {
			Files: map[string]string{
				"~/.clickhouse-client/config.yml": `host: localhost
user: default`,
			},
			ExpectedCandidates: nil,
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
