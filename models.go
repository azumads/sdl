package sdl

type Auth struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type Session struct {
	Id             string `json:"sessionId"`
	LastUpdateTime int    `json:"lastUpdateTime"`
}

type File struct {
	FullName string `json:"fullName"`
}

type Task struct {
	Id           int    `json:"id"`
	TargetLocale Locale `json:"targetLocale"`
	Status       Status `json:"status"`
}
type Project struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Tasks []*Task `json:"tasks"`
}

type ProjectGroupInput struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	ProjectTypeId int          `json:"projectTypeId"`
	ClientId      int          `json:"clientId"`
	SystemFiles   []string     `json:"systemFiles"`
	Attributes    []*Attribute `json:"attributes,omitempty"`
	Locales       []*Locale    `json:"locales"`
}

type CreateProjectGroupResponse struct {
	Status   string `json:"status"`
	Response int    `json:"response"`
}

type CreateProjectGroupResult struct {
	Status   string                        `json:"status"`
	Response []*CreateProjectGroupResponse `json:"response"`
}

type AttributeType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Attribute struct {
	Attribute *AttributeType `json:"attribute"`
	Value     interface{}    `json:"value"`
}

type ProjectGroup struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Projects []*Project `json:"projects"`
}

type Locale struct {
	Id       int       `json:"id"`
	Name     string    `json:"name,omitempty"`
	Language *Language `json:"language,omitempty"`
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

func (e Error) Error() string {
	return e.Message
}

type Errors struct {
	Errors []*Error `json:"errors"`
}
