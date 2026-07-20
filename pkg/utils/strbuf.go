//lint:file-ignore SA6002 allocations cannot be avoided

package utils

import (
	"strings"
	"sync"
)

var intBytesPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 20)
	},
}

func WriteInt64ToStringBuilder(sb *strings.Builder, value int64) { _ = "STUB: not implemented"; return }
