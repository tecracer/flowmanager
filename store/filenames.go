package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

func ConvertFlowName(name string) string {
	// convert all spaces to underscores
	name = strings.ReplaceAll(name, " ", "_")
	return name
}

func StoreFlow(outPutDir string, flowName string, flow *types.ContactFlow) error {
	// store the flow to the output directory
	flowFileName := flowName + ".json"
	contentFileName := filepath.Join(outPutDir + flowFileName)

	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(*flow.Content), "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)

		return err
	}

	err = os.WriteFile(contentFileName, prettyJSON.Bytes(), 0644)
	if err != nil {
		fmt.Println("Got an error storing flow:")
		fmt.Println(err)
		return err
	}
	return nil
}
