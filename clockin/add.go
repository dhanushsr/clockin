package clockin

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	PROJECT_FILE string = "projects.yml"
)

type Project struct {
	ProjectName string
	Modules     []Module
}

type Module string

func AddProject(c *Config, project, module string) error {
	sanitizedProjectFilePath, err := SanitizePath(path.Join(c.BaseDir, PROJECT_FILE))
	if err != nil {
		return err
	}
	f, err := os.OpenFile(sanitizedProjectFilePath, os.O_WRONLY|os.O_CREATE, DEFAULT_FILE_MODE)
	if err != nil {
		return err
	}
	defer f.Close()

	var allProjects map[string]Project = make(map[string]Project)
	var marshalledProjects []byte
	if marshalledProjects, err = os.ReadFile(sanitizedProjectFilePath); err != nil {
		return nil
	}
	err = yaml.Unmarshal(marshalledProjects, &allProjects)
	if err != nil {
		return err
	}
	currentProject, isPresent := allProjects[project]
	if !isPresent {
		currentProject.ProjectName = project
	}
	if len(module) > 0 {
		currentProject.Modules = append(currentProject.Modules, Module(module))
	}
	allProjects[project] = currentProject

	if marshalledProjects, err = yaml.Marshal(allProjects); err != nil {
		return err
	}
	if _, err = f.WriteString(string(marshalledProjects)); err != nil {
		return err
	}
	return nil
}

func PrintAllProjects(c *Config) error {
	sanitizedProjectFilePath, err := SanitizePath(path.Join(c.BaseDir, PROJECT_FILE))
	if err != nil {
		return err
	}
	f, err := os.OpenFile(sanitizedProjectFilePath, os.O_RDONLY|os.O_CREATE, DEFAULT_FILE_MODE)
	if err != nil {
		return err
	}
	defer f.Close()

	var allProjects map[string]Project
	var marshalledProjects []byte
	if marshalledProjects, err = os.ReadFile(sanitizedProjectFilePath); err != nil {
		return nil
	}
	err = yaml.Unmarshal(marshalledProjects, &allProjects)
	if err != nil {
		return err
	}
	for _, project := range allProjects {
		marshalledProject, err := yaml.Marshal(project)
		if err != nil {
			return err
		}
		fmt.Println(string(marshalledProject))
	}
	return nil
}
