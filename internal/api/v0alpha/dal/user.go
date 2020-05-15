package dal

import (
	"context"
	"errors"
	"net/http"

	"github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/model"
	"github.com/alexbednarczyk/crud-gopher/internal/common"
	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

// CreateUser godoc
// @Summary creates a database entry for the user
// @Description creates user
// @Tags user crud
// @Accept json
// @Produce json
// @Param UserRequest body model.UserRequest false "User's data"
// @Success 201 {object} model.UserResponse
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users [post]
func CreateUser(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Acquire database connection
		conn, err := db.Acquire(c)
		if err != nil {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
			return
		}
		defer conn.Release()

		newRecord := model.UserResponse{}

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
		newRecord.GUID = uuid.NewV4().String()

		if _, err := conn.Exec(context.Background(), "insert into users(guid, display_name, first_name, last_name, created_at, updated_at) values($1, $2, $3, $4, $5, $6)", newRecord.GUID, newRecord.DisplayName, newRecord.FirstName, newRecord.LastName, newRecord.CreatedAt, newRecord.UpdatedAt); err == nil {
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
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [get]
func GetSpecificUser(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {

		if uuid, uuidFound := c.Params.Get("userId"); uuidFound {
			// Acquire database connection
			conn, err := db.Acquire(c)
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
			defer conn.Release()

			record := model.UserResponse{}

			if err := conn.QueryRow(context.Background(), "select * from users where guid=$1", uuid).Scan(&record.GUID, &record.DisplayName, &record.FirstName, &record.LastName, &record.CreatedAt, &record.UpdatedAt); err == nil {
				c.JSON(http.StatusOK, record)
			} else {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			log.Warn(errors.New("User GUID not found"))
			common.NewError(c, http.StatusBadRequest, errors.New("User GUID not found"))
		}
	}
}

// GetAllUsers godoc
// @Summary fetches all users from the database
// @Description fetch all users
// @Tags user crud
// @Produce json
// @Success 200 {object} model.UserResponse[]
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users [get]
func GetAllUsers(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Acquire database connection
		conn, err := db.Acquire(c)
		if err != nil {
			log.Warn(err.Error())
			common.NewError(c, http.StatusInternalServerError, err)
			return
		}
		defer conn.Release()

		records := []model.UserResponse{}

		if rows, err := conn.Query(context.Background(), "select * from users"); err == nil {
			for rows.Next() {
				record := model.UserResponse{}
				err := rows.Scan(&record.GUID, &record.DisplayName, &record.FirstName, &record.LastName, &record.CreatedAt, &record.UpdatedAt)
				if err != nil {
					log.Warn(err.Error())
					common.NewError(c, http.StatusInternalServerError, err)
					return
				}
				records = append(records, record)
			}
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
// @Param UserRequest body model.UserRequest false "User's data"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [patch]
func UpdateUser(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {

		if uuid, uuidFound := c.Params.Get("userId"); uuidFound {
			// Acquire database connection
			conn, err := db.Acquire(c)
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
			defer conn.Release()

			existingRecord := model.UserResponse{}

			if err := conn.QueryRow(context.Background(), "select * from users where guid=$1", uuid).Scan(&existingRecord.GUID, &existingRecord.DisplayName, &existingRecord.FirstName, &existingRecord.LastName, &existingRecord.CreatedAt, &existingRecord.UpdatedAt); err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			record := model.UserResponse{}

			if err := c.ShouldBindJSON(&record); err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			if len(record.DisplayName) == 0 {
				record.DisplayName = existingRecord.DisplayName
			}
			if len(record.FirstName) == 0 {
				record.FirstName = existingRecord.FirstName
			}
			if len(record.LastName) == 0 {
				record.LastName = existingRecord.LastName
			}

			currentTime, err := common.CurrentTime()
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

			record.CreatedAt = existingRecord.CreatedAt
			record.UpdatedAt = currentTime
			record.GUID = uuid

			if _, err := db.Exec(context.Background(), "update users set display_name=$2, first_name=$3, last_name=$4, updated_at=$5 where guid=$1", record.GUID, record.DisplayName, record.FirstName, record.LastName, record.UpdatedAt); err == nil {
				c.JSON(http.StatusOK, record)
			} else {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
		} else {
			log.Warn(errors.New("User GUID not found"))
			common.NewError(c, http.StatusBadRequest, errors.New("User GUID not found"))
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
// @Param UserRequest body model.UserRequest false "User's data"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} common.HTTPError
// @Failure 500 {object} common.HTTPError
// @Router /v0alpha/crud/users/{user_id} [put]
func ReplaceUser(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {

		if uuid, uuidFound := c.Params.Get("userId"); uuidFound {
			// Acquire database connection
			conn, err := db.Acquire(c)
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
			defer conn.Release()

			record := model.UserResponse{}

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
			record.GUID = uuid

			if _, err := db.Exec(context.Background(), "update users set display_name=$2, first_name=$3, last_name=$4, created_at=$5, updated_at=$6 where guid=$1", uuid, record.DisplayName, record.FirstName, record.LastName, record.CreatedAt, record.UpdatedAt); err == nil {
				c.JSON(http.StatusOK, record)
			} else {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
			}
		} else {
			log.Warn(errors.New("User GUID not found"))
			common.NewError(c, http.StatusBadRequest, errors.New("User GUID not found"))
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
func DeleteUser(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {

		if uuid, uuidFound := c.Params.Get("userId"); uuidFound {
			// Acquire database connection
			conn, err := db.Acquire(c)
			if err != nil {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}
			defer conn.Release()

			if _, err := db.Exec(context.Background(), "delete from users where guid=$1", uuid); err == nil {
				c.JSON(http.StatusNoContent, gin.H{})
			} else {
				log.Warn(err.Error())
				common.NewError(c, http.StatusInternalServerError, err)
				return
			}

		} else {
			log.Warn(errors.New("User GUID not found"))
			common.NewError(c, http.StatusBadRequest, errors.New("User GUID not found"))
		}
	}
}
