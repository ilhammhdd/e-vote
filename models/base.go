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
	Value sql.NullBool
}

type NullFloat64 struct {
	Value sql.NullFloat64
}

type NullInt64 struct {
	Value sql.NullInt64
}

type NullString struct {
	Value sql.NullString
}

type NullTime struct {
	Value mysql.NullTime
}

type Model interface {
	PrimaryKey() map[string]uint
	ForeignKey() map[string]uint
	Columns() map[string]interface{}
}

func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Value.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Value.Bool)
}

func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Value.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Value.Float64)
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Value.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Value.Int64)
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Value.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Value.String)
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Value.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Value.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
