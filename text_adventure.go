package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/Racinettee/explorer/pkg/explorer"
	ui "github.com/VladimirMarkelov/clui"
)

func main() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	initPath, err := resolveInitialPath()

	if err != nil {
		log.Println(err)
	}
	log.Printf("Starting explorer @ %v\n", initPath)
	fexp, err := explorer.CreateFileExplorer(0, 0, 30, 20, initPath)

	textWin := ui.AddWindow(30, 1, 20, 10, "Text view")
	textView := ui.CreateTextView(textWin, 40, 10, 1)

	fexp.OnItemClicked = func(path string) {
		content := getFileContent(path)
		textView.SetText(content)
	}
	ui.MainLoop()
}

func getFileContent(path string) []string {
	f, _ := os.Open(path)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func resolveInitialPath() (result string, err error) {
	result, err = os.Getwd()
	if err != nil {
		return ".", err
	}
	switch {
	case len(os.Args) == 1:
		return
	case len(os.Args) > 1:
		if os.Args[1] == "." {
			return result, nil
		}
		if os.Args[1] == ".." {
			result, err = filepath.Abs(os.Args[1])
		}
	}
	return
}
