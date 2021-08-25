// Package config reasd configurations from a YAML file and load them into a AppConfig type to save the configuration
// information for the application.
// Configuration for different environment can be saved in files with different suffix, for example [Dev], [Prod]
package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// AppConfig represents the application config
type AppConfig struct {
	GatewayGrpcConfig       DataStoreConfig `yaml:"gatewayGrpcConfig"`
	FileandfolderGrpcConfig DataStoreConfig `yaml:"fileandfolderGrpcConfig"`
	ZapConfig               LogConfig       `yaml:"zapConfig"`
	LogrusConfig            LogConfig       `yaml:"logrusConfig"`
	Log                     LogConfig       `yaml:"logConfig"`
	UseCase                 UseCaseConfig   `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	CopyFile     CopyFileConfig     `yaml:"copyFile"`
	CreateFile   CreateFileConfig   `yaml:"createFile"`
	DeleteFile   DeleteFileConfig   `yaml:"deleteFile"`
	CopyFolder   CopyFolderConfig   `yaml:"copyFolder"`
	CreateFolder CreateFolderConfig `yaml:"createFolder"`
	DeleteFolder DeleteFolderConfig `yaml:"deleteFolder"`
}

// CopyFileConfig represents add request use case
type CopyFileConfig struct {
	Code string `yaml:"code"`
}

// CreateFileConfig represents add request use case
type CreateFileConfig struct {
	Code string `yaml:"code"`
}
type DeleteFileConfig struct {
	Code string `yaml:"code"`
}

// CopyFolderConfig represents add request use case
type CopyFolderConfig struct {
	Code string `yaml:"code"`
}

// CreateFolderConfig represents add request use case
type CreateFolderConfig struct {
	Code string `yaml:"code"`
}

// DeleteFolderConfig represents add request use case
type DeleteFolderConfig struct {
	Code string `yaml:"code"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataConfig represents handlers for data store. It can be a database or a gRPC connection
type DataStoreConfig struct {
	Code string `yaml:"code"`
	// Only database has a driver name, for grpc it is "tcp" ( network) for server
	DriverName string `yaml:"driverName"`
	// For database, this is datasource name; for grpc, it is target url
	UrlAddress string `yaml:"urlAddress"`
	// Only some databases need this database name
	DbName string `yaml:"dbName"`
}

// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

// ReadConfig reads the file of the filename (in the same folder) and put it into the AppConfig
func ReadConfig(filename string) (*AppConfig, error) {

	var ac AppConfig
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	fmt.Println("appConfig:", ac)
	return &ac, nil
}
