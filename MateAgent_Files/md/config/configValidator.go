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
	FILE_LIST          string = "fileList"
	CREATE_DIRECTORY   string = "createDirectory"
	DELETE_FILE_FOLDER string = "deleteFileFolder"
	RENAME_FILE_FOLDER string = "renameFileFolder"
	UPLOAD_FILE        string = "uploadFile"
	DOWNLOAD_FILE      string = "downloadFile"
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
	err := validateFileList(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateCreateDirectory(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDeleteFileFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRenameFileFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateUploadFile(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDownloadFile(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateFileList(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.FileList
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if FILE_LIST != key {
		errMsg := FILE_LIST + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}
func validateCreateDirectory(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.CreateDirectory
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if CREATE_DIRECTORY != key {
		errMsg := CREATE_DIRECTORY + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}
func validateDeleteFileFolder(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.DeleteFileFolder
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DELETE_FILE_FOLDER != key {
		errMsg := DELETE_FILE_FOLDER + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}
func validateRenameFileFolder(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.RenameFileFolder
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if RENAME_FILE_FOLDER != key {
		errMsg := RENAME_FILE_FOLDER + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateUploadFile(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.UploadFile
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if UPLOAD_FILE != key {
		errMsg := UPLOAD_FILE + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDownloadFile(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.DownloadFile
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DOWNLOAD_FILE != key {
		errMsg := DOWNLOAD_FILE + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}
