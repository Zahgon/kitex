package descriptor

func countParams(path string) uint16 { _ = "STUB: not implemented"; return 0 }

type nodeType uint8

const (
	static nodeType = iota
	param
	catchAll
	paramLabel = byte(':')
	anyLabel   = byte('*')
	slash      = "/"
	nilString  = ""
)

type (
	node struct {
		nType    nodeType
		label    byte
		prefix   string
		parent   *node
		children children

		ppath string

		pnames     []string
		function   *FunctionDescriptor
		paramChild *node
		anyChild   *node

		isLeaf bool
	}
	children []*node
)

func checkPathValid(path string) { _ = "STUB: not implemented"; return }

func (n *node) addRoute(path string, function *FunctionDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (n *node) insert(path string, function *FunctionDescriptor, t nodeType, ppath string, pnames []string) {
	_ = "STUB: not implemented"
	return
}

func (n *node) getValue(path string, params func() *Params, unescape bool) (function *FunctionDescriptor, ps *Params, tsr bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (n *node) findChild(l byte) *node { _ = "STUB: not implemented"; return nil }

func (n *node) findChildWithLabel(l byte) *node { _ = "STUB: not implemented"; return nil }

func newNode(t nodeType, pre string, p *node, child children, f *FunctionDescriptor, ppath string, pnames []string, paramChildren, anyChildren *node) *node {
	_ = "STUB: not implemented"
	return nil
}
