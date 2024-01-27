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

// GetRecordHandler godoc
// @Summary Get Record
// @Description Get Record
// @Tags Record
// @Accept  json
// @Produce  json
// @Param id path string true "Record ID"
// @Success 200 {object} GetRecordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /records/{id} [get]
func (h *RecordHandler) GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := ctx.Value("ctx").(*models.Request)

	vars := mux.Vars(r)
	id := vars["id"]

	record, err := h.recordRepo.FindById(req.User.ID, id)

	if err != nil {
		message := fmt.Sprintf(constants.EntityNotFound, "Record", id)
		helpers.SendErrorResponse(w, message, constants.NotFound, err)
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

// SaveRecordHandler godoc
// @Summary Save Record
// @Description Save Record
// @Tags Record
// @Accept  json
// @Produce  json
// @Param id path string true "Record ID"
// @Param body body SaveRecordRequest true "Record"
// @Success 200 {object} GetRecordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /records [post]
func (h *RecordHandler) SaveRecordHandler(w http.ResponseWriter, r *http.Request) {
	var data models.SaveRecordRequest

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.BadRequest, err)
		return
	}

	helpers.ValidateStruct(w, &data)

	if data.TTL == 0 {
		data.TTL = config.AppConfig.DefaultTTL
	}

	if data.TTL < config.AppConfig.DefaultTTL {
		helpers.SendErrorResponse(w, fmt.Sprintf("TTL cannot be less than %d", config.AppConfig.DefaultTTL), constants.BadRequest, nil)
		return
	}

	ctx := r.Context().Value("ctx").(*models.Request)

	tmp := &models.RecordModel{
		ID:        data.Key,
		Value:     data.Value,
		ExpiresAt: time.Now().Local().Add(time.Duration(data.TTL) * time.Second).UTC(),
		TenantId:  ctx.User.ID,
	}

	err = h.recordRepo.Save(tmp)

	if err != nil {
		helpers.SendErrorResponse(w, "Unable to save record. Please try again.", constants.InternalServerError, err)
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
