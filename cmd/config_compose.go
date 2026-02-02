package cmd

import "bnbot/config"

func ReadAWSAppLog() (*config.Config, error) {
	config, err := config.ReadAWSAppConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}
