package download

import (
	"fmt"
	"testing"
)

func TestDownloader(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url: "https://www.fifa.com/worlcup/matches/?#groupphase"},     // 小组赛
		{url: "https://www.fifa.com/worldcup/matches/?#knockoutphase"}, // 淘汰赛
	}
	for _, test := range tests {
		doc, err := Downloader(test.url)
		newDoc, _ := doc.Html()
		fmt.Println(newDoc, err)
	}
}

/*func TestDownloaderBySelenium(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url: "https://www.fifa.com/worlcup/matches/?#groupphase"},     // 小组赛
		{url: "https://www.fifa.com/worldcup/matches/?#knockoutphase"}, // 淘汰赛
	}
	for _, test := range tests {
		fmt.Println(DownloaderBySelenium(test.url))
	}
}
*/
