// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package isp

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

func easyjson947c98ebDecodeGithubComDewepproGoAuthProviderIsp(in *jlexer.Lexer, out *tempUserGoogle) {
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
		case "sub":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "picture":
			out.Icon = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "email_verified":
			out.EmailVerified = bool(in.Bool())
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
func easyjson947c98ebEncodeGithubComDewepproGoAuthProviderIsp(out *jwriter.Writer, in tempUserGoogle) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"sub\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"picture\":"
		out.RawString(prefix)
		out.String(string(in.Icon))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"email_verified\":"
		out.RawString(prefix)
		out.Bool(bool(in.EmailVerified))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v tempUserGoogle) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson947c98ebEncodeGithubComDewepproGoAuthProviderIsp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v tempUserGoogle) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson947c98ebEncodeGithubComDewepproGoAuthProviderIsp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *tempUserGoogle) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson947c98ebDecodeGithubComDewepproGoAuthProviderIsp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *tempUserGoogle) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson947c98ebDecodeGithubComDewepproGoAuthProviderIsp(l, v)
}
