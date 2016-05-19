package config

import (
	"log"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl"
)

// Config type
type Config struct {
	Region      string            `hcl:"region"`
	AccessKey   string            `hcl:"access_key"`
	SecretKey   string            `hcl:"secret_key"`
	Bucket      string            `hcl:"bucket"`
	Directories []DirectoryConfig `hcl:"directory"`
}

// DirectoryConfig type
type DirectoryConfig struct {
	Name                  string   `hcl:",key"`
	SourceDirectory       string   `hcl:"source_dir"`
	DestinationPrefix     string   `hcl:"dest_prefix"`
	ExcludePatterns       []string `hcl:"exclude"`
	PreBackupScriptPath   string   `hcl:"pre_backup_script"`
	PostBackupScriptPath  string   `hcl:"post_backup_script"`
	PreRestoreScriptPath  string   `hcl:"pre_restore_script"`
	PostRestoreScriptPath string   `hcl:"post_restore_script"`
}

// ParseConfig parse the given HCL string into a Config struct.
func ParseConfig(hclText string) (*Config, error) {
	result := &Config{}
	var errors *multierror.Error

	hclParseTree, err := hcl.Parse(hclText)
	if err != nil {
		return nil, err
	}

	if err := hcl.DecodeObject(&result, hclParseTree); err != nil {
		return nil, err
	}

	log.Printf("%+v\n", result)

	return result, errors.ErrorOrNil()
}
