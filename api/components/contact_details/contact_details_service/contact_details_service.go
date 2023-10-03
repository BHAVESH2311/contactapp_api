package contactdetailsservice

import (
	"github.com/google/uuid"
	"errors"
)

type ContactDetail struct {
	Id           string
	ContactType  string
	ContactValue string
	IsActive     bool
}

var contactDetails = []*ContactDetail{}

func CreateContactDetail(contactType, contactValue string) *ContactDetail {
	contactDetailsId := uuid.NewString()
	
	newContactDetail := &ContactDetail{
		Id:           contactDetailsId,
		ContactType:  contactType,
		ContactValue: contactValue,
		IsActive:     true,
	}

	contactDetails = append(contactDetails, newContactDetail)
	return newContactDetail
}

func GetContactDetailById(contactDetailId string) (*ContactDetail, error) {
	for _, contactDetail := range contactDetails {
		if contactDetail.Id == contactDetailId {
			return contactDetail, nil
		}
	}

	return nil, errors.New("contact detail does not exist")
}

func UpdateContactDetail(body *ContactDetail, contactDetail *ContactDetail) {
	if body.ContactType != "" && body.ContactType != contactDetail.ContactType {
		contactDetail.ContactType = body.ContactType
	}

	if body.ContactValue != "" && body.ContactValue != contactDetail.ContactValue {
		contactDetail.ContactValue = body.ContactValue
	}
}

func DeleteContactDetail(contactDetail *ContactDetail) {
	contactDetail.IsActive = false
}