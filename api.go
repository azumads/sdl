package sdl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/sendgrid/rest"
)

const host = "https://asics.sdlproducts.com/ws-api/v1/"

var session *Session = &Session{}

var Config = struct {
	UserName string
	PassWord string
}{}

func API(endpoint string, method rest.Method, headers, queryParams map[string]string, body []byte) (res *rest.Response, err error) {
	baseURL := host + endpoint
	request := rest.Request{
		BaseURL:     baseURL,
		Method:      method,
		QueryParams: queryParams,
		Headers:     headers,
		Body:        body,
	}
	res, err = rest.API(request)
	return
}

func Login() *Session {
	if session != nil && session.Id != "" && (int(time.Now().Unix())-session.LastUpdateTime/1000)/60 < 60 {
		return session
	}
	endpoint := "/login"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	// auth := Auth{UserName: "SDK_TEST", PassWord: "p@ssw0rd"}
	auth := Auth{UserName: Config.UserName, PassWord: Config.PassWord}
	body, _ := ffjson.Marshal(&auth)
	response, err := API(endpoint, rest.Post, headers, nil, body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return nil
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), session)
	fmt.Println(session.Id)
	return session
}

func GetProjectGroup(id int) (projectGroup *ProjectGroup) {
	token := Login()
	endpoint := "/projectGroups/" + strconv.Itoa(id)
	queryParams := make(map[string]string)
	queryParams["token"] = token.Id
	response, err := API(endpoint, rest.Get, nil, queryParams, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &projectGroup)
	fmt.Printf("%+v\n", projectGroup)
	return
}

func GetProject(id int) (project *Project) {
	token := Login()
	endpoint := "/projects/" + strconv.Itoa(id)
	queryParams := make(map[string]string)
	queryParams["token"] = token.Id
	response, err := API(endpoint, rest.Get, nil, queryParams, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &project)
	fmt.Printf("%+v\n", project)
	return
}

func GetTask(id int) (task *Task) {
	token := Login()
	endpoint := "/tasks/" + strconv.Itoa(id)
	queryParams := make(map[string]string)
	queryParams["token"] = token.Id
	response, err := API(endpoint, rest.Get, nil, queryParams, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &task)
	fmt.Printf("%+v\n", task)
	return
}
