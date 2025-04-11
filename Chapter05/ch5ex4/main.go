package main

import (
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

	// The paths to these binaries will be different on your machine!

	// const (
	// 	seleniumPath    = "c/Users/Ivan/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/selenium-server.jar"
	// 	geckoDriverPath = "с/Users/Ivan/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/geckodriver.tar.gz"
	// )

	// service, err := selenium.NewSeleniumService(
	// 	seleniumPath,
	// 	8080,
	// 	selenium.GeckoDriver(geckoDriverPath))

	// if err != nil {
	// 	panic(err)
	// }
	// defer service.Stop()

	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver/chromedriver.exe", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get("https://www.packtpub.com/networking-and-servers/mastering-go")
	if err != nil {
		panic(err)
	}

	var elems []selenium.WebElement
	wd.Wait(func(wd2 selenium.WebDriver) (bool, error) {
		elems, err = wd.FindElements(selenium.ByCSSSelector, "div.product-reviews-review div.review-body")
		if err != nil {
			return false, err
		} else {
			return len(elems) > 0, nil
		}
	})

	for _, review := range elems {
		body, err := review.Text()
		if err != nil {
			panic(err)
		}
		println(body)
	}
}
