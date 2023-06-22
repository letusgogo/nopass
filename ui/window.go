package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Window struct {
	ui fyne.App
}

func NewWindow() *Window {
	return &Window{ui: app.New()}
}

func (w *Window) Run() {
	w.ui.Run()
}

func (w *Window) PrintPass(pass string) {
	myWindow := w.ui.NewWindow("Pass")
	copyableText := widget.NewEntry()
	copyableText.SetText(pass)

	closeTime := 8
	countdownDuration := time.Duration(closeTime) * time.Second
	countdownLabel := widget.NewLabel(fmt.Sprintf("Time remaining: %v", countdownDuration))

	// 添加复制按钮
	copyButton := widget.NewButton("Copy", func() {
		clipboard := myWindow.Clipboard()
		clipboard.SetContent(copyableText.Text)
		myWindow.Close()
	})

	myWindow.SetContent(
		container.NewVBox(
			widget.NewLabel("Please copy pass and exit:"),
			copyableText,
			countdownLabel,
			copyButton,
		),
	)

	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.Show()
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.NewTimer(countdownDuration)

	// 更新倒计时标签
	go func() {
		for {
			select {
			case <-ticker.C:
				countdownDuration -= 1 * time.Second
				countdownLabel.SetText(fmt.Sprintf("Time remaining: %v", countdownDuration))
			case <-timeout.C:
				ticker.Stop()
				myWindow.Close()
				return
			}
		}
	}()
}

//func (w *Window) ShowConfig(context string) (updateContent string, err error) {
//
//}
