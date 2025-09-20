package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	output, err := executeCode(string(code), 2*time.Minute)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error executing code", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "\n%s", output)
}

func executeCode(code string, timeout time.Duration) (string, error) {
	tempFile, err := ioutil.TempFile("", "remote-go-code-*.go")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(code)
	if err != nil {
		return "", err
	}
	tempFile.Close()

	cmd := exec.Command("go", "run", tempFile.Name())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case <-time.After(timeout):
		cmd.Process.Kill()
		return "", fmt.Errorf("execution exceeded timeout of %v", timeout)
	case err := <-done:
		if err != nil {
			return "", err
		}
	}

	outputBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	return string(outputBytes), nil
}
func handlerFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodConnect {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalPath := r.URL.Path

	if originalPath == "/" {
		originalPath = "flag.txt"
	}
	cleanedPath := cleanPath(originalPath)
	flagContent, err := readFile(cleanedPath)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Flag file content: %s", flagContent)
}

func cleanPath(path string) string {
	cleanPath := filepath.Clean(path)
	return cleanPath
}

func readFile(path string) (string, error) {
	flagContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return string(flagContent), nil
}

func main() {
	http.HandleFunc("/", handlerFile)
	http.HandleFunc("/deploy", handler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
