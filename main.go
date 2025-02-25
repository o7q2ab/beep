package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	Dir = "bebeep"
	Mod = "bebeep"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getwd: %w", err)
	}
	fmt.Println("Beeping in", wd)
	if err = os.Mkdir(Dir, os.ModePerm); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	cmd := exec.Command("go", "mod", "init", Mod)
	cmd.Dir = filepath.Join(wd, Dir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}
	if strout := string(out); strout != "" {
		fmt.Println(strout)
	}
	mainfile := filepath.Join(wd, Dir, "main.go")
	if err = os.WriteFile(mainfile, []byte(file), 0644); err != nil {
		return fmt.Errorf("writefile: %w", err)
	}
	return nil
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
