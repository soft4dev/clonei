package projects_handler

type IProjectHandler interface {
	Install(projectDir string) error
}
