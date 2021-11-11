package config

import (
	"github.com/pkg/errors"
)

// attribute code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	LOG_JSON_DATA                  string = "logJsonData"
	QUEUE_JSON_DATA                string = "queueJsonData"
	AUTOM_CENTER                   string = "automCenter"
	EXCEL_GRPC                     string = "excelGrpc"
	AD_GRPC                        string = "adGrpc"
	DATABASE_GRPC                  string = "databaseGrpc"
	SERVICEDESKPLUS_GRPC           string = "serviceDeskPlusGrpc"
	FILEANDFOLDER_GRPC             string = "fileAndFolderGrpc"
	PDF_GRPC                       string = "pdfGrpc"
	OCR_GRPC                       string = "ocrGrpc"
	EWS_GRPC                       string = "ewsGrpc"
	MATEDRIVE_GRPC                 string = "matedriveGrpc"
	DESKTOPAGENT_GRPC              string = "desktopAgentGrpc"
	EXCEL_MICROSERVICE             string = "excelMicroservice"
	ACTIVEDIRECTORY_MICROSERVICE   string = "activeDirectoryMicroservice"
	RESTFULWEBSERVICE_MICROSERVICE string = "restfulWebServiceMicroservice"
	DATABASE_MICROSERVICE          string = "databaseMicroservice"
	SERVICEDESKPLUS_MICROSERVICE   string = "serviceDeskPlusMicroservice"
	RECORDINGAGENT_MICROSERVICE    string = "recordingAgentMicroservice"
	FILEANDFOLDER_MICROSERVICE     string = "fileAndFolderMicroservice"
	PDF_MICROSERVICE               string = "pdfMicroservice"
	OCR_MICROSERVICE               string = "ocrMicroservice"
	EWS_MICROSERVICE               string = "ewsMicroservice"
	MATEDRIVE_MICROSERVICE         string = "matedriveMicroservice"
	DESKTOPAGENT_MICROSERVICE      string = "desktopAgentMicroservice"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	LOG_JSON          string = "logJson"
	QUEUE             string = "queue"
	EXCEL             string = "excel"
	ACTIVEDIRECTORY   string = "activeDirectory"
	RESTFULWEBSERVICE string = "restfulWebService"
	DATABASE          string = "database"
	SERVICEDESKPLUS   string = "serviceDeskPlus"
	LOOP              string = "loop"
	IFELSE            string = "ifElse"
	RECORDINGAGENT    string = "recordingAgent"
	FILEANDFOLDER     string = "fileAndFolder"
	PDF               string = "pdf"
	OCR               string = "ocr"
	EWS               string = "ews"
	MATEDRIVE         string = "matedrive"
	DESKTOPAGENT      string = "desktopAgent"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	LOG_DATA string = "logData"
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
	err = validateExcelMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateActiveDirectoryMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRestfulWebServiceMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRecordingAgentMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDatabaseMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateServiceDeskPlusMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateFileAndFolderMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validatePdfMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateOcrMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateEwsMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateMatedriveMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDesktopAgentMicroservice(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateExcelMicroservice(appConfig AppConfig) error {
	em := appConfig.ExcelMicroserviceConfig
	key := em.Code
	emMsg := " in validateExcelMicroservice doesn't match key = "
	if EXCEL_MICROSERVICE != key {
		errMsg := EXCEL_MICROSERVICE + emMsg + key
		return errors.New(errMsg)
	}
	key = em.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + emMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateActiveDirectoryMicroservice(appConfig AppConfig) error {
	am := appConfig.ActiveDirectoryMicroserviceConfig
	key := am.Code
	amMsg := " in validateActiveDirectoryMicroservice doesn't match key = "
	if ACTIVEDIRECTORY_MICROSERVICE != key {
		errMsg := ACTIVEDIRECTORY_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateRestfulWebServiceMicroservice(appConfig AppConfig) error {
	am := appConfig.RestfulWebServiceMicroserviceConfig
	key := am.Code
	amMsg := " in validateRestfulWebServiceMicroservice doesn't match key = "
	if RESTFULWEBSERVICE_MICROSERVICE != key {
		errMsg := RESTFULWEBSERVICE_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateRecordingAgentMicroservice(appConfig AppConfig) error {
	am := appConfig.RecordingAgentMicroserviceConfig
	key := am.Code
	amMsg := " in validateRecordingAgentMicroservice doesn't match key = "
	if RECORDINGAGENT_MICROSERVICE != key {
		errMsg := RECORDINGAGENT_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDatabaseMicroservice(appConfig AppConfig) error {
	am := appConfig.DatabaseMicroserviceConfig
	key := am.Code
	amMsg := " in validateDatabaseMicroservice doesn't match key = "
	if DATABASE_MICROSERVICE != key {
		errMsg := DATABASE_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateServiceDeskPlusMicroservice(appConfig AppConfig) error {
	am := appConfig.ServiceDeskPlusMicroserviceConfig
	key := am.Code
	amMsg := " in validateSDPMicroservice doesn't match key = "
	if SERVICEDESKPLUS_MICROSERVICE != key {
		errMsg := SERVICEDESKPLUS_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateFileAndFolderMicroservice(appConfig AppConfig) error {
	am := appConfig.FileAndFolderMicroserviceConfig
	key := am.Code
	amMsg := " in validateFileAndFolderMicroservice doesn't match key = "
	if FILEANDFOLDER_MICROSERVICE != key {
		errMsg := FILEANDFOLDER_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validatePdfMicroservice(appConfig AppConfig) error {
	am := appConfig.PdfMicroserviceConfig
	key := am.Code
	amMsg := " in validatePDFMiroservice doesn't match key = "
	if PDF_MICROSERVICE != key {
		errMsg := PDF_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateOcrMicroservice(appConfig AppConfig) error {
	am := appConfig.OcrMicroserviceConfig
	key := am.Code
	amMsg := " in validateOCRMicroservice doesn't match key = "
	if OCR_MICROSERVICE != key {
		errMsg := OCR_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateEwsMicroservice(appConfig AppConfig) error {
	am := appConfig.EwsMicroserviceConfig
	key := am.Code
	amMsg := " in validateExchangeWebServerMicroservice doesn't match key = "
	if EWS_MICROSERVICE != key {
		errMsg := EWS_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateMatedriveMicroservice(appConfig AppConfig) error {
	am := appConfig.MatedriveMicroserviceConfig
	key := am.Code
	amMsg := " in validateMateDriveMicroservice doesn't match key = "
	if MATEDRIVE_MICROSERVICE != key {
		errMsg := MATEDRIVE_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDesktopAgentMicroservice(appConfig AppConfig) error {
	am := appConfig.DesktopAgentMicroserviceConfig
	key := am.Code
	amMsg := " in validateDesktopAgentMicroservice doesn't match key = "
	if DESKTOPAGENT_MICROSERVICE != key {
		errMsg := DESKTOPAGENT_MICROSERVICE + amMsg + key
		return errors.New(errMsg)
	}
	key = am.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + amMsg + key
		return errors.New(errMsg)
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
	cc := appConfig.LogJsonDataConfig
	key := cc.Code
	if LOG_JSON_DATA != key {
		errMsg := LOG_JSON_DATA + scMsg + key
		return errors.New(errMsg)
	}

	qgc := appConfig.QueueJsonDataConfig
	key = qgc.Code
	if QUEUE_JSON_DATA != key {
		errMsg := QUEUE_JSON_DATA + scMsg + key
		return errors.New(errMsg)
	}

	acc := appConfig.AutomCenterConfig
	key = acc.Code
	if AUTOM_CENTER != key {
		errMsg := AUTOM_CENTER + scMsg + key
		return errors.New(errMsg)
	}

	egc := appConfig.ExcelGrpcConfig
	key = egc.Code
	if EXCEL_GRPC != key {
		errMsg := EXCEL_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	adgc := appConfig.ADGrpcConfig
	key = adgc.Code
	if AD_GRPC != key {
		errMsg := AD_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	dbgc := appConfig.DatabaseGrpcConfig
	key = dbgc.Code
	if DATABASE_GRPC != key {
		errMsg := DATABASE_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	sdpgc := appConfig.ServiceDeskPlusGrpcConfig
	key = sdpgc.Code
	if SERVICEDESKPLUS_GRPC != key {
		errMsg := SERVICEDESKPLUS_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	fafgc := appConfig.FileAndFolderGrpcConfig
	key = fafgc.Code
	if FILEANDFOLDER_GRPC != key {
		errMsg := FILEANDFOLDER_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	pdfgc := appConfig.PdfGrpcConfig
	key = pdfgc.Code
	if PDF_GRPC != key {
		errMsg := PDF_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	ocrgc := appConfig.OcrGrpcConfig
	key = ocrgc.Code
	if OCR_GRPC != key {
		errMsg := OCR_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	ewsgc := appConfig.EwsGrpcConfig
	key = ewsgc.Code
	if EWS_GRPC != key {
		errMsg := EWS_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	matedrivegc := appConfig.MatedriveGrpcConfig
	key = matedrivegc.Code
	if MATEDRIVE_GRPC != key {
		errMsg := MATEDRIVE_GRPC + scMsg + key
		return errors.New(errMsg)
	}

	dagc := appConfig.DesktopAgentGrpcConfig
	key = dagc.Code
	if DESKTOPAGENT_GRPC != key {
		errMsg := DESKTOPAGENT_GRPC + scMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateUseCase(useCase UseCaseConfig) error {
	err := validateLogJson(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateQueue(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateExcel(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateActiveDirectory(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRestfulWebService(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateRecordingAgent(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDatabase(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateServiceDeskplus(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateFileAndFolder(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validatePdf(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateOcr(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateEws(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLoop(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateIfElse(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateMatedrive(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateDesktopAgent(useCase)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateExcel(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Excel
	key := rc.Code
	rcMsg := " in validateExcel doesn't match key = "
	if EXCEL != key {
		errMsg := EXCEL + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.MicroserviceConfig.Code
	if EXCEL_MICROSERVICE != key {
		errMsg := EXCEL_MICROSERVICE + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateQueue(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Queue
	key := rc.Code
	rcMsg := " in validateLogJson doesn't match key = "
	if QUEUE != key {
		errMsg := QUEUE + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateLogJson(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.LogJson
	key := rc.Code
	rcMsg := " in validateLogJson doesn't match key = "
	if LOG_JSON != key {
		errMsg := LOG_JSON + rcMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateActiveDirectory(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.ActiveDirectory
	key := ac.Code
	acMsg := " in validateActiveDirectory doesn't match key = "
	if ACTIVEDIRECTORY != key {
		errMsg := ACTIVEDIRECTORY + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if ACTIVEDIRECTORY_MICROSERVICE != key {
		errMsg := ACTIVEDIRECTORY_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateRestfulWebService(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.RestfulWebService
	key := ac.Code
	acMsg := " in validateRestfulWebService doesn't match key = "
	if RESTFULWEBSERVICE != key {
		errMsg := RESTFULWEBSERVICE + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if RESTFULWEBSERVICE_MICROSERVICE != key {
		errMsg := RESTFULWEBSERVICE_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateDatabase(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Database
	key := ac.Code
	acMsg := " in validateDatabase doesn't match key = "
	if DATABASE != key {
		errMsg := DATABASE + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if DATABASE_MICROSERVICE != key {
		errMsg := DATABASE_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateRecordingAgent(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.RecordingAgent
	key := ac.Code
	acMsg := " in validateRecordingAgent doesn't match key = "
	if RECORDINGAGENT != key {
		errMsg := RECORDINGAGENT + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if RECORDINGAGENT_MICROSERVICE != key {
		errMsg := RECORDINGAGENT_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateServiceDeskplus(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.ServiceDeskPlus
	key := ac.Code
	acMsg := " in validateSDP doesn't match key = "
	if SERVICEDESKPLUS != key {
		errMsg := SERVICEDESKPLUS + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if SERVICEDESKPLUS_MICROSERVICE != key {
		errMsg := SERVICEDESKPLUS_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateFileAndFolder(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.FileAndFolder
	key := ac.Code
	acMsg := " in validateFileAndFolder doesn't match key = "
	if FILEANDFOLDER != key {
		errMsg := FILEANDFOLDER + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if FILEANDFOLDER_MICROSERVICE != key {
		errMsg := FILEANDFOLDER_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validatePdf(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Pdf
	key := ac.Code
	acMsg := " in validatePDF doesn't match key = "
	if PDF != key {
		errMsg := PDF + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if PDF_MICROSERVICE != key {
		errMsg := PDF_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateOcr(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Ocr
	key := ac.Code
	acMsg := " in validateOCR doesn't match key = "
	if OCR != key {
		errMsg := OCR + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if OCR_MICROSERVICE != key {
		errMsg := OCR_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateEws(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Ews
	key := ac.Code
	acMsg := " in validateExchangeWebServer doesn't match key = "
	if EWS != key {
		errMsg := EWS + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if EWS_MICROSERVICE != key {
		errMsg := EWS_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateMatedrive(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Matedrive
	key := ac.Code
	acMsg := " in validateMateDrive doesn't match key = "
	if MATEDRIVE != key {
		errMsg := MATEDRIVE + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if MATEDRIVE_MICROSERVICE != key {
		errMsg := MATEDRIVE_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDesktopAgent(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.DesktopAgent
	key := ac.Code
	acMsg := " in validateDesktopAgent doesn't match key = "
	if DESKTOPAGENT != key {
		errMsg := DESKTOPAGENT + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.MicroserviceConfig.Code
	if DESKTOPAGENT_MICROSERVICE != key {
		errMsg := DESKTOPAGENT_MICROSERVICE + acMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateLoop(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.Loop
	key := ac.Code
	acMsg := " in validateLoop doesn't match key = "
	if LOOP != key {
		errMsg := LOOP + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}

func validateIfElse(useCaseConfig UseCaseConfig) error {
	ac := useCaseConfig.IfElse
	key := ac.Code
	acMsg := " in validateIfElse doesn't match key = "
	if IFELSE != key {
		errMsg := IFELSE + acMsg + key
		return errors.New(errMsg)
	}
	key = ac.LogDataConfig.Code
	if LOG_DATA != key {
		errMsg := LOG_DATA + acMsg + key
		return errors.New(errMsg)
	}

	return nil
}
