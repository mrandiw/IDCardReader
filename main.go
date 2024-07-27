package main

import (
	"fmt"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

type Person struct {
	NIK       string
	Name      string
	Born      string
	Sex       string
	BloodType string
	Address   string
	RTRW      string
	Village   string
	District  string
}

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage("image/watermark.png")

	text, _ := client.Text()

	person, err := GetKTP(text)
	if err != nil {
		fmt.Printf("error when GetKTP : %s", err)
	}

	fmt.Printf("%+v\n", person)
}

func GetKTP(text string) (Person, error) {
	textArr := strings.Split(strings.ToUpper(text), ":")
	fmt.Println("textArr :", textArr)
	var person Person
	for idx, val := range textArr {
		if strings.Contains(val, "NIK") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var nik string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "NAMA") {
					nik = nik + val2
					person.NIK = nik
				}
			}
		}
		if strings.Contains(val, "NAMA") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var nama string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "TEMPAT") {
					nama = nama + " " + val2
					person.Name = nama
				} else {
					nama = nama + " " + val2
				}
			}
		}
		if strings.Contains(val, "TEMPAT") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var tempat string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "JENIS") {
					tempat = tempat + " " + val2
					person.Born = tempat
				} else {
					tempat = tempat + " " + val2
				}
			}
		}
		if strings.Contains(val, "JENIS") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var sex string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "GOL.") {
					sex = sex + " " + val2
					person.Sex = sex
				} else {
					sex = sex + " " + val2
				}
			}
		}
		if strings.Contains(val, "GOL.") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var darah string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "ALAMAT") {
					darah = darah + " " + val2
					person.BloodType = darah
				} else {
					darah = darah + " " + val2
				}
			}
		}
		if strings.Contains(val, "ALAMAT") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var alamat string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "RT/RW") {
					alamat = alamat + " " + val2
					person.Address = alamat
				} else {
					alamat = alamat + " " + val2
				}
			}
		}
		if strings.Contains(val, "RT/RW") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var rtrw string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "KEL/DESA") {
					rtrw = rtrw + " " + val2
					person.RTRW = rtrw
				} else {
					rtrw = rtrw + " " + val2
				}
			}
		}
		if strings.Contains(val, "KEL/DESA") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var desa string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "KECAMATAN") {
					desa = desa + " " + val2
					person.Village = desa
				} else {
					desa = desa + " " + val2
				}
			}
		}
		if strings.Contains(val, "KECAMATAN") {
			textArrSpec := strings.Fields(strings.ToUpper(textArr[idx+1]))
			var kecamatan string
			for idx2, val2 := range textArrSpec {
				if len(textArrSpec) > idx2+1 && strings.Contains(textArrSpec[idx2+1], "AGAMA") {
					kecamatan = kecamatan + " " + val2
					person.District = kecamatan
				} else {
					kecamatan = kecamatan + " " + val2
				}
			}
		}
	}

	return person, nil
}
