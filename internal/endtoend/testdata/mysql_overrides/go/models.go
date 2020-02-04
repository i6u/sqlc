// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"database/sql"

	"example.com/mysql"
)

type JobStatusType string

const (
	APPLIED  JobStatusType = "APPLIED"
	PENDING  JobStatusType = "PENDING"
	ACCEPTED JobStatusType = "ACCEPTED"
	REJECTED JobStatusType = "REJECTED"
)

func (e *JobStatusType) Scan(src interface{}) error {
	*e = JobStatusType(src.([]byte))
	return nil
}

type Order struct {
	ID     mysql.ID
	Price  float64
	UserID int
}

type User struct {
	ID        mysql.ID
	FirstName string
	LastName  sql.NullString
	Age       int
	JobStatus JobStatusType
	Created   mysql.Timestamp
}
