package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

type VaultConfig struct {
	VaultAddr     string `hcl:"VAULT_ADDR"`
	SSHMountPoint string `hcl:"SSH_MOUNT_POINT"`
	CACert        string `hcl:"CA_CERT"`
	CAPath        string `hcl:"CA_PATH"`
	TLSSkipVerify bool   `hcl:"TLS_SKIP_VERIFY"`
}

func LoadConfig(path string) (*VaultConfig, error) {
	var config VaultConfig
	contents, err := ioutil.ReadFile(path)
	if !os.IsNotExist(err) {
		obj, err := hcl.Parse(string(contents))
		if err != nil {
			return nil, fmt.Errorf("Error parsing Vault address")
		}

		if err := hcl.DecodeObject(&config, obj); err != nil {
			return nil, fmt.Errorf("Error decoding Vault address")
		}
	} else {
		return nil, fmt.Errorf("Error finding vault agent config file")
	}
	return &config, nil
}
