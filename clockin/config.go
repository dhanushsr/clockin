package clockin

import (
	"fmt"
	"os"
	"path"
	"unicode"

	"gopkg.in/yaml.v3"
)

const (
	DefaultBaseDirPath = "~/.clockin/"
	ConfigFileName     = "config.yml"
)

type Config struct {
	BaseDir string
}

func (c *Config) GetBaseDir() (string, error) {
	var baseDir string
	if len(c.BaseDir) > 0 {
		baseDir = c.BaseDir
	} else {
		baseDir = DefaultBaseDirPath
	}
	baseDir, err := SanitizePath(baseDir)
	if err != nil {
		return "", err
	}
	return baseDir, nil
}

func (c *Config) SetBaseDir(dirPath string) error {
	if len(dirPath) == 0 {
		dirPath = DefaultBaseDirPath
	}
	dirPath, err := SanitizePath(dirPath)
	if err != nil {
		return err
	}
	err = CreateIfNotPresent(dirPath)
	if err != nil {
		return err
	}
	c.BaseDir = dirPath
	return nil
}

func (c *Config) Save() error {
	baseDir, err := SanitizePath(c.BaseDir)
	if err != nil {
		return err
	}
	err = CreateIfNotPresent(baseDir)
	if err != nil {
		return err
	}
	marshalledConfig, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	os.WriteFile(path.Join(baseDir, ConfigFileName), marshalledConfig, 0644)
	return nil
}

func (c *Config) Print() {
	baseDir, err := c.GetBaseDir()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ClockIn: Current Config")
	fmt.Println()
	fmt.Printf("Base Directory : %s\n", baseDir)
	fmt.Println()
}

func (c *Config) ReadAndSave() {
	var baseDir string
	fmt.Println()
	fmt.Print("Enter path to base directory: ")
	fmt.Scanln(&baseDir)
	c.SetBaseDir(baseDir)
	fmt.Println()
	if err := c.Save(); err != nil {
		fmt.Println(err)
	}
}

func (c *Config) LoadDefaultConfig() error {
	defaultBaseDir, err := c.GetBaseDir()
	if err != nil {
		return err
	}
	c.BaseDir = defaultBaseDir
	return nil
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	for i := 0; i < len(ValidConfigPaths); i++ {
		validConfigPath, err := SanitizePath(ValidConfigPaths[i])
		if err != nil {
			return nil, err
		}
		isExist, err := DoesPathExist(validConfigPath)
		if err != nil {
			return nil, err
		}
		if isExist {
			marshalledConfig, err := os.ReadFile(validConfigPath)
			if err != nil {
				return nil, err
			}
			err = yaml.Unmarshal(marshalledConfig, &config)
			if err != nil {
				return nil, err
			}
			return config, nil
		}
	}
	return nil, nil
}

func InitializeConfig() (*Config, error) {
	config := &Config{}
	fmt.Print("Load Default Config (Y/n) ? ")
	var shouldLoadDefaultConfig rune
	fmt.Scanf("%c", &shouldLoadDefaultConfig)
	fmt.Scanln() // Flush stdin for next input.
	if unicode.ToLower(shouldLoadDefaultConfig) == 'y' {
		err := config.LoadDefaultConfig()
		if err != nil {
			return nil, err
		}
	} else {
		config.ReadAndSave()
	}
	config.Print()
	return config, nil
}

func LoadOrInitialiseConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	if config == nil {
		config, err = InitializeConfig()
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
