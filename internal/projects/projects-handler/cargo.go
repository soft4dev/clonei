package projects_handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/soft4dev/clonei/internal/color"
)

type CargoProjectHandler struct{}

func (cargoProjectHandler *CargoProjectHandler) Install(projectDir string) error {
	if _, err := exec.LookPath("cargo"); err != nil {
		return fmt.Errorf("cargo not found; please install cargo and ensure it's on your PATH")
	}

	color.PrintSuccess("  → Running cargo ci...")
	init := exec.Command("cargo", "fetch")
	init.Dir = projectDir
	init.Stdout = os.Stdout
	init.Stderr = os.Stderr
	init.Stdin = os.Stdin
	if err := init.Run(); err != nil {
		return fmt.Errorf("error initializing project (cargo fetch): %w", err)
	}

	return nil
}

type CargoProjectType struct{}

func (cargoProjectType *CargoProjectType) Name() string {
	return "cargo"
}

func (cargoProjectType *CargoProjectType) Detect(projectPath string) (IProjectHandler, error) {
	cargoProjectConfigPath := filepath.Join(projectPath, "Cargo.toml")
	if _, err := os.Stat(cargoProjectConfigPath); err == nil {
		return &CargoProjectHandler{}, nil
	}
	return nil, nil
}

func (cargoProjectType *CargoProjectType) DefaultProjectHandler() IProjectHandler {
	return &CargoProjectHandler{}
}
