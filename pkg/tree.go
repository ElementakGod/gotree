package pkg

import (
	"fmt"
	"go/build"
	"os"
)

const (
	Cmd        = "cmd"
	Pkg        = "pkg"
	Internal   = "internal"
	Api        = "api"
	Web        = "web"
	Configs    = "configs"
	Init       = "init"
	Scripts    = "scripts"
	Build      = "build"
	Deploy     = "deploy"
	Test       = "test"
	Docs       = "docs"
	Examples   = "examples"
	Tools      = "tools"
	ThirdParty = "third_party"
	GitHooks   = "githooks"
	Assets     = "assets"
	WebSite    = "website"
)

var projectCatalogSlice = [...]string{Cmd, Pkg, Internal, Api, Web, Configs, Init, Scripts, Build,
	Deploy, Test, Docs, Examples, Tools, ThirdParty, GitHooks, Assets, WebSite}

type ProjectTree struct {
	ProjectName string // 项目名
	ProjectPath string // 项目路径，如果为空 则为当前路径
}

func NewProjectTree(name, path string) *ProjectTree {
	var fPath string
	if path == "" {
		fPath = name
	} else {
		fPath = fmt.Sprintf("%s/%s", path, name)
	}
	return &ProjectTree{ProjectName: name, ProjectPath: fPath}
}

// Setup
// @Description: 初始化项目的工程结构
// @receiver p
// @return error
func (p *ProjectTree) Setup() error {
	err := p.createFileCatalog()
	if err != nil {
		return err
	}
	err = p.createMod()
	if err != nil {
		return err
	}
	return nil
}

// createFileCatalog
// @Description: 创建文件目录
// @receiver p
// @return error
func (p *ProjectTree) createFileCatalog() error {
	err := os.Mkdir(p.ProjectPath, os.ModePerm)
	if err != nil {
		return err
	}
	for _, tmp := range projectCatalogSlice {
		err = os.Mkdir(fmt.Sprintf("%s/%s", p.ProjectPath, tmp), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// createMod
// @Description: init go mod
// @receiver p
// @return error
func (p *ProjectTree) createMod() error {
	err := os.Chdir(p.ProjectPath)
	if err != nil {
		return err
	}

	f, err := os.Create("go.mod")
	if err != nil {
		return err
	}
	tags := build.Default.ReleaseTags
	version := tags[len(tags)-1]
	str := fmt.Sprintf("module %s\n\ngo %s", p.ProjectName, version[2:])
	_, err = f.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}
