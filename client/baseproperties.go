package client

import (
	"github.com/contactlab/contacthub-sdk-go/enums"
	"github.com/guregu/null"
)

// BaseProperties represent the base properties of a Customer.
type BaseProperties struct {
	PictureURL    null.String    `json:"pictureUrl,omitempty"`
	Title         null.String    `json:"title,omitempty"`
	Prefix        null.String    `json:"prefix,omitempty"`
	FirstName     null.String    `json:"firstName,omitempty"`
	LastName      null.String    `json:"lastName,omitempty"`
	MiddleName    null.String    `json:"middleName,omitempty"`
	Gender        null.String    `json:"gender,omitempty"`
	Dob           SimpleDate     `json:"dob,omitempty"`
	Locale        null.String    `json:"locale,omitempty"`
	TimeZone      null.String    `json:"timezone,omitempty"`
	Contacts      *Contacts      `json:"contacts,omitempty"`
	Address       *Address       `json:"address,omitempty"`
	Credential    *Credential    `json:"credential,omitempty"`
	Educations    []Education    `json:"education,omitempty"`
	Likes         []Like         `json:"likes,omitempty"`
	SocialProfile *SocialProfile `json:"socialProfile,omitempty"`
	Jobs          []Job          `json:"jobs,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// Contacts are the contact info of a Customer.
type Contacts struct {
	Email         null.String    `json:"email,omitempty"`
	Fax           null.String    `json:"fax,omitempty"`
	MobilePhone   null.String    `json:"mobilePhone,omitempty"`
	Phone         null.String    `json:"Phone,omitempty"`
	OtherContacts []OtherContact `json:"otherContacts,omitempty"`
	MobileDevices []MobileDevice `json:"mobileDevices,omitempty"`
}

// OtherContact is a generic contact info of a Customer.
type OtherContact struct {
	Name  string            `json:"name"`
	Type  enums.ContactType `json:"type,omitempty"`
	Value string            `json:"valuetype"`
}

// MobileDevice contains info about a mobile device of a Customer.
type MobileDevice struct {
	Identifier          string                        `json:"identifier"`
	AppID               string                        `json:"appId"`
	Name                string                        `json:"name"`
	Type                enums.MobileDeviceType        `json:"type,omitempty"`
	NotificationService enums.NotificationServiceType `json:"notificationService,omitempty"`
}

// Address contains an address of a Customer.
type Address struct {
	Street   null.String `json:"street,omitempty"`
	City     null.String `json:"city,omitempty"`
	Country  null.String `json:"country,omitempty"`
	Province null.String `json:"province,omitempty"`
	Zip      null.String `json:"zip,omitempty"`
	Geo      Geo         `json:"geo,omitempty"`
}

// Geo contains the coordinate of an Address.
type Geo struct {
	Lat null.Float `json:"lat,omitempty"`
	Lon null.Float `json:"lon,omitempty"`
}

// Credential contains the credentials of a Customer.
// I don't think storing the password would be a good idea
type Credential struct {
	Password null.String `json:"password,omitempty"`
	Username null.String `json:"username,omitempty"`
}

// Education contains the Education info of a Customer
type Education struct {
	ID                  string           `json:"id,required"`
	SchoolType          enums.SchoolType `json:"schoolType,omitempty"`
	SchoolName          null.String      `json:"schoolName,omitempty"`
	SchoolConcentration null.String      `json:"schoolConcentration,omitempty"`
	StartYear           null.Int         `json:"startYear,omitempty"`
	EndYear             null.Int         `json:"endYear,omitempty"`
	IsCurrent           null.Bool        `json:"isCurrent,omitempty"`
}

// Like represents a thing the Customer liked
type Like struct {
	ID          string      `json:"id,required"`
	Category    null.String `json:"category,omitempty"`
	Name        null.String `json:"name,omitempty"`
	CreatedTime CustomDate  `json:"createdTime,omitempty"`
}

// SocialProfile contains all social profile of the Customer
type SocialProfile struct {
	Facebook  null.String `json:"facebook,omitempty"`
	Google    null.String `json:"google,omitempty"`
	Instagram null.String `json:"instagram,omitempty"`
	Linkedin  null.String `json:"linkedin,omitempty"`
	Qzone     null.String `json:"qzone,omitempty"`
	Twitter   null.String `json:"twitter,omitempty"`
}

// Job contains info about the Customer job
type Job struct {
	ID              string      `json:"id,required"`
	CompanyIndustry null.String `json:"companyIndustry,omitempty"`
	CompanyName     null.String `json:"companyName,omitempty"`
	JobTitle        null.String `json:"jobTitle,omitempty"`
	StartDate       SimpleDate  `json:"startDate,omitempty"`
	EndDate         SimpleDate  `json:"endDate,omitempty"`
	IsCurrent       null.Bool   `json:"isCurrent,omitempty"`
}

// Subscription contains info about the Customer subscriptions
type Subscription struct {
	ID           string                 `json:"id,required"`
	Name         null.String            `json:"name,omitempty"`
	Type         null.String            `json:"type,omitempty"`
	Kind         enums.SubscriptionKind `json:"kind,omitempty"`
	Subscribed   null.Bool              `json:"subscribed,omitempty"`
	StartDate    SimpleDate             `json:"startDate,omitempty"`
	EndDate      SimpleDate             `json:"endDate,omitempty"`
	SubscriberID null.String            `json:"subscriberId,required"`
	RegisteredAt CustomDate             `json:"registeredAt,omitempty"`
	UpdatedAt    CustomDate             `json:"updatedAt,omitempty"`
	Preferences  map[string]string      `json:"preferences,omitempty"`
}
