package common

import (
	"encoding/json"

	"github.com/bcc-code/mediabank-bridge/log"
	"gopkg.in/guregu/null.v4"
)

// LocaleString is a map of strings to nullable strings
type LocaleString LocaleMap[null.String]
type nnLocaleString LocaleMap[string]

// Get from a translation map based on the fallbacks
func (localeString LocaleString) Get(languages []string) string {
	languages = append(languages, DefaultLanguages...) // We force the DefaultLanguages as the last languages regardless if they have been specified before already
	for _, l := range languages {
		if val, ok := localeString[l]; ok && val.Valid {
			return val.String
		}
	}

	//log.L.Warn().
	//	Strs("languages", languages).
	//	Str("localeString", spew.Sdump(localeString)).
	//	Msg("Failed to find any localeString for specified languages")

	return ""
}

// GetValueOrNil returns either the value for selected languages or nil
func (localeString LocaleString) GetValueOrNil(languages []string) *string {
	languages = append(languages, DefaultLanguages...) // We force the DefaultLanguages as the last languages regardless if they have been specified before already
	for _, l := range languages {
		if val, ok := localeString[l]; ok && val.Valid {
			return &val.String
		}
	}

	return nil
}

// Any returns true if there's any valid value in the map
func (localeString LocaleString) Any() bool {
	for _, v := range localeString {
		if v.Valid && v.String != "" {
			return true
		}
	}
	return false
}

func (localeString LocaleString) Has(key string) bool {
	_, ok := localeString[key]
	return ok
}

// AsJSON returns the values as a JSON string
//
// For example:
// ```
//
//	{
//		"de": "Wilkommen",
//	 "en": "Welcome",
//	}
func (localeString LocaleString) AsJSON() []byte {
	out := nnLocaleString{}

	for l, s := range localeString {
		if !s.Valid {
			continue
		}

		out[l] = s.String
	}

	j, err := json.Marshal(out)
	if err != nil {
		log.L.Error().Err(err).Send()
	}

	return j
}
