package retry

const (
	maxBackupRetryTimes     = 2
	defaultBackupRetryTimes = 1
)

func NewBackupPolicy(delayMS uint32) *BackupPolicy { _ = "STUB: not implemented"; return nil }

func (p *BackupPolicy) WithMaxRetryTimes(retryTimes int) { _ = "STUB: not implemented"; return }

func (p *BackupPolicy) DisableChainRetryStop() { _ = "STUB: not implemented"; return }

func (p *BackupPolicy) WithRetryBreaker(errRate float64) { _ = "STUB: not implemented"; return }

func (p *BackupPolicy) WithRetrySameNode() { _ = "STUB: not implemented"; return }

func (p *BackupPolicy) String() string { _ = "STUB: not implemented"; return "" }

func (p *BackupPolicy) Equals(np *BackupPolicy) bool { _ = "STUB: not implemented"; return false }

func (p *BackupPolicy) DeepCopy() *BackupPolicy { _ = "STUB: not implemented"; return nil }
