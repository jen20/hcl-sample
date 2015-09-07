package config

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
	hclObj "github.com/hashicorp/hcl/hcl"
	"log"
)

type Config struct {
	Region      string
	AccessKey   string
	SecretKey   string
	Bucket      string
	Directories []DirectoryConfig
}

type DirectoryConfig struct {
	Name                  string
	SourceDirectory       string
	DestinationPrefix     string
	ExcludePatterns       []string
	PreBackupScriptPath   string
	PostBackupScriptPath  string
	PreRestoreScriptPath  string
	PostRestoreScriptPath string
}

func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}
	var errors *multierror.Error

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if rawRegion := hclParseTree.Get("region", false); rawRegion != nil {
		if rawRegion.Len() > 1 {
			errors = multierror.Append(errors, fmt.Errorf("Region was specified more than once in the configuration"))
		} else {
			if rawRegion.Type != hclObj.ValueTypeString {
				errors = multierror.Append(errors, fmt.Errorf("Region was specified as an invalid type in the config - expected string, found %s", rawRegion.Type))
			} else {
				result.Region = rawRegion.Value.(string)
			}
		}
	} else {
		errors = multierror.Append(errors, fmt.Errorf("No region was specified in the configuration"))
	}

	if rawAccessKey := hclParseTree.Get("access_key", false); rawAccessKey != nil {
		if rawAccessKey.Len() > 1 {
			errors = multierror.Append(errors, fmt.Errorf("Access Key was specified more than once in the configuration"))
		} else {
			if rawAccessKey.Type != hclObj.ValueTypeString {
				errors = multierror.Append(errors, fmt.Errorf("Access Key was specified as an invalid type in the config - expected string, found %s", rawAccessKey.Type))
			} else {
				result.AccessKey = rawAccessKey.Value.(string)
			}
		}
	} else {
		errors = multierror.Append(errors, fmt.Errorf("No access key was specified in the configuration"))
	}

	log.Printf("%+v\n", result)

	return result, errors.ErrorOrNil()
}
