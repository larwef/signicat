package signicat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Available redirection modes.
const (
	RedirectModeDoNotRedirect                     = "donot_redirect"
	RedirectModeRedirect                          = "redirect"
	RedirectModeIframeWithWebMessaging            = "iframe_with_webmessaging"
	RedirectModeIframeWithRedirect                = "iframe_with_redirect"
	RedirectModeIframeWithRedirectAndWebMessaging = "iframe_with_redirect_and_webmessaging"

	// Available signature mechanisms.
	MechanismsPkiSignature                  = "pkisignature"
	MechanismsIdentification                = "identification"
	MechanismsHandwritten                   = "handwritten"
	MechanismsHandWrittenWithIdentification = "handwritten_with_identification"

	// Available auth mechanism
	AuthMechanismOff          = "off"
	AuthMechanismEid          = "eid"
	AuthMechanismSmsOtp       = "smsOtp"
	AuthMechanismEidAndSmsOtp = "eidAndSmsOtp"

	// Available notification setups.
	NotificationSetupOff       = "off"
	NotificationSetupSendSms   = "sendSms"
	NotificationSetupSendEmail = "sendEmail"
	NotificationSetupSendBoth  = "sendBoth"

	// Available signature methods.
	SignatureMethodNoBankIDMobile     = "no_bankid_mobile"
	SignatureMethodNoBankIDNetCentric = "no_bankid_netcentric"
	SignatureMethodNoBuypass          = "no_buypass"
	SignatureMethodSeBankID           = "se_bankid"
	SignatureMethodDkNemID            = "dk_nemid"
	SignatureMethodFiTupas            = "fi_tupas"
	SignatureMethodFiMobiilivarmenne  = "fi_mobiilivarmenne"
	SignatureMethodFiEid              = "fi_eid"
	SignatureMethodSmsOtp             = "sms_otp"
	SignatureMethodUnknown            = "unknown"

	// Avalable options for personalInfoOrigin field.
	PersonalInfoOriginUnknown       = "unknown"
	PersonalInfoOriginEid           = "eid"
	PersonalInfoOriginUserFormInput = "userFormInput"

	// Available document statuses.
	DocumentStatusUnsigned              = "unsigned"
	DocumentStatusWaitingForAttachments = "waiting_for_attachments"
	DocumentStatusPartialSigned         = "partialsigned"
	DocumentStatusSigned                = "signed"
	DocumentStatusCanceled              = "canceled"
	DocumentStatusExpired               = "expired"

	// Available file formats.
	FileFormatUnsigned          = "unsigned"
	FileFormatNative            = "native"
	FileFormatStandardPackaging = "standard_packaging"
	FileFormatPades             = "pades"
	FileFormatXades             = "xades"

	// Available languages.
	LanguageEnglish   = "EN"
	LanguageNorwegian = "NO"
	LanguageDanish    = "DA"
	LanguageSweedish  = "SV"
	LanguageFinnish   = "FI"
)

// SignatureService handles communication with the Signature API.
type SignatureService service

