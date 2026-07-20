package retry

const maxMixRetryTimes = 3

func NewMixedPolicy(delayMS uint32) *MixedPolicy { _ = "STUB: not implemented"; return nil }

func NewMixedPolicyWithResultRetry(delayMS uint32, rr *ShouldResultRetry) *MixedPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MixedPolicy) String() string { _ = "STUB: not implemented"; return "" }

func (p *MixedPolicy) Equals(np *MixedPolicy) bool { _ = "STUB: not implemented"; return false }

func (p *MixedPolicy) DeepCopy() *MixedPolicy { _ = "STUB: not implemented"; return nil }
