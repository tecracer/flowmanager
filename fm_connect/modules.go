package fm_connect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

func ListModules(ctx context.Context, clientConnect *connect.Client, instanceId *string) (*[]types.ContactFlowModuleSummary, error) {

	parms := &connect.ListContactFlowModulesInput{
		InstanceId: instanceId,
	}
	resp, err := clientConnect.ListContactFlowModules(ctx, parms)
	if err != nil {
		fmt.Println("Got an error listing modules:")
		fmt.Println(err)
		return nil, err
	}

	return &resp.ContactFlowModulesSummaryList, nil
}
