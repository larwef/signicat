package signature

import (
	"time"
)

func (s *Signature) CreateDocument(req *CreateDocumentRequest) (*Document, error) {
	return nil, ErrNotImplemented
}

func (s *Signature) RetrieveDocument(documentId string) (*Document, error) {
	return nil, ErrNotImplemented
}

func (s *Signature) RetrieveDocumentStatus(documentId string) (*Status, error) {
	return nil, ErrNotImplemented
}

type CreateDocumentRequest struct {
	Title          string           `json:"title"`
	Signers        []*SignerRequest `json:"signers"`
	DataToSign     *DataToSign      `json:"dataToSign"`
	ContactDetails *ContactDetails  `json:"contactDetails"`
	ExternalId     string           `json:"externalId"`
	Description    string           `json:"description,omitempty"`
}

type SignerRequest struct {
	ExternalSignerId string            `json:"externalSignerId"`
	RedirectSettings *RedirectSettings `json:"redirectSettings"`
	SignatureType    *SignatureType    `json:"signatureType"`
	SignerInfo       *SignerInfo       `json:"signerInfo,omitempty"`
	Notifications    *Notifications    `json:"notifications,omitempty"`
}

type RedirectSettings struct {
	RedirectMode RedirectMode `json:"redirectMode"`
	Domain       string       `json:"domain,omitempty"`
	Error        string       `json:"error,omitempty"`
	Cancel       string       `json:"cancel,omitempty"`
	Success      string       `json:"success,omitempty"`
}

type RedirectMode string

const (
	RedirectModeDoNotRedirect                     RedirectMode = "donot_redirect"
	RedirectModeRedirect                                       = "redirect"
	RedirectModeIframeWithWebMessaging                         = "iframe_with_webmessaging"
	RedirectModeIframeWithRedirect                             = "iframe_with_redirect"
	RedirectModeIframeWithRedirectAndWebMessaging              = "iframe_with_redirect_and_webmessaging"
)

type SignatureType struct {
	Mechanism Mechanisms `json:"mechanism"`
}

type Mechanisms string

const (
	MechanismsPkiSignature                  Mechanisms = "pkisignature"
	MechanismsIdentification                           = "identification"
	MechanismsHandwritten                              = "handwritten"
	MechanismsHandWrittenWithIdentification            = "handwritten_with_identification"
)

type SignerInfo struct {
	FirstName            string            `json:"firstName,omitempty"`
	LastName             string            `json:"lastName,omitempty"`
	Email                string            `json:"email,omitempty"`
	SocialSecurityNumber string            `json:"socialSecurityNumber,omitempty"`
	Mobile               *Mobile           `json:"mobile,omitempty"`
	OrganizationInfo     *OrganizationInfo `json:"organizationInfo,omitempty"`
}

type Mobile struct {
	CountryCode string `json:"countryCode,omitempty"`
	Number      string `json:"number,omitempty"`
}

type OrganizationInfo struct {
	OrgNo       string `json:"orgNo,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type DataToSign struct {
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Base64Content string `json:"base64Content"`
	FileName      string `json:"fileName"`
	ConvertToPDF  bool   `json:"convertToPdf,omitempty"`
}

type ContactDetails struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email"`
	URL   string `json:"url,omitempty"`
}

type Notifications struct {
	Setup *Setup `json:"setup,omitempty"`
}

type Setup struct {
	Request          NotificationSetup `json:"request,omitempty"`
	Reminder         NotificationSetup `json:"reminder,omitempty"`
	SignatureReceipt NotificationSetup `json:"signatureReceipt,omitempty"`
	FinalReceipt     NotificationSetup `json:"finalReceipt,omitempty"`
	Canceled         NotificationSetup `json:"canceled,omitempty"`
	Expired          NotificationSetup `json:"expired,omitempty"`
}

type NotificationSetup string

const (
	NotificationSetupOff       NotificationSetup = "off"
	NotificationSetupSendSms                     = "sendSms"
	NotificationSetupSendEmail                   = "sendEmail"
	NotificationSetupSendBoth                    = "sendBoth"
)

