package main

import (
	"fmt"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

// refer：go-fyne-webserver

const (
	startBKText = "Start BK" // start button 开始按钮
	stopBKText  = "Stop BK"  // stop button 停止按钮
)

var BKStatus int

const (
	BKStatusRunning = 1 // status running 运行中
	BKStatusStop    = 0 // status stopped 停止
)

func main() {
	myApp := app.New()
	r, _ := fyne.LoadResourceFromPath("./Icon.png")
	myApp.SetIcon(r)
	myWindow := myApp.NewWindow("BluetoothKeeper") // set title

	// set app icon
	myWindow.SetIcon(r)

	startChan := make(chan bool, 1) // send start command
	stopChan := make(chan bool, 1)  // send start command
	go myFunc(startChan, stopChan)  // run your job

	var button *widget.Button // init buttons
	button = widget.NewButtonWithIcon(startBKText, theme.MediaPlayIcon(), func() {
		if BKStatus == BKStatusStop {
			fmt.Println("Press Start")
			button.SetText(stopBKText)
			button.SetIcon(theme.MediaPauseIcon())
			startChan <- true
			BKStatus = BKStatusRunning
		} else {
			fmt.Println("Press Stop")
			button.SetText(startBKText)
			button.SetIcon(theme.MediaPlayIcon())
			stopChan <- true
			BKStatus = BKStatusStop
		}
	})
	buttonBox := container.NewHBox(button)
	content := container.NewVBox(
		buttonBox,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})

	myWindow.ShowAndRun()

}

func myFunc(start, stop chan bool) {
	done := make(chan bool, 1)
	for {
		running := false
		select {
		case <-stop:
			running = false
			done <- true
			continue
		case <-start:
			if running == false {
				running = true
				go func() {
					for {
						select {
						case <-done:
							return
						default:
							doTask()
							fmt.Println("do task...")
						}
						time.Sleep(5 * time.Second)
					}
				}()
			}
			continue
		}
	}
}

func doTask() {
	if IsScreenClose() {
		screenSleep()
	} else {
		screenOn()
	}
	fmt.Println("my Func is running, don't worry!")
}

//func envCheck() {
//	if runtime.GOOS != "darwin" {
//		panic("This platform is not supported")
//	}
//}

func IsScreenClose() bool {
	return screenshot.NumActiveDisplays() == 0
}

// 处理关闭盖子事件
func screenSleep() {
	// 关闭蓝牙
	cmd := exec.Command("blueutil", "--power", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to turn off bluetooth:", err)
	}

	// 打印日志
	fmt.Println("Screen is now sleeping")
}

// 处理关闭盖子事件
func screenOn() {
	// 关闭蓝牙
	cmd := exec.Command("blueutil", "--power", "1")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to turn on bluetooth:", err)
	}

	// 打印日志
	fmt.Println("Screen is not sleeping now")
}
