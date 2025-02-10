package main

import (
	"context"
	"flag"
	"fmt"

	"log"
	"os"

	"github.com/tecracer/flowmanager/fm_connect"
	"github.com/tecracer/flowmanager/store"
)

func main() {
	// This is the main entry point for the Flow Manager
	// This is a placeholder for the main function
	instanceIDParm := flag.String("id", "", "Instance ID for the flow")

	// Parse the flags
	flag.Parse()

	// Check if instance ID is provided
	if *instanceIDParm == "" {
		log.Fatal("Instance ID is required --id")
	}
	instanceId := *instanceIDParm
	ctx := context.Background()
	clientConnect := fm_connect.InitClient(ctx)
	flows, err := fm_connect.ListFlows(ctx, clientConnect, instanceId)
	if err != nil {
		log.Fatal(err)
	}
	stage := "dev"
	outPutBaseDir := "output/"
	outPutStageDir := "output/" + stage
	flowDir := outPutStageDir + "/flows/"
	modulesDir := outPutStageDir + "/modules/"

	// mkdir output, ignore if it already exists
	err = os.Mkdir(outPutBaseDir, 0755)
	err = os.Mkdir(outPutStageDir, 0755)
	err = os.Mkdir(stage, 0755)
	err = os.Mkdir(flowDir, 0755)
	err = os.Mkdir(modulesDir, 0755)
	for _, flow := range *flows {
		flowName := store.ConvertFlowName(*flow.Name)
		flowId := *flow.Id
		flow, err := fm_connect.GetFlow(ctx, clientConnect, instanceId, flowId)
		if err != nil {
			log.Fatal(err)
		}
		// store the flow to the output directory
		fmt.Printf("Save flow: %v\n", *flow.Name)
		store.StoreFlow(flowDir, flowName, flow)
	}

	// 	modules, err := fm_connect.ListModules(ctx, clientConnect, instanceId)
}
