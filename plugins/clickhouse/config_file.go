package clickhouse

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TryClickHouseConfigFile() sdk.Importer {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = "~/.config"
	}

	basePaths := []string{
		"clickhouse-client",
		xdgConfigHome + "/clickhouse/config",
		"~/.clickhouse-client/config",
		"/etc/clickhouse-client/config",
	}

	var importers []sdk.Importer
	for _, base := range basePaths {
		for _, ext := range []string{".xml", ".yaml", ".yml"} {
			path := base + ext
			importers = append(importers, importer.TryFile(path, importConfigFile(path)))
		}
	}
	return importer.TryAll(importers...)
}

type clickHouseConfigFile struct {
	Host     string         `xml:"host" yaml:"host"`
	Port     flexibleString `xml:"port" yaml:"port"`
	User     string         `xml:"user" yaml:"user"`
	Password string         `xml:"password" yaml:"password"`
	Secure   flexibleString `xml:"secure" yaml:"secure"`
}

func importConfigFile(path string) func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
	return func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		var config clickHouseConfigFile

		var err error
		if strings.HasSuffix(path, ".xml") {
			err = contents.ToXML(&config)
		} else {
			err = contents.ToYAML(&config)
		}
		if err != nil {
			out.AddError(err)
			return
		}

		if config.Password == "" {
			return
		}

		fields := map[sdk.FieldName]string{
			fieldname.Password: config.Password,
		}
		if config.Host != "" {
			fields[fieldname.Host] = config.Host
		}
		if config.Port != "" {
			fields[fieldname.Port] = string(config.Port)
		}
		if config.User != "" {
			fields[fieldname.User] = config.User
		}
		if isTruthy(string(config.Secure)) {
			fields[fieldname.Secure] = "true"
		}

		out.AddCandidate(sdk.ImportCandidate{
			Fields: fields,
		})
	}
}

type flexibleString string

func (s *flexibleString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value interface{}
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		*s = ""
		return nil
	}
	*s = flexibleString(fmt.Sprintf("%v", value))
	return nil
}
