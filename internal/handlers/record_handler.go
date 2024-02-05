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
	"kv-store/pkg/storage/redis"
)

type RecordHandler struct {
	redis *redis.Redis
}

var err error

func NewRecordHandler(_redis *redis.Redis) *RecordHandler {
	return &RecordHandler{
		redis: _redis,
	}
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
	key := vars["key"]

	compoundKey := fmt.Sprintf("%s:%s", req.User.ID, key)

	record, err := h.redis.GetData(compoundKey)

	if err != nil {
		message := fmt.Sprintf(constants.EntityNotFound, "Record", compoundKey)
		helpers.SendErrorResponse(w, message, constants.NotFound, err)
		return
	}

	res := &models.GetRecordResponse{
		Key:   key,
		Value: record,
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

	key := fmt.Sprintf("%s:%s", ctx.User.ID, data.Key)

	err = h.redis.SetData(key, data.Value, time.Duration(data.TTL)*time.Second)

	if err != nil {
		helpers.SendErrorResponse(w, "Unable to save record. Please try again.", constants.InternalServerError, err)
		return
	}

	res := &models.GetRecordResponse{
		Key:     data.Key,
		Value:   data.Value,
		Expires: time.Now().Add(time.Duration(data.TTL) * time.Second).Format(time.RFC3339),
	}

	helpers.SendSuccessResponse(w, "Successfully saved record.", res)
	return
}
