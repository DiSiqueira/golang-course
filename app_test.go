package main
//
//import (
//	"errors"
//	"github.com/golang/mock/gomock"
//	"testing"
//)
//
//func TestApp_Run(t *testing.T) {
//	url := "https://www.youtube.com/watch?v=nBxnNxrAQHc"
//
//	ctrl := gomock.NewController(t)
//	goodDownloader := NewMockDownloader(ctrl)
//	goodDownloader.EXPECT().Download(url).Return(nil)
//
//	badDownloader := NewMockDownloader(ctrl)
//	badDownloader.EXPECT().Download(url).Return(errors.New("download error"))
//
//	type fields struct {
//		downloader Downloader
//	}
//	type args struct {
//		url string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "download works",
//			fields: fields{
//				downloader: goodDownloader,
//			},
//			args:    args{
//				url: url,
//			},
//			wantErr: false,
//		},
//		{
//			name: "download fails",
//			fields: fields{
//				downloader: badDownloader,
//			},
//			args:    args{
//				url: url,
//			},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			a := &App{
//				downloader: tt.fields.downloader,
//			}
//			if err := a.Run(tt.args.url); (err != nil) != tt.wantErr {
//				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
