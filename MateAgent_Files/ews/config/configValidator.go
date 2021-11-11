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
	ADD_ACCOUNT                 string = "addAccount"
	REMOVE_USER_FROM_MAIL_GROUP string = "removeUserFromMailGroup"
	ADD_USER_TO_MAİL_GROUP      string = "addUserToMailGroup"
	DISABLE                     string = "disable"
	DELETE                      string = "delete"
	CHANGE_MAILBOX_SIZE         string = "changeMailboxSize"
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
	err := validateAddAccount(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRemoveUserFromMailGroup(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateAddUserToMailGroup(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDisable(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDelete(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateChangeMailboxSize(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateAddAccount(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.AddAccount
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if ADD_ACCOUNT != key {
		errMsg := ADD_ACCOUNT + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateRemoveUserFromMailGroup(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.RemoveUserFromMailGroup
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if REMOVE_USER_FROM_MAIL_GROUP != key {
		errMsg := REMOVE_USER_FROM_MAIL_GROUP + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateAddUserToMailGroup(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.AddUserToMailGroup
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if ADD_USER_TO_MAİL_GROUP != key {
		errMsg := ADD_USER_TO_MAİL_GROUP + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDisable(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Disable
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DISABLE != key {
		errMsg := DISABLE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDelete(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Delete
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if DELETE != key {
		errMsg := DELETE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateChangeMailboxSize(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.ChangeMailboxSize
	key := rc.Code
	rcMsg := " in validateRegistration doesn't match key = "
	if CHANGE_MAILBOX_SIZE != key {
		errMsg := CHANGE_MAILBOX_SIZE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}
