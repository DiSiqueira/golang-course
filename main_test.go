package main

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPrintWithSuffix(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrinter := NewMockPrinter(ctrl)

	mockPrinter.EXPECT().Print("suffix:hello!").Return()
	PrintWithSuffix(mockPrinter, "hello!")

	mockPrinter.EXPECT().Print("suffix:test").Return()
	PrintWithSuffix(mockPrinter, "tes")
}

