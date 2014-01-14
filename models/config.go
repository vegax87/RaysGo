package models

import (
	"RaysGo/helpers"
)

type ConfigForm struct {
	AppName        string `valid:"MaxSize(30)"`
	AppDescription string
	AppKeywords    string
	MailUser       string `valid:"Email"`
	MailFrom       string
}

func (this *ConfigForm) Save() bool {
	configs := make([]Variable, 5)
	configs[0].Name = "AppName"
	configs[0].Value = this.AppName

	configs[1].Name = "AppDescription"
	configs[1].Value = this.AppDescription

	configs[2].Name = "AppKeywords"
	configs[2].Value = this.AppKeywords

	configs[3].Name = "MailUser"
	configs[3].Value = this.MailUser

	configs[4].Name = "MailFrom"
	configs[4].Value = this.MailFrom

	Engine.In("Name", "AppName", "AppDescription", "AppKeywords", "MailUser", "MailFrom").Delete(&Variable{})
	if _, err := Engine.Insert(&configs); err == nil {
		return true
	}
	return false
}

func LoadConfig() *ConfigForm {
	var configs []Variable
	result := &ConfigForm{}
	if err := Engine.In("Name", "AppName", "AppDescription", "AppKeywords", "MailUser", "MailFrom").Find(&configs); err == nil {
		for _, config := range configs {
			switch config.Name {
			case "AppName":
				result.AppName = config.Value
			case "AppDescription":
				result.AppDescription = config.Value
			case "AppKeywords":
				result.AppKeywords = config.Value
			case "MailUser":
				result.MailUser = config.Value
			case "MailFrom":
				result.MailFrom = config.Value
			}
		}
	}

	return result
}

func LoadAndSetConfig() {
	config := LoadConfig()
	if config.AppName != "" {
		helpers.AppName = config.AppName
	}
	if config.AppDescription != "" {
		helpers.AppDescription = config.AppDescription
	}
	if config.AppName != "" {
		helpers.AppKeywords = config.AppKeywords
	}
	if config.MailUser != "" {
		helpers.MailUser = config.MailUser
	}
	if config.AppName != "" {
		helpers.MailFrom = config.MailFrom
	}
}
