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

type Flagrule map[string]interface{}

type Audset struct {
	Combine    string     `json:"combine"`
	Conditions []CondInst `json:"conditions"`
}

type CondInst struct {
	*ConditionEmbedded
	Vals []string `json:"vals"`
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
		db.Preload("Audiences").First(&flag, flags[ind].ID)
		flagrule["status"] = flag.Status

		for i, _ := range flag.Audiences {
			flagrule[flag.Audiences[i].Key] = *buildAudrule(flag.Audiences[i], db)
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

func getEmbeddedConds(aud Audience, db *gorm.DB) []CondInst {
	conds := []CondInst{}
	for ind, _ := range aud.Conditions {
		cond := aud.Conditions[ind]
		var attr Attribute
		db.Find(&attr, cond.AttributeID)
		db.Find(&cond)
		cond.Attribute = attr
		conds = append(conds, CondInst{
			ConditionEmbedded: &ConditionEmbedded{
				Condition: &cond,
				Attribute: attr.Key,
			},
			Vals: strings.Split(cond.Vals, ", "),
		})
	}
	return conds
}
