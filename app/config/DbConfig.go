package config

import (
	"fmt"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Driver   string
	DbName   string
}

func GetSourceCfg() DbConfig {
	return DbConfig{
		Host:     AppCfg.SourceDbHost,
		Port:     AppCfg.SourceDbPort,
		User:     AppCfg.SourceDbUser,
		Password: AppCfg.SourceDbPassword,
		Driver:   AppCfg.SourceDbDriver,
		DbName:   AppCfg.SourceDbDbname,
	}
}

func (dbConf *DbConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.DbName)
}

func (dbConf *DbConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.DbName)
}
