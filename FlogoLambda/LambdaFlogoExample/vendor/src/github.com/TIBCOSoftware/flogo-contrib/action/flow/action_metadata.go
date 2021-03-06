package flow

var jsonMetadata = `{
  "name": "flow",
  "type": "flogo:action",
  "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
  "version": "0.0.1",
  "title": "Simple Flow",
  "description": "Simple Flow Action",
  "homepage": "https://github.com/TIBCOSoftware/flogo-contrib/tree/master/action/flow",
  "run_options":[
    {
      "name": "op",
      "type": "int"
    },
    {
      "name": "flowURI",
      "type": "string"
    },
    {
      "name": "returnId",
      "type": "boolean"
    },
    {
      "name": "initialState",
      "type": "any"
    }
  ]
}
`

func getJsonMetadata() string {
	return jsonMetadata
}
