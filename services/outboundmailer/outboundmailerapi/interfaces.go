package outboundmailerapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/samjegal/fincloud-sdk-for-go/services/outboundmailer"
	"io"
)

// MailClientAPI contains the set of methods on the MailClient type.
type MailClientAPI interface {
	Create(ctx context.Context, parameter outboundmailer.MailRequestParameter) (result outboundmailer.MailResponseParameter, err error)
	Get(ctx context.Context, mailID string, xNCPLANG string) (result outboundmailer.SetObject, err error)
	GetListByRequestID(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (result outboundmailer.SetObject, err error)
	GetRequestList(ctx context.Context, xNCPLANG string, dispatchType string, endDateTime string, endUtc *int64, mailID string, page *int32, recipientAddress string, senderAddress string, sendStatus []string, size *int32, sortParameter string, startDateTime string, startUtc *int64, templateSid *int32, title string) (result outboundmailer.SetObject, err error)
	GetStatusByRequestID(ctx context.Context, requestID string, xNCPLANG string, mailID string, page *int32, recipientAddress string, sendStatus []string, size *int32, sortParameter []string, title string) (result outboundmailer.SetObject, err error)
}

var _ MailClientAPI = (*outboundmailer.MailClient)(nil)

// FileClientAPI contains the set of methods on the FileClient type.
type FileClientAPI interface {
	Create(ctx context.Context, fileList io.ReadCloser, xNCPLANG string) (result outboundmailer.FileUploadResponse, err error)
	Delete(ctx context.Context, tempRequestID string, xNCPLANG string) (result outboundmailer.SetObject, err error)
	Get(ctx context.Context, tempRequestID string, xNCPLANG string) (result outboundmailer.SetObject, err error)
}

var _ FileClientAPI = (*outboundmailer.FileClient)(nil)

// AddressBookClientAPI contains the set of methods on the AddressBookClient type.
type AddressBookClientAPI interface {
	Create(ctx context.Context, requestBody outboundmailer.AddressBookGenerateRequest, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
	Delete(ctx context.Context, xNCPLANG string) (result outboundmailer.AddressBookInitResponse, err error)
	DeleteAddress(ctx context.Context, requestBody outboundmailer.AddressBookDeleteAddressRequest, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
	DeleteEmptyRecipientGroup(ctx context.Context, groupName string, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
	DeleteRecipientGroup(ctx context.Context, groupName string, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
	DeleteRecipientGroupByAddress(ctx context.Context, requestBody outboundmailer.AddressBookDeleteRelationRequest, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
	Get(ctx context.Context, xNCPLANG string) (result outboundmailer.AddressBookResponse, err error)
}

var _ AddressBookClientAPI = (*outboundmailer.AddressBookClient)(nil)

// UnsubscriberClientAPI contains the set of methods on the UnsubscriberClient type.
type UnsubscriberClientAPI interface {
	Get(ctx context.Context, xNCPLANG string, emailAddress string, endDateTime string, endUtc *int64, isRegByManager *bool, page *int32, size *int32, sortParameter string, startDateTime string, startUtc *int64) (result outboundmailer.SetObject, err error)
}

var _ UnsubscriberClientAPI = (*outboundmailer.UnsubscriberClient)(nil)
