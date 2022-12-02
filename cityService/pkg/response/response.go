package response

import (
	glog "cityService/pkg/glog"
	"encoding/json"
	"net/http"
)

var (
	log = glog.Init()
)

type Response struct {
	Status int
	Res    map[string]interface{}
	RW     http.ResponseWriter
}

func (r *Response) Err(err string, code int) {
	r.Status = code
	r.Res["err"] = err
	log.Printf("error: code: %d msg:%s", code, err)
	r.Send()
}

func (r *Response) Send() {
	r.RW.Header().Set("Content-Type", "application/json; charset=UTF-8")
	r.RW.WriteHeader(r.Status)

	jData, err := json.Marshal(r.Res)
	if err != nil {
		log.Fatalln(err)
	}
	r.RW.Write([]byte(jData))
}

func (r *Response) Init(w http.ResponseWriter) {
	r.Status = http.StatusMethodNotAllowed
	r.Res = make(map[string]interface{})
	r.RW = w
}
func (r *Response) RetId(id uint64) map[string]uint64 {
	return map[string]uint64{"id": id}
}