type Document struct {
	DocumentID     string            `json:"documentId,omitempty"`
	Signers        []*SignerResponse `json:"signers,omitempty"`
	Status         *Status           `json:"status,omitempty"`
	Title          string            `json:"title,omitempty"`
	Description    string            `json:"description,omitempty"`
	ExternalID     string            `json:"externalId,omitempty"`
	DataToSign     *DataToSign       `json:"dataToSign,omitempty"`
	ContactDetails *ContactDetails   `json:"contactDetails,omitempty"`
}

type SignerResponse struct {
	ID                      string             `json:"id,omitempty"`
	URL                     string             `json:"url,omitempty"`
	DocumentSignature       *DocumentSignature `json:"documentSignature,omitempty"`
	ExternalSignerId        string             `json:"externalSignerId,omitempty"`
	RedirectSettings        *RedirectSettings  `json:"redirectSettings,omitempty"`
	SignatureType           *SignatureType     `json:"signatureType,omitempty"`
	SignerInfo              *SignerInfo        `json:"signerInfo,omitempty"`
	Notifications           *Notifications     `json:"notifications,omitempty"`
	Order                   int32              `json:"order,omitempty"`
	Required                bool               `json:"required,omitempty"`
	SignUrlExpires          *time.Time         `json:"signUrlExpires,omitempty"`
	GetSocialSecurityNumber bool               `json:"getSocialSecurityNumber,omitempty"`
}

type DocumentSignature struct {
	SignatureMethod         SignatureMethod       `json:"signatureMethod"`
	FullName                string                `json:"fullName,omitempty"`
	FirstName               string                `json:"firstName,omitempty"`
	LastName                string                `json:"lastName,omitempty"`
	MiddleName              string                `json:"middleName,omitempty"`
	SignedTime              *time.Time            `json:"signedTime,omitempty"`
	DateOfBirth             string                `json:"dateOfBirth,omitempty"`
	SignatureMethodUniqueId string                `json:"signatureMethodUniqueId,omitempty"`
	SocialSecurityNumber    *SocialSecurityNumber `json:"socialSecurityNumber,omitempty"`
	ClientIp                string                `json:"clientIp,omitempty"`
	Mechanism               Mechanisms            `json:"mechanism,omitempty"`
	PersonalInfoOrigin      PersonalInfoOrigin    `json:"personalInfoOrigin,omitempty"`
}

type SignatureMethod string

const (
	SignatureMethodNoBankIdMobile     SignatureMethod = "no_bankid_mobile"
	SignatureMethodNoBankIdNetCentric                 = "no_bankid_netcentric"
	SignatureMethodNoBuypass                          = "no_buypass"
	SignatureMethodSeBankId                           = "se_bankid"
	SignatureMethodDkNemId                            = "dk_nemid"
	SignatureMethodFiTupas                            = "fi_tupas"
	SignatureMethodFiMobiilivarmenne                  = "fi_mobiilivarmenne"
	SignatureMethodFiEid                              = "fi_eid"
	SignatureMethodSmsOtp                             = "sms_otp"
	SignatureMethodUnknown                            = "unknown"
)

type SocialSecurityNumber struct {
	Value       string `json:"value,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type PersonalInfoOrigin string

const (
	PersonalInfoOriginUnknown       PersonalInfoOrigin = "unknown"
	PersonalInfoOriginEid                              = "eid"
	PersonalInfoOriginUserFormInput                    = "userFormInput"
)

type Status struct {
	DocumentStatus    DocumentStatus `json:"documentStatus,omitempty"`
	CompletedPackages []FileFormat   `json:"completedPackages,omitempty"`
}

type DocumentStatus string

const (
	DocumentStatusUnsigned              DocumentStatus = "unsigned"
	DocumentStatusWaitingForAttachments                = "waiting_for_attachments"
	DocumentStatusPartialSigned                        = "partialsigned"
	DocumentStatusSigned                               = "signed"
	DocumentStatusCanceled                             = "canceled"
	DocumentStatusExpired                              = "expired"
)

type FileFormat string

const (
	FileFormatUnsigned          FileFormat = "unsigned"
	FileFormatNative                       = "native"
	FileFormatStandardPackaging            = "standard_packaging"
	FileFormatPades                        = "pades"
	FileFormatXades                        = "xades"
)
