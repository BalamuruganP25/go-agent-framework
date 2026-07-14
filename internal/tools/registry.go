package tools

type Registry struct {
	tools map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

func (r *Registry) Register(
	tool Tool,
) {
	r.tools[tool.Name()] = tool
}


func (r *Registry) Get(
	name string,
) (Tool, bool) {

	tool, ok := r.tools[name]

	return tool, ok
}