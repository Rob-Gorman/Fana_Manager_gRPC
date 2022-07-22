package models

import (
	"encoding/json"
	"fmt"
	"manager/db"
	"strings"
)

type Flagset struct {
	Sdkkeys []string            `json:"sdkKeys"`
	Flags   map[string]Flagrule `json:"flags"`
}

type Flagrule struct {
	Status    bool     `json:"status"`
	Audiences []Audset `json:"audiences"`
}

type Audset struct {
	Combine    string     `json:"combine"`
	Conditions []CondInst `json:"conditions"`
}

type CondInst struct {
	Attribute string   `json:"attribute"`
	Operator  string   `json:"operator"`
	Vals      []string `json:"vals"`
	Negate    bool     `json:"negate"`
}

func BuildFlagset() (fs *Flagset) {
	sdks := buildSdkkeys()
	flrules := buildFlagrules()

	fs = &Flagset{Sdkkeys: *sdks, Flags: flrules}
	result, _ := json.Marshal(&fs)
	fmt.Println(string(result))
	return fs
}

func buildSdkkeys() *[]string {
	type SDK struct {
		Key    string
		Status bool
	}
	r := []string{}

	var sdks []SDK
	db.DB.Model(Sdkkey{}).Select("key", "status").Find(&sdks)

	for i, _ := range sdks {
		r = append(r, sdks[i].Key)
	}

	return &r
}

func buildFlagrules() (frs map[string]Flagrule) {
	// type FlagMsg struct {
	// 	ID     uint
	// 	Status bool
	// 	// Audiences []Audience `gorm:"foreignKey:AudienceID; references:ID"`
	// }
	var flags []Flag
	frs = map[string]Flagrule{}
	db.DB.Model(Flag{}).Select("id", "status").Find(&flags)

	for ind, _ := range flags {
		flag := Flag{}
		flagrule := Flagrule{}
		audiences := []Audset{}
		db.DB.Preload("Audiences").First(&flag, flags[ind].ID)
		for i, _ := range flag.Audiences {
			audiences = append(audiences, *buildAudrule(flag.Audiences[i]))
		}
		flagrule = Flagrule{
			Status:    flag.Status,
			Audiences: audiences,
		}

		frs[flags[ind].Key] = flagrule

		return frs
	}

	result, _ := json.Marshal(&flags)
	fmt.Println("FLAGS AS BUILDING", string(result))

	return frs
}

func buildAudrule(aud Audience) (ar *Audset) {
	db.DB.Preload("Conditions").First(&aud)
	conditions := []CondInst{}
	for ind, _ := range aud.Conditions {
		conditions = append(conditions, *buildCondinst(aud.Conditions[ind]))
	}
	ar = &Audset{
		Combine:    aud.Combine,
		Conditions: conditions,
	}
	return ar
}

func buildCondinst(cond Condition) (ci *CondInst) {

	db.DB.Preload("Attribute").First(&cond)
	return &CondInst{
		Attribute: cond.Attribute.Key,
		Operator:  cond.Operator,
		Vals:      strings.Split(cond.Vals, " "),
		Negate:    cond.Negate,
	}
}
