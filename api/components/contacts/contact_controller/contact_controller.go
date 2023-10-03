package contactcontroller

import (
	contact_service "contactapp_api/components/contacts/contact_service"
	user_service "contactapp_api/components/user/user_service"
	"contactapp_api/validators"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func GetContactsForUser(w http.ResponseWriter, r *http.Request) {

	uId := mux.Vars(r)["uId"]

	user,err := user_service.GetUserById(uId)

	if err {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	contacts := user.Contacts
	w.Header().Set("ContentType", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func AddNewContact(w http.ResponseWriter, r *http.Request) {

	uId := mux.Vars(r)["uId"]
	var contact *contact_service.Contact

	json.NewDecoder(r.Body).Decode(&contact)

	user, err := user_service.GetUserById(uId)

	if err {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Does Not Exist"))
		return
	}

	if !validators.ValidateName(contact.FirstName) {
		panic("Enter valid contact Firstname")
	}

	if !validators.ValidateName(contact.LastName) {
		panic("Enter valid contact Lastname")
	}

	newContact := contact_service.NewContact(contact.FirstName, contact.LastName)
	user_service.InsertContact(user, newContact)
	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact Successfully Created"})
}

func UpdateExistingContact(w http.ResponseWriter, r *http.Request) {
	cId := mux.Vars(r)["cId"]
	contact, err := contact_service.GetContactById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact does Not Exist"))
		return
	}

	var body *contact_service.Contact
	json.NewDecoder(r.Body).Decode(&body)

	contact_service.UpdateContact(body, contact)

	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact Updated Successfully"})
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {

	cId := mux.Vars(r)["cId"]

	contact, err := contact_service.GetContactById(cId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Contact Does not Exist"))
		return
	}

	contact_service.DeleteContact(contact)

	w.Header().Set("ContentType", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact Deleted Successfully"})
}