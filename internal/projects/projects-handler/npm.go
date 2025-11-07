package projects_handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/soft4dev/clonei/internal/color"
)

/* NPM */
type npmProjectHandler struct{}

func (n npmProjectHandler) Install(projectDir string) error {
	if _, err := exec.LookPath("npm"); err != nil {
		return fmt.Errorf("npm not found; please install npm and ensure it's on your PATH")
	}

	color.PrintSuccess("  → Running npm ci...")
	init := exec.Command("npm", "ci")
	init.Dir = projectDir
	init.Stdout = os.Stdout
	init.Stderr = os.Stderr
	init.Stdin = os.Stdin
	if err := init.Run(); err != nil {
		return fmt.Errorf("error initializing project (npm ci): %w", err)
	}

	return nil
}

type NpmProjectType struct{}

func (npmProjectType *NpmProjectType) Name() string {
	return "npm"
}

func (npmProjectType *NpmProjectType) Detect(projectPath string) (IProjectHandler, error) {
	npmLockPath := filepath.Join(projectPath, "package-lock.json")
	if _, err := os.Stat(npmLockPath); err == nil {
		return npmProjectHandler{}, nil
	}
	return nil, nil
}

func (npmProjectType *NpmProjectType) DefaultProjectHandler() IProjectHandler {
	return npmProjectHandler{}
}
