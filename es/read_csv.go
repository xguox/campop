package es

import (
	. "campop/commonutil"
	"campop/model"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const (
	CAMERA_DATA_FILE = "cameras.csv"
	LENS_DATA_FILE   = "lenses.csv"
	FILE_DELIMITER   = ','
)

func ReadCamerasCsv() {
	file, err := os.Open(CAMERA_DATA_FILE)
	CheckError("", err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = FILE_DELIMITER
	reader.LazyQuotes = true

	rows, err := reader.ReadAll()
	CheckError("", err)
	var cameras []model.Camera
	for _, row := range rows {
		camera := model.Camera{}
		s := reflect.ValueOf(&camera).Elem()
		id, err := strconv.ParseInt(row[0], 10, 64)
		CheckError("", err)
		s.Field(0).SetInt(id)
		for i := 1; i < s.NumField(); i++ {
			f := s.Field(i)
			f.SetString(row[i])
		}

		cameras = append(cameras, camera)
		insertDoc(CameraIndexName, CameraTypeName, fmt.Sprintf("%d", camera.ID), camera)
	}

	fmt.Printf("Camera docs: %d\n", len(cameras))

}

func ReadLensesCsv() {
	file, err := os.Open(LENS_DATA_FILE)
	CheckError("", err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = FILE_DELIMITER
	reader.LazyQuotes = true

	rows, err := reader.ReadAll()
	CheckError("", err)
	var lenses []model.Lens
	for _, row := range rows {
		lens := model.Lens{}
		s := reflect.ValueOf(&lens).Elem()
		// typeOfLens := s.Type()
		id, err := strconv.ParseInt(row[0], 10, 64)
		CheckError("", err)
		s.Field(0).SetInt(id)
		for i := 1; i < s.NumField(); i++ {
			f := s.Field(i)
			f.SetString(row[i])

			// fmt.Printf("%d: %s %s = %v\n", i, typeOfLens.Field(i).Name, f.Type(), f.Interface())
		}

		lenses = append(lenses, lens)
		insertDoc(LensIndexName, LensTypeName, fmt.Sprintf("%d", lens.ID), lens)
	}

	fmt.Printf("Lens docs: %d\n", len(lenses))
}