// CreateDocument creates a new document. In the response you will receive a document ID to retrieve info about the document at a
// later time. You also receive a URL and unique identifier per signer.
func (s *SignatureService) CreateDocument(ctx context.Context, createReq *CreateDocumentRequest) (*Document, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/signature/documents", createReq)
	if err != nil {
		return nil, err
	}

	response := new(Document)
	if err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveDocument retrieves details of a single document.
func (s *SignatureService) RetrieveDocument(ctx context.Context, documentID string) (*Document, error) {
	u := fmt.Sprintf("/signature/documents/%s", documentID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	response := new(Document)
	if err := s.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveDocumentStatus gets the status of a document.
func (s *SignatureService) RetrieveDocumentStatus(ctx context.Context, documentID string) (*Status, error) {
	u := fmt.Sprintf("/signature/documents/%s/status", documentID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	response := new(Status)
	if err := s.client.Do(ctx, req, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RetrieveFile retrieves the signed document file and stored it in the value pointed to by v. v can implement io.Writer. Eg.
// write to a file.
func (s *SignatureService) RetrieveFile(ctx context.Context, documentID, format string, originalFileName bool, v interface{}) error {
	u, err := url.Parse(fmt.Sprintf("/signature/documents/%s/files", documentID))
	if err != nil {
		return err
	}

	params := u.Query()
	params.Set("fileFormat", format)
	params.Set("originalFileName", strconv.FormatBool(originalFileName))
	u.RawQuery = params.Encode()

	req, err := s.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	if err := s.client.Do(ctx, req, v); err != nil {
		return err
	}

	return nil
}

// CreateDocumentRequest is ...
type CreateDocumentRequest struct {
	Title          string           `json:"title"`
	Signers        []*SignerRequest `json:"signers"`
	DataToSign     *DataToSign      `json:"dataToSign"`
	ContactDetails *ContactDetails  `json:"contactDetails"`
	ExternalID     string           `json:"externalId"`
	Description    string           `json:"description,omitempty"`
	Notification   *Notification    `json:"notification"`
}

// SignerRequest is ...
type SignerRequest struct {
	ExternalSignerID string            `json:"externalSignerId"`
	RedirectSettings *RedirectSettings `json:"redirectSettings"`
	SignatureType    *SignatureType    `json:"signatureType"`
	Authentication   *Authentication   `json:"authentication"`
	SignerInfo       *SignerInfo       `json:"signerInfo,omitempty"`
	Notifications    *Notifications    `json:"notifications,omitempty"`
}

// RedirectSettings is ...
type RedirectSettings struct {
	RedirectMode string `json:"redirectMode"`
	Domain       string `json:"domain,omitempty"`
	Error        string `json:"error,omitempty"`
	Cancel       string `json:"cancel,omitempty"`
	Success      string `json:"success,omitempty"`
}

// SignatureType is ...
type SignatureType struct {
	Mechanism string `json:"mechanism"`
}

// Authentication is ...
type Authentication struct {
	Mechanism               string `json:"mechanism"`
	SocialSecurityNumber    string `json:"socialSecurityNumber"`
	SignatureMethodUniqueId string `json:"signatureMethodUniqueId"`
}

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

// Notification is ...
type Notification struct {
	SignRequest      *SignRequest      `json:"signRequest,omitEmpty"`
	Reminder         *Reminder         `json:"reminder,omitEmpty"`
	SignatureReceipt *SignatureReceipt `json:"signatureReceipt,omitEmpty"`
	FinalReceipt     *FinalReceipt     `json:"finalReceipt,omitEmpty"`
	CanceledReceipt  *CanceledReceipt  `json:"canceledReceipt,omitEmpty"`
	ExpiredReceipt   *ExpiredReceipt   `json:"expiredReceipt,omitEmpty"`
}

// SignRequest is ...
type SignRequest struct {
	IncludeOriginalFile bool     `json:"includeOriginalFile,omitEmpty"`
	Email               []*Email `json:"email,omitEmpty"`
	Sms                 []*Sms   `json:"sms,omitEmpty"`
}

// Email is ...
type Email struct {
	Language   string `json:"language"`
	Subject    string `json:"subject,omitEmpty"`
	Text       string `json:"text,omitEmpty"`
	SenderName string `json:"senderName,omitEmpty"`
}

// Sms is ...
type Sms struct {
	Language string `json:"language"`
	Text     string `json:"text,omitEmpty"`
	Sender   string `json:"sender,omitEmpty"`
}

// Reminder is ...
type Reminder struct {
	ChronSchedule string   `json:"chronSchedule"`
	MaxReminders  int32    `json:"maxReminders,omitempty"`
	Email         []*Email `json:"email,omitempty"`
	Sms           []*Sms   `json:"sms,omitempty"`
}

// SignatureReceipt is ...
type SignatureReceipt struct {
	Email []*Email `json:"email,omitempty"`
	Sms   []*Sms   `json:"sms,omitempty"`
}

// FinalReceipt is ...
type FinalReceipt struct {
	AdditionalRecipients []*AdditionalRecipient `json:"additionalRecipients,omitempty"`
	IncludeSignedFile    bool                   `json:"includeSignedFile,omitempty"`
	Email                []*Email               `json:"email,omitempty"`
	Sms                  []*Sms                 `json:"sms,omitempty"`
}

// AdditionalRecipient is ...
type AdditionalRecipient struct {
	Language          string            `json:"language,omitempty"`
	Email             string            `json:"email"`
	SustomMergeFields map[string]string `json:"sustomMergeFields,omitempty"`
}

// CanceledReceipt is ...
type CanceledReceipt struct {
	Email []*Email `json:"email,omitempty"`
	Sms   []*Sms   `json:"sms,omitempty"`
}

// ExpiredReceipt is ...
type ExpiredReceipt struct {
	Email []*Email `json:"email,omitempty"`
	Sms   []*Sms   `json:"sms,omitempty"`
}

// Notifications is ...
type Notifications struct {
	Setup *Setup `json:"setup,omitempty"`
}

// Setup is ...
type Setup struct {
	Request          string `json:"request,omitempty"`
	Reminder         string `json:"reminder,omitempty"`
	SignatureReceipt string `json:"signatureReceipt,omitempty"`
	FinalReceipt     string `json:"finalReceipt,omitempty"`
	Canceled         string `json:"canceled,omitempty"`
	Expired          string `json:"expired,omitempty"`
}

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
	SignatureMethod         string                `json:"signatureMethod"`
	FullName                string                `json:"fullName,omitempty"`
	FirstName               string                `json:"firstName,omitempty"`
	LastName                string                `json:"lastName,omitempty"`
	MiddleName              string                `json:"middleName,omitempty"`
	SignedTime              *time.Time            `json:"signedTime,omitempty"`
	DateOfBirth             string                `json:"dateOfBirth,omitempty"`
	SignatureMethodUniqueID string                `json:"signatureMethodUniqueId,omitempty"`
	SocialSecurityNumber    *SocialSecurityNumber `json:"socialSecurityNumber,omitempty"`
	ClientIP                string                `json:"clientIp,omitempty"`
	Mechanism               string                `json:"mechanism,omitempty"`
	PersonalInfoOrigin      string                `json:"personalInfoOrigin,omitempty"`
}

// SocialSecurityNumber is ...
type SocialSecurityNumber struct {
	Value       string `json:"value,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

// Status is ...
type Status struct {
	DocumentStatus    string   `json:"documentStatus,omitempty"`
	CompletedPackages []string `json:"completedPackages,omitempty"`
}
