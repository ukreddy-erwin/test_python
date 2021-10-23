package flowclient

import (
	fspb "gitlab.com/ithero-arge/grpc.gateway/adapter/server/flowclient/generatedclient"
	"gitlab.com/ithero-arge/grpc.gateway/model"
	"gitlab.com/ithero-arge/grpc.gateway/pkg/utils"
)

// GrpcToFlow converts from grpc Flow type to domain Model flow type
func GrpcToFlow(flow []*fspb.Action) ([]*model.Flow, error) {
	if flow == nil {
		return []*model.Flow{}, nil
	}
	var resultFlow []*model.Flow
	for _, item := range flow {
		resultFlowCur := &model.Flow{}
		resultFlowCur.Id = item.Id
		resultFlowCur.Name = item.Name
		resultFlowCur.Params = utils.DictionaryToMap(item.Params)
		resultFlowCur.Library.V = item.Library.V
		resultFlowCur.Library.Name = item.Library.Name
		resultFlowCur.Definition = item.Definition
		if item.GetChildren() != nil {
			resultFlowCur.Children, _ = GrpcToFlow(item.GetChildren())
		}

		resultFlow = append(resultFlow, resultFlowCur)
	}
	return resultFlow, nil
}

// FlowToGrpc converts from domain Model Flow type to grpc flow type
func FlowToGrpc(flow []*model.Flow) ([]*fspb.Action, error) {
	if flow == nil {
		return nil, nil
	}

	var resultFlow []*fspb.Action
	for _, item := range flow {
		resultFlowCur := fspb.Action{}
		resultFlowCur.Id = item.Id
		resultFlowCur.Name = item.Name
		resultFlowCur.Params = utils.ObjectToDictionary(item.Params)
		resultFlowCur.Library = &fspb.Library{
			Name: item.Library.V,
			V:    item.Library.Name,
		}
		resultFlowCur.Definition = item.Definition
		resultFlowCur.Children, _ = FlowToGrpc(item.Children)
		resultFlow = append(resultFlow, &resultFlowCur)
	}

	return resultFlow, nil
}
