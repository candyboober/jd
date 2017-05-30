package models

import "testing"


func TestVacancy(t *testing.T) {
	// init params
	db := DB{}
	db.Init("user=candy dbname=jd_test password=1")
	v := Vacancy{}
	// test create table
	v.CreateTable(*db.Connect)
	// test create obj

	// test update obj

	// test get obj

	// test delete obj
}