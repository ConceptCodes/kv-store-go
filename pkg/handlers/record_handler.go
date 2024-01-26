package handlers

import (
	"encoding/json"
	"fmt"
	"kv-store/pkg/constants"
	"kv-store/pkg/helpers"
	"kv-store/pkg/models"
	repository "kv-store/pkg/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

type RecordHandler struct {
	recordRepo repository.RecordRepository
}

var err error

func NewRecordHandler(recordRepo repository.RecordRepository) *RecordHandler {
	return &RecordHandler{recordRepo: recordRepo}
}

// get record
func (h *RecordHandler) GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context().Value("ctx").(*models.Request)

	vars := mux.Vars(r)
	id := vars["id"]

	record, err := h.recordRepo.FindById(ctx.Id, id)

	if err != nil {
		message := fmt.Sprintf(constants.EntityNotFound, "Tenant", id)
		helpers.SendErrorResponse(w, message, constants.NotFound)
	}

	res := &models.GetRecordResponse{
		Key:   record.ID,
		Value: record.Value,
	}

	helpers.SendSuccessResponse(w, "Tenant Found Successfully", res)
}

// save record
func (h *RecordHandler) SaveRecordHandler(w http.ResponseWriter, r *http.Request) {
	var data models.SaveRecordRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.BadRequest)
	}

	helpers.ValidateStruct(w, &data)

	ctx := r.Context().Value("ctx").(*models.Request)

	tmp := &models.RecordModel{
		ID:       data.Key,
		Value:    data.Value,
		TTL:      uint64(data.TTL),
		TenantId: ctx.User.ID,
	}

	err = h.recordRepo.Save(tmp)

	if err != nil {
		helpers.SendErrorResponse(w, "Unable to save record. Please try again.", constants.InternalServerError)
	}

	helpers.SendSuccessResponse(w, "Successfully saved record.", tmp)

}
