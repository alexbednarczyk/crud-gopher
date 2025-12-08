package handlers

import (
	"errors"
	"net/http"

	"github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/models"
	"github.com/alexbednarczyk/crud-gopher/internal/common"
	"github.com/alexbednarczyk/crud-gopher/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// CreateUser godoc
// @Summary creates a database entry for the user
// @Description creates user
// @Tags user crud
// @Accept json
// @Produce json
// @Param UserRequest body models.UserRequest false "User's data"
// @Success 201 {object} models.UserResponse
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users [post]
func CreateUser(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		newRecord := models.UserResponse{}

		if err := c.ShouldBindJSON(&newRecord); err != nil {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
			return
		}

		currentTime, err := common.CurrentTime()
		if err != nil {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
			return
		}
		newRecord.CreatedAt = currentTime
		newRecord.UpdatedAt = currentTime
		newRecord.GUID = uuid.NewString()

		if err := repo.Create(c.Request.Context(), &newRecord); err == nil {
			c.JSON(http.StatusCreated, newRecord)
		} else {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
		}
	}
}

// GetSpecificUser godoc
// @Summary fetches a user from the database given the user's GUID
// @Description fetch a user's data
// @Tags user crud
// @Accept json
// @Produce json
// @Param user_id path string true "User GUID"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [get]
func GetSpecificUser(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		if uuidParam, uuidFound := c.Params.Get("userId"); uuidFound {
			record, err := repo.GetByGUID(c.Request.Context(), uuidParam)
			if err == nil {
				c.JSON(http.StatusOK, record)
			} else {
				if errors.Is(err, repository.ErrNotFound) {
					common.NewError(c, http.StatusNotFound, err)
					return
				}
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			err := errors.New("User GUID not found")
			log.Warn(err)
			common.NewError(c, http.StatusBadRequest, err)
		}
	}
}

// GetAllUsers godoc
// @Summary fetches all users from the database
// @Description fetch all users
// @Tags user crud
// @Produce json
// @Success 200 {object} models.UserResponse[]
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users [get]
func GetAllUsers(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		records, err := repo.GetAll(c.Request.Context())
		if err == nil {
			c.JSON(http.StatusOK, records)
		} else {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
		}
	}
}

// UpdateUser godoc
// @Summary updates a database entry for the user
// @Description update a user's database entry
// @Tags user crud
// @Accept json
// @Produce json
// @Param user_id path string true "User GUID"
// @Param UserRequest body models.UserRequest false "User's data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [patch]
func UpdateUser(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		if uuidParam, uuidFound := c.Params.Get("userId"); uuidFound {
			existingRecord, err := repo.GetByGUID(c.Request.Context(), uuidParam)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					common.NewError(c, http.StatusNotFound, err)
					return
				}
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			var inputRecord models.UserResponse
			if err := c.ShouldBindJSON(&inputRecord); err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			if len(inputRecord.DisplayName) > 0 {
				existingRecord.DisplayName = inputRecord.DisplayName
			}
			if len(inputRecord.FirstName) > 0 {
				existingRecord.FirstName = inputRecord.FirstName
			}
			if len(inputRecord.LastName) > 0 {
				existingRecord.LastName = inputRecord.LastName
			}

			currentTime, err := common.CurrentTime()
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
			existingRecord.UpdatedAt = currentTime

			if err := repo.Update(c.Request.Context(), existingRecord); err == nil {
				c.JSON(http.StatusOK, existingRecord)
			} else {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			err := errors.New("User GUID not found")
			log.Warn(err)
			common.NewError(c, http.StatusBadRequest, err)
		}
	}
}

// ReplaceUser godoc
// @Summary replaces a database entry for the user
// @Description replace user
// @Tags user crud
// @Accept json
// @Produce json
// @Param user_id path string true "User GUID"
// @Param UserRequest body models.UserRequest false "User's data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [put]
func ReplaceUser(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		if uuidParam, uuidFound := c.Params.Get("userId"); uuidFound {
			var record models.UserResponse
			if err := c.ShouldBindJSON(&record); err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			currentTime, err := common.CurrentTime()
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			record.CreatedAt = currentTime
			record.UpdatedAt = currentTime
			record.GUID = uuidParam

			// Note: Replace implies complete overwrite.
			// However, in standard REST PUT, if resource exists, update it. If not, create?
			// The original code did an UPDATE on all fields.
			// But wait, if it's an UPDATE, we might need to know if it existed.
			// The original code returned success if UPDATE succeeded.
			// My repository Update returns ErrNotFound if not found.
			// So if Update fails with NotFound, should we Create?
			// The original code: "update users set ... where guid=$1". If 0 rows, it didn't error?
			// Wait, original code:
			// if _, err := db.Exec(...); err == nil { ... }
			// db.Exec doesn't return error on 0 rows affected.
			// So the original code reported success even if ID didn't exist?
			// "update" returns CommandTag.
			// Let's assume strict update for now as per "ReplaceUser" naming usually on existing resource or we'd handle upsert.
			// I'll stick to calling Update and handling NotFound.

			err = repo.Update(c.Request.Context(), &record)
			if err == nil {
				c.JSON(http.StatusOK, record)
			} else {
				if errors.Is(err, repository.ErrNotFound) {
					common.NewError(c, http.StatusNotFound, err)
					return
				}
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			err := errors.New("User GUID not found")
			log.Warn(err)
			common.NewError(c, http.StatusBadRequest, err)
		}
	}
}

// DeleteUser godoc
// @Summary delete a database entry for the user
// @Description delete a user
// @Tags user crud
// @Accept json
// @Produce json
// @Param user_id path string true "User GUID"
// @Success 204 ""
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [delete]
func DeleteUser(repo repository.UserRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		if uuidParam, uuidFound := c.Params.Get("userId"); uuidFound {
			if err := repo.Delete(c.Request.Context(), uuidParam); err == nil {
				c.JSON(http.StatusNoContent, gin.H{})
			} else {
				if errors.Is(err, repository.ErrNotFound) {
					common.NewError(c, http.StatusNotFound, err)
					return
				}
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			err := errors.New("User GUID not found")
			log.Warn(err)
			common.NewError(c, http.StatusBadRequest, err)
		}
	}
}
