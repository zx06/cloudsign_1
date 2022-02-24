package provider

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Provider = (*Sample)(nil)

type SampleCfg struct {
	URL           string
	Method        string
	Header        map[string]string
	Body          string
	RespSuccess   string
	RespFailure   string
	RespDuplicate string
}

type Sample struct {
	SignConfig
	SampleCfg
}

func NewSample(name string, cfg *SampleCfg) (*Sample, error) {
	s := &Sample{
		SignConfig: SignConfig{
			Name:         name,
			ProviderName: "sample",
		},
		SampleCfg: *cfg,
	}
	err := s.MarshalConfig()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (c *Sample) Sign(ctx context.Context) (resp *SignResp, err error) {
	resp = &SignResp{}
	client := &http.Client{}
	req, err := http.NewRequest(c.Method, c.URL, strings.NewReader(c.Body))
	if err != nil {
		return nil, err
	}
	for k, v := range c.Header {
		req.Header.Add(k, v)
	}
	httpResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	resp.Body, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	resp.StatusCode = httpResp.StatusCode
	return

}

func (c *Sample) CheckSuccess(ctx context.Context, resp *SignResp) CheckState {
	if c.RespSuccess != "" && strings.Contains(string(resp.Body), c.RespSuccess) {
		return CheckStateSuccess
	}
	if c.RespFailure != "" && strings.Contains(string(resp.Body), c.RespFailure) {
		return CheckStateFailure
	}
	if c.RespDuplicate != "" && strings.Contains(string(resp.Body), c.RespDuplicate) {
		return CheckStateDuplicate
	}
	return CheckStateUnknown
}

func (c *Sample) MarshalConfig() error {
	var err error
	c.Data, err = json.Marshal(c.SampleCfg)
	if err != nil {
		return err
	}
	return nil

}

func (c *Sample) UnmarshalConfig() error {
	return json.Unmarshal(c.Data, &c.SampleCfg)
}
