package kms

// FINCLOUD_APACHE_NO_VERSION

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/kms"

// CreateCustomKeyRequest ...
type CreateCustomKeyRequest struct {
	// RequestPlainKey - 생성된 키를 평문 형태로 반환받을지 선택하는 플래그
	RequestPlainKey *bool `json:"requestPlainKey,omitempty"`
	// Bits - 생성될 키의 비트 (128, 256, 512) (기본 256)
	Bits *int32 `json:"bits,omitempty"`
}

// DecryptRequestKey ...
type DecryptRequestKey struct {
	// Ciphertext - 복호화할 데이터 (kms prefix와 Base64 Encoding된 암호문으로 구성되어야합니다)
	Ciphertext *string `json:"ciphertext,omitempty"`
}

// EncryptRequestKey ...
type EncryptRequestKey struct {
	// Plaintext - 암호화할 데이터 (최대 32KB) (Base64 Encoding이 필요합니다)
	Plaintext *string `json:"plaintext,omitempty"`
}

// ReencryptRequestKey ...
type ReencryptRequestKey struct {
	// Ciphertext - 재암호화할 데이터 (kms prefix와 Base64 Encoding된 암호문으로 구성되어야합니다)
	Ciphertext *string `json:"ciphertext,omitempty"`
}

// SignRequest ...
type SignRequest struct {
	// Data - 서명값을 생성할 데이터 (최대 8KB) (Base64 Encoding이 필요합니다)
	Data *string `json:"data,omitempty"`
}

// VerifyRequest ...
type VerifyRequest struct {
	// Data - 서명값을 생성할 데이터 (최대 8KB) (Base64 Encoding이 필요합니다)
	Data *string `json:"data,omitempty"`
	// Signature - 서명값 데이터 (kms prefix와 Base64 Encoding된 서명값으로 구성되어 있습니다)
	Signature *string `json:"signature,omitempty"`
}
