// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package ovnmodel

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5f408f88DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(in *jlexer.Lexer, out *HAChassisGroup) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "UUID":
			out.UUID = string(in.String())
		case "ExternalIDs":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ExternalIDs = make(map[string]string)
				} else {
					out.ExternalIDs = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.ExternalIDs)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "HaChassis":
			if in.IsNull() {
				in.Skip()
				out.HaChassis = nil
			} else {
				in.Delim('[')
				if out.HaChassis == nil {
					if !in.IsDelim(']') {
						out.HaChassis = make([]string, 0, 4)
					} else {
						out.HaChassis = []string{}
					}
				} else {
					out.HaChassis = (out.HaChassis)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.HaChassis = append(out.HaChassis, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Name":
			out.Name = string(in.String())
		case "ExternalIDsMeta":
			(out.ExternalIDsMeta).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5f408f88EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(out *jwriter.Writer, in HAChassisGroup) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"UUID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	if len(in.ExternalIDs) != 0 {
		const prefix string = ",\"ExternalIDs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.ExternalIDs {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				out.String(string(v3Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.HaChassis) != 0 {
		const prefix string = ",\"HaChassis\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.HaChassis {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.ExternalIDsMeta) != 0 {
		const prefix string = ",\"ExternalIDsMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExternalIDsMeta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HAChassisGroup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5f408f88EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HAChassisGroup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5f408f88EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HAChassisGroup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5f408f88DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HAChassisGroup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5f408f88DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(l, v)
}