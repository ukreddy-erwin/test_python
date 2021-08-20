package excelclient

import (
	espb "gitlab.com/ithero-arge/service.excel/adapter/server/generatedclient"
	"gitlab.com/ithero-arge/service.excel/model"
)

// GrpcToExcelOpen converts from grpc ExcelOpen type to domain Model ExcelOpen type
func GrpcToExcelOpen(excel *espb.ExcelOpen) (*model.ExcelOpen, error) {
	if excel == nil {
		return &model.ExcelOpen{}, nil
	}

	resultExcelOpen := &model.ExcelOpen{}
	resultExcelOpen.Name = excel.Name
	resultExcelOpen.Id = excel.Id
	resultExcelOpen.File = excel.File
	resultExcelOpen.FileType = excel.FileType
	resultExcelOpen.IsHeader = excel.IsHeader
	resultExcelOpen.OpenAs = excel.OpenAs

	return resultExcelOpen, nil
}

// ExcelOpenToGrpc converts from domain Model ExcelOpen type to grpc ExcelOpen type
func ExcelOpenToGrpc(excel *model.ExcelOpen) (*espb.ExcelOpen, error) {
	if excel == nil {
		return nil, nil
	}

	resultExcelOpen := &espb.ExcelOpen{}
	resultExcelOpen.OpenAs = excel.OpenAs
	resultExcelOpen.IsHeader = excel.IsHeader
	resultExcelOpen.FileType = excel.FileType
	resultExcelOpen.Id = excel.Id
	resultExcelOpen.Name = excel.Name

	return resultExcelOpen, nil
}

// GrpcToReadData converts from grpc ReadData type to domain Model ReadData type
func GrpcToReadData(excel *espb.ReadData) (*model.ReadData, error) {
	if excel == nil {
		return &model.ReadData{}, nil
	}

	resultReadData := &model.ReadData{}
	resultReadData.FromRow = excel.From
	resultReadData.ToRow = excel.To
	resultReadData.RowData = excel.RowData

	return resultReadData, nil
}

// ReadDataToGrpc converts from grpc ReadData type to domain Model ReadData type
func ReadDataToGrpc(excel *model.ReadData) (*espb.ReadData, error) {
	if excel == nil {
		return &espb.ReadData{}, nil
	}

	resultReadData := &espb.ReadData{}
	resultReadData.From = excel.FromRow
	resultReadData.To = excel.ToRow
	resultReadData.RowData = excel.RowData

	return resultReadData, nil
}