package page

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/plant99/telegraphcl/pkg/user"
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
	fmt.Println("Index | URL | Title")
	fmt.Println("--------------------")
	for i := 0; i < len(pageList.Pages); i++ {
		fmt.Println(i, ")", pageList.Pages[i].URL, "|", pageList.Pages[i].Title)
	}
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

func CreatePage(path string, title string) {
	// get access_token
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println(err)
	}

	// get []Nodes from markdown file in path
	nodes, err := MarkdownFileToNodes(path)

	if err != nil {
		panic(err)
	}

	userInfo := user.GetCurrentUserNameAndURL()

	createPageRequestInstance := createPageRequest{
		Title:         title,
		AuthorName:    userInfo[0],
		AuthorUrl:     userInfo[1],
		Content:       nodes,
		ReturnContent: true,
		AccessToken:   accessToken,
	}

	createPageResponseInstance := new(Page)
	// get total views on page
	data, err := util.MakeRequest("createPage", createPageRequestInstance)
	parser := jsoniter.ConfigFastest
	if err = parser.Unmarshal(data, &createPageResponseInstance); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response. Is the Telegra.ph path correct?", err)
	}
	fmt.Println(createPageResponseInstance.URL)

}

func GetPage(path string) {

	requestGetPage := map[string]string{
		"path": path,
	}

	responseGetPage := new(Page)
	// get total views on page
	data, err := util.MakeRequest("getPage", requestGetPage)
	parser := jsoniter.ConfigFastest
	if err = parser.Unmarshal(data, &responseGetPage); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response. Is the Telegra.ph path correct?", err)
	}
	fmt.Println(responseGetPage.Title, responseGetPage.URL)
}

func EditPage(path string) {
	// get access_token
	accessToken, err := util.FetchAccessToken()
	if err != nil {
		fmt.Println(err)
	}

	editPageRequestInstance := editPageRequest{
		Title:         "Some dummy title",
		AuthorName:    "Some dummy name edited",
		AuthorUrl:     "https://one.two.three",
		Content:       []Node{"some dummy content but edited"},
		ReturnContent: true,
		AccessToken:   accessToken,
		Path:          path,
	}

	editPageResponseInstance := new(Page)
	// get total views on page
	data, err := util.MakeRequest("editPage", editPageRequestInstance)
	parser := jsoniter.ConfigFastest
	if err = parser.Unmarshal(data, &editPageResponseInstance); err != nil {
		fmt.Println("Couldn't handle api.telegra.ph response. Is the Telegra.ph path correct?", err)
	}
	fmt.Println(editPageResponseInstance.Title)

}
