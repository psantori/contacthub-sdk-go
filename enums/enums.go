package enums

type ContactType int
type MobileDeviceType int
type NotificationServiceType int
type SchoolType int
type SubscriptionKind int

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
