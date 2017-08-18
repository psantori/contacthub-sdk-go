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

package client

import (
	"github.com/contactlab/contacthub-sdk-go/enums"
	"github.com/guregu/null"
)

// BaseProperties represent the base properties of a Customer.
type BaseProperties struct {
	PictureURL    *null.String   `json:"pictureUrl,omitempty"`
	Title         *null.String   `json:"title,omitempty"`
	Prefix        *null.String   `json:"prefix,omitempty"`
	FirstName     *null.String   `json:"firstName,omitempty"`
	LastName      *null.String   `json:"lastName,omitempty"`
	MiddleName    *null.String   `json:"middleName,omitempty"`
	Gender        *null.String   `json:"gender,omitempty"`
	Dob           *SimpleDate    `json:"dob,omitempty"`
	Locale        *null.String   `json:"locale,omitempty"`
	TimeZone      *null.String   `json:"timezone,omitempty"`
	Contacts      *Contacts      `json:"contacts,omitempty"`
	Address       *Address       `json:"address,omitempty"`
	Credential    *Credential    `json:"credential,omitempty"`
	Educations    []Education    `json:"educations,omitempty"`
	Likes         []Like         `json:"likes,omitempty"`
	SocialProfile *SocialProfile `json:"socialProfile,omitempty"`
	Jobs          []Job          `json:"jobs,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

type BasePropertiesResponse struct {
	PictureURL    null.String            `json:"pictureUrl,required"`
	Title         null.String            `json:"title,required"`
	Prefix        null.String            `json:"prefix,required"`
	FirstName     null.String            `json:"firstName,required"`
	LastName      null.String            `json:"lastName,required"`
	MiddleName    null.String            `json:"middleName,required"`
	Gender        null.String            `json:"gender,required"`
	Dob           SimpleDate             `json:"dob,required"`
	Locale        null.String            `json:"locale,required"`
	TimeZone      null.String            `json:"timezone,required"`
	Contacts      *ContactsResponse      `json:"contacts,required"`
	Address       *AddressResponse       `json:"address,required"`
	Credential    *CredentialResponse    `json:"credential,required"`
	Educations    []EducationResponse    `json:"educations,required"`
	Likes         []LikeResponse         `json:"likes,required"`
	SocialProfile *SocialProfileResponse `json:"socialProfile,required"`
	Jobs          []JobResponse          `json:"jobs,required"`
	Subscriptions []SubscriptionResponse `json:"subscriptions,required"`
}

// Contacts are the contact info of a Customer.
type Contacts struct {
	Email         *null.String   `json:"email,omitempty"`
	Fax           *null.String   `json:"fax,omitempty"`
	MobilePhone   *null.String   `json:"mobilePhone,omitempty"`
	Phone         *null.String   `json:"phone,omitempty"`
	OtherContacts []OtherContact `json:"otherContacts,omitempty"`
	MobileDevices []MobileDevice `json:"mobileDevices,omitempty"`
}

type ContactsResponse struct {
	Email         null.String            `json:"email,required"`
	Fax           null.String            `json:"fax,required"`
	MobilePhone   null.String            `json:"mobilePhone,required"`
	Phone         null.String            `json:"phone,required"`
	OtherContacts []OtherContactResponse `json:"otherContacts,required"`
	MobileDevices []MobileDeviceResponse `json:"mobileDevices,required"`
}

// OtherContact is a generic contact info of a Customer.
type OtherContact struct {
	Name  string             `json:"name,required"`
	Type  *enums.ContactType `json:"type,omitempty"`
	Value string             `json:"value,required"`
}

type OtherContactResponse struct {
	Name  string            `json:"name,required"`
	Type  enums.ContactType `json:"type,required"`
	Value string            `json:"value,required"`
}

// MobileDevice contains info about a mobile device of a Customer.
type MobileDevice struct {
	Identifier          string                         `json:"identifier,required"`
	AppID               string                         `json:"appId,required"`
	Name                string                         `json:"name,required"`
	Type                *enums.MobileDeviceType        `json:"type,omitempty"`
	NotificationService *enums.NotificationServiceType `json:"notificationService,omitempty"`
}

type MobileDeviceResponse struct {
	Identifier          string                        `json:"identifier,required"`
	AppID               string                        `json:"appId,required"`
	Name                string                        `json:"name,required"`
	Type                enums.MobileDeviceType        `json:"type,required"`
	NotificationService enums.NotificationServiceType `json:"notificationService,required"`
}

// Address contains an address of a Customer.
type Address struct {
	Street   *null.String `json:"street,omitempty"`
	City     *null.String `json:"city,omitempty"`
	Country  *null.String `json:"country,omitempty"`
	Province *null.String `json:"province,omitempty"`
	Zip      *null.String `json:"zip,omitempty"`
	Geo      *Geo         `json:"geo,omitempty"`
}

type AddressResponse struct {
	Street   null.String `json:"street,required"`
	City     null.String `json:"city,required"`
	Country  null.String `json:"country,required"`
	Province null.String `json:"province,required"`
	Zip      null.String `json:"zip,required"`
	Geo      *Geo        `json:"geo,required"`
}

// Geo contains the coordinate of an Address.
type Geo struct {
	Lat *null.Float `json:"lat,omitempty"`
	Lon *null.Float `json:"lon,omitempty"`
}
type GeoResponse struct {
	Lat null.Float `json:"lat,required"`
	Lon null.Float `json:"lon,required"`
}

// Credential contains the credentials of a Customer.
// I don't think storing the password would be a good idea
type Credential struct {
	Password *null.String `json:"password,omitempty"`
	Username *null.String `json:"username,omitempty"`
}

type CredentialResponse struct {
	Password null.String `json:"password,required"`
	Username null.String `json:"username,required"`
}

// Education contains the Education info of a Customer
type Education struct {
	ID                  string            `json:"id,required"`
	SchoolType          *enums.SchoolType `json:"schoolType,omitempty"`
	SchoolName          *null.String      `json:"schoolName,omitempty"`
	SchoolConcentration *null.String      `json:"schoolConcentration,omitempty"`
	StartYear           *null.Int         `json:"startYear,omitempty"`
	EndYear             *null.Int         `json:"endYear,omitempty"`
	IsCurrent           *null.Bool        `json:"isCurrent,omitempty"`
}

type EducationResponse struct {
	ID                  string            `json:"id,required"`
	SchoolType          *enums.SchoolType `json:"schoolType,required"`
	SchoolName          null.String       `json:"schoolName,required"`
	SchoolConcentration null.String       `json:"schoolConcentration,required"`
	StartYear           null.Int          `json:"startYear,required"`
	EndYear             null.Int          `json:"endYear,required"`
	IsCurrent           null.Bool         `json:"isCurrent,required"`
}

// SocialProfile contains all social profile of the Customer
type SocialProfile struct {
	Facebook  *null.String `json:"facebook,omitempty"`
	Google    *null.String `json:"google,omitempty"`
	Instagram *null.String `json:"instagram,omitempty"`
	Linkedin  *null.String `json:"linkedin,omitempty"`
	Qzone     *null.String `json:"qzone,omitempty"`
	Twitter   *null.String `json:"twitter,omitempty"`
}

type SocialProfileResponse struct {
	Facebook  null.String `json:"facebook,required"`
	Google    null.String `json:"google,required"`
	Instagram null.String `json:"instagram,required"`
	Linkedin  null.String `json:"linkedin,required"`
	Qzone     null.String `json:"qzone,required"`
	Twitter   null.String `json:"twitter,required"`
}

// Job contains info about the Customer job
type Job struct {
	ID              string       `json:"id,required"`
	CompanyIndustry *null.String `json:"companyIndustry,omitempty"`
	CompanyName     *null.String `json:"companyName,omitempty"`
	JobTitle        *null.String `json:"jobTitle,omitempty"`
	StartDate       *SimpleDate  `json:"startDate,omitempty"`
	EndDate         *SimpleDate  `json:"endDate,omitempty"`
	IsCurrent       *null.Bool   `json:"isCurrent,omitempty"`
}

type JobResponse struct {
	ID              string      `json:"id,required"`
	CompanyIndustry null.String `json:"companyIndustry,required"`
	CompanyName     null.String `json:"companyName,required"`
	JobTitle        null.String `json:"jobTitle,required"`
	StartDate       SimpleDate  `json:"startDate,required"`
	EndDate         SimpleDate  `json:"endDate,required"`
	IsCurrent       null.Bool   `json:"isCurrent,required"`
}
