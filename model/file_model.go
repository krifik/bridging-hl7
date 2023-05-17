package model

import "time"

type Patient struct {
	MRN          string `json:"mrn"`
	IDNumber     string `json:"idNumber"`
	FullName     string `json:"fullname"`
	Gender       string `json:"gender"`
	DateOfBirth  string `json:"dob"`
	PlaceOfBirth string `json:"placeOfBirth"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
}

type Demographics struct {
	RegNumber        string  `json:"regNumber"`
	VisitNumber      string  `json:"visitNumber"`
	Type             string  `json:"type"`
	IsCyto           bool    `json:"isCyto"`
	Diagnose         string  `json:"diagnose"`
	RegistrationDate string  `json:"registrationDate"`
	CollectDate      string  `json:"collectDate"`
	DoctorID         string  `json:"doctorId"`
	DoctorName       string  `json:"doctorName"`
	PartnerID        string  `json:"partnerId"`
	PartnerName      string  `json:"partnerName"`
	Patient          Patient `json:"patient"`
	CommentsSample   string  `json:"commentsSample"`
	NoOrder          string  `json:"noOrder"`
	OrderTestID      string  `json:"orderTestId"`
	SourceID         string  `json:"sourceId"`
	SourceName       string  `json:"sourceName"`
	SourceType       string  `json:"sourceType"`
}

type Response struct {
	Count          int            `json:"count"`
	NoLaboratorium string         `json:"noLaboratorium"`
	NoOrder        string         `json:"noOrder"`
	Detail         []interface{}  `json:"detail"`
	Demographics   Demographics   `json:"demographics"`
	Examinations   []Examinations `json:"examinations"`
}

type MetaData struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type JSONRequest struct {
	Response Response `json:"response"`
	MetaData MetaData `json:"metaData"`
}

type Examinations struct {
	TestID     int        `json:"testId"`
	ParentID   string     `json:"parentId"`
	Position   string     `json:"position"`
	TestName   string     `json:"testName"`
	ExternalID string     `json:"externalId"`
	Children   []Children `json:"children"`
}

type Children struct {
	ID              string     `json:"id"`
	ExternalID      string     `json:"externalId"`
	TestID          int        `json:"testId"`
	ParentID        int        `json:"parentId"`
	Position        string     `json:"position"`
	TestName        string     `json:"testName"`
	PanelID         int        `json:"panelId"`
	ExamValue       string     `json:"examValue"`
	ExamValueType   string     `json:"examValueType"`
	ExamValueFlag   string     `json:"examValueFlag"`
	NormalValueText string     `json:"normalValueText"`
	Indicator       string     `json:"indicator"`
	UnitName        string     `json:"unitName"`
	ConfirmedAt     time.Time  `json:"confirmedAt"`
	ConfirmedBy     string     `json:"confirmedBy"`
	VerifiedAt      time.Time  `json:"verifiedAt"`
	VerifiedBy      string     `json:"verifiedBy"`
	ValidatedAt     time.Time  `json:"validatedAt"`
	ValidatedBy     string     `json:"validatedBy"`
	Comment         string     `json:"comment"`
	Decimal         int        `json:"decimal"`
	Metode          string     `json:"metode"`
	Children        []Children `json:"children"`
}

type MSH struct {
	Type      string
	MessageID string
	MessageDT string
	Version   string
}

type OBR struct {
	Type        string
	PID         string
	Apid        string
	Pname       string
	Pidentityno string
	Pmobileno   string
	Street      string
	Title       string
	Ptype       string
	BirthDt     string
	Sex         string
	Ono         string
	Lno         string
	RequestDt   string
	SpecimenDt  string
	Source      string
	Clinician   string
	Priority    string
	Pstatus     string
	Visitno     string
	OrderTestID string
	Comment     string
}

type OBX struct {
	Type  string
	Items []string
	Item  string
}

type FileResult struct {
	Msh MSH
	Obr OBR
	Obx OBX
}

type FileRequest struct {
	ID        int
	FileName  string
	CreatedAt time.Time
}
