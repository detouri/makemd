package readme

type ProjectConfig struct {
	Title       string
	Description string
	Template    string
	Owner       string
	Repo        string
	Module      string
	BinaryName  string

	InstallCommand  string
	RunCommand      string
	TestCommand     string
	CoverageCommand string

	DocsURL string
	DemoURL string

	IssuesURL string

	Audience       string
	Status         string
	MinimumVersion string
	CIProvider     string

	LicenseName string

	Private         bool
	IncludeBadges   bool
	IncludeTOC      bool
	IncludeFAQ      bool
	IncludeRoadmap  bool
	IncludeSecurity bool
	IncludeContrib  bool
	IncludeArch     bool
	IncludeAPI      bool
	IncludeConfig   bool
}

type Template struct {
	Name        string
	Summary     string
	Description string
	Build       func(ProjectConfig) (string, error)
}
