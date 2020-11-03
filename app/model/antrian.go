package model

import (
	"fmt"
	"log"

	"firebase.google.com/go/db"
)

// Antrian describe queue
type Antrian struct {
	ID     string `json: "id"`
	Status bool   `json: "status"`
}

func GetAntrian() (bool, []map[string]interface{}, error) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("error reading from database: ", err)
		return false, nil, err
	}

	return true, data, nil
}

func AddAntrian() (bool, error) {
	_, dataAntrian, _ := GetAntrian()
	var ID string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		ID = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		ID = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("&d", len(dataAntrian)))
	}

	// data = append(data, Antrian{
	// 	ID:		ID,
	// 	Status: false,
	// })
	antrian := Antrian{
		ID:     ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func UpdateAntrian(idAntrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := string.Split(idAntrian, "-") // B-0 => [B, 0]
	childRef := ref.Child(id[1])
	antrian := Antrian{
		ID:     idAntrian,
		Status: true,
	}

	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func DeleteAntrian(idAantrian string) (bool, error) {
	ref := client.NewRef("antrian")
	id := string.Split(idAntrian, "-") // B-0 => [B, 0]
	childRef := ref.Child(id[1])

	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
