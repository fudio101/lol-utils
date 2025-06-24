package main

import (
	"embed"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"lol-utils/internal/app"
)

//go:embed all:frontend-dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	println("Requesting file:", requestedFilename)
	
	// Define a safe directory
	const safeDir = "./safe-files/"
	
	// Resolve the path relative to the safe directory
	absPath, err := filepath.Abs(filepath.Join(safeDir, requestedFilename))
	if err != nil || !strings.HasPrefix(absPath, safeDir) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Invalid file path: %s", html.EscapeString(requestedFilename))))
		return
	}
	
	// Read the file from the resolved safe path
	fileData, err := os.ReadFile(absPath)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", html.EscapeString(requestedFilename))))
		return
	}
	
	res.Write(fileData)
}

func main() {
	// Create an instance of the app structure
	appInstance := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "LoL Utils",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appInstance.Startup,
		Bind: []interface{}{
			appInstance,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
