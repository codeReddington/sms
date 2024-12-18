package structs

type TasksResponse struct {
	HostHeaderInfo HostHeaderInfo `json:"hostHeaderInfo"`
	Tasks          []Tasks        `json:"tasks"`
}

type TaskResponse struct {
	HostHeaderInfo HostHeaderInfo `json:"hostHeaderInfo"`
	Tasks          Tasks          `json:"tasks"`
}

type HostHeaderInfo struct {
	Channel         string `json:"channel"`
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type Tasks struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
