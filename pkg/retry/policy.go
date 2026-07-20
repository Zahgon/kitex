package retry

import (
	"context"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type Type int

const (
	FailureType Type = iota
	BackupType
	MixedType
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func BuildFailurePolicy(p *FailurePolicy) Policy { _ = "STUB: not implemented"; return *new(Policy) }

func BuildBackupRequest(p *BackupPolicy) Policy { _ = "STUB: not implemented"; return *new(Policy) }

func BuildMixedPolicy(p *MixedPolicy) Policy { _ = "STUB: not implemented"; return *new(Policy) }

type Policy struct {
	Enable bool `json:"enable"`

	Type Type `json:"type"`

	FailurePolicy *FailurePolicy `json:"failure_policy,omitempty"`
	BackupPolicy  *BackupPolicy  `json:"backup_policy,omitempty"`
	MixedPolicy   *MixedPolicy   `json:"mixed_policy,omitempty"`
}

func (p *Policy) DeepCopy() *Policy { _ = "STUB: not implemented"; return nil }

type FailurePolicy struct {
	StopPolicy        StopPolicy         `json:"stop_policy"`
	BackOffPolicy     *BackOffPolicy     `json:"backoff_policy,omitempty"`
	RetrySameNode     bool               `json:"retry_same_node"`
	ShouldResultRetry *ShouldResultRetry `json:"-"`

	Extra string `json:"extra"`
}

type BackupPolicy struct {
	RetryDelayMS  uint32     `json:"retry_delay_ms"`
	StopPolicy    StopPolicy `json:"stop_policy"`
	RetrySameNode bool       `json:"retry_same_node"`
}

type MixedPolicy struct {
	RetryDelayMS uint32 `json:"retry_delay_ms"`
	FailurePolicy
}

type StopPolicy struct {
	MaxRetryTimes    int      `json:"max_retry_times"`
	MaxDurationMS    uint32   `json:"max_duration_ms"`
	DisableChainStop bool     `json:"disable_chain_stop"`
	DDLStop          bool     `json:"ddl_stop"`
	CBPolicy         CBPolicy `json:"cb_policy"`
}

const (
	defaultCBErrRate = 0.1
	cbMinSample      = 10
)

type CBPolicy struct {
	ErrorRate float64 `json:"error_rate"`
}

type BackOffPolicy struct {
	BackOffType BackOffType               `json:"backoff_type"`
	CfgItems    map[BackOffCfgKey]float64 `json:"cfg_items,omitempty"`
}

type BackOffType string

const (
	NoneBackOffType   BackOffType = "none"
	FixedBackOffType  BackOffType = "fixed"
	RandomBackOffType BackOffType = "random"
)

type BackOffCfgKey string

const (
	FixMSBackOffCfgKey      BackOffCfgKey = "fix_ms"
	MinMSBackOffCfgKey      BackOffCfgKey = "min_ms"
	MaxMSBackOffCfgKey      BackOffCfgKey = "max_ms"
	InitialMSBackOffCfgKey  BackOffCfgKey = "initial_ms"
	MultiplierBackOffCfgKey BackOffCfgKey = "multiplier"
)

type ShouldResultRetry struct {
	ErrorRetryWithCtx func(ctx context.Context, err error, ri rpcinfo.RPCInfo) bool

	RespRetryWithCtx func(ctx context.Context, resp interface{}, ri rpcinfo.RPCInfo) bool

	ErrorRetry func(err error, ri rpcinfo.RPCInfo) bool

	RespRetry func(resp interface{}, ri rpcinfo.RPCInfo) bool

	NotRetryForTimeout bool
}

func (p Policy) Equals(np Policy) bool { _ = "STUB: not implemented"; return false }

func (p *BackOffPolicy) Equals(np *BackOffPolicy) bool { _ = "STUB: not implemented"; return false }

func (p *BackOffPolicy) DeepCopy() *BackOffPolicy { _ = "STUB: not implemented"; return nil }

func (p *BackOffPolicy) copyCfgItems() map[BackOffCfgKey]float64 {
	_ = "STUB: not implemented"
	return nil
}

func (rr *ShouldResultRetry) IsValid() bool { _ = "STUB: not implemented"; return false }

func checkCBErrorRate(p *CBPolicy) error { _ = "STUB: not implemented"; return nil }

func checkStopPolicy(sp *StopPolicy, maxRetryTimes int, retryer Retryer) error {
	_ = "STUB: not implemented"
	return nil
}
