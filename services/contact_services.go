package services

import (
	"log"

	"github.com/aidensV/gin_example/dtos"
	"github.com/aidensV/gin_example/models"
	"github.com/aidensV/gin_example/repositories"
	"github.com/google/uuid"
)

func CreateContact(contact *models.Contact, repository repositories.ContactRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}
	contact.ID = uuidResult.String()
	operationResult := repository.Save(contact)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}

func FindAllContacts(repository repositories.ContactRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	var datas = operationResult.Result.(*models.Contacts)

	return dtos.Response{Success: true, Data: datas}
}
func FindByIdContact(id string, repository repositories.ContactRepository) dtos.Response {
	operationResult := repository.FindById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}

func UpdateContactById(id string, contact *models.Contact, repository repositories.ContactRepository) dtos.Response {
	existingContactResponse := FindByIdContact(id, repository)

	if !existingContactResponse.Success {
		return existingContactResponse
	}
	existingContact := existingContactResponse.Data.(*models.Contact)

	existingContact.Name = contact.Name
	existingContact.Email = contact.Email
	existingContact.Address = contact.Address

	operationResult := repository.Save(existingContact)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	return dtos.Response{Success: true, Data: operationResult.Result}

}

func DeleteContactById(id string, repository repositories.ContactRepository) dtos.Response {
	operationResult := repository.DeleteById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
