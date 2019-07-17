package main

import (
  "github.com/nlopes/slack"
  "./scraping"
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

  token := "YOUR_TOKEN"
  api := slack.New(token)
  attachment := slack.Attachment{
    Text: number_info,
  }

  channelID, timestamp, err := api.PostMessage("twitter", slack.MsgOptionText(date_info, false), slack.MsgOptionAttachments(attachment))
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
