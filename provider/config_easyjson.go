// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package provider

import (
	json "encoding/json"
	isp "github.com/deweppro/go-auth/provider/isp"
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

func easyjson6615c02eDecodeGithubComDewepproGoAuthProvider(in *jlexer.Lexer, out *Config) {
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
		case "auth_provider":
			if in.IsNull() {
				in.Skip()
				out.Provider = nil
			} else {
				in.Delim('[')
				if out.Provider == nil {
					if !in.IsDelim(']') {
						out.Provider = make([]isp.Config, 0, 1)
					} else {
						out.Provider = []isp.Config{}
					}
				} else {
					out.Provider = (out.Provider)[:0]
				}
				for !in.IsDelim(']') {
					var v1 isp.Config
					(v1).UnmarshalEasyJSON(in)
					out.Provider = append(out.Provider, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6615c02eEncodeGithubComDewepproGoAuthProvider(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"auth_provider\":"
		out.RawString(prefix[1:])
		if in.Provider == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Provider {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeGithubComDewepproGoAuthProvider(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeGithubComDewepproGoAuthProvider(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeGithubComDewepproGoAuthProvider(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeGithubComDewepproGoAuthProvider(l, v)
}
