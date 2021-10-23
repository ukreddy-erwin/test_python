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
	LogJsonDataConfig                   LogJsonStoreConfig                  `yaml:"logJsonDataConfig"`
	QueueJsonDataConfig                 QueueJsonStoreConfig                `yaml:"queueJsonDataConfig"`
	GatewayGrpcConfig                   DataStoreConfig                     `yaml:"gatewayGrpcConfig"`
	ExcelGrpcConfig                     DataStoreConfig                     `yaml:"excelGrpcConfig"`
	ADGrpcConfig                        DataStoreConfig                     `yaml:"adGrpcConfig"`
	RestfulWebServiceGrpcConfig         DataStoreConfig                     `yaml:"restfulWebServiceGrpcConfig"`
	DatabaseGrpcConfig                  DataStoreConfig                     `yaml:"databaseGrpcConfig"`
	RecordingAgentGrpcConfig            DataStoreConfig                     `yaml:"recordingAgentGrpcConfig"`
	ServiceDeskPlusGrpcConfig           DataStoreConfig                     `yaml:"serviceDeskPlusGrpcConfig"`
	FileAndFolderGrpcConfig             DataStoreConfig                     `yaml:"fileAndFolderGrpcConfig"`
	PdfGrpcConfig                       DataStoreConfig                     `yaml:"pdfGrpcConfig"`
	OcrGrpcConfig                       DataStoreConfig                     `yaml:"ocrGrpcConfig"`
	EwsGrpcConfig                       DataStoreConfig                     `yaml:"ewsGrpcConfig"`
	ZapConfig                           LogConfig                           `yaml:"zapConfig"`
	LogrusConfig                        LogConfig                           `yaml:"logrusConfig"`
	Log                                 LogConfig                           `yaml:"logConfig"`
	UseCase                             UseCaseConfig                       `yaml:"useCaseConfig"`
	ExcelMicroserviceConfig             ExcelMicroserviceConfig             `yaml:"excelMicroserviceConfig"`
	ActiveDirectoryMicroserviceConfig   ActiveDirectoryMicroserviceConfig   `yaml:"activeDirectoryMicroserviceConfig"`
	RestfulWebServiceMicroserviceConfig RestfulWebServiceMicroserviceConfig `yaml:"restfulWebServiceMicroserviceConfig"`
	RecordingAgentMicroserviceConfig    RecordingAgentMicroserviceConfig    `yaml:"recordingAgentMicroserviceConfig"`
	DatabaseMicroserviceConfig          DatabaseMicroserviceConfig          `yaml:"databaseMicroserviceConfig"`
	ServiceDeskPlusMicroserviceConfig   ServiceDeskPlusMicroserviceConfig   `yaml:"serviceDeskPlusMicroserviceConfig"`
	FileAndFolderMicroserviceConfig     FileAndFolderMicroserviceConfig     `yaml:"fileAndFolderMicroserviceConfig"`
	PdfMicroserviceConfig               PdfMicroserviceConfig               `yaml:"pdfMicroserviceConfig"`
	OcrMicroserviceConfig               OcrMicroserviceConfig               `yaml:"ocrMicroserviceConfig"`
	EwsMicroserviceConfig               EwsMicroserviceConfig               `yaml:"ewsMicroserviceConfig"`
}

// ExcelMicroserviceConfig represents excel use case
type ExcelMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// ActiveDirectoryMicroserviceConfig represents excel use case
type ActiveDirectoryMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// RestfulWebServiceMicroserviceConfig represents excel use case
type RestfulWebServiceMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// DatabaseMicroserviceConfig represents database use case
type DatabaseMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

type RecordingAgentMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// ServiceDeskPlusMicroserviceConfig represents database use case
type ServiceDeskPlusMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}
type FileAndFolderMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}
type PdfMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}
type OcrMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}
type EwsMicroserviceConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	LogJson           LogJsonConfig           `yaml:"logJson"`
	Queue             QueueConfig             `yaml:"queue"`
	Excel             ExcelConfig             `yaml:"excel"`
	ActiveDirectory   ActiveDirectoryConfig   `yaml:"activeDirectory"`
	RestfulWebService RestfulWebServiceConfig `yaml:"restfulWebService"`
	RecordingAgent    RecordingAgentConfig    `yaml:"recordingAgent"`
	Loop              LoopConfig              `yaml:"loop"`
	IfElse            IfElseConfig            `yaml:"ifelse"`
	Database          DatabaseConfig          `yaml:"database"`
	ServiceDeskPlus   ServiceDeskPlusConfig   `yaml:"serviceDeskPlus"`
	FileAndFolder     FileAndFolderConfig     `yaml:"fileAndFolder"`
	Pdf               PdfConfig               `yaml:"pdf"`
	Ocr               OcrConfig               `yaml:"ocr"`
	Ews               EwsConfig               `yaml:"ews"`
}

// LogJsonConfig represents excel use case
type LogJsonConfig struct {
	Code string `yaml:"code"`
}

// QueueConfig represents excel use case
type QueueConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// ExcelConfig represents excel use case
type ExcelConfig struct {
	Code               string                  `yaml:"code"`
	MicroserviceConfig ExcelMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig              `yaml:"logDataConfig"`
}

// ActiveDirectoryConfig represents active directory use case
type ActiveDirectoryConfig struct {
	Code               string                            `yaml:"code"`
	MicroserviceConfig ActiveDirectoryMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                        `yaml:"logDataConfig"`
}

// RestfulWebServiceConfig represents active directory use case
type RestfulWebServiceConfig struct {
	Code               string                              `yaml:"code"`
	MicroserviceConfig RestfulWebServiceMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                          `yaml:"logDataConfig"`
}

// DatabaseConfig represents database use case
type DatabaseConfig struct {
	Code               string                     `yaml:"code"`
	MicroserviceConfig DatabaseMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                 `yaml:"logDataConfig"`
}

type RecordingAgentConfig struct {
	Code               string                           `yaml:"code"`
	MicroserviceConfig RecordingAgentMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                       `yaml:"logDataConfig"`
}
type ServiceDeskPlusConfig struct {
	Code               string                            `yaml:"code"`
	MicroserviceConfig ServiceDeskPlusMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                        `yaml:"logDataConfig"`
}
type FileAndFolderConfig struct {
	Code               string                          `yaml:"code"`
	MicroserviceConfig FileAndFolderMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig                      `yaml:"logDataConfig"`
}
type PdfConfig struct {
	Code               string                `yaml:"code"`
	MicroserviceConfig PdfMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig            `yaml:"logDataConfig"`
}
type OcrConfig struct {
	Code               string                `yaml:"code"`
	MicroserviceConfig OcrMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig            `yaml:"logDataConfig"`
}
type EwsConfig struct {
	Code               string                `yaml:"code"`
	MicroserviceConfig EwsMicroserviceConfig `yaml:"microserviceConfig"`
	LogDataConfig      DataConfig            `yaml:"logDataConfig"`
}

// LoopConfig represents excel use case
type LoopConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// LoopConfig represents excel use case
type IfElseConfig struct {
	Code          string     `yaml:"code"`
	LogDataConfig DataConfig `yaml:"logDataConfig"`
}

// DataConfig represents data service
type DataConfig struct {
	Code            string          `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// LogStoreConfig represents handlers for data store. It can be a database or a gRPC connection
type LogJsonStoreConfig struct {
	Code string `yaml:"code"`
	Path string `yaml:"path"`
	Ext  string `yaml:"ext"`
}

// LogStoreConfig represents handlers for data store. It can be a database or a gRPC connection
type QueueJsonStoreConfig struct {
	Code string `yaml:"code"`
	Path string `yaml:"path"`
	Ext  string `yaml:"ext"`
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
