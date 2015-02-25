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

func OnDemandTableHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

			products, _total := models.GetTableProduct(queryMap, from, to)
			total = _total

			size := len(products)
			to = from + (size - 1)

			for index, product := range products {
				buffer.WriteString(fmt.Sprintf(`{"productId":"%s","productName":"%s"}`, product.ProductId, product.ProductName))
				if index+1 < size {
					buffer.WriteString(",")
				}
			}
		case "agreement":
			agreements, _total := models.GetTableAgreement(queryMap, from, to)
			total = _total

			size := len(agreements)
			to = from + (size - 1)

			for index, agreement := range agreements {
				buffer.WriteString(fmt.Sprintf(
					`{"agreementId":"%s","productId":"%s","partyIdFrom":"%s","partyIdTo":"%s","roleTypeIdTo":"%s","fromDate":"%s","thruDate":"%s","description":"%s"}`,
					agreement.AgreementId, agreement.ProductId, agreement.PartyIdFrom, agreement.PartyIdTo, agreement.RoleTypeIdTo, agreement.FromDate, agreement.ThruDate, agreement.Description))
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
