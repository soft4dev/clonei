package projects

type IProjectHandler interface {
	Install(projectDir string) error
}
