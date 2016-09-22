package sdl

type Auth struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Session struct {
	Id             string `json:"sessionId"`
	LastUpdateTime int    `json:"lastUpdateTime"`
}

type Task struct {
	Id           int    `json:"id"`
	TargetLocale Locale `json:"targetLocale"`
	Status       Status `json:"status"`
}
type Project struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type ProjectGroup struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

type Locale struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Language struct {
	Id           int    `json:"id"`
	LanguageCode string `json:"languageCode"`
	CountryCode  string `json:"countryCode"`
}

type Status struct {
	Status      string `json:"status"`
	DisplayText string `json:"displayText"`
}
type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
type Errors struct {
	Errors []Error `json:"errors"`
}
