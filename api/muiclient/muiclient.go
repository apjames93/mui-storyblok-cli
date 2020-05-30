package muiclient

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"

	"github.com/apjames93/mui-storyblok-cli/pkg/config"
)

func clone(filedir string) error {
	cmd := exec.Command("git", "clone", "https://github.com/apjames93/mui-storyblok-scaffold.git", filedir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func moveFile(filedir string) error {
	cmd := exec.Command("mv", ".env", filedir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func makeEnv(cm string, filedir string) error {
	d1 := []byte(cm)
	f, err := os.Create(".env")
	defer f.Close()

	wf := ioutil.WriteFile(".env", d1, 0644)
	if wf != nil {
		return err
	}

	mv := moveFile(filedir)
	if mv != nil {
		return mv
	}

	return nil
}

// GitRepo will clone mui-storyblok-scaffold
func GitRepo(c *config.Config) error {
	home, err := homeDir()
	if err != nil {
		return err
	}

	filedir := fmt.Sprintf("%s/storyblok", home)
	clone(filedir)

	e := makeEnv(c.GenerateEnvString(), filedir)
	if e != nil {
		return e
	}

	return nil
}

func homeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
