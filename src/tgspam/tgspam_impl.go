package tgspam

import (
	"github.com/go-vgo/robotgo"
	"time"
)

func OpenTelegram() error {
	//set default sleep time
	robotgo.MouseSleep = 1000
	//Open telegram
	robotgo.Move(0, 540)
	robotgo.Click()
	return nil
}
func SendMessage(username, message string) error {
	//set default sleep time
	robotgo.KeySleep = 100
	robotgo.MouseSleep = 100

	//Open saved messages
	robotgo.Move(100, 160)
	robotgo.Click()

	//Send username to the saved messages
	robotgo.Move(960, 1060)
	robotgo.Click()
	robotgo.TypeStr(username)
	if err := robotgo.KeyTap("enter"); err != nil {
		return err
	}

	//Open user profile
	robotgo.Move(760, 1010)
	robotgo.Click()

	//Find send message button
	robotgo.MouseSleep = 1
	for i := 0; i < 500; i++ {
		robotgo.Move(900, 800-i)
		color := robotgo.GetPixelColor(900, 800-i)
		if color == "4c7aa5" {
			time.Sleep(1 * time.Second)
			robotgo.Click()
			break
		}
	}
	robotgo.MouseSleep = 100
	//Send the message
	time.Sleep(1 * time.Second)
	robotgo.TypeStr(message)
	if err := robotgo.KeyTap("enter"); err != nil {
		return err
	}
	return nil
}

func ChangeAccount(accountNumber int) error {
	//set default sleep time
	robotgo.MouseSleep = 1000

	//Change account
	robotgo.Move(30, 80)
	robotgo.Click()
	robotgo.Move(80, 170+40*accountNumber)
	robotgo.Click()
	return nil
}
