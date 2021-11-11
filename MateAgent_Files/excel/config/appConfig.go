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
	GatewayGrpcConfig DataStoreConfig `yaml:"gatewayGrpcConfig"`
	ExcelGrpcConfig   DataStoreConfig `yaml:"excelGrpcConfig"`
	ZapConfig         LogConfig       `yaml:"zapConfig"`
	LogrusConfig      LogConfig       `yaml:"logrusConfig"`
	Log               LogConfig       `yaml:"logConfig"`
	UseCase           UseCaseConfig   `yaml:"useCaseConfig"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	Excel       ExcelConfig       `yaml:"excel"`
	Copy        CopyConfig        `yaml:"copy"`
	Paste       PasteConfig       `yaml:"paste"`
	CreatePivot CreatePivotConfig `yaml:"createPivot"`
	ExcelToCsv  ExcelToCsvConfig  `yaml:"excelToCsv"`
	VLookoup    VLookupConfig     `yaml:"vlookup"`
	ReadSheet   ReadSheetConfig   `yaml:"readSheet"`
}

// ExcelConfig represents excel use case
type ExcelConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

type VLookupConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

type ReadSheetConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// CopyConfig represents excel use case
type CopyConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// PasteConfig represents excel use case
type PasteConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// CreatePivotConfig represents excel use case
type CreatePivotConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// CreatePivotConfig represents excel use case
type ExcelToCsvConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
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
