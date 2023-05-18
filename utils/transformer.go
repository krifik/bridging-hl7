package utils

import (
	"strings"

	"github.com/krifik/bridging-hl7/model"
)

/*
here is function for transform raw data from file to 'right' json
data from file is an unordered array
so i (krifik) make a condition based on index and assign it
and transform to 'right' json format
*/

func TransformToRightJson(data map[string]interface{}) model.Json {
	var rightJson model.Json
	var detailRujukan []model.Rujukan
	var resultOrders []model.Order
	for index, item := range data {

		if index == "pid" {
			rightJson.NoRm = item.(string)
		}
		if index == "pname" {
			rightJson.NamaPasien = item.(string)
		}
		if index == "pidentityno" {
			rightJson.Nik = item.(string)
		}
		if index == "pmobileno" {
			rightJson.Phone = item.(string)
		}
		if index == "street" {
			rightJson.Alamat = item.(string)
		}
		if index == "ptype" {
			rightJson.RujukanAsal = "1"
		}
		if index == "birth_dt" {
			year := item.(string)[0:4]
			month := item.(string)[4:6]
			date := item.(string)[6:8]
			hour := item.(string)[8:10]
			minute := item.(string)[10:12]
			fullDate := year + "-" + month + "-" + date + " " + hour + ":" + minute
			rightJson.TglLahir = fullDate
		}
		if index == "sex" {
			if item == 1 {
				rightJson.Jk = "P"
			} else {
				rightJson.Jk = "L"
			}
		}

		if index == "ono" {
			rightJson.NoOrder = item.(string)
		}
		if index == "lno" {
			rightJson.NoPendaftaran = ""
		}

		//! dont know this is  necessary for bridging or not, for now, i just commented it LOL

		// if index == "request_dt" {
		// 	year := item.(string)[0:4]
		// 	month := item.(string)[4:6]
		// 	date := item.(string)[6:8]
		// 	hour := item.(string)[8:10]
		// 	minute := item.(string)[10:12]
		// 	fullDate := year + "-" + month + "-" + date + " " + hour + ":" + minute

		// 	pp.Print(fullDate)

		// }

		if index == "priority" {
			if item.(string) == "CITO" {
				rightJson.Cito = "true"
			} else {
				rightJson.Cito = "false"
			}
		}
		if index == "comment" {
			rightJson.Diagnosa = item.(string)
		}
		if index == "pstatus" {
			penjamin := strings.Split(item.(string), "|")[0]
			idPenjamin := strings.Split(item.(string), "|")[1]
			rightJson.Penjamin = penjamin
			rightJson.IdPenjamin = idPenjamin
			rightJson.JenisPasien = "ASURANSI"
			rightJson.IdJenisPasien = "2"

		} else {
			rightJson.JenisPasien = "UMUM"
			rightJson.IdJenisPasien = "1"
		}

		if index == "source" {
			ward := strings.Split(item.(string), "|")[0]
			idWard := strings.Split(item.(string), "|")[1]
			rujukanSource := model.Rujukan{
				Ward:   ward,
				IdWard: idWard,
			}
			if len(detailRujukan) > 0 {
				for i := 0; i < len(detailRujukan); i++ {
					detailRujukan[i].Ward = rujukanSource.Ward
					detailRujukan[i].IdWard = rujukanSource.IdWard
				}
			} else {
				detailRujukan = append(detailRujukan, rujukanSource)
			}
		}
		if index == "visitno" {
			rightJson.NoPendaftaran = item.(string)
		}
		if index == "clinician" {
			dokter := strings.Split(item.(string), "|")[0]
			idDokter := strings.Split(item.(string), "|")[1]
			rujukanClinician := model.Rujukan{
				IdDokter:   idDokter,
				NamaDokter: dokter,
			}
			if len(detailRujukan) > 0 {
				for i := 0; i < len(detailRujukan); i++ {
					detailRujukan[i].NamaDokter = rujukanClinician.NamaDokter
					detailRujukan[i].IdDokter = rujukanClinician.IdDokter
				}
			} else {
				detailRujukan = append(detailRujukan, rujukanClinician)
			}

			// if index == "phblebotomis" {
			// 	// just leave it here, in case in the future needs it
			// }
		}
		if index == "order_test_id" {
			orders := strings.Split(item.(string), "~")
			for _, test := range orders {
				var testName string
				testSplit := strings.Split(test, "^")
				testId := testSplit[0]
				if len(testSplit) > 1 {
					testName = testSplit[1]
				} else {
					testName = ""
				}
				// testName := strings.Split(test, "^")[1]
				modelOrder := model.Order{
					IdTest:   testId,
					NamaTest: testName,
				}
				resultOrders = append(resultOrders, modelOrder)
			}

		}
		if index == "order_control" {
			if item.(string) == "CA" {
				resultOrders = make([]model.Order, 0)
			}
		}
	}

	rightJson.Order = resultOrders
	rightJson.DetailRujukan = detailRujukan
	return rightJson

}
