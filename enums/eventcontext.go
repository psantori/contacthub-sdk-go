/**
 * This file is part of contacthub-sdk-go.
 *
 * contacthub-sdk-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * contacthub-sdk-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with contacthub-sdk-go. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2017 Arduino AG
 *
 * @author Luca Osti
 *
 */

package enums

import (
	"encoding/json"
	"fmt"
)

var (
	_EventContextNameToValue = map[string]EventContext{
		"WEB":              Web,
		"MOBILE":           MobileCtx,
		"ECOMMERCE":        Ecommerce,
		"RETAIL":           Retail,
		"SOCIAL":           Social,
		"DIGITAL_CAMPAIGN": DigitalCampaign,
		"CONTACT_CENTER":   ContactCenter,
		"IOT":              IOT,
		"OTHER":            Other,
	}

	_EventContextValueToName = map[EventContext]string{
		Web:             "WEB",
		MobileCtx:       "MOBILE",
		Ecommerce:       "ECOMMERCE",
		Retail:          "RETAIL",
		Social:          "SOCIAL",
		DigitalCampaign: "DIGITAL_CAMPAIGN",
		ContactCenter:   "CONTACT_CENTER",
		IOT:             "IOT",
		Other:           "OTHER",
	}
)

func init() {
	var v EventContext
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_EventContextNameToValue = map[string]EventContext{
			interface{}(Web).(fmt.Stringer).String():             Web,
			interface{}(Ecommerce).(fmt.Stringer).String():       Ecommerce,
			interface{}(Retail).(fmt.Stringer).String():          Retail,
			interface{}(Social).(fmt.Stringer).String():          Social,
			interface{}(DigitalCampaign).(fmt.Stringer).String(): DigitalCampaign,
			interface{}(ContactCenter).(fmt.Stringer).String():   ContactCenter,
			interface{}(IOT).(fmt.Stringer).String():             IOT,
			interface{}(Other).(fmt.Stringer).String():           Other,
		}
	}
}

// MarshalJSON is generated so EventContext satisfies json.Marshaler.
func (r EventContext) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _EventContextValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid EventContext: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so EventContext satisfies json.Unmarshaler.
func (r *EventContext) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EventContext should be a string, got %s", data)
	}
	v, ok := _EventContextNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid EventContext %q", s)
	}
	*r = v
	return nil
}
