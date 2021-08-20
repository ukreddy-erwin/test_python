package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	GATEWAY_GRPC string = "gatewayGrpc"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	CONNECT              string = "connect"
	SELECT               string = "select"
	INSERT_UPDATE_DELETE string = "insertUpdateDelete"
	DISCONNECT           string = "disconnect"
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

	adgc := appConfig.GatewayGrpcConfig
	key := adgc.Code
	if GATEWAY_GRPC != key {
		errMsg := GATEWAY_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateConnect(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateSelect(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateInsertUpdateDelete(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDisconnect(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateConnect(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Connect
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if CONNECT != key {
		errMsg := CONNECT + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateSelect(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Select
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if SELECT != key {
		errMsg := SELECT + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateInsertUpdateDelete(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.InsertUpdateDelete
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if INSERT_UPDATE_DELETE != key {
		errMsg := INSERT_UPDATE_DELETE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDisconnect(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Disconnect
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if DISCONNECT != key {
		errMsg := DISCONNECT + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}
