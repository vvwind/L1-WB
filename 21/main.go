package main

import "fmt"

type Server interface {
	Call(url string) string
}

type Client struct{}

func (v *Client) Call(o Server, url string) string {
	return o.Call(url)
}

type Mozilla struct{}

func (v *Mozilla) CallA(url string) string {
	return fmt.Sprintf("You called %v through Mozilla", url)
}

type AdapterMozilla struct {
	Mozilla
}

func (v AdapterMozilla) Call(url string) string {
	return v.CallA(url)
}

type Chrome struct{}

func (v *Chrome) CallB(url string) string {
	return fmt.Sprintf("You called %v through Chrome", url)
}

type AdapterChrome struct {
	Chrome
}

func (v AdapterChrome) Call(url string) string {
	return v.CallB(url)
}

func main() {
	client := Client{}

	mz := Mozilla{}
	fmt.Println(client.Call(AdapterMozilla{mz}, "wildberries.com"))

	ch := Chrome{}
	fmt.Println(client.Call(AdapterChrome{ch}, "wildberries.com"))
}
