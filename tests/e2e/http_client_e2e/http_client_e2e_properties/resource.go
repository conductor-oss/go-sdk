package http_client_e2e_properties

import (
	"fmt"
	"os"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	log "github.com/sirupsen/logrus"
)

type TreasureChest struct {
	ImportantValue string `json:"importantValue"`
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

var (
	TASK_OUTPUT = map[string]interface{}{
		"hello": "world",
	}

	WORKER_THREAD_COUNT     = 50
	WORKER_POLLING_INTERVAL = 500 * time.Millisecond

	WORKFLOW_EXECUTION_AMOUNT = 15

	WORKFLOW_DEFINITIONS = []http_model.WorkflowDef{
		WORKFLOW_DEFINITION,
		TREASURE_WORKFLOW_DEFINITION,
	}
	TASK_DEFINITIONS = []http_model.TaskDef{
		TASK_DEFINITION,
		TREASURE_TASK_DEFINITION,
	}
)

var (
	WORKFLOW_NAME = "workflow_with_go_task_example_from_code"

	TASK_NAME = "GO_TASK_OF_SIMPLE_TYPE"

	WORKFLOW_DEFINITION = http_model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        WORKFLOW_NAME,
		Description: "Workflow with go task example from code",
		Version:     1,
		Tasks: []http_model.WorkflowTask{
			{
				Name:              TASK_NAME,
				TaskReferenceName: TASK_NAME,
				Type_:             "SIMPLE",
				StartDelay:        0,
				Optional:          false,
				AsyncComplete:     false,
			},
		},
		OutputParameters: map[string]interface{}{
			"workerOutput": "${go_task_example_from_code_ref_0.output}",
		},
		SchemaVersion:                 2,
		Restartable:                   true,
		WorkflowStatusListenerEnabled: false,
		OwnerEmail:                    "gustavo.gardusi@orkes.io",
		TimeoutPolicy:                 "ALERT_ONLY",
		TimeoutSeconds:                0,
	}

	TASK_DEFINITION = http_model.TaskDef{
		Name:                        TASK_NAME,
		Description:                 "Go task example from code",
		RetryCount:                  3,
		TimeoutSeconds:              300,
		InputKeys:                   make([]string, 0),
		OutputKeys:                  make([]string, 0),
		TimeoutPolicy:               "TIME_OUT_WF",
		RetryLogic:                  "FIXED",
		RetryDelaySeconds:           10,
		ResponseTimeoutSeconds:      180,
		InputTemplate:               make(map[string]interface{}),
		RateLimitPerFrequency:       0,
		RateLimitFrequencyInSeconds: 1,
		OwnerEmail:                  "gustavo.gardusi@orkes.io",
		BackoffScaleFactor:          1,
	}
)

var (
	TREASURE_CHEST_WORKFLOW_NAME = "treasure_chest_workflow"
	TREASURE_CHEST_TASK_NAME     = "treasure_chest_task"

	TREASURE_WORKFLOW_DEFINITION = http_model.WorkflowDef{
		UpdateTime:  1650595431465,
		Name:        TREASURE_CHEST_WORKFLOW_NAME,
		Description: "What's inside the treasure chest?",
		Version:     1,
		Tasks: []http_model.WorkflowTask{
			{
				Name:              TREASURE_CHEST_TASK_NAME,
				TaskReferenceName: TREASURE_CHEST_TASK_NAME,
				Type_:             "SIMPLE",
				StartDelay:        0,
				Optional:          false,
				AsyncComplete:     false,
				InputParameters: map[string]interface{}{
					"importantValue": "${workflow.input.importantValue}",
				},
			},
		},
		InputParameters: []string{"importantValue"},
		OutputParameters: map[string]interface{}{
			"workerOutput": fmt.Sprintf("${%s.output}", TREASURE_CHEST_TASK_NAME),
		},
		SchemaVersion:                 2,
		Restartable:                   true,
		WorkflowStatusListenerEnabled: false,
		OwnerEmail:                    "gustavo.gardusi@orkes.io",
		TimeoutPolicy:                 "ALERT_ONLY",
		TimeoutSeconds:                0,
	}

	TREASURE_TASK_DEFINITION = http_model.TaskDef{
		Name:                        TREASURE_CHEST_TASK_NAME,
		Description:                 "Go task example from code",
		RetryCount:                  3,
		TimeoutSeconds:              300,
		InputKeys:                   []string{"importantValue"},
		OutputKeys:                  make([]string, 0),
		TimeoutPolicy:               "TIME_OUT_WF",
		RetryLogic:                  "FIXED",
		RetryDelaySeconds:           10,
		ResponseTimeoutSeconds:      180,
		RateLimitPerFrequency:       0,
		RateLimitFrequencyInSeconds: 1,
		OwnerEmail:                  "gustavo.gardusi@orkes.io",
		BackoffScaleFactor:          1,
	}

	IMPORTANT_VALUE = "Go is really nice :)"

	TREASURE_WORKFLOW_INPUT = &TreasureChest{
		ImportantValue: IMPORTANT_VALUE,
	}
)
