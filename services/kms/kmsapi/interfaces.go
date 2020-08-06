package kmsapi

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/kms"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateCustomKey(ctx context.Context, keyTag string, parameters kms.CreateCustomKeyRequest) (result autorest.Response, err error)
	Descrypt(ctx context.Context, keyTag string, parameters kms.DecryptRequestKey) (result autorest.Response, err error)
	Encrypt(ctx context.Context, keyTag string, parameters kms.EncryptRequestKey) (result autorest.Response, err error)
	Reencrypt(ctx context.Context, keyTag string, parameters kms.ReencryptRequestKey) (result autorest.Response, err error)
	Sign(ctx context.Context, keyTag string, parameters kms.SignRequest) (result autorest.Response, err error)
	Verify(ctx context.Context, keyTag string, parameters kms.VerifyRequest) (result autorest.Response, err error)
}

var _ ClientAPI = (*kms.Client)(nil)
