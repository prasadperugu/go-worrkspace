package main

import (
	"encoding/json"
	"net/http"
)

type Orgs struct {
	Devices []Device `json:"Orgs"`
}

type Device struct {
	SerialNumber string  `json:"serialNumber"`
	CareArea     string  `json:"careArea"`
	Connectivity bool    `json:"connectivity"`
	Therapy      bool    `json:"therapy"`
	Alarms       []Alarm `json:"alarms"`
}

type Alarm struct {
	AlarmName   string `json:"alarmName"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
}

func main() {
	http.HandleFunc("/devices", getDevices)
	http.ListenAndServe(":8080", nil)
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	devices := []Device{
		{
			SerialNumber: "HVM23456789-T",
			CareArea:     "NICU",
			Connectivity: false,
			Therapy:      false,
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
	}

	orgs := Orgs{
		Devices: devices,
	}

	json.NewEncoder(w).Encode(orgs)
}
