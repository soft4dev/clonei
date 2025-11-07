package projects_handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/soft4dev/clonei/internal/color"
)

/* PNPM */
type pnpmProjectHandler struct{}

func (n pnpmProjectHandler) Install(projectDir string) error {
	if _, err := exec.LookPath("pnpm"); err != nil {
		return fmt.Errorf("pnpm not found; please install pnpm and ensure it's on your PATH")
	}

	color.PrintSuccess("  → Running pnpm install --frozen-lockfile...")
	init := exec.Command("pnpm", "install", "--frozen-lockfile")
	init.Dir = projectDir
	init.Stdout = os.Stdout
	init.Stderr = os.Stderr
	init.Stdin = os.Stdin
	if err := init.Run(); err != nil {
		return fmt.Errorf("error initializing project (pnpm install): %w", err)
	}

	return nil
}

type PnpmProjectType struct{}

func (pnpmProjectType *PnpmProjectType) Name() string {
	return "pnpm"
}

func (pnpmProjectType *PnpmProjectType) Detect(projectPath string) (IProjectHandler, error) {
	pnpmLockPath := filepath.Join(projectPath, "pnpm-lock.yaml")
	if _, err := os.Stat(pnpmLockPath); err == nil {
		return pnpmProjectHandler{}, nil
	}
	return nil, nil
}

func (pnpmProjectType *PnpmProjectType) DefaultProjectHandler() IProjectHandler {
	return pnpmProjectHandler{}
}
