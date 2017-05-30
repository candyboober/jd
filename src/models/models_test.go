package models

import (
	"testing"
	"strconv"
)

func TestVacancy(t *testing.T) {
	// init params
	db := DB{}
	db.Init("user=candy dbname=jd_test password=1")
	v := Vacancy{}
	// test create table
	v.CreateTable(*db.Connect)
	// test create obj
	v = Vacancy{Title: "Some title", Body: "So long text body"}
	db.Connect.Create(&v)
	// test update obj
	db.Connect.Model(&v).Update("title", "Another title")
	// test get obj
	id := strconv.FormatUint(uint64(v.ID), 16)
	db.Connect.Where("id = ?", id).First(&v)
	// test delete obj

}
