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
	COPY_FILE     string = "copyFile"
	CREATE_FILE   string = "createFile"
	DELETE_FILE   string = "deleteFile"
	COPY_FOLDER   string = "copyFolder"
	CREATE_FOLDER string = "createFolder"
	DELETE_FOLDER string = "deleteFolder"
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

	fafgc := appConfig.GatewayGrpcConfig
	key := fafgc.Code
	if GATEWAY_GRPC != key {
		errMsg := GATEWAY_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateCopyFile(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateCreateFile(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDeleteFile(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateCopyFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateCreateFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDeleteFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateCopyFile(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.CopyFile
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if COPY_FILE != key {
		errMsg := COPY_FILE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateCreateFile(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.CreateFile
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if CREATE_FILE != key {
		errMsg := CREATE_FILE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDeleteFile(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.DeleteFile
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DELETE_FILE != key {
		errMsg := DELETE_FILE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateCopyFolder(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.CopyFolder
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if COPY_FOLDER != key {
		errMsg := COPY_FOLDER + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateCreateFolder(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.CreateFolder
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if CREATE_FOLDER != key {
		errMsg := CREATE_FOLDER + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDeleteFolder(useCaseConfig UseCaseConfig) error {
	cf := useCaseConfig.DeleteFolder
	key := cf.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DELETE_FOLDER != key {
		errMsg := DELETE_FOLDER + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}
