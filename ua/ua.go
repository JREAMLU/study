package main

import (
	"fmt"

	"github.com/mssola/user_agent"
)

func main() {
	// u_a := "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0"
	// u_a := "Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5"
	// u_a := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
	// u_a := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	u_a := "Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, like Gecko) Version/6.0.0.337 Mobile Safari/534.1+"

	ua := user_agent.New(u_a)

	fmt.Printf("moble: %v\n", ua.Mobile())
	fmt.Printf("bot: %v\n", ua.Bot())
	fmt.Printf("mozilla: %v\n", ua.Mozilla())

	fmt.Printf("platform: %v\n", ua.Platform())
	fmt.Printf("os: %v\n", ua.OS())

	fmt.Printf("Localization: %v\n", ua.Localization())

	name, version := ua.Engine()
	fmt.Printf("engine name: %v\n", name)
	fmt.Printf("engine version: %v\n", version)

	name, version = ua.Browser()
	fmt.Printf("browser name: %v\n", name)
	fmt.Printf("browser version: %v\n", version)
}
