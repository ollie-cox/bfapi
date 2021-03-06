// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bfapi

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

func easyjson5d2f9a13DecodeGithubComTarbBfapi(in *jlexer.Lexer, out *RORunner) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "selectionId":
			out.SelectionID = SelectionID(in.Uint64())
		case "handicap":
			out.Handicap = float64(in.Float64())
		case "description":
			easyjson5d2f9a13Decode(in, &out.Description)
		case "state":
			easyjson5d2f9a13Decode1(in, &out.State)
		case "sp":
			easyjson5d2f9a13Decode2(in, &out.Sp)
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi(out *jwriter.Writer, in RORunner) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"selectionId\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.SelectionID))
	}
	{
		const prefix string = ",\"handicap\":"
		out.RawString(prefix)
		out.Float64(float64(in.Handicap))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode(out, in.Description)
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode1(out, in.State)
	}
	{
		const prefix string = ",\"sp\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode2(out, in.Sp)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RORunner) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RORunner) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RORunner) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RORunner) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi(l, v)
}
func easyjson5d2f9a13Decode2(in *jlexer.Lexer, out *struct {
	ActualStartingPrice PriceVol `json:"actualStartingPrice"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "actualStartingPrice":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ActualStartingPrice).UnmarshalJSON(data))
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
func easyjson5d2f9a13Encode2(out *jwriter.Writer, in struct {
	ActualStartingPrice PriceVol `json:"actualStartingPrice"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"actualStartingPrice\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.ActualStartingPrice))
	}
	out.RawByte('}')
}
func easyjson5d2f9a13Decode1(in *jlexer.Lexer, out *struct {
	SortPriority int    `json:"sortPriority"`
	Status       string `json:"status"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "sortPriority":
			out.SortPriority = int(in.Int())
		case "status":
			out.Status = string(in.String())
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
func easyjson5d2f9a13Encode1(out *jwriter.Writer, in struct {
	SortPriority int    `json:"sortPriority"`
	Status       string `json:"status"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"sortPriority\":"
		out.RawString(prefix[1:])
		out.Int(int(in.SortPriority))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}
func easyjson5d2f9a13Decode(in *jlexer.Lexer, out *struct {
	RunnerName string `json:"runnerName"`
	Metadata   struct {
		RunnerID uint64 `json:"runnerId,string"`
	} `json:"metadata"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "runnerName":
			out.RunnerName = string(in.String())
		case "metadata":
			easyjson5d2f9a13Decode3(in, &out.Metadata)
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
func easyjson5d2f9a13Encode(out *jwriter.Writer, in struct {
	RunnerName string `json:"runnerName"`
	Metadata   struct {
		RunnerID uint64 `json:"runnerId,string"`
	} `json:"metadata"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"runnerName\":"
		out.RawString(prefix[1:])
		out.String(string(in.RunnerName))
	}
	{
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode3(out, in.Metadata)
	}
	out.RawByte('}')
}
func easyjson5d2f9a13Decode3(in *jlexer.Lexer, out *struct {
	RunnerID uint64 `json:"runnerId,string"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "runnerId":
			out.RunnerID = uint64(in.Uint64Str())
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
func easyjson5d2f9a13Encode3(out *jwriter.Writer, in struct {
	RunnerID uint64 `json:"runnerId,string"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"runnerId\":"
		out.RawString(prefix[1:])
		out.Uint64Str(uint64(in.RunnerID))
	}
	out.RawByte('}')
}
func easyjson5d2f9a13DecodeGithubComTarbBfapi1(in *jlexer.Lexer, out *ROResult) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventTypes":
			if in.IsNull() {
				in.Skip()
				out.Events = nil
			} else {
				in.Delim('[')
				if out.Events == nil {
					if !in.IsDelim(']') {
						out.Events = make([]ROEventType, 0, 2)
					} else {
						out.Events = []ROEventType{}
					}
				} else {
					out.Events = (out.Events)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ROEventType
					(v1).UnmarshalEasyJSON(in)
					out.Events = append(out.Events, v1)
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi1(out *jwriter.Writer, in ROResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventTypes\":"
		out.RawString(prefix[1:])
		if in.Events == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Events {
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
func (v ROResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ROResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ROResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ROResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi1(l, v)
}
func easyjson5d2f9a13DecodeGithubComTarbBfapi2(in *jlexer.Lexer, out *ROMarketNode) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "marketId":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.MarketID).UnmarshalJSON(data))
			}
		case "runners":
			if in.IsNull() {
				in.Skip()
				out.Runners = nil
			} else {
				in.Delim('[')
				if out.Runners == nil {
					if !in.IsDelim(']') {
						out.Runners = make([]RORunner, 0, 1)
					} else {
						out.Runners = []RORunner{}
					}
				} else {
					out.Runners = (out.Runners)[:0]
				}
				for !in.IsDelim(']') {
					var v4 RORunner
					(v4).UnmarshalEasyJSON(in)
					out.Runners = append(out.Runners, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "description":
			(out.Description).UnmarshalEasyJSON(in)
		case "rates":
			easyjson5d2f9a13Decode4(in, &out.Rates)
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi2(out *jwriter.Writer, in ROMarketNode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"marketId\":"
		out.RawString(prefix[1:])
		out.Raw((in.MarketID).MarshalJSON())
	}
	{
		const prefix string = ",\"runners\":"
		out.RawString(prefix)
		if in.Runners == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Runners {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		(in.Description).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"rates\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode4(out, in.Rates)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ROMarketNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ROMarketNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ROMarketNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ROMarketNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi2(l, v)
}
func easyjson5d2f9a13Decode4(in *jlexer.Lexer, out *struct {
	MarketBaseRate  float64 `json:"marketBaseRate"`
	DiscountAllowed bool    `json:"discountAllowed"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "marketBaseRate":
			out.MarketBaseRate = float64(in.Float64())
		case "discountAllowed":
			out.DiscountAllowed = bool(in.Bool())
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
func easyjson5d2f9a13Encode4(out *jwriter.Writer, in struct {
	MarketBaseRate  float64 `json:"marketBaseRate"`
	DiscountAllowed bool    `json:"discountAllowed"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"marketBaseRate\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.MarketBaseRate))
	}
	{
		const prefix string = ",\"discountAllowed\":"
		out.RawString(prefix)
		out.Bool(bool(in.DiscountAllowed))
	}
	out.RawByte('}')
}
func easyjson5d2f9a13DecodeGithubComTarbBfapi3(in *jlexer.Lexer, out *ROMarketDescription) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "marketName":
			out.MarketName = string(in.String())
		case "marketTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.MarketTime).UnmarshalJSON(data))
			}
		case "suspendTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.SuspendTime).UnmarshalJSON(data))
			}
		case "settleTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.SettleTime).UnmarshalJSON(data))
			}
		case "marketType":
			out.MarketType = string(in.String())
		case "regulator":
			out.Regulator = string(in.String())
		case "eachWayDivisor":
			out.EachWayDivisor = float64(in.Float64())
		case "priceLadderDescription":
			easyjson5d2f9a13Decode5(in, &out.PriceLadderDescription)
		case "bettingType":
			out.BettingType = string(in.String())
		case "persistenceEnabled":
			out.PersistenceEnabled = bool(in.Bool())
		case "bspMarket":
			out.BspMarket = bool(in.Bool())
		case "turnInPlayEnabled":
			out.TurnInPlayEnabled = bool(in.Bool())
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi3(out *jwriter.Writer, in ROMarketDescription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"marketName\":"
		out.RawString(prefix[1:])
		out.String(string(in.MarketName))
	}
	{
		const prefix string = ",\"marketTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.MarketTime))
	}
	{
		const prefix string = ",\"suspendTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.SuspendTime))
	}
	{
		const prefix string = ",\"settleTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.SettleTime))
	}
	{
		const prefix string = ",\"marketType\":"
		out.RawString(prefix)
		out.String(string(in.MarketType))
	}
	{
		const prefix string = ",\"regulator\":"
		out.RawString(prefix)
		out.String(string(in.Regulator))
	}
	{
		const prefix string = ",\"eachWayDivisor\":"
		out.RawString(prefix)
		out.Float64(float64(in.EachWayDivisor))
	}
	{
		const prefix string = ",\"priceLadderDescription\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode5(out, in.PriceLadderDescription)
	}
	{
		const prefix string = ",\"bettingType\":"
		out.RawString(prefix)
		out.String(string(in.BettingType))
	}
	{
		const prefix string = ",\"persistenceEnabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.PersistenceEnabled))
	}
	{
		const prefix string = ",\"bspMarket\":"
		out.RawString(prefix)
		out.Bool(bool(in.BspMarket))
	}
	{
		const prefix string = ",\"turnInPlayEnabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.TurnInPlayEnabled))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ROMarketDescription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ROMarketDescription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ROMarketDescription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ROMarketDescription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi3(l, v)
}
func easyjson5d2f9a13Decode5(in *jlexer.Lexer, out *struct {
	Type string `json:"type"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
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
func easyjson5d2f9a13Encode5(out *jwriter.Writer, in struct {
	Type string `json:"type"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	out.RawByte('}')
}
func easyjson5d2f9a13DecodeGithubComTarbBfapi4(in *jlexer.Lexer, out *ROEventType) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventTypeId":
			out.EventTypeID = uint64(in.Uint64())
		case "eventNodes":
			if in.IsNull() {
				in.Skip()
				out.EventNodes = nil
			} else {
				in.Delim('[')
				if out.EventNodes == nil {
					if !in.IsDelim(']') {
						out.EventNodes = make([]ROEventNode, 0, 1)
					} else {
						out.EventNodes = []ROEventNode{}
					}
				} else {
					out.EventNodes = (out.EventNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v7 ROEventNode
					(v7).UnmarshalEasyJSON(in)
					out.EventNodes = append(out.EventNodes, v7)
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi4(out *jwriter.Writer, in ROEventType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventTypeId\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.EventTypeID))
	}
	{
		const prefix string = ",\"eventNodes\":"
		out.RawString(prefix)
		if in.EventNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.EventNodes {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ROEventType) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ROEventType) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ROEventType) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ROEventType) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi4(l, v)
}
func easyjson5d2f9a13DecodeGithubComTarbBfapi5(in *jlexer.Lexer, out *ROEventNode) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventId":
			out.EventID = uint64(in.Uint64())
		case "marketNodes":
			if in.IsNull() {
				in.Skip()
				out.MarketNodes = nil
			} else {
				in.Delim('[')
				if out.MarketNodes == nil {
					if !in.IsDelim(']') {
						out.MarketNodes = make([]ROMarketNode, 0, 1)
					} else {
						out.MarketNodes = []ROMarketNode{}
					}
				} else {
					out.MarketNodes = (out.MarketNodes)[:0]
				}
				for !in.IsDelim(']') {
					var v10 ROMarketNode
					(v10).UnmarshalEasyJSON(in)
					out.MarketNodes = append(out.MarketNodes, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "event":
			easyjson5d2f9a13Decode6(in, &out.Event)
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
func easyjson5d2f9a13EncodeGithubComTarbBfapi5(out *jwriter.Writer, in ROEventNode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventId\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.EventID))
	}
	{
		const prefix string = ",\"marketNodes\":"
		out.RawString(prefix)
		if in.MarketNodes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.MarketNodes {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix)
		easyjson5d2f9a13Encode6(out, in.Event)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ROEventNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5d2f9a13EncodeGithubComTarbBfapi5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ROEventNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5d2f9a13EncodeGithubComTarbBfapi5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ROEventNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5d2f9a13DecodeGithubComTarbBfapi5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ROEventNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5d2f9a13DecodeGithubComTarbBfapi5(l, v)
}
func easyjson5d2f9a13Decode6(in *jlexer.Lexer, out *struct {
	EventName   string `json:"eventName"`
	CountryCode string `json:"countryCode"`
	Timezone    string `json:"timezone"`
	Venue       string `json:"venue"`
	OpenDate    Time   `json:"openDate"`
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventName":
			out.EventName = string(in.String())
		case "countryCode":
			out.CountryCode = string(in.String())
		case "timezone":
			out.Timezone = string(in.String())
		case "venue":
			out.Venue = string(in.String())
		case "openDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.OpenDate).UnmarshalJSON(data))
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
func easyjson5d2f9a13Encode6(out *jwriter.Writer, in struct {
	EventName   string `json:"eventName"`
	CountryCode string `json:"countryCode"`
	Timezone    string `json:"timezone"`
	Venue       string `json:"venue"`
	OpenDate    Time   `json:"openDate"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventName\":"
		out.RawString(prefix[1:])
		out.String(string(in.EventName))
	}
	{
		const prefix string = ",\"countryCode\":"
		out.RawString(prefix)
		out.String(string(in.CountryCode))
	}
	{
		const prefix string = ",\"timezone\":"
		out.RawString(prefix)
		out.String(string(in.Timezone))
	}
	{
		const prefix string = ",\"venue\":"
		out.RawString(prefix)
		out.String(string(in.Venue))
	}
	{
		const prefix string = ",\"openDate\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.OpenDate))
	}
	out.RawByte('}')
}
