package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	root := "drills"
	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		fmt.Fprintf(os.Stderr, "no encuentro la carpeta %s (ejecuta desde la raíz del módulo: go run ./cmd)\n", root)
		os.Exit(1)
	}

	for {
		categories, err := listDirs(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "no puedo leer categorías: %v\n", err)
			os.Exit(1)
		}
		if len(categories) == 0 {
			fmt.Println("No hay categorías en drills/")
			return
		}

		fmt.Println("Categorías")
		for i, name := range categories {
			fmt.Printf("  %d. %s\n", i+1, name)
		}
		fmt.Print("> ")

		choice, ok := readChoice()
		if !ok || choice == "q" {
			return
		}
		n, err := strconv.Atoi(choice)
		if err != nil || n < 1 || n > len(categories) {
			fmt.Println("opción no válida")
			fmt.Println()
			continue
		}
		if !runCategory(filepath.Join(root, categories[n-1]), categories[n-1]) {
			return
		}
	}
}

// runCategory returns false when the user quits.
func runCategory(dir, name string) bool {
	for {
		drills, err := listGoFiles(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "no puedo leer %s: %v\n", name, err)
			return true
		}

		fmt.Printf("\n%s\n", name)
		if len(drills) == 0 {
			fmt.Println("  (sin ejercicios)")
		}
		for i, file := range drills {
			fmt.Printf("  %d. %s\n", i+1, strings.TrimSuffix(file, ".go"))
		}
		fmt.Println("  0. volver")
		fmt.Print("> ")

		choice, ok := readChoice()
		if !ok || choice == "q" {
			return false
		}
		if choice == "0" {
			fmt.Println()
			return true
		}
		n, err := strconv.Atoi(choice)
		if err != nil || n < 1 || n > len(drills) {
			fmt.Println("opción no válida")
			continue
		}

		fmt.Println()
		path := filepath.Join(dir, drills[n-1])
		if err := runDrill(path); err != nil {
			fmt.Fprintf(os.Stderr, "error al ejecutar %s: %v\n", drills[n-1], err)
		}
		fmt.Println()
	}
}

func listDirs(root string) ([]string, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			names = append(names, entry.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}

func listGoFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasPrefix(name, "_") {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

func runDrill(path string) error {
	cmd := exec.Command("go", "run", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// readChoice reads a single line without buffering ahead, so a drill can use stdin afterwards.
func readChoice() (string, bool) {
	var buf []byte
	tmp := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(tmp)
		if n > 0 {
			if tmp[0] == '\n' {
				return strings.ToLower(strings.TrimSpace(string(buf))), true
			}
			if tmp[0] != '\r' {
				buf = append(buf, tmp[0])
			}
		}
		if err != nil {
			if err == io.EOF && len(buf) > 0 {
				return strings.ToLower(strings.TrimSpace(string(buf))), true
			}
			return "", false
		}
	}
}
