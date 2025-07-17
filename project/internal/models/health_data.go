package models

type HealthData struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	RDW    float64 `json:"rdw,omitempty"`
	WBC    float64 `json:"wbc,omitempty"`
	RBC    float64 `json:"rbc,omitempty"`
	HGB    float64 `json:"hgb,omitempty"`
	HCT    float64 `json:"hct,omitempty"`
	MCV    float64 `json:"mcv,omitempty"`
	MCH    float64 `json:"mch,omitempty"`
	MCHC   float64 `json:"mchc,omitempty"`
	PLT    float64 `json:"plt,omitempty"`
	NEU    float64 `json:"neu,omitempty"`
	EOS    float64 `json:"eos,omitempty"`
	BAS    float64 `json:"bas,omitempty"`
	LYM    float64 `json:"lym,omitempty"`
	MON    float64 `json:"mon,omitempty"`
	SOE    float64 `json:"soe,omitempty"`
	CHOL   float64 `json:"chol,omitempty"`
	GLU    float64 `json:"glu,omitempty"`
	HDL    float64 `json:"hdl,omitempty"`
	TG     float64 `json:"tg,omitempty"`
	CRP    float64 `json:"crp,omitempty"`
}
