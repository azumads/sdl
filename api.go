package sdl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/sendgrid/rest"
)

const host = "https://asics.sdlproducts.com/ws-api/v1"

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
	if session.Id != "" && (int(time.Now().Unix())-session.LastUpdateTime/1000)/60 < 60 {
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
	if response.StatusCode >= 300 {
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

func UploadFile(path string) (file *File, err error) {
	token := Login()
	baseURL := host + "/files"
	queryParams := make(map[string]string)
	queryParams["token"] = token.Id
	baseURL = rest.AddQueryParameters(baseURL, queryParams)
	request, err := NewfileUploadRequest(baseURL, nil, "file", path)
	if err != nil {
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	response.Body.Close()
	if response.StatusCode >= 300 {
		var errs Errors
		ffjson.Unmarshal(body, &errs)
		fmt.Printf("%+v", errs.Errors)
		if len(errs.Errors) > 0 {
			err = errs.Errors[0]
		}
		return
	}
	ffjson.Unmarshal(body, &file)
	fmt.Printf("%+v\n", file)
	return
}

func CreateProjectGroup(name, desc string, projectTypeId, clientId int, files []string, locales []*Locale, att *Attribute) (result *CreateProjectGroupResult, err error) {
	token := Login()
	endpoint := "/projectGroups/create"
	queryParams := make(map[string]string)
	queryParams["token"] = token.Id
	input := ProjectGroupInput{Name: name,
		Description:   desc,
		ProjectTypeId: projectTypeId,
		ClientId:      clientId,
		SystemFiles:   files,
		Attributes:    []*Attribute{att},
		Locales:       locales,
	}
	body, _ := ffjson.Marshal([]*ProjectGroupInput{&input})
	fmt.Printf("%s", body)
	response, err := API(endpoint, rest.Post, nil, queryParams, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		if len(errs.Errors) > 0 {
			err = errs.Errors[0]
		}
		return nil, err
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &result)
	fmt.Printf("%+v\n", result.Response[0])
	return
}

func GetProjectGroup(id int) (projectGroup *ProjectGroup, err error) {
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
	if response.StatusCode >= 300 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		if len(errs.Errors) > 0 {
			err = errs.Errors[0]
		}
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &projectGroup)
	fmt.Printf("%+v\n", projectGroup)
	return
}

func GetProject(id int) (project *Project, err error) {
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
	if response.StatusCode >= 300 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		if len(errs.Errors) > 0 {
			err = errs.Errors[0]
		}
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &project)
	fmt.Printf("%+v\n", project)
	return
}

func GetTask(id int) (task *Task, err error) {
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
	if response.StatusCode >= 300 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		if len(errs.Errors) > 0 {
			err = errs.Errors[0]
		}
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &task)
	fmt.Printf("%+v\n", task)
	return
}
