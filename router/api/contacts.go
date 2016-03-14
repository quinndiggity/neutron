package api

import (
	"gopkg.in/macaron.v1"

	"github.com/emersion/neutron/backend"
)

type ContactsResp struct {
	Resp
	Contacts []*backend.Contact
}

func GetContacts(ctx *macaron.Context) (err error) {
	contacts, err := backend.GetContacts(userId)
	if err != nil {
		return
	}

	ctx.JSON(200, &ContactsResp{
		Resp: Resp{1000},
		Contacts: contacts,
	})
	return
}