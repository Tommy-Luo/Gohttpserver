package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var DB Database

func (DB *Database)GetConf() *Database{
	basePath, err:=os.Getwd()
	if err !=nil{
		fmt.Println("base path error")
	}
	fileName:=filepath.Join(basePath,"config","config.yaml")

	yamlFile,err:=ioutil.ReadFile(fileName)
	if err !=nil{
		fmt.Println("load conf error")
	}

	err = yaml.Unmarshal(yamlFile, DB)

	if err !=nil{
		fmt.Println(err.Error())
	}
	return DB
}

func (DB *Database)GetDatabaseType() string{
	if DB.Driver == ""{
		fmt.Println("Database struct is empty")
	}
		return DB.Driver

}