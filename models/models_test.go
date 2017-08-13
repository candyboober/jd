package models

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestVacancy(t *testing.T) {
	// init params
	db := DB{}
	db.Init("user=candy dbname=jd_test password=1")
	v := Vacancy{}

	// test create obj
	data := map[string]string{"title": "Some title", "body": "So long text body"}
	v = Vacancy{Title: data["title"], Body: data["body"]}
	db.Connect.Create(&v)
	if v.ID == 0 {
		t.Error("ID of object is `0`")
	}

	// test update obj
	updatedField := "Another title"
	db.Connect.Model(&v).Update("title", updatedField)
	if v.Title != updatedField {
		t.Error("Field wasn't updated")
	}

	// test get obj
	id := strconv.FormatUint(uint64(v.ID), 16)
	db.Connect.Where("id = ?", id).First(&v)
	if v.ID == 0 {
		t.Error("Object wasn't found")
	}

	// test get collection of objects
	var vacansies []Vacancy
	//db.Connect.Find(db.Connect, &vacansies)
	db.Connect.Find(&vacansies)
	vacansyJson, err := json.Marshal(vacansies)
	if err != nil {
		panic(err)
	}
	if len(vacansyJson) == 0 {
		t.Error("Collection of Vacansy is empty")
	}

	//test delete obj
	db.Connect.Delete(&v)

	newV := Vacancy{}
	db.Connect.Where("id = ?", id).First(&newV)
	if newV.ID != 0 {
		t.Error("Object was found, but has been deleted")
	}
}
