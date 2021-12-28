package config

import (
	"gorm.io/gorm"
)

type Configuration struct {
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
}


type Application struct {
	Config Configuration
	DB *gorm.DB

}

var App = new(Application)