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
			rightJson.OrderJson.NoRm = item.(string)
		}
		if index == "pname" {
			rightJson.OrderJson.NamaPasien = item.(string)
		}
		if index == "pidentityno" {
			rightJson.OrderJson.Nik = item.(string)
		}
		if index == "pmobileno" {
			rightJson.OrderJson.Phone = item.(string)
		}
		if index == "street" {
			rightJson.OrderJson.Alamat = item.(string)
		}
		if index == "ptype" {
			rightJson.OrderJson.RujukanAsal = "1"
		}
		if index == "birth_dt" {
			year := item.(string)[0:4]
			month := item.(string)[4:6]
			date := item.(string)[6:8]
			hour := item.(string)[8:10]
			minute := item.(string)[10:12]
			fullDate := year + "-" + month + "-" + date + " " + hour + ":" + minute
			rightJson.OrderJson.TglLahir = fullDate
		}
		if index == "sex" {
			if item == 1 {
				rightJson.OrderJson.Jk = "L"
			} else {
				rightJson.OrderJson.Jk = "P"
			}
		}

		if index == "ono" {
			rightJson.OrderJson.NoOrder = item.(string)
		}
		if index == "lno" {
			rightJson.OrderJson.NoPendaftaran = ""
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
				rightJson.OrderJson.Cito = "true"
			} else {
				rightJson.OrderJson.Cito = "false"
			}
		}
		if index == "comment" {
			rightJson.OrderJson.Diagnosa = item.(string)
		}
		if index == "pstatus" {
			penjaminLen := strings.Split(item.(string), "|")
			var penjamin string
			var idPenjamin string
			if len(penjaminLen) > 1 {
				idPenjamin = strings.Split(item.(string), "|")[1]
				penjamin = strings.Split(item.(string), "|")[0]
			} else {
				penjamin = strings.Split(item.(string), "|")[0]
			}
			rightJson.OrderJson.Penjamin = penjamin
			rightJson.OrderJson.IdPenjamin = idPenjamin
			rightJson.OrderJson.JenisPasien = "ASURANSI"
			rightJson.OrderJson.IdJenisPasien = "2"

		} else {
			rightJson.OrderJson.JenisPasien = "UMUM"
			rightJson.OrderJson.IdJenisPasien = "1"
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
			rightJson.OrderJson.NoPendaftaran = item.(string)
		}
		if index == "clinician" {
			var dokter string
			var idDokter string
			if strings.Contains(item.(string), "|") {
				clinicianLen := strings.Split(item.(string), "|")
				if len(clinicianLen) > 1 {
					idDokter = strings.Split(item.(string), "|")[1]
				}
				dokter = strings.Split(item.(string), "|")[0]
			}

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
		if index == "pob" {
			rightJson.OrderJson.TempatLahir = item.(string)
		} else {
			rightJson.OrderJson.TempatLahir = "-"
		}
	}

	rightJson.OrderJson.Icd10 = []model.Icd10{}
	rightJson.OrderJson.Order = resultOrders
	rightJson.OrderJson.DetailRujukan = detailRujukan
	rightJson.Pattern = "bridging_order"
	return rightJson

}
