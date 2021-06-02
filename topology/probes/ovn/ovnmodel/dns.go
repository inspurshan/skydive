//go:generate go run github.com/skydive-project/skydive/graffiti/gendecoder -package github.com/skydive-project/skydive/topology/probes/ovn/ovnmodel
//go:generate go run github.com/mailru/easyjson/easyjson $GOFILE

// Code generated by ovnmetagen
// DO NOT EDIT.

package ovnmodel

import (
	"encoding/json"
	"fmt"

	"github.com/skydive-project/skydive/graffiti/getter"
	"github.com/skydive-project/skydive/graffiti/graph"
)

// DNS defines the type used by both libovsdb and skydive for table DNS
// easyjson:json
// gendecoder
type DNS struct {
	UUID        string            `ovsdb:"_uuid" json:",omitempty" `
	ExternalIDs map[string]string `ovsdb:"external_ids" json:",omitempty" `
	Records     map[string]string `ovsdb:"records" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	RecordsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *DNS) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.RecordsMeta = graph.NormalizeValue(t.Records).(map[string]interface{})

	return graph.Metadata{
		"Type":    "DNS",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *DNS) GetUUID() string {
	return t.UUID
}

func (t *DNS) GetName() string {
	if name := t.UUID; name != "" {
		return name
	}
	return t.GetUUID()
}

// DNSDecoder implements a json message raw decoder
func DNSDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t DNS
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal DNS metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
