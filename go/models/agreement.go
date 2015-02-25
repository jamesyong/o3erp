// agreement
package models

import (
	"database/sql"
)

type Agreement struct {
	AgreementId  []byte
	ProductId    []byte
	PartyIdFrom  []byte
	PartyIdTo    []byte
	RoleTypeIdTo []byte
	FromDate     []byte
	ThruDate     []byte
	Description  []byte
}

func GetTableAgreement(queryMap map[string]string, offset int, limit int) ([]*Agreement, int) {
	var total int
	agreements := []*Agreement{}

	dbCountFunc := func(queryMap *map[string]string) error {
		var err error
		if (*queryMap)["query"] == "" {
			err = DbMap.QueryRow("select count(agreement_id) from AGREEMENT").Scan(&total)
		} else {
			err = DbMap.QueryRow("select count(agreement_id) from AGREEMENT where agreement_id like ?", (*queryMap)["query"]).Scan(&total)
		}
		return err
	}

	dbRowsFunc := func(queryMap *map[string]string) (*sql.Rows, error) {
		var err error
		var rows *sql.Rows
		if (*queryMap)["query"] == "" {
			rows, err = DbMap.Query("select agreement_id, product_id, party_id_from, party_id_to, role_type_id_to, from_date, thru_date, description from AGREEMENT limit ?, ?", offset, limit)
		} else {
			rows, err = DbMap.Query("select agreement_id, product_id, party_id_from, party_id_to, role_type_id_to, from_date, thru_date, description from AGREEMENT where agreement_id like ? limit ?, ?", (*queryMap)["query"], offset, limit)
		}
		return rows, err
	}

	dbListFunc := func(rows *sql.Rows) error {
		var err error
		var p Agreement
		err = rows.Scan(&p.AgreementId, &p.ProductId, &p.PartyIdFrom, &p.PartyIdTo, &p.RoleTypeIdTo, &p.FromDate, &p.ThruDate, &p.Description)
		if err == nil {
			agreements = append(agreements, &p)
		}
		return err
	}

	GetDbListTemplate(queryMap, dbCountFunc, dbRowsFunc, dbListFunc)

	return agreements, total
}
