package main

import (
	"fmt"
)

type (
	App struct {
		downloader Downloader
	}

	Downloader interface {
		Download(url string) error
	}
)

func newApp(downloader Downloader) *App {
	return &App{downloader: downloader}
}

func (a *App) Run(url string) error {
	fmt.Printf("Downloading %s", url)
	err := a.downloader.Download(url)
	if err != nil {
		fmt.Println("oooo no - it failed :(")
		return err
	}

	return nil
}