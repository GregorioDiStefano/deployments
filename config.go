package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mendersoftware/artifacts/config"
)

const (
	SettingHttps            = "https"
	SettingHttpsCertificate = SettingHttps + ".certificate"
	SettingHttpsKey         = SettingHttps + ".key"

	SettingListen        = "listen"
	SettingListenDefault = ":8080"

	SettingsAws               = "aws"
	SettingAwsS3Region        = SettingsAws + ".region"
	SettingAwsS3RegionDefault = "eu-west-1"
	SettingAweS3Bucket        = SettingsAws + ".bucket"
	SettingAwsS3BucketDefault = "mender-artifact-storage"

	SettingsAwsAuth      = SettingsAws + ".auth"
	SettingAwsAuthKeyId  = SettingsAwsAuth + ".key"
	SettingAwsAuthSecret = SettingsAwsAuth + ".secret"
	SettingAwsAuthToken  = SettingsAwsAuth + ".token"
)

// ValidateAwsAuth validates configuration of SettingsAwsAuth section if provided.
func ValidateAwsAuth(c config.ConfigReader) error {

	if c.IsSet(SettingsAwsAuth) {
		required := []string{SettingAwsAuthKeyId, SettingAwsAuthSecret}
		for _, key := range required {
			if !c.IsSet(key) {
				return MissingOptionError(key)
			}

			if c.GetString(key) == "" {
				return MissingOptionError(key)
			}
		}
	}

	return nil
}

// ValidateHttps validates configuration of SettingHttps section if provided.
func ValidateHttps(c config.ConfigReader) error {

	if c.IsSet(SettingHttps) {
		required := []string{SettingHttpsCertificate, SettingHttpsKey}
		for _, key := range required {
			if !c.IsSet(key) {
				return MissingOptionError(key)
			}

			value := c.GetString(key)
			if value == "" {
				return MissingOptionError(key)
			}

			if _, err := os.Stat(value); err != nil {
				return err
			}
		}
	}

	return nil
}

// Generate error with missing reuired option message.
func MissingOptionError(option string) error {
	return errors.New(fmt.Sprintf("Required option: '%s'", option))
}