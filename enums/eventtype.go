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
	_EventTypeNameToValue = map[string]EventType{
		"abandonedCart":        AbandonedCart,
		"addedCompare":         AddedCompare,
		"addedProduct":         AddedProduct,
		"addedWishlist":        AddedWishlist,
		"campaignBlacklisted":  CampaignBlacklisted,
		"campaignBounced":      CampaignBounced,
		"campaignLinkClicked":  CampaignLinkClicked,
		"campaignMarkedSpam":   CampaignMarkedSpam,
		"campaignOpened":       CampaignOpened,
		"campaignSent":         CampaignSent,
		"campaignSubscribed":   CampaignSubscribed,
		"campaignUnsubscribed": CampaignUnsubscribed,
		"changedSetting":       ChangedSetting,
		"clickedLink":          ClickedLink,
		"closedTicket":         ClosedTicket,
		"completedOrder":       CompletedOrder,
		"eventConfirmed":       EventConfirmed,
		"eventDeclined":        EventDeclined,
		"eventEligible":        EventEligible,
		"eventInvited":         EventInvited,
		"eventNotShow":         EventNotShow,
		"eventNotInvited":      EventNotInvited,
		"eventParticipated":    EventParticipated,
		"formCompiled":         FormCompiled,
		"genericActiveEvent":   GenericActiveEvent,
		"genericPassiveEvent":  GenericPassiveEvent,
		"loggedIn":             LoggedIn,
		"loggedOut":            LoggedOut,
		"openedTicket":         OpenedTicket,
		"orderShipped":         OrderShipped,
		"removedCompare":       RemovedCompare,
		"removedProduct":       RemovedProduct,
		"removedWishlist":      RemovedWishlist,
		"repliedTicket":        RepliedTicket,
		"reviewedProduct":      ReviewedProduct,
		"searched":             Searched,
		"serviceSubscribed":    ServiceSubscribed,
		"serviceUnsubscribed":  ServiceUnsubscribed,
		"viewedPage":           ViewedPage,
		"viewedProduct":        ViewedProduct,
	}

	_EventTypeValueToName = map[EventType]string{
		AbandonedCart:        "abandonedCart",
		AddedCompare:         "addedCompare",
		AddedProduct:         "addedProduct",
		AddedWishlist:        "addedWishlist",
		CampaignBlacklisted:  "campaignBlacklisted",
		CampaignBounced:      "campaignBounced",
		CampaignLinkClicked:  "campaignLinkClicked",
		CampaignMarkedSpam:   "campaignMarkedSpam",
		CampaignOpened:       "campaignOpened",
		CampaignSent:         "campaignSent",
		CampaignSubscribed:   "campaignSubscribed",
		CampaignUnsubscribed: "campaignUnsubscribed",
		ChangedSetting:       "changedSetting",
		ClickedLink:          "clickedLink",
		ClosedTicket:         "closedTicket",
		CompletedOrder:       "completedOrder",
		EventConfirmed:       "eventConfirmed",
		EventDeclined:        "eventDeclined",
		EventEligible:        "eventEligible",
		EventInvited:         "eventInvited",
		EventNotShow:         "eventNotShow",
		EventNotInvited:      "eventNotInvited",
		EventParticipated:    "eventParticipated",
		FormCompiled:         "formCompiled",
		GenericActiveEvent:   "genericActiveEvent",
		GenericPassiveEvent:  "genericPassiveEvent",
		LoggedIn:             "loggedIn",
		LoggedOut:            "loggedOut",
		OpenedTicket:         "openedTicket",
		OrderShipped:         "orderShipped",
		RemovedCompare:       "removedCompare",
		RemovedProduct:       "removedProduct",
		RemovedWishlist:      "removedWishlist",
		RepliedTicket:        "repliedTicket",
		ReviewedProduct:      "reviewedProduct",
		Searched:             "searched",
		ServiceSubscribed:    "serviceSubscribed",
		ServiceUnsubscribed:  "serviceUnsubscribed",
		ViewedPage:           "viewedPage",
		ViewedProduct:        "viewedProduct",
	}
)

