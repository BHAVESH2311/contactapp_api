package contactservice

import (
	"errors"
	"github.com/google/uuid"
	contact_details_service "contactapp_api/components/contact_details/contact_details_service"
)

type Contact struct {
	Id             string
	FirstName      string
	LastName       string
	IsActive       bool
	ContactDetails []*contact_details_service.ContactDetail
}

var Contacts []*Contact



func NewContact(Fname, LName string) *Contact {

	contactId := uuid.NewString()
	newContact := &Contact{
		Id:        contactId,
		FirstName: Fname,
		LastName:  LName,
		IsActive:  true,
	}

	Contacts = append(Contacts, newContact)
	return newContact
}

func UpdateContact(body *Contact, contact *Contact) {
	if body.FirstName != "" && body.FirstName != contact.FirstName {
		contact.FirstName = body.FirstName
	}

	if body.LastName != "" && body.LastName != contact.LastName {
		contact.LastName = body.LastName
	}
}

func GetContactById(contactId string) (*Contact, error) {
	for _, contact := range Contacts {
		if contact.Id == contactId {
			return contact, nil
		}
	}

	return nil, errors.New("Contact Does Not Exist")
}

func DeleteContact(contact *Contact) {
	contact.IsActive = false
}

func AppendContactDetail(contact *Contact, contactDetail *contact_details_service.ContactDetail) {
	contact.ContactDetails = append(contact.ContactDetails, contactDetail)
}