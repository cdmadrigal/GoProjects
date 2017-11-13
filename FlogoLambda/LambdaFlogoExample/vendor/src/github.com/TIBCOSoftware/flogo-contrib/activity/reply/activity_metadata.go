package reply

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "tibco-reply",
  "type": "flogo:activity",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/reply",
  "version": "0.0.1",
  "title": "Reply To Trigger",
  "description": "Simple Reply Activity",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/activity/reply",
  "inputs":[
    {
      "name": "code",
      "type": "integer",
      "required": true
    },
    {
      "name": "data",
      "type": "any"
    }
  ],
  "outputs": [
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