func init() {
	var v EventType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_EventTypeNameToValue = map[string]EventType{
			interface{}(AbandonedCart).(fmt.Stringer).String():        AbandonedCart,
			interface{}(AddedCompare).(fmt.Stringer).String():         AddedCompare,
			interface{}(AddedProduct).(fmt.Stringer).String():         AddedProduct,
			interface{}(AddedWishlist).(fmt.Stringer).String():        AddedWishlist,
			interface{}(CampaignBlacklisted).(fmt.Stringer).String():  CampaignBlacklisted,
			interface{}(CampaignBounced).(fmt.Stringer).String():      CampaignBounced,
			interface{}(CampaignLinkClicked).(fmt.Stringer).String():  CampaignLinkClicked,
			interface{}(CampaignMarkedSpam).(fmt.Stringer).String():   CampaignMarkedSpam,
			interface{}(CampaignOpened).(fmt.Stringer).String():       CampaignOpened,
			interface{}(CampaignSent).(fmt.Stringer).String():         CampaignSent,
			interface{}(CampaignSubscribed).(fmt.Stringer).String():   CampaignSubscribed,
			interface{}(CampaignUnsubscribed).(fmt.Stringer).String(): CampaignUnsubscribed,
			interface{}(ChangedSetting).(fmt.Stringer).String():       ChangedSetting,
			interface{}(ClickedLink).(fmt.Stringer).String():          ClickedLink,
			interface{}(ClosedTicket).(fmt.Stringer).String():         ClosedTicket,
			interface{}(CompletedOrder).(fmt.Stringer).String():       CompletedOrder,
			interface{}(EventConfirmed).(fmt.Stringer).String():       EventConfirmed,
			interface{}(EventDeclined).(fmt.Stringer).String():        EventDeclined,
			interface{}(EventEligible).(fmt.Stringer).String():        EventEligible,
			interface{}(EventInvited).(fmt.Stringer).String():         EventInvited,
			interface{}(EventNotShow).(fmt.Stringer).String():         EventNotShow,
			interface{}(EventNotInvited).(fmt.Stringer).String():      EventNotInvited,
			interface{}(EventParticipated).(fmt.Stringer).String():    EventParticipated,
			interface{}(FormCompiled).(fmt.Stringer).String():         FormCompiled,
			interface{}(GenericActiveEvent).(fmt.Stringer).String():   GenericActiveEvent,
			interface{}(GenericPassiveEvent).(fmt.Stringer).String():  GenericPassiveEvent,
			interface{}(LoggedIn).(fmt.Stringer).String():             LoggedIn,
			interface{}(LoggedOut).(fmt.Stringer).String():            LoggedOut,
			interface{}(OpenedTicket).(fmt.Stringer).String():         OpenedTicket,
			interface{}(OrderShipped).(fmt.Stringer).String():         OrderShipped,
			interface{}(RemovedCompare).(fmt.Stringer).String():       RemovedCompare,
			interface{}(RemovedProduct).(fmt.Stringer).String():       RemovedProduct,
			interface{}(RemovedWishlist).(fmt.Stringer).String():      RemovedWishlist,
			interface{}(RepliedTicket).(fmt.Stringer).String():        RepliedTicket,
			interface{}(ReviewedProduct).(fmt.Stringer).String():      ReviewedProduct,
			interface{}(Searched).(fmt.Stringer).String():             Searched,
			interface{}(ServiceSubscribed).(fmt.Stringer).String():    ServiceSubscribed,
			interface{}(ServiceUnsubscribed).(fmt.Stringer).String():  ServiceUnsubscribed,
			interface{}(ViewedPage).(fmt.Stringer).String():           ViewedPage,
			interface{}(ViewedProduct).(fmt.Stringer).String():        ViewedProduct,
		}
	}
}

// MarshalJSON is generated so EventType satisfies json.Marshaler.
func (r EventType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _EventTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid EventType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so EventType satisfies json.Unmarshaler.
func (r *EventType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EventType should be a string, got %s", data)
	}
	v, ok := _EventTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid EventType %q", s)
	}
	*r = v
	return nil
}
