package provider

import (
	"context"
)

type CheckState uint

const (
	CheckStateUnknown   CheckState = iota
	CheckStateSuccess              // 签到成功
	CheckStateFailure              // 签到失败
	CheckStateDuplicate            // 重复签到

)

type SignConfig struct {
	Name         string
	ProviderName string
	// 额外数据
	Data []byte
}

type SignResp struct {
	StatusCode int
	Body       []byte
	// 额外数据
	ExtraData map[string]interface{}
}

type Provider interface {
	// 签到
	Sign(ctx context.Context) (resp *SignResp, err error)
	// 检测签到状态
	CheckSuccess(ctx context.Context, resp *SignResp) CheckState
	// 序列化配置
	MarshalConfig() error
	// 反序列化配置
	UnmarshalConfig() error
}
