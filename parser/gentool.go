// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const antlrVersion = `4.7.1`
const antlrDownloadURL = `http://www.antlr.org/download`

func main() {
	defer func() {
		// allow scripting helpers to panic ;)
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no command")
	}

	switch args[0] {
	case "antlr":
		antlr(args[1:])

	case "replace":
		replace(args[1:])
	}
}

func antlr(args []string) {
	genDir := absPath(".build")
	if !existsDir(genDir) {
		mkdir(genDir)
	}

	antlrJar := fmt.Sprintf("antlr-%s-complete.jar", antlrVersion)
	antlrJarPath := filepath.Join(genDir, antlrJar)
	if !existsFile(antlrJarPath) {
		fmt.Println("Download antlr")

		n := downloadFile(antlrJarPath, fmt.Sprintf("%v/%v", antlrDownloadURL, antlrJar))
		fmt.Printf("Finished downloading antlr. %v bytes have been downloaded.\n", n)
	}

	antlrArgs := []string{"-jar", antlrJarPath, "-Dlanguage=Go", "-Werror"}
	runExternal("java", append(antlrArgs, args...))
}

func replace(args []string) {
	if len(args) != 3 {
		log.Fatal("required arguments <old string> <new string> <file>")
	}

	oldString := []byte(args[0])
	newString := []byte(args[1])
	filename := args[2]

	contents := readFile(filename)
	contents = bytes.Replace(contents, oldString, newString, -1)
	writeFile(filename, contents)
}

func runExternal(path string, args []string) {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func existsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

func existsFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.Mode().IsRegular()
}

func mkdir(name string) {
	if err := os.MkdirAll(name, 0755); err != nil {
		panic(err)
	}
}

func absPath(name string) string {
	p, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}
	return p
}

func downloadFile(file, url string) int64 {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body := resp.Body
	defer body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("download failed with http status %v", resp.StatusCode))
	}

	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	numBytes, err := io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	return numBytes
}

func readFile(file string) []byte {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return contents
}

func writeFile(file string, contents []byte) {
	err := ioutil.WriteFile(file, contents, 0644)
	if err != nil {
		panic(err)
	}
}
