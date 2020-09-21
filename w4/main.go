package main

import (
	"fmt"
	"os"
	_ "path/filepath"
)

func main() {
	youtubeDownloader := newYoutubeDownloader()
	app := newApp(youtubeDownloader)
	url := os.Args[1]

	err := app.Run(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Download completed!")
}
