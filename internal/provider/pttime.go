package provider

import "net/http"

type PTTime struct {
	Sample
}

func NewPTTime(name string, cookies string) (*PTTime, error) {
	cfg := &SampleCfg{
		URL:    "https://www.pttime.org/attendance.php",
		Method: http.MethodGet,
		Header: map[string]string{
			"cookie": cookies,
		},
		Body:          "",
		RespSuccess:   "",
		RespFailure:   "",
		RespDuplicate: "你今天已经签到过了，请勿重复刷新。",
	}
	s, err := NewSample(name, cfg)
	if err != nil {
		return nil, err
	}
	return &PTTime{
		Sample: *s,
	}, nil
}
