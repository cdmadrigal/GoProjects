package activity

import (
	"encoding/json"

	"github.com/TIBCOSoftware/flogo-lib/core/data"
)

// Metadata is the metadata for the Activity
type Metadata struct {
	ID      string
	Input   map[string]*data.Attribute
	Output  map[string]*data.Attribute
}

// NewMetadata creates the metadata object from its json representation
func NewMetadata(jsonMetadata string) *Metadata {
	md := &Metadata{}
	err := json.Unmarshal([]byte(jsonMetadata), md)
	if err != nil {
		panic("Unable to parse activity metadata: " + err.Error())
	}

	return md
}

// UnmarshalJSON overrides the default UnmarshalJSON for TaskEnv
func (md *Metadata) UnmarshalJSON(b []byte) error {

	ser := &struct {
		Name    string            `json:"name"`
		Ref     string            `json:"ref"`
		Input   []*data.Attribute `json:"input"`
		Output  []*data.Attribute `json:"output"`
		//for backwards compatibility
		Inputs  []*data.Attribute `json:"inputs"`
		Outputs []*data.Attribute `json:"outputs"`
	}{}

	if err := json.Unmarshal(b, ser); err != nil {
		return err
	}

	if len(ser.Ref) > 0 {
		md.ID = ser.Ref
	} else {
		// Added for backwards compatibility
		// TODO remove and add a proper error once the BC is removed
		md.ID = ser.Name
	}

	md.Input = make(map[string]*data.Attribute, len(ser.Inputs))
	md.Output = make(map[string]*data.Attribute, len(ser.Outputs))

	if len(ser.Input) > 0 {
		for _, attr := range ser.Input {
			md.Input[attr.Name] = attr
		}
	} else {
		// for backwards compatibility
		for _, attr := range ser.Inputs {
			md.Input[attr.Name] = attr
		}
	}

	if len(ser.Output) > 0 {
		for _, attr := range ser.Output {
			md.Output[attr.Name] = attr
		}
	} else {
		// for backwards compatibility
		for _, attr := range ser.Outputs {
			md.Output[attr.Name] = attr
		}
	}

	return nil
}
