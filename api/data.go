package api

import "github.com/gorilla/mux"

type Value map[string]bool

type Flags map[string]Value

var Rules = Flags{
	"style": {"userabc": true, "user123": false},
	"title": {"user123": false, "userabc": true},
}

type Server struct {
	*mux.Router
	Ruleset *Flags
}
