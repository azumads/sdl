package main

import "github.com/azumads/sdl"

func main() {
	sdl.Config.UserName = "SDK_TEST"
	sdl.Config.PassWord = "p@ssw0rd"

	// file, err := sdl.UploadFile("Test.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// locales := []*sdl.Locale{&sdl.Locale{Id: 1018}, &sdl.Locale{Id: 1026}}
	// attTpe := sdl.AttributeType{
	// 	Id:   1018,
	// 	Name: "testURL",
	// 	Type: "URL",
	// }
	// att := sdl.Attribute{
	// 	Attribute: &attTpe,
	// 	Value:     "https://asicstiger-draft.theplant-dev.com/us/en-us/ronniefiegdev context",
	// }

	// sdl.CreateProjectGroup("Test2", "It's a go versoin test", 1039, 1120, []string{"upload_4438840343460206409.Test.html"}, locales, &att)
	// sdl.GetProjectGroup(1094)
	// sdl.GetProject(1120)
	// sdl.GetProject(1121)
	// sdl.GetTask(1125)
	// sdl.GetTask(1126)
	sdl.DownloadTaskFile(1157)
}
