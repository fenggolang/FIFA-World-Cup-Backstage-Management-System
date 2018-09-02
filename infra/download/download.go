package download

import (
	"FIFA-World-Cup-Backstage-Management-System/infra/config"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"net/http"
)

var (
	ErrDownloader      = errors.New("download html failed")
	ErrSeleniumService = errors.New("selenium service failed")
	ErrWebDriver       = errors.New("web driver failed")
	ErrWebDriverGet    = errors.New("web driver get url failed")
)

// 下载器，传入一个url地址
// 使用内置的net/http获取网页源代码
func Downloader(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, ErrDownloader
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	//client:=http.DefaultClient
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrDownloader
	}
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}

// 如果web页面是动态加载数据，不想费劲分析网页，对速度要求也不高，可以使用selenium
func DownloaderBySelenium(url string) (string, error) {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	imageCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imageCaps,
		Path:  "",
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)

	service, err := selenium.NewChromeDriverService(
		config.ChromeDriverPath, 9515,
	)
	defer service.Stop()

	if err != nil {
		fmt.Println(ErrSeleniumService)
		return "", ErrSeleniumService
	}
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))

	if err != nil {
		fmt.Println(ErrWebDriver)
		return "", ErrWebDriver
	}

	err = webDriver.Get(url)

	if err != nil {
		fmt.Println(ErrWebDriverGet)
		return "", ErrWebDriverGet
	}
	return webDriver.PageSource()
}
