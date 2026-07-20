package versions

import (
	"errors"
	"regexp"
)

const (
	CompatibilityCheckExitCode = 3
)

var dm *dependencyManager

func init() {
	dm = &dependencyManager{}
}

func RegisterMinDepVersion(depVer *MinDepVersion) error { _ = "STUB: not implemented"; return nil }

func CheckDependency() *CheckResult { _ = "STUB: not implemented"; return nil }

func DefaultCheckDependencyAndProcess() error { _ = "STUB: not implemented"; return nil }

func defaultParseCheckResult(cr *CheckResult) (prompt string, shouldExit bool) {
	_ = "STUB: not implemented"
	return "", false
}

var defaultPrompt = `# Kitex Cmd Tool %s is not compatible with %s %s in your go.mod
# You can upgrade %s to latest version
go get %s@latest
# Or upgrade %s to %s version
go get %s@%s
# Or downgrade Kitex Cmd Tool to %s version
go install github.com/cloudwego/kitex/tool/cmd/kitex@%s`

func defaultPromptWithCheckResult(cr *CheckResult) string { _ = "STUB: not implemented"; return "" }

type MinDepVersion struct {
	RefPath string

	Version string

	ver      *version
	goModVer *version
}

func (m *MinDepVersion) init() error {
	if m == nil {
		return errors.New("nil MinDepVersion")
	}
	if m.RefPath == "" {
		return errors.New("empty RefPath")
	}
	if m.Version == "" {
		return errors.New("empty Version")
	}
	ver, err := newVersion(m.Version)
	if err != nil {
		return err
	}
	m.ver = ver

	return nil
}

func (m *MinDepVersion) parseGoModVersion() error { _ = "STUB: not implemented"; return nil }

func (m *MinDepVersion) getGoModVersion() string { _ = "STUB: not implemented"; return "" }

func (m *MinDepVersion) isCompatible() bool { _ = "STUB: not implemented"; return false }

var (
	ErrGoCmdNotFound                   = errors.New("go cmd not found")
	ErrGoModNotFound                   = errors.New("go.mod file not found in current directory or any parent directory")
	ErrDependencyNotFound              = errors.New("dependency not found")
	ErrDependencyVersionNotCompatible  = errors.New("dependency not compatible")
	ErrDependencyVersionNotSemantic    = errors.New("dependency version is not semantic version")
	ErrDependencyReplacedWithLocalRepo = errors.New("dependency replaced with local repo")
)

type CheckResult struct {
	ver      *MinDepVersion
	goModVer string
	err      error
}

func (cr *CheckResult) MinDepVersion() *MinDepVersion { _ = "STUB: not implemented"; return nil }

func (cr *CheckResult) Err() error { _ = "STUB: not implemented"; return nil }

func (cr *CheckResult) GoModVersion() string { _ = "STUB: not implemented"; return "" }

type dependencyManager struct {
	depVer *MinDepVersion
}

func (dm *dependencyManager) Register(depVer *MinDepVersion) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *dependencyManager) CheckDependency() *CheckResult { _ = "STUB: not implemented"; return nil }

func runGoListCmd(refPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func runCommand(input string) (string, error) { _ = "STUB: not implemented"; return "", nil }

var goListVersionRegexp = regexp.MustCompile(`^(\S+)\s+(\S+)(\s+=>\s+)?(\S+)?(\s+)?(\S+)?$`)

func parseGoListVersion(str string) string { _ = "STUB: not implemented"; return "" }

func generateCheckResult(depVer *MinDepVersion, err error) *CheckResult {
	_ = "STUB: not implemented"
	return nil
}
