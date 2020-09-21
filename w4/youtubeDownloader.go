package main

import (
	"github.com/kkdai/youtube"
)

type (
	youtubeDownloader struct {

	}
)

func newYoutubeDownloader() *youtubeDownloader {
	return &youtubeDownloader{}
}

func (yd youtubeDownloader) Download(url string) error {
	dir := "."

	y := youtube.NewYoutube(false)
	if err := y.DecodeURL(url); err != nil {
		return err
	}
	if err := y.StartDownload(dir, "dl.mp4", "medium",0); err != nil {
		return err
	}
	return nil
}