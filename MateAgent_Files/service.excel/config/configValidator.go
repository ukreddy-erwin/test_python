package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	GATEWAY_GRPC  		string = "gatewayGrpc"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	EXCEL	string = "excel"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	LOG_DATA	string = "logData"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLogger(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	useCase := appConfig.UseCase
	err = validateUseCase(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateLogger(appConfig AppConfig) error {
	zc := appConfig.ZapConfig
	key := zc.Code
	zcMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + zcMsg + key
		return errors.New(errMsg)
	}
	lc := appConfig.LogrusConfig
	key = lc.Code
	if LOGRUS != lc.Code {
		errMsg := LOGRUS + zcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	scMsg := " in validateDataStore doesn't match key = "
	ggc := appConfig.GatewayGrpcConfig
	key := ggc.Code
	if GATEWAY_GRPC != key {
		errMsg := GATEWAY_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateADAction(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}


func validateADAction(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Excel
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if EXCEL != key {
		errMsg := EXCEL + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}
