// nolint
package main

import (
	"fmt"
	"os"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/info"
	"github.com/LagrangeDev/LagrangeGo/message"
)

func main() {
	appInfo := info.AppList["linux"]
	deviceInfo := &info.DeviceInfo{
		Guid:          "cfcd208495d565ef66e7dff9f98764da",
		DeviceName:    "Lagrange-DCFCD07E",
		SystemKernel:  "Windows 10.0.22631",
		KernelVersion: "10.0.22631",
	}
	sig := info.LoadSig("./sig.bin")
	qqclient := client.NewQQclient(0, "https://sign.lagrangecore.org/api/sign", appInfo, deviceInfo, sig)

	qqclient.GroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		if event.ToString() == "114514" {
			img, _ := os.ReadFile("./testgroup.png")
			_, err := client.SendGroupMessage(event.GroupCode, []message.IMessageElement{&message.GroupImageElement{Stream: img}})
			if err != nil {
				return
			}
		}
	})

	qqclient.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		img, _ := os.ReadFile("testprivate.png")
		_, err := client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{&message.FriendImageElement{Stream: img}})
		if err != nil {
			return
		}
	})

	err := qqclient.Loop()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = qqclient.Login("", "./qrcode.png")
	if err != nil {
		fmt.Println(err)
	}
	info.SaveSig(sig, "./sig.bin")

	select {}
}
