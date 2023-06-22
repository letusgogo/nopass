package ui

var DefaultWindow *Window

func init() {
	DefaultWindow = NewWindow()
}

func Run() {
	DefaultWindow.Run()
}
