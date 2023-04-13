package response

type Response struct {
	Data   interface{} `json:"data"`
	Result result      `json:"result"`
}

type result struct {
	ResultCode  string `json:"resultCode"`
	ResultError string `json:"resultError"`
}

type Page struct {
	data       interface{}
	pageNum    int32
	pageSize   int32
	totalCount int64
	totalPage  int32
}
