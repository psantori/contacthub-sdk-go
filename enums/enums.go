package enums

type ContactType int
type MobileDeviceType int
type NotificationServiceType int
type SchoolType int
type SubscriptionKind int
type EventType int

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
