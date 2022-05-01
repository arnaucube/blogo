package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		color.Red(path)
	}
	check(err)
	return string(dat)
}

func writeFile(path string, newContent string) {
	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	check(err)

	color.Green(path)
	//color.Blue(newContent)
}

func getLines(text string) []string {
	lines := strings.Split(text, "\n")
	return lines
}

func concatStringsWithJumps(lines []string) string {
	var r string
	for _, l := range lines {
		r = r + l + "\n"
	}
	return r
}

func copyRaw(original string, destination string) {
	color.Green(original + " --> to --> " + destination)
	_, err := exec.Command("cp", "-rf", original, destination).Output()
	check(err)
}

var watcherInputFiles, watcherPublic *fsnotify.Watcher

func watch(dir string) {
	var err error
	var watcher *fsnotify.Watcher
	if dir == "./blogo-input" {
		watcherInputFiles, err = fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		watcher = watcherInputFiles
	} else {
		watcherPublic, err = fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		watcher = watcherPublic
	}

	defer watcher.Close()

	if dir == "./blogo-input" {
		if err := filepath.Walk(dir, watchInputFilesDir); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := filepath.Walk(dir, watchPublicDir); err != nil {
			log.Fatal(err)
		}
	}

	for {
		select {
		case event := <-watcher.Events:
			fmt.Printf("file system event: %#v\n", event)
			if dir == "./blogo-input" {
				generateHTML()
			}
		case err := <-watcher.Errors:
			log.Fatal("file system watcher error:", err)
		}
	}
}

// watchInputFilesDir gets run as a walk func, searching for directories to add watchers to
func watchInputFilesDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcherInputFiles.Add(path)
	}

	return nil
}

// watchPublicDir gets run as a walk func, searching for directories to add watchers to
func watchPublicDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcherPublic.Add(path)
	}

	return nil
}
