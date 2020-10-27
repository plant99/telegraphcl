package page

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/plant99/telegraphcl/pkg/util"
)

func ListPages() {
	// get access_token
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println(err)
	}
	// make request parameters
	getPageListRequest := GetPageListRequest{
		AccessToken: accessToken,
		Offset:      0,
		Limit:       200,
	}
	// get page list
	data, err := util.MakeRequest("getPageList", getPageListRequest)
	if err != nil {
		fmt.Println(err)
	}

	parser := jsoniter.ConfigFastest

	pageList := new(PageList)

	if err = parser.Unmarshal(data, pageList); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response.")
	}
	fmt.Println(pageList)
}

func GetViews(path string) {
	// make request parameters
	requestPageViews := map[string]string{
		"path": path,
	}

	responsePageViews := new(PageViews)
	// get total views on page
	data, err := util.MakeRequest("getViews", requestPageViews)

	parser := jsoniter.ConfigFastest

	if err = parser.Unmarshal(data, &responsePageViews); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response. Is the Telegra.ph path correct?")
	}
	fmt.Println(responsePageViews.Views)

}
