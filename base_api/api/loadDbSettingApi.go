package api

import (
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Dbconfig struct {
	Developments Development `yaml:"development"`
}

type Development struct {
	Dialect    string `yaml:"dialect"`
	Datasource string `yaml:"datasource"`
	Dir        string `yaml:"dir"`
	Table      string `yaml:"table"`
}

type DbSettingModel struct {
	Dbname   string
	Host     string
	Password string
	Port     string
	Sslmode  string
	User     string
}

func LoadDbSetting() *DbSettingModel {
	buf, err := ioutil.ReadFile("setting/dbconfig.yml")
	if err != nil {
		return nil
	}

	var dbConfig Dbconfig
	err = yaml.Unmarshal(buf, &dbConfig)
	if err != nil {
		return nil
	}

	dataSource := dbConfig.Developments.Datasource
	dataName := strings.Fields(dataSource)
	settingMap := make(map[string]string)
	for _, d := range dataName {
		settingValue := strings.Split(d, "=")
		settingMap[settingValue[0]] = settingValue[1]
	}

	return &DbSettingModel{
		Dbname:   settingMap["dbname"],
		Host:     settingMap["host"],
		Password: settingMap["password"],
		Port:     settingMap["port"],
		Sslmode:  settingMap["sslmode"],
		User:     settingMap["user"],
	}
}
