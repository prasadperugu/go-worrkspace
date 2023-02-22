package main

import (
	"encoding/json"
	"net/http"
)

type HealthData struct {
	HR          int `json:"hr"`
	Spo2        int `json:"spo2"`
	RR          int `json:"rr"`
	LastUpdated int `json:"lastUpdated"`
}

type Alarm struct {
	AlarmName   string `json:"alarmName"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	CreatedAt   int    `json:"createdAt"`
}

type Patient struct {
	MRN        int        `json:"MRN"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	HealthData HealthData `json:"healthData"`
	ECA        string     `json:"eca"`
	Alarms     []Alarm    `json:"alarms"`
}

type PatientList struct {
	Patients []Patient `json:"patients"`
}

func main() {
	http.HandleFunc("/patients", getPatientList)
	http.ListenAndServe(":8080", nil)
}

func getPatientList(w http.ResponseWriter, r *http.Request) {
	patientList := PatientList{
		Patients: []Patient{
			{
				MRN:       12345678,
				FirstName: "Jane",
				LastName:  "Doe",
				HealthData: HealthData{
					HR:          71,
					Spo2:        98,
					RR:          16,
					LastUpdated: 1676351833,
				},
				ECA: "NICU",
				Alarms: []Alarm{
					{
						AlarmName:   "alarm",
						Severity:    "critical",
						Description: "Air Gas sensor fault",
						CreatedAt:   1516325422,
					},
					{
						AlarmName:   "alarm",
						Severity:    "critical",
						Description: "Power Disconnected",
						CreatedAt:   1516325422,
					},
				},
			},
			{
				MRN:       12345678,
				FirstName: "Smith",
				LastName:  "John",
				HealthData: HealthData{
					HR:          71,
					Spo2:        98,
					RR:          16,
					LastUpdated: 1676351833,
				},
				ECA: "ICU",
				Alarms: []Alarm{
					{
						AlarmName:   "alarm",
						Severity:    "critical",
						Description: "Air Gas sensor fault",
						CreatedAt:   1516325422,
					},
				},
			},
		},
	}

	response, err := json.Marshal(patientList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
