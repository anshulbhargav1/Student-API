package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/anshulbhargav1/student-api/Internal/storage"
	"github.com/anshulbhargav1/student-api/Internal/types"
	"github.com/anshulbhargav1/student-api/Internal/utiles/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		slog.Info("creating the student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return

		}

		if err != nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation..

		if err:= validator.New().Struct(student); err != nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validateErrs))
			return
		}

		// use createstudent method
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created succesfully", slog.String("userId", fmt.Sprint(lastId)) )

		if err != nil{
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}



		response.WriteJson(w, http.StatusCreated, map[string] int64{"id":lastId})
		
	}
}