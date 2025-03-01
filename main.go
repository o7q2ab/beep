package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/o7q2ab/beep/config"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	createModule := len(args) == 0 || (len(args) == 1 && args[0] != "pkg")
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getwd: %w", err)
	}

	fmt.Printf("Beeping in %s. ", wd)

	dir, err := mkdir(config.Dir)
	if err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	fmt.Printf("Directory %s.\n", dir)

	fullpath := filepath.Join(wd, dir)

	if createModule {
		cmd := exec.Command("go", "mod", "init", config.Mod)
		cmd.Dir = fullpath
		out, err := cmd.CombinedOutput()
		if len(out) != 0 {
			fmt.Println(string(out))
		}
		if err != nil {
			return fmt.Errorf("run: %w", err)
		}
	}

	mainfile := filepath.Join(fullpath, "main.go")
	if err = os.WriteFile(mainfile, []byte(file), 0644); err != nil {
		return fmt.Errorf("writefile: %w", err)
	}
	return nil
}

func mkdir(name string) (string, error) {
	for i := 0; ; i++ {
		if i != 0 {
			name = name + strconv.Itoa(i)
		}
		err := os.Mkdir(name, os.ModePerm)
		if err == nil {
			return name, nil
		}
		if errors.Is(err, os.ErrExist) {
			continue
		}
		return "", err
	}
}

const file = `package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("bebeep [%s]\n", runtime.Version())

	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
`
