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
	GET_METHOD    string = "getMethod"
	POST_METHOD   string = "postMethod"
	PUT_METHOD    string = "putMethod"
	PATCH_METHOD  string = "patchMethod"
	DELETE_METHOD string = "deleteMethod"
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
	err := validateGetMethod(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validatePostMethod(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validatePutMethod(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validatePatchMethod(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDeleteMethod(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateGetMethod(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.GetMethod
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if GET_METHOD != key {
		errMsg := GET_METHOD + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validatePostMethod(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.PostMethod
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if POST_METHOD != key {
		errMsg := PATCH_METHOD + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validatePutMethod(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.PutMethod
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if PUT_METHOD != key {
		errMsg := PUT_METHOD + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validatePatchMethod(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.PatchMethod
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if PATCH_METHOD != key {
		errMsg := PATCH_METHOD + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDeleteMethod(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.DeleteMethod
	key := ac.Code
	acMsg := " in validateListUser doesn't match key = "
	if DELETE_METHOD != key {
		errMsg := DELETE_METHOD + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}
