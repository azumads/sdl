package sdl

import (
	"fmt"
	"strconv"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/sendgrid/rest"
)

const host = "https://asics.sdlproducts.com/ws-api/v1/"

var token string

func init() {
	token = Login().Id
}

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

func Login() (session Session) {
	endpoint := "/login"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	auth := Auth{UserName: "SDK_TEST", PassWord: "p@ssw0rd"}
	body, _ := ffjson.Marshal(&auth)
	// fmt.Printf("%s", body)
	// body := []byte(`{"username":"SDK_TEST","password":"p@ssw0rd"}`)
	response, err := API(endpoint, rest.Post, headers, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &session)
	fmt.Println(session.Id)
	return
}

func GetTask(id int) (task *Task) {
	endpoint := "/tasks/" + strconv.Itoa(id)
	queryParams := make(map[string]string)
	queryParams["token"] = token
	response, err := API(endpoint, rest.Get, nil, queryParams, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 400 {
		var errs Errors
		ffjson.Unmarshal([]byte(response.Body), &errs)
		fmt.Printf("%+v", errs.Errors)
		return
	}
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	ffjson.Unmarshal([]byte(response.Body), &task)
	fmt.Printf("%+v", task)
	return
}
