package readme

import "fmt"

type Service struct {
	templates map[string]Template
}

func NewService() Service {
	return Service{templates: Templates()}
}

func (s Service) List() []Template {
	out := make([]Template, 0, len(s.templates))
	for _, tpl := range s.templates {
		out = append(out, tpl)
	}
	return out
}

func (s Service) Generate(cfg ProjectConfig) (string, error) {
	if cfg.Template == "" {
		return "", fmt.Errorf("template is required")
	}
	if cfg.Title == "" {
		return "", fmt.Errorf("title is required")
	}

	applyDerivedDefaults(&cfg)
	tpl, ok := s.templates[cfg.Template]
	if !ok {
		return "", fmt.Errorf("unknown template %q", cfg.Template)
	}
	return tpl.Build(cfg)
}
