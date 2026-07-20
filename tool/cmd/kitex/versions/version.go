package versions

import (
	"regexp"
)

const (
	partsNum = 3
)

var gitSemanticVersionRegexp = regexp.MustCompile(`^v(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

type version struct {
	parts []int

	original string

	preRelease string

	buildMetadata string
}

func newVersion(verStr string) (*version, error) { _ = "STUB: not implemented"; return nil, nil }

func (ver *version) lessThan(other *version) bool { _ = "STUB: not implemented"; return false }

func (ver *version) greatOrEqual(other *version) bool { _ = "STUB: not implemented"; return false }

func (ver *version) String() string { _ = "STUB: not implemented"; return "" }
