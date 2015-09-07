package config

import (
	"reflect"
	"testing"
)

func TestConfigParsing(t *testing.T) {
	expected := &Config{
		Region:    "us-west-2",
		AccessKey: "something",
		SecretKey: "something_else",
		Bucket:    "backups",
		Directories: []DirectoryConfig{
			DirectoryConfig{
				Name:                  "config",
				SourceDirectory:       "/etc/eventstore",
				DestinationPrefix:     "escluster/config",
				ExcludePatterns:       []string{},
				PreBackupScriptPath:   "before_backup.sh",
				PostBackupScriptPath:  "after_backup.sh",
				PreRestoreScriptPath:  "before_restore.sh",
				PostRestoreScriptPath: "after_restore.sh",
			},
			DirectoryConfig{
				Name:                  "data",
				SourceDirectory:       "/var/lib/eventstore",
				DestinationPrefix:     "escluster/a/data",
				ExcludePatterns:       []string{"*.merging"},
				PreBackupScriptPath:   "",
				PostBackupScriptPath:  "",
				PreRestoreScriptPath:  "before_restore.sh",
				PostRestoreScriptPath: "after_restore.sh",
			},
		},
	}

	config, err := ParseConfig(testConfig)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, expected) {
		t.Error("Config structure differed from expectation")
	}
}

const testConfig = `region = "us-west-2"
region = "something"
secret_key = "something_else"
bucket = "backups"

directory "config" {
	source_dir = "/etc/eventstore"
	dest_prefix = "escluster/config"
	exclude = []
	pre_backup_script = "before_backup.sh"
	post_backup_script = "after_backup.sh"
	pre_restore_script = "before_restore.sh"
	post_restore_script = "after_restore.sh"
}

directory "data" {
	source_dir = "/var/lib/eventstore"
	dest_prefix = "escluster/a/data"
	exclude = [
	"*.merging"
	]
	pre_restore_script = "before_restore.sh"
	post_restore_script = "after_restore.sh"
}`
