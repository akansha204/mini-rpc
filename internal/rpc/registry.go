package rpc

type HandlerFunc func(params []interface{}) (interface{}, error)

type Registry struct {
	methods map[string]HandlerFunc
}

func NewRegistry() *Registry {
	return &Registry{
		methods: make(map[string]HandlerFunc),
	}
}

func (r *Registry) Register(name string, handler HandlerFunc) {
	r.methods[name] = handler
}

func (r *Registry) Get(name string) (HandlerFunc, bool) {
	h, ok := r.methods[name]
	return h, ok
}
