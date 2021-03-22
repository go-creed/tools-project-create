package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	CMD      = "cmd"
	INTERNAL = "internal"
	PKG      = "pkg"
	CONFIG   = "configs"
	SCRIPTS  = "scripts"
	BUILD    = "build"
	GITHOOKS = "githooks"
	TEST     = "test"
)

var contents = []string{CMD, INTERNAL, PKG, CONFIG, SCRIPTS, BUILD, GITHOOKS, TEST}

type Generate struct {
	packageName string
	packagePath string

	withGit   bool
	withGoMod bool
}

func NewGenerate(packageName, packagePath string, withGit, withGoMod bool) *Generate {
	return &Generate{
		packageName: packageName,
		packagePath: packagePath,
		withGit:     withGit,
		withGoMod:   withGoMod,
	}
}

func (g *Generate) Output() {
	err := g.mkdirPackage()
	if err != nil {
		log.Fatalf("mkdirPackage err:%s", err)
	}

	err = g.mkdirContents()
	if err != nil {
		log.Fatalf("mkdirContents err:%s", err)
	}

	err = g.goModAndGitInit()
	if err != nil {
		log.Fatalf("goModAndGitInit err:%s", err)
	}

	err = g.golangCi()
	if err != nil {
		log.Fatalf("golangCi err:%s", err)
	}
}

func (g *Generate) mkdirPackage() (err error) {
	g.packagePath, err = filepath.Abs(g.packagePath)
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if g.packagePath == "" {
		g.packagePath = wd
	}

	err = os.MkdirAll(filepath.Join(g.packagePath, g.packageName), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generate) mkdirContents() (err error) {
	for i := range contents {
		err = os.Mkdir(filepath.Join(g.packagePath, g.packageName, contents[i]), os.ModePerm)
		if err != nil {
			return err
		}
		if contents[i] == GITHOOKS {
			restoreAssets("assets/pre-commit", filepath.Join(g.packagePath, g.packageName, contents[i]), "pre-commit")
			restoreAssets("assets/commit-msg", filepath.Join(g.packagePath, g.packageName, contents[i]), "commit-msg")
		}
	}

	return err
}

func (g *Generate) goModAndGitInit() error {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf(`
	#!/bin/bash
	cd %s
	go mod init %s
	git init
	cp githooks/commit-msg .git/hooks/commit-msg
	cp githooks/pre-commit .git/hooks/pre-commit`, filepath.Join(g.packagePath, g.packageName), g.packageName))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return g.gitIgnore()
}

func (g *Generate) gitIgnore() error {
	return restoreAssets("assets/.gitignore", filepath.Join(g.packagePath, g.packageName), ".gitignore")
}

func (g *Generate) golangCi() error {
	return restoreAssets("assets/.golangci.yaml", filepath.Join(g.packagePath, g.packageName), ".golangci.yaml")
}

func restoreAssets(assetName, packagePath, fileName string) error {
	data, err := Asset(assetName)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(packagePath, fileName), data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
