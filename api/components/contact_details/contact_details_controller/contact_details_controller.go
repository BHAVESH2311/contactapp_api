package contactdetailscontroller

import (
	contact_details_service "contactapp_api/components/contact_details/contact_details_service"
	contact_service "contactapp_api/components/contacts/contact_service"
	user_service "contactapp_api/components/user/user_service"
	"contactapp_api/validators"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	userId := mux.Vars(r)["userId"]
	contactId := mux.Vars(r)["contactId"]

	_, err := user_service.GetUserById(userId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	contact, err := contact_service.GetContactById(contactId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Does Not Exist"))
		return
	}
	var container *contact_details_service.ContactDetail

	json.NewDecoder(r.Body).Decode(&container)

	isValid, err := validators.ValidateContactDetails(container.ContactType, container.ContactValue)
	if !isValid {
		panic("please enter valid contact details")
	}

	newContact := contact_details_service.CreateContactDetail(container.ContactType, container.ContactValue)

	contact_service.AppendContactDetail(contact, newContact)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Contact Detail Created Successfully"))
}

func GetContactDetailById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	uId := mux.Vars(r)["uId"]
	cId := mux.Vars(r)["cId"]
	cdId := mux.Vars(r)["cdId"]

	_, err := user_service.GetUserById(uId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does not Exist"))
	}

	_, err = contact_service.GetContactById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Does not Exist"))
	}

	contactDetail, err := contact_details_service.GetContactDetailById(cdId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Detail Does not Exist"))
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(contactDetail)
}

func UpdateContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	userId := mux.Vars(r)["userId"]
	contactId := mux.Vars(r)["contactId"]
	contactDetailId := mux.Vars(r)["contactDetailId"]

	_, err := user_service.GetUserById(userId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
	}

	_, err = contact_service.GetContactById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Does not Exist"))
	}

	contactDetail, err := contact_details_service.GetContactDetailById(cdId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Detail Does not Exist"))
	}

	var container *contact_details_service.ContactDetail
	json.NewDecoder(r.Body).Decode(&container)

	contact_details_service.UpdateContactDetail(container, contactDetail)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Contact Detail Updated Successfully"))
}

func DeleteContactDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "application/json")

	uId := mux.Vars(r)["uId"]
	cId := mux.Vars(r)["cId"]
	cdId := mux.Vars(r)["cdId"]

	_, err := user_service.GetUserById(uId)

	if err {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does not Exist"))
	}

	_,err = contact_service.GetContactById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Does not Exist"))
	}

	contactDetail, err := contact_details_service.GetContactDetailById(cdId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Detail not found"))
	}

	contact_details_service.DeleteContactDetail(contactDetail)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Contact Detail Deleted Successfully"))
}