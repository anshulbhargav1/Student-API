package storage

import (
	"github.com/anshulbhargav1/student-api/Internal/types"
)

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student , error)
}
