package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct {}

func (w *Windows) InsertIntoUsbPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adaptor converts Lightning signal to USB.")
	w.windowsMachine.InsertIntoUsbPort()
}

type Client struct {}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightning connector into computer")
	com.InsertIntoLightningPort()
}

func main() {
	client := &Client{}

	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}