package scraping

import (
  "github.com/PuerkitoBio/goquery"
)

func main() {
  url := "https://twitter.com/NuinkTSUKUBA"
  // get method
  doc, _ := goquery.NewDocument(url)
  // css selecter
  selecter_following := "#page-container > div.ProfileCanopy.ProfileCanopy--withNav.js-variableHeightTopBar > div > div.ProfileCanopy-navBar.u-boxShadow > div > div > div.Grid-cell.u-size2of3.u-lg-size3of4 > div > div > ul > li.ProfileNav-item.ProfileNav-item--following > a > span.ProfileNav-value"
  selecter_followers := "#page-container > div.ProfileCanopy.ProfileCanopy--withNav.js-variableHeightTopBar > div > div.ProfileCanopy-navBar.u-boxShadow > div > div > div.Grid-cell.u-size2of3.u-lg-size3of4 > div > div > ul > li.ProfileNav-item.ProfileNav-item--followers > a > span.ProfileNav-value"
  // selected texts
  following := doc.Find(selecter_following).Text()
  followers := doc.Find(selecter_followers).Text()
  return following, followers
}
