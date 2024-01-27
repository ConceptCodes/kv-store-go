package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"kv-store/config"
	"kv-store/internal/constants"
	"kv-store/internal/helpers"
	"kv-store/internal/models"
	repository "kv-store/internal/repositories"
)

type RecordHandler struct {
	recordRepo repository.RecordRepository
}

var err error

func NewRecordHandler(recordRepo repository.RecordRepository) *RecordHandler {
	return &RecordHandler{recordRepo: recordRepo}
}

func (h *RecordHandler) GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := ctx.Value("ctx").(*models.Request)

	vars := mux.Vars(r)
	id := vars["id"]

	record, err := h.recordRepo.FindById(req.User.ID, id)

	if err != nil {
		message := fmt.Sprintf(constants.EntityNotFound, "Record", id)
		helpers.SendErrorResponse(w, message, constants.NotFound)
		return
	}

	res := &models.GetRecordResponse{
		Key:     record.ID,
		Value:   record.Value,
		Expires: record.ExpiresAt.Format(constants.TimeFormat),
	}

	helpers.SendSuccessResponse(w, "Record Found Successfully", res)
	return
}

func (h *RecordHandler) SaveRecordHandler(w http.ResponseWriter, r *http.Request) {
	var data models.SaveRecordRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.BadRequest)
		return
	}

	helpers.ValidateStruct(w, &data)

	if data.TTL == 0 {
		data.TTL = config.AppConfig.DefaultTTL
	}

	if data.TTL < config.AppConfig.DefaultTTL {
		helpers.SendErrorResponse(w, "TTL cannot be less than default TTL", constants.BadRequest)
	}

	ctx := r.Context().Value("ctx").(*models.Request)

	tmp := &models.RecordModel{
		ID:        data.Key,
		Value:     data.Value,
		ExpiresAt: time.Now().Add(time.Duration(data.TTL) * time.Second),
		TenantId:  ctx.User.ID,
	}

	err = h.recordRepo.Save(tmp)

	if err != nil {
		helpers.SendErrorResponse(w, "Unable to save record. Please try again.", constants.InternalServerError)
		return
	}

	res := &models.GetRecordResponse{
		Key:     tmp.ID,
		Value:   tmp.Value,
		Expires: tmp.ExpiresAt.Format(constants.TimeFormat),
	}

	helpers.SendSuccessResponse(w, "Successfully saved record.", res)
	return
}
