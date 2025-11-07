package projects

import (
	projectHandler "github.com/soft4dev/clonei/internal/projects/projects-handler"
)

type ProjectType interface {
	Name() string
	Detect(projectPath string) (projectHandler.IProjectHandler, error)
	DefaultProjectHandler() projectHandler.IProjectHandler
}

type ProjectDetector struct {
	projectTypes []ProjectType
}

func (projectDetector *ProjectDetector) FindProjectHandler(projectPath string) (projectHandler.IProjectHandler, error) {
	for _, projectType := range projectDetector.projectTypes {
		projectHandler, err := projectType.Detect(projectPath)
		if err != nil {
			return nil, err
		}
		if projectHandler != nil {
			return projectHandler, nil
		}
	}
	return nil, nil
}

func (projectDetector *ProjectDetector) FindProjectHandlerFromName(name string) projectHandler.IProjectHandler {
	for _, projectType := range projectDetector.projectTypes {
		if name == projectType.Name() {
			return projectType.DefaultProjectHandler()
		}
	}
	return nil
}

func (projectDetector *ProjectDetector) RegisterDetector(projectType ProjectType) {
	projectDetector.projectTypes = append(projectDetector.projectTypes, projectType)
}

func (projectDetector *ProjectDetector) GetAvailableProjectTypes() []string {
	var projectTypes = []string{}
	for _, projectType := range projectDetector.projectTypes {
		projectTypes = append(projectTypes, projectType.Name())
	}
	return projectTypes
}

func DefaultDetector() ProjectDetector {
	projectDetector := ProjectDetector{}
	projectDetector.RegisterDetector(&projectHandler.PnpmProjectType{})
	projectDetector.RegisterDetector(&projectHandler.NpmProjectType{})
	projectDetector.RegisterDetector(&projectHandler.CargoProjectType{})
	return projectDetector
}
