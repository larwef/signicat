package signicat

import (
	"fmt"
	"net/http"
	"time"
)

// SignatureService handles communication with the Signature API.
type SignatureService service

// CreateDocument creates a new document. In the response you will receive a document ID to retrieve info about the document at a
// later time. You also receive a URL and unique identifier per signer.
func (s *SignatureService) CreateDocument(createReq *CreateDocumentRequest) (*Document, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/signature/documents", createReq)
	if err != nil {
		return nil, err
	}

	var response *Document
	if err := s.client.Do(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveDocument retrieves details of a single document.
func (s *SignatureService) RetrieveDocument(documentID string) (*Document, error) {
	u := fmt.Sprintf("/signature/documents/%s", documentID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var response *Document
	if err := s.client.Do(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveDocumentStatus gets the status of a document.
func (s *SignatureService) RetrieveDocumentStatus(documentID string) (*Status, error) {
	u := fmt.Sprintf("/signature/documents/%s/status", documentID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var response *Status
	if err := s.client.Do(req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateDocumentRequest is ...
type CreateDocumentRequest struct {
	Title          string           `json:"title"`
	Signers        []*SignerRequest `json:"signers"`
	DataToSign     *DataToSign      `json:"dataToSign"`
	ContactDetails *ContactDetails  `json:"contactDetails"`
	ExternalID     string           `json:"externalId"`
	Description    string           `json:"description,omitempty"`
}

// SignerRequest is ...
type SignerRequest struct {
	ExternalSignerID string            `json:"externalSignerId"`
	RedirectSettings *RedirectSettings `json:"redirectSettings"`
	SignatureType    *SignatureType    `json:"signatureType"`
	SignerInfo       *SignerInfo       `json:"signerInfo,omitempty"`
	Notifications    *Notifications    `json:"notifications,omitempty"`
}

// RedirectSettings is ...
type RedirectSettings struct {
	RedirectMode RedirectMode `json:"redirectMode"`
	Domain       string       `json:"domain,omitempty"`
	Error        string       `json:"error,omitempty"`
	Cancel       string       `json:"cancel,omitempty"`
	Success      string       `json:"success,omitempty"`
}

// RedirectMode is ...
type RedirectMode string

// Available redirection modes.
const (
	RedirectModeDoNotRedirect                     RedirectMode = "donot_redirect"
	RedirectModeRedirect                                       = "redirect"
	RedirectModeIframeWithWebMessaging                         = "iframe_with_webmessaging"
	RedirectModeIframeWithRedirect                             = "iframe_with_redirect"
	RedirectModeIframeWithRedirectAndWebMessaging              = "iframe_with_redirect_and_webmessaging"
)

// SignatureType is ...
type SignatureType struct {
	Mechanism Mechanisms `json:"mechanism"`
}

// Mechanisms is ...
type Mechanisms string

// Available signature mechanisms.
const (
	MechanismsPkiSignature                  Mechanisms = "pkisignature"
	MechanismsIdentification                           = "identification"
	MechanismsHandwritten                              = "handwritten"
	MechanismsHandWrittenWithIdentification            = "handwritten_with_identification"
)

// SignerInfo is ...
type SignerInfo struct {
	FirstName            string            `json:"firstName,omitempty"`
	LastName             string            `json:"lastName,omitempty"`
	Email                string            `json:"email,omitempty"`
	SocialSecurityNumber string            `json:"socialSecurityNumber,omitempty"`
	Mobile               *Mobile           `json:"mobile,omitempty"`
	OrganizationInfo     *OrganizationInfo `json:"organizationInfo,omitempty"`
}

// Mobile is ...
type Mobile struct {
	CountryCode string `json:"countryCode,omitempty"`
	Number      string `json:"number,omitempty"`
}

// OrganizationInfo is ...
type OrganizationInfo struct {
	OrgNo       string `json:"orgNo,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// DataToSign is ...
type DataToSign struct {
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Base64Content string `json:"base64Content"`
	FileName      string `json:"fileName"`
	ConvertToPDF  bool   `json:"convertToPdf,omitempty"`
}

// ContactDetails is ...
type ContactDetails struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email"`
	URL   string `json:"url,omitempty"`
}

// Notifications is ...
type Notifications struct {
	Setup *Setup `json:"setup,omitempty"`
}

// Setup is ...
type Setup struct {
	Request          NotificationSetup `json:"request,omitempty"`
	Reminder         NotificationSetup `json:"reminder,omitempty"`
	SignatureReceipt NotificationSetup `json:"signatureReceipt,omitempty"`
	FinalReceipt     NotificationSetup `json:"finalReceipt,omitempty"`
	Canceled         NotificationSetup `json:"canceled,omitempty"`
	Expired          NotificationSetup `json:"expired,omitempty"`
}

// NotificationSetup is ...
type NotificationSetup string

// Available notification setups.
const (
	NotificationSetupOff       NotificationSetup = "off"
	NotificationSetupSendSms                     = "sendSms"
	NotificationSetupSendEmail                   = "sendEmail"
	NotificationSetupSendBoth                    = "sendBoth"
)

// Document is ...
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

// SignerResponse is ...
type SignerResponse struct {
	ID                      string             `json:"id,omitempty"`
	URL                     string             `json:"url,omitempty"`
	DocumentSignature       *DocumentSignature `json:"documentSignature,omitempty"`
	ExternalSignerID        string             `json:"externalSignerId,omitempty"`
	RedirectSettings        *RedirectSettings  `json:"redirectSettings,omitempty"`
	SignatureType           *SignatureType     `json:"signatureType,omitempty"`
	SignerInfo              *SignerInfo        `json:"signerInfo,omitempty"`
	Notifications           *Notifications     `json:"notifications,omitempty"`
	Order                   int32              `json:"order,omitempty"`
	Required                bool               `json:"required,omitempty"`
	SignURLExpires          *time.Time         `json:"signUrlExpires,omitempty"`
	GetSocialSecurityNumber bool               `json:"getSocialSecurityNumber,omitempty"`
}

// DocumentSignature is ...
type DocumentSignature struct {
	SignatureMethod         SignatureMethod       `json:"signatureMethod"`
	FullName                string                `json:"fullName,omitempty"`
	FirstName               string                `json:"firstName,omitempty"`
	LastName                string                `json:"lastName,omitempty"`
	MiddleName              string                `json:"middleName,omitempty"`
	SignedTime              *time.Time            `json:"signedTime,omitempty"`
	DateOfBirth             string                `json:"dateOfBirth,omitempty"`
	SignatureMethodUniqueID string                `json:"signatureMethodUniqueId,omitempty"`
	SocialSecurityNumber    *SocialSecurityNumber `json:"socialSecurityNumber,omitempty"`
	ClientIP                string                `json:"clientIp,omitempty"`
	Mechanism               Mechanisms            `json:"mechanism,omitempty"`
	PersonalInfoOrigin      PersonalInfoOrigin    `json:"personalInfoOrigin,omitempty"`
}

// SignatureMethod is ...
type SignatureMethod string

// Available signature methods.
const (
	SignatureMethodNoBankIDMobile     SignatureMethod = "no_bankid_mobile"
	SignatureMethodNoBankIDNetCentric                 = "no_bankid_netcentric"
	SignatureMethodNoBuypass                          = "no_buypass"
	SignatureMethodSeBankID                           = "se_bankid"
	SignatureMethodDkNemID                            = "dk_nemid"
	SignatureMethodFiTupas                            = "fi_tupas"
	SignatureMethodFiMobiilivarmenne                  = "fi_mobiilivarmenne"
	SignatureMethodFiEid                              = "fi_eid"
	SignatureMethodSmsOtp                             = "sms_otp"
	SignatureMethodUnknown                            = "unknown"
)

// SocialSecurityNumber is ...
type SocialSecurityNumber struct {
	Value       string `json:"value,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// PersonalInfoOrigin is ...
type PersonalInfoOrigin string

// Avalable options for personalInfoOrigin field.
const (
	PersonalInfoOriginUnknown       PersonalInfoOrigin = "unknown"
	PersonalInfoOriginEid                              = "eid"
	PersonalInfoOriginUserFormInput                    = "userFormInput"
)

// Status is ...
type Status struct {
	DocumentStatus    DocumentStatus `json:"documentStatus,omitempty"`
	CompletedPackages []FileFormat   `json:"completedPackages,omitempty"`
}

// DocumentStatus is ...
type DocumentStatus string

// Available document statuses.
const (
	DocumentStatusUnsigned              DocumentStatus = "unsigned"
	DocumentStatusWaitingForAttachments                = "waiting_for_attachments"
	DocumentStatusPartialSigned                        = "partialsigned"
	DocumentStatusSigned                               = "signed"
	DocumentStatusCanceled                             = "canceled"
	DocumentStatusExpired                              = "expired"
)

// FileFormat is ...
type FileFormat string

// Available file formats.
const (
	FileFormatUnsigned          FileFormat = "unsigned"
	FileFormatNative                       = "native"
	FileFormatStandardPackaging            = "standard_packaging"
	FileFormatPades                        = "pades"
	FileFormatXades                        = "xades"
)
