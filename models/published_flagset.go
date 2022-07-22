package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
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
	Combine    string              `json:"combine"`
	Conditions []ConditionEmbedded `json:"conditions"`
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

	fs = &Flagset{
		Sdkkeys: *sdks,
		Flags:   flrules,
	}
	result, _ := json.Marshal(&fs)
	fmt.Println(string(result))
	return fs
}

func buildSdkkeys(db *gorm.DB) *[]string {
	var sdks []string
	db.Model(Sdkkey{}).Select("key").Find(&sdks)

	return &sdks
}

func buildFlagrules(db *gorm.DB) (frs map[string]Flagrule) {
	var flags []Flag
	frs = map[string]Flagrule{}
	db.Model(Flag{}).Select("id", "key", "status").Find(&flags)

	for ind, _ := range flags {
		flag := Flag{}
		flagrule := Flagrule{}
		audiences := []Audset{}
		db.Preload("Audiences").First(&flag, flags[ind].ID)
		for i, _ := range flag.Audiences {
			audiences = append(audiences, *buildAudrule(flag.Audiences[i], db))
		}
		flagrule = Flagrule{
			Status:    flag.Status,
			Audiences: audiences,
		}
		printrule, _ := json.Marshal(&flagrule)
		fmt.Printf("%s: %s\n\n\n", flags[ind].Key, string(printrule))

		frs[flags[ind].Key] = flagrule
	}

	result, _ := json.Marshal(&frs)
	fmt.Println("FLAGS AS BUILDING", string(result))

	return frs
}

func buildAudrule(aud Audience, db *gorm.DB) (ar *Audset) {
	db.Preload("Conditions").First(&aud)
	conds := getEmbeddedConds(aud, db)
	ar = &Audset{
		Combine:    aud.Combine,
		Conditions: conds,
	}
	return ar
}

func getEmbeddedConds(aud Audience, db *gorm.DB) (conds []ConditionEmbedded) {
	for ind, _ := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr Attribute
		db.Find(&attr, cond.AttributeID)
		db.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, ConditionEmbedded{
			Condition: &cond,
			Attribute: attr.Key,
			Vals:      strings.Split(cond.Vals, ", "),
		})
	}
	return conds
}
