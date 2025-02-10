package fm_connect

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

// List all flows of ine instance
func ListFlows(ctx context.Context, client *connect.Client, instanceId string) (*[]types.ContactFlowSummary, error) {
	input := &connect.ListContactFlowsInput{
		InstanceId: aws.String(instanceId),
	}

	resp, err := client.ListContactFlows(ctx, input)
	// ###TODO: logging
	if err != nil {
		fmt.Println("Got an error listing flows:")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Flows:")
	for _, flow := range resp.ContactFlowSummaryList {
		fmt.Println(*flow.Name)
	}
	return &resp.ContactFlowSummaryList, nil
}

// Get the flow
func GetFlow(ctx context.Context, client *connect.Client, instanceId string, flowId string) (*types.ContactFlow, error) {
	params := &connect.DescribeContactFlowInput{
		InstanceId:    aws.String(instanceId),
		ContactFlowId: aws.String(flowId),
	}

	resp, err := client.DescribeContactFlow(ctx, params)
	if err != nil {
		fmt.Println("Got an error getting flow:")
		fmt.Println(err)
		return nil, err
	}

	return resp.ContactFlow, nil
}
