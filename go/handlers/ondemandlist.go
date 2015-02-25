package handlers

import (
	"bytes"
	"fmt"
	"github.com/jamesyong/o3erp/go/models"
	"github.com/jamesyong/o3erp/go/sessions"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func OnDemandListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	name := ps.ByName("name")

	session, _ := sessions.SessionStore.Get(r, "session-name")
	userLoginId := session.Values[sessions.USER_LOGIN_ID]
	if userLoginId != nil {

		query := r.URL.Query()["name"]

		from, to := getRangeFromReq(r)

		queryMap := make(map[string]string)
		if query != nil && query[0] != "" {
			queryMap["query"] = strings.Replace(query[0], "*", "%", -1)
		} else {
			queryMap["query"] = ""
		}

		var buffer bytes.Buffer
		buffer.WriteString("[")

		var total int

		switch name {
		case "product":

			products, _total := models.GetListProduct(queryMap, from, to)
			total = _total

			size := len(products)
			to = from + (size - 1)

			for index, product := range products {
				buffer.WriteString(fmt.Sprintf(`{id:"%s",name:"%s"}`, product.ProductId, product.ProductName))
				if index+1 < size {
					buffer.WriteString(",")
				}
			}

		}
		buffer.WriteString("]")

		w.Header().Set("Content-Range", fmt.Sprintf("items %d-%d/%d", from, to, total))
		w.Write([]byte(buffer.String()))

	}
}
