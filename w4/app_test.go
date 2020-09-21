package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

//import (
//	"errors"
//	"github.com/golang/mock/gomock"
//	"testing"
//)
//
//func TestAppRun(t *testing.T){
//	ctrl := gomock.NewController(t)
//
//	mockDownloader := NewMockDownloader(ctrl)
//	app := newApp(mockDownloader)
//
//	url := "https://www.youtube.com/watch?v=mytestvideo"
//	badURL := "@@@@@@@@@@@@@"
//	mockDownloader.EXPECT().Download(url).Return(nil)
//	mockDownloader.EXPECT().Download(badURL).Return(errors.New("can't download video"))
//
//
//	err := app.Run(url)
//	if err != nil {
//		t.Errorf("app.Run() = nil, got = %s", err)
//	}
//
//	err = app.Run(badURL)
//	if err == nil {
//		t.Errorf("app.Run() = %s, got = nil", err)
//	}
//
//	ctrl.Finish()
//}

func TestApp_Run(t *testing.T) {
	goodURL := "https://www.youtube.com/watch?v=mytestvideo"
	badURL := "@@@@@@@@@@@@@"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	goodDownloader := NewMockDownloader(ctrl)
	badDownloader := NewMockDownloader(ctrl)

	goodDownloader.EXPECT().Download(goodURL).Return(nil)
	badDownloader.EXPECT().Download(badURL).Return(errors.New("cant download video"))


	type fields struct {
		downloader Downloader
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "fail to download video",
			fields:  fields{
				downloader: badDownloader,
			},
			args:    args{
				url: badURL,
			},
			wantErr: true,
		},
		{
			name:    "downloaded video",
			fields:  fields{
				downloader: goodDownloader,
			},
			args:    args{
				url: goodURL,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				downloader: tt.fields.downloader,
			}
			if err := a.Run(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
