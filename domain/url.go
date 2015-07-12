package domain

import (
	"github.com/solher/zest/domain"
)

func init() {
	relations := []domain.DBRelation{
		{Related: "accounts", Fk: "accountId"},
	}

	domain.ModelDirectory.Register(Url{}, "urls", relations)
}

type Url struct {
	domain.GormModel
	ShortHandle string         `json:"shortHandle" sql:"unique"`
	LongUrl     string         `json:"longUrl"`
	Notes       string         `json:"notes"`
	Enabled     bool           `json:"enabled"`
	AccountID   int            `json:"accountId" sql:"index"`
	Account     domain.Account `json:"account,omitempty"`
}

func (m *Url) SetRelatedID(idKey string, id int) {
	switch idKey {
	case "accountID":
		m.AccountID = id
	}
}

func (m *Url) BeforeRender() {
	m.Account.BeforeRender()
}
