// Run speed test on various possible hosts port and provide results
package main

import (
  "fmt"
  "time"
  "os"
  "flag"
  "sort"

  "github.com/tebeka/selenium"
)

type Pair struct {
  Key string
  Value time.Duration
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

func run_test(remote string) time.Duration{
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
  for j := 0; j < 10; j++ {
    elem.SendKeys("a")
  }
  diff := time.Since(start)
  
  wd.Quit()
  return diff
}

func main() {
  var profile = flag.String("region", "default", "region to run test [us, usw, eu] \n\toptional, if not passed will route to nearest")
  flag.Parse()

  remotes := make(map[string][]string)

  remotes["default"] = []string{
    "http://hub.browserstack.com",
    "https://hub.browserstack.com",
    "http://hub.browserstack.com:4444",
    "http://hub-cloud.browserstack.com",
    "https://hub-cloud.browserstack.com",
  }

  remotes["us"] = []string{
    "http://hub-us.browserstack.com",
    "https://hub-us.browserstack.com",
    "http://hub-us.browserstack.com:4444",
    "http://hub-cloud-us.browserstack.com",
    "https://hub-cloud-us.browserstack.com",
  }

  remotes["usw"] = []string{
    "http://hub-usw.browserstack.com",
    "https://hub-usw.browserstack.com",
    "http://hub-usw.browserstack.com:4444",
    "http://hub-cloud-usw.browserstack.com",
    "https://hub-cloud-usw.browserstack.com",
  }

  remotes["eu"] = []string{
    "http://hub-eu.browserstack.com",
    "https://hub-eu.browserstack.com",
    "http://hub-eu.browserstack.com:4444",
    "http://hub-cloud-eu.browserstack.com",
    "https://hub-cloud-eu.browserstack.com",
  }

  rl := len(remotes[*profile])
  pl := make(PairList, rl)

  fmt.Printf("%d/%d Completed\n", 0, rl)
  for i := 0; i < rl; i++ {
    pl[i] = Pair{remotes[*profile][i], run_test(remotes[*profile][i])}
    fmt.Printf("%d/%d Completed\n", i+1, rl)
  }

  fmt.Println("\n\nResults:\n")
  sort.Sort(pl)
  for i := 0; i < rl; i++ {
    fmt.Printf("%s %s\n", pl[i].Key, pl[i].Value)
  }
}
