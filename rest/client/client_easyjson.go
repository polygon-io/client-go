// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package client

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

func easyjsonC0e5e3f1DecodeGithubComPolygonIoClientGolangRestClient(in *jlexer.Lexer, out *BaseResponse) {
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
		case "status":
			out.Status = string(in.String())
		case "request_id":
			out.RequestID = string(in.String())
		case "count":
			out.Count = int(in.Int())
		case "error":
			out.Error = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "next_url":
			out.NextURL = string(in.String())
		case "previous_url":
			out.PreviousURL = string(in.String())
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
func easyjsonC0e5e3f1EncodeGithubComPolygonIoClientGolangRestClient(out *jwriter.Writer, in BaseResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"request_id\":"
		out.RawString(prefix)
		out.String(string(in.RequestID))
	}
	if in.Count != 0 {
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Int(int(in.Count))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.NextURL != "" {
		const prefix string = ",\"next_url\":"
		out.RawString(prefix)
		out.String(string(in.NextURL))
	}
	if in.PreviousURL != "" {
		const prefix string = ",\"previous_url\":"
		out.RawString(prefix)
		out.String(string(in.PreviousURL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC0e5e3f1EncodeGithubComPolygonIoClientGolangRestClient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC0e5e3f1EncodeGithubComPolygonIoClientGolangRestClient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC0e5e3f1DecodeGithubComPolygonIoClientGolangRestClient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC0e5e3f1DecodeGithubComPolygonIoClientGolangRestClient(l, v)
}
