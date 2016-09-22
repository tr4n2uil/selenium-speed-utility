// Run speed test on various possible hosts port and provide results
package main

import (
  "fmt"
  "time"
  "os"

  "github.com/tebeka/selenium"
)

func run_test(remote string) {
  caps := selenium.Capabilities{
    "browserName": "chrome",
    "os": "OSX",
    "os_version": "Yosemite",
    "build": "selenium speed test",
    "name": remote,
    "browserstack.user": os.Getenv("BROWSERSTACK_USERNAME"),
    "browserstack.key": os.Getenv("BROWSERSTACK_ACCESS_KEY"),
  }

  wd, err := selenium.NewRemote(caps, remote + "/wd/hub")
  if err != nil {
    panic(err)
  }
  defer wd.Quit()

  wd.Get("https://www.google.com/ncr")
  elem, _ := wd.FindElement(selenium.ByName, "q")
  elem.Clear()
  start := time.Now()
  for j := 0; j < 100; j++ {
    elem.SendKeys("a")
  }
  diff := time.Since(start)
  
  elem.Submit()
  wd.Title()

  wd.Quit()
  fmt.Printf("%s %s\n", remote, diff)
}

func main() {
  remotes := [20]string{
    "http://hub.browserstack.com",
    "https://hub.browserstack.com",
    "http://hub.browserstack.com:4444",
    "http://hub-cloud.browserstack.com",
    "https://hub-cloud.browserstack.com",

    "http://hub-eu.browserstack.com",
    "https://hub-eu.browserstack.com",
    "http://hub-eu.browserstack.com:4444",
    "http://hub-cloud-eu.browserstack.com",
    "https://hub-cloud-eu.browserstack.com",

    "http://hub-us.browserstack.com",
    "https://hub-us.browserstack.com",
    "http://hub-us.browserstack.com:4444",
    "http://hub-cloud-us.browserstack.com",
    "https://hub-cloud-us.browserstack.com",

    "http://hub-usw.browserstack.com",
    "https://hub-usw.browserstack.com",
    "http://hub-usw.browserstack.com:4444",
    "http://hub-cloud-usw.browserstack.com",
    "https://hub-cloud-usw.browserstack.com",
  }

  for i := 0; i < len(remotes); i++ {
    run_test(remotes[i])
  }
}
