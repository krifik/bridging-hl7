package model

type Icd10 struct {
	Type string `json:"type"`
	Code string `json:"code"`
	Text string `json:"text"`
}
