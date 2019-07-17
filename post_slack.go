package main

import (
  "github.com/nlopes/slack"
  "./scraping"
  "./key"
  "fmt"
  "time"
)

func main() {
  // scraping the numbers of followings and followers
  following, followers := scraping.Scraping()
  // make message
  number_info := fmt.Sprintf("フォロー：%s、フォロワー：%s", following, followers)
  // make date info
  date := time.Now()
  date_info := fmt.Sprintf("%d年%d月%d日の情報",date.Year(),int(date.Month()),int(date.Day()))
  text_info := fmt.Sprintf("%s\n%s\n", date_info, number_info)

  api := slack.New(key.API_KEY())
  params := slack.PostMessageParameters{
    Username: "twi_bot",
  }

  api.PostMessages("twitter", text_info, params)
}
