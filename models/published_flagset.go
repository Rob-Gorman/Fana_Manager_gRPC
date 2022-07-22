package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Flagset struct {
	Sdkkeys map[string]bool     `json:"sdkKeys"`
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

func BuildFlagset(db *gorm.DB) (fs *Flagset) {
	sdks := buildSdkkeys(db)
	flrules := buildFlagrules(db)

	fs = &Flagset{Sdkkeys: *sdks, Flags: flrules}
	result, _ := json.Marshal(&fs)
	fmt.Println(string(result))
	return fs
}

func buildSdkkeys(db *gorm.DB) *map[string]bool {
	type SDK struct {
		Key    string
		Status bool
	}
	r := map[string]bool{}

	var sdks []SDK
	db.Model(&Sdkkey{}).Select("key", "status").Find(&sdks)

	for i, _ := range sdks {
		r[sdks[i].Key] = sdks[i].Status
	}

	return &r
}

func buildFlagrules(db *gorm.DB) (frs map[string]Flagrule) {
	// type FlagMsg struct {
	// 	ID     uint
	// 	Status bool
	// 	// Audiences []Audience `gorm:"foreignKey:AudienceID; references:ID"`
	// }
	var flags []Flag
	frs = map[string]Flagrule{}
	db.Model(&Flag{}).Select("id", "status").Find(&flags)

	for ind, _ := range flags {
		flag := Flag{}
		flagrule := Flagrule{}
		audiences := []Audset{}
		db.Preload("Audiences").First(&flag, flags[ind].ID)
		for i, _ := range flag.Audiences {
			audiences = append(audiences, *buildAudrule(db, flag.Audiences[i]))
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

func buildAudrule(db *gorm.DB, aud Audience) (ar *Audset) {
	db.Preload("Conditions").First(&aud)
	conditions := []CondInst{}
	for ind, _ := range aud.Conditions {
		conditions = append(conditions, *buildCondinst(db, aud.Conditions[ind]))
	}
	ar = &Audset{
		Combine:    aud.Combine,
		Conditions: conditions,
	}
	return ar
}

func buildCondinst(db *gorm.DB, cond Condition) (ci *CondInst) {

	db.Preload("Attribute").First(&cond)
	return &CondInst{
		Attribute: cond.Attribute.Key,
		Operator:  cond.Operator,
		Vals:      strings.Split(cond.Vals, " "),
		Negate:    cond.Negate,
	}
}
