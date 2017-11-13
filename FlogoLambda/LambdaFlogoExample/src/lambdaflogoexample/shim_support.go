// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/app"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/logger"

)

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "LambdaFlogoExample",
  "type": "flogo:app",
  "version": "0.0.1",
  "description": "",
  "triggers": [
    {
      "name": "Start Flow as a function in Lambda",
      "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/lambda",
      "description": "Simple Lambda Trigger",
      "settings": {},
      "id": "start_flow_as_a_function_in_lambda",
      "handlers": [
        {
          "settings": {},
          "actionId": "lambda_toggle"
        }
      ]
    }
  ],
  "actions": [
    {
      "name": "LambdaToggle",
      "data": {
        "flow": {
          "type": 1,
          "attributes": [],
          "explicitReply": true,
          "rootTask": {
            "id": 1,
            "type": 1,
            "tasks": [
              {
                "id": 2,
                "name": "Log Message",
                "description": "Simple Log Activity",
                "type": 1,
                "activityType": "github-com-tibco-software-flogo-contrib-activity-log",
                "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/log",
                "attributes": [
                  {
                    "name": "message",
                    "value": "Calling mqtt activity to toggle light",
                    "required": false,
                    "type": "string"
                  },
                  {
                    "name": "flowInfo",
                    "value": "false",
                    "required": false,
                    "type": "boolean"
                  },
                  {
                    "name": "addToFlow",
                    "value": "true",
                    "required": false,
                    "type": "boolean"
                  }
                ],
                "inputMappings": [
                  {
                    "type": 1,
                    "value": "{T.evt}.request.intent.slots.OnOff.value",
                    "mapTo": "message"
                  }
                ]
              },
              {
                "id": 7,
                "name": "Invoke REST Service",
                "description": "Simple REST Activity",
                "type": 1,
                "activityType": "tibco-rest",
                "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/rest",
                "attributes": [
                  {
                    "name": "method",
                    "value": "POST",
                    "required": true,
                    "type": "string"
                  },
                  {
                    "name": "uri",
                    "value": "https://api.particle.io/v1/devices/48001f000351353530373132/setLightOn?access_token=db16715f6d23cddbca63ac665384abce09702710",
                    "required": true,
                    "type": "string"
                  },
                  {
                    "name": "pathParams",
                    "value": null,
                    "required": false,
                    "type": "params"
                  },
                  {
                    "name": "queryParams",
                    "value": null,
                    "required": false,
                    "type": "params"
                  },
                  {
                    "name": "content",
                    "value": "{\"arg\": \"on\"}",
                    "required": false,
                    "type": "any"
                  }
                ]
              },
              {
                "id": 8,
                "name": "Reply To Trigger",
                "description": "Simple Reply Activity",
                "type": 1,
                "activityType": "tibco-reply",
                "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/reply",
                "attributes": [
                  {
                    "name": "code",
                    "value": "200",
                    "required": true,
                    "type": "integer"
                  },
                  {
                    "name": "data",
                    "value": { "version": "1.0", "response": { "outputSpeech": { "type": "PlainText", "text": "Okay, the light has been turned on." }, "shouldEndSession": true }, "sessionAttributes": {} },
                    "required": false,
                    "type": "any"
                  }
                ]
              },
              {
                "id": 9,
                "name": "Invoke REST Service",
                "description": "Simple REST Activity",
                "type": 1,
                "activityType": "tibco-rest",
                "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/rest",
                "attributes": [
                  {
                    "name": "method",
                    "value": "POST",
                    "required": true,
                    "type": "string"
                  },
                  {
                    "name": "uri",
                    "value": "https://api.particle.io/v1/devices/48001f000351353530373132/setLightOff?access_token=db16715f6d23cddbca63ac665384abce09702710",
                    "required": true,
                    "type": "string"
                  },
                  {
                    "name": "pathParams",
                    "value": null,
                    "required": false,
                    "type": "params"
                  },
                  {
                    "name": "queryParams",
                    "value": null,
                    "required": false,
                    "type": "params"
                  },
                  {
                    "name": "content",
                    "value": "{\"arg\": \"off\"}",
                    "required": false,
                    "type": "any"
                  }
                ]
              },
              {
                "id": 10,
                "name": "Reply To Trigger (2)",
                "description": "Simple Reply Activity",
                "type": 1,
                "activityType": "tibco-reply",
                "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/reply",
                "attributes": [
                  {
                    "name": "code",
                    "value": "200",
                    "required": true,
                    "type": "integer"
                  },
                  {
                    "name": "data",
                    "value": { "version": "1.0", "response": { "outputSpeech": { "type": "PlainText", "text": "Okay, the light has been turned off." }, "shouldEndSession": true }, "sessionAttributes": {} },
                    "required": false,
                    "type": "any"
                  }
                ]
              }
            ],
            "links": [
              {
                "id": 1,
                "from": 2,
                "to": 7,
                "type": 1,
                "value": "${A2.message}  == \"on\""
              },
              {
                "id": 2,
                "from": 7,
                "to": 8,
                "type": 0
              },
              {
                "id": 3,
                "from": 2,
                "to": 9,
                "type": 1,
                "value": "${A2.message}  == \"off\""
              },
              {
                "id": 4,
                "from": 9,
                "to": 10,
                "type": 0
              }
            ],
            "attributes": []
          }
        }
      },
      "id": "lambda_toggle",
      "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow"
    }
  ]
}
`

func init() {
	config.SetDefaultLogLevel("ERROR")
	logger.SetLogLevel(logger.ErrorLevel)

	var cp app.ConfigProvider

	if flogoJSON != "" {
		cp = EmbeddedProvider()
	} else {
		cp = app.DefaultConfigProvider()
	}

	app, err := cp.GetApp()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	e, err := engine.New(app)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	e.Init(true)
}

// embeddedConfigProvider implementation of ConfigProvider
type embeddedProvider struct {
}

//EmbeddedProvider returns an app config from a compiled json file
func EmbeddedProvider() (app.ConfigProvider){
	return &embeddedProvider{}
}

// GetApp returns the app configuration
func (d *embeddedProvider) GetApp() (*app.Config, error){

	appCfg := &app.Config{}
	err := json.Unmarshal([]byte(flogoJSON), appCfg)
	if err != nil {
		return nil, err
	}
	return appCfg, nil
}
