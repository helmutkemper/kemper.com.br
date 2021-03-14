package endpoint

type Restful struct {
	Meta   Metadata    `json:"meta"`
	Object interface{} `json:"data"`
}
