package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type NullBool struct {
	sql.NullBool
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullInt64 struct {
	sql.NullInt64
}

type NullString struct {
	sql.NullString
}

type NullTime struct {
	mysql.NullTime
}

type Model interface {
	PrimaryKey() map[string]uint
	ForeignKey() map[string]uint
	Columns() map[string]interface{}
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
