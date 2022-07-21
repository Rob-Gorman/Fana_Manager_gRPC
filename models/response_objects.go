package models

// this is some clever shit, I'll be honest. From here:
// https://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type omit bool // this is a custom type to make the code intent more obvious

// what's going on here is that we're using a Flag object we get back
// from our DB. we're then inserting that Flag into _another_ struct.
// By doing that, we can then _shadow_ the fields we want to manipulate
// for the response body. (hence that struct is called `FlagResponse`)
// the `omit` type is basically just an explicit way to set the field to
// `empty` for the json tag `omitempty`. (It could just be a bool (it is just a bool)
// but calling it `omit` is clearer--the value is always false because that's
// the boolean zero-value in Go. What this means is that field is always omitted
// when converting (marshalling) this response struct into JSON. It will
// simply not encode that field (because of `omitempty` specified in the tag)

type FlagResponse struct {
	*Flag
	SdkKey    omit `json:"sdkKey,omitempty"`
	DeletedAt omit `json:"deleted_at,omitempty"`
}

type FlagNoAudsResponse struct {
	*Flag
	SdkKey    omit `json:"sdkKey,omitempty"`
	DeletedAt omit `json:"deleted_at,omitempty"`
	Audiences omit `json:"audiences,omitempty"`
}

type AudienceResponse struct {
	*Audience
	Flags     omit `json:"flags,omitempty"`
	DeletedAt omit `json:"deleted_at,omitempty"`
}

type AudienceNoCondsResponse struct {
	*Audience
	Flags      omit `json:"flags,omitempty"`
	DeletedAt  omit `json:"deleted_at,omitempty"`
	Conditions omit `json:"conditions,omitempty"`
}
