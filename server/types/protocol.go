package types

import "strings"

type ResultCode int

const (
	Success ResultCode = 0
	Failed  ResultCode = 1 // 요청이 실패하였습니다.

)

func (r ResultCode) toString() string {
	switch r {
	case Success:
		return "success"
	case Failed:
		return "failed"
	}

	return ""
}

type RespHeader struct {
	Result       ResultCode `json:"result"`
	ResultString string     `json:"resultString"`
	Desc         string     `json:"desc"`
}

type RespWithData struct {
	*RespHeader
	Data interface{} `json:"data"`
}

func NewRespHeader(resultCode ResultCode, desc ...string) *RespHeader {
	return &RespHeader{
		Result:       resultCode,
		ResultString: resultCode.toString(),
		Desc:         strings.Join(desc, ","),
	}
}
