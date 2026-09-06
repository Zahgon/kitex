package generator

type feature int

const (
	placeHolder feature = iota
)

var (
	featureMap = make(map[string]feature)
	maxFeature = placeHolder
)

func HasFeature(list []feature, key string) bool { _ = "STUB: not implemented"; return false }

func RegisterFeature(key string) { _ = "STUB: not implemented"; return }

func getFeature(key string) (feature, bool) { _ = "STUB: not implemented"; return *new(feature), false }
