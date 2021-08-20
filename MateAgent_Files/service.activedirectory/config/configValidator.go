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
	CONNECT      string = "connect"
	CREATE_USER  string = "createUser"
	UPDATE_USER  string = "updateUser"
	DISABLE_USER string = "disableUser"
	SEARCH_USER  string = "searchUser"
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
	err = validateCreateUser(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateUpdateUser(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDisableUser(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateSearchUser(useCase)
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

func validateCreateUser(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.CreateUser
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if CREATE_USER != key {
		errMsg := CREATE_USER + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUpdateUser(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.UpdateUser
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if UPDATE_USER != key {
		errMsg := UPDATE_USER + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDisableUser(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.DisableUser
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if DISABLE_USER != key {
		errMsg := DISABLE_USER + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateSearchUser(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.SearchUser
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if SEARCH_USER != key {
		errMsg := SEARCH_USER + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}
