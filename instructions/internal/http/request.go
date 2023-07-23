package http

type ExecuteRequest struct {
	Param string `json:"param"`
	Path  string `json:"path"`
}
