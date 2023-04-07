package helper

type Track struct {
	ID    string `json:"Id"`
	Audio string `json:"Audio"`
}
type Tracks []Track

type Response struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

type Result struct {
	Title string `json:"title"`
}
