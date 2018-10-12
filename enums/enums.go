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

// Those enums are generated by jsonenums, as a nice alternative to manually set and check the string values
// Note that during marshaling/unmarshaling type errors are logged

type ContactType int
type MobileDeviceType int
type NotificationServiceType int
type SchoolType int
type SubscriptionKind int
type EventType int
type BringBackPropertyType int
type EventContext int

const (
	Mobile ContactType = iota
	Phone
	Email
	Fax
	OtherContact
)
const (
	IOS MobileDeviceType = iota
	Android
	WindowsPhone
	FireOS
)

const (
	APN NotificationServiceType = iota
	GCM
	WNS
	ADM
	SNS
)

const (
	PrimarySchool SchoolType = iota
	SecondarySchool
	HighSchool
	College
	OtherSchool
)

const (
	DigitalMessage SubscriptionKind = iota
	Service
	OtherSubscription
)

const (
	AbandonedCart EventType = iota
	AddedCompare
	AddedProduct
	AddedWishlist
	CampaignBlacklisted
	CampaignBounced
	CampaignLinkClicked
	CampaignMarkedSpam
	CampaignOpened
	CampaignSent
	CampaignSubscribed
	CampaignUnsubscribed
	ChangedSetting
	ClickedLink
	ClosedTicket
	CompletedOrder
	EventConfirmed
	EventDeclined
	EventEligible
	EventInvited
	EventNotShow
	EventNotInvited
	EventParticipated
	FormCompiled
	GenericActiveEvent
	GenericPassiveEvent
	LoggedIn
	LoggedOut
	OpenedTicket
	OrderShipped
	RemovedCompare
	RemovedProduct
	RemovedWishlist
	RepliedTicket
	ReviewedProduct
	Searched
	ServiceSubscribed
	ServiceUnsubscribed
	ViewedPage
	ViewedProduct
)

const (
	SessionId BringBackPropertyType = iota
	ExternalId
)

const (
	Web EventContext = iota
	Ecommerce
	Retail
	Social
	DigitalCampaign
	ContactCenter
	IOT
	Other
	MobileCtx
)
