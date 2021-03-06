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
	GatewayGrpcConfig   DataStoreConfig `yaml:"gatewayGrpcConfig"`
	MateDriveGrpcConfig DataStoreConfig `yaml:"matedriveGrpcConfig"`
	ZapConfig           LogConfig       `yaml:"zapConfig"`
	LogrusConfig        LogConfig       `yaml:"logrusConfig"`
	Log                 LogConfig       `yaml:"logConfig"`
	UseCase             UseCaseConfig   `yaml:"useCaseConfig"`
}

// DownloadFileConfig represents matedrive use case
type DownloadFileConfig struct {
	Code string `yaml:"code"`
}

// UploadFileConfig  represents matedrive use case
type UploadFileConfig struct {
	Code string `yaml:"code"`
}

// FileListConfig represents matedrive use case
type FileListConfig struct {
	Code string `yaml:"code"`
}

// DeleteFileFolderConfig represents matedrive use case
type DeleteFileFolderConfig struct {
	Code string `yaml:"code"`
}

// CreateDirConfig represents matedrive use case
type CreateDirectoryConfig struct {
	Code string `yaml:"code"`
}

// RenameFileFolder represents matedrive use case
type RenameFileFolder struct {
	Code string `yaml:"code"`
}

// SearchConfig represents matedrive use case
type SearchConfig struct {
	Code string `yaml:"code"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	DownloadFile     DownloadFileConfig     `yaml:"downloadFile"`
	UploadFile       UploadFileConfig       `yaml:"uploadFile"`
	FileList         FileListConfig         `yaml:"fileList"`
	DeleteFileFolder DeleteFileFolderConfig `yaml:"deleteFileFolder"`
	CreateDirectory  CreateDirectoryConfig  `yaml:"createDirectory"`
	RenameFileFolder RenameFileFolder       `yaml:"renameFileFolder"`
	Search           SearchConfig           `yaml:"search"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataStoreConfig represents handlers for data store. It can be a database or a gRPC connection
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
