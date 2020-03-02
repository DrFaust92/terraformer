// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GenerateDataKeyWithoutPlaintextInput struct {
	_ struct{} `type:"structure"`

	// Specifies the encryption context that will be used when encrypting the data
	// key.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional
	// when encrypting with a symmetric CMK, but it is highly recommended.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// The identifier of the symmetric customer master key (CMK) that encrypts the
	// data key.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// The length of the data key. Use AES_128 to generate a 128-bit symmetric key,
	// or AES_256 to generate a 256-bit symmetric key.
	KeySpec DataKeySpec `type:"string" enum:"true"`

	// The length of the data key in bytes. For example, use the value 64 to generate
	// a 512-bit data key (64 bytes is 512 bits). For common key lengths (128-bit
	// and 256-bit symmetric keys), we recommend that you use the KeySpec field
	// instead of this one.
	NumberOfBytes *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GenerateDataKeyWithoutPlaintextInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataKeyWithoutPlaintextInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataKeyWithoutPlaintextInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.NumberOfBytes != nil && *s.NumberOfBytes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBytes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GenerateDataKeyWithoutPlaintextOutput struct {
	_ struct{} `type:"structure"`

	// The encrypted data key. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not Base64-encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// The identifier of the CMK that encrypted the data key.
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GenerateDataKeyWithoutPlaintextOutput) String() string {
	return awsutil.Prettify(s)
}

const opGenerateDataKeyWithoutPlaintext = "GenerateDataKeyWithoutPlaintext"

// GenerateDataKeyWithoutPlaintextRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Generates a unique symmetric data key. This operation returns a data key
// that is encrypted under a customer master key (CMK) that you specify. To
// request an asymmetric data key pair, use the GenerateDataKeyPair or GenerateDataKeyPairWithoutPlaintext
// operations.
//
// GenerateDataKeyWithoutPlaintext is identical to the GenerateDataKey operation
// except that returns only the encrypted copy of the data key. This operation
// is useful for systems that need to encrypt data at some point, but not immediately.
// When you need to encrypt the data, you call the Decrypt operation on the
// encrypted copy of the key.
//
// It's also useful in distributed systems with different levels of trust. For
// example, you might store encrypted data in containers. One component of your
// system creates new containers and stores an encrypted data key with each
// container. Then, a different component puts the data into the containers.
// That component first decrypts the data key, uses the plaintext data key to
// encrypt data, puts the encrypted data into the container, and then destroys
// the plaintext data key. In this system, the component that creates the containers
// never sees the plaintext data key.
//
// GenerateDataKeyWithoutPlaintext returns a unique data key for each request.
// The bytes in the keys are not related to the caller or CMK that is used to
// encrypt the private key.
//
// To generate a data key, you must specify the symmetric customer master key
// (CMK) that is used to encrypt the data key. You cannot use an asymmetric
// CMK to generate a data key. To get the type of your CMK, use the DescribeKey
// operation.
//
// If the operation succeeds, you will find the encrypted copy of the data key
// in the CiphertextBlob field.
//
// You can use the optional encryption context to add additional security to
// the encryption operation. If you specify an EncryptionContext, you must specify
// the same encryption context (a case-sensitive exact match) when decrypting
// the encrypted data key. Otherwise, the request to decrypt fails with an InvalidCiphertextException.
// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the AWS Key Management Service Developer Guide.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using GenerateDataKeyWithoutPlaintextRequest.
//    req := client.GenerateDataKeyWithoutPlaintextRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/GenerateDataKeyWithoutPlaintext
func (c *Client) GenerateDataKeyWithoutPlaintextRequest(input *GenerateDataKeyWithoutPlaintextInput) GenerateDataKeyWithoutPlaintextRequest {
	op := &aws.Operation{
		Name:       opGenerateDataKeyWithoutPlaintext,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataKeyWithoutPlaintextInput{}
	}

	req := c.newRequest(op, input, &GenerateDataKeyWithoutPlaintextOutput{})
	return GenerateDataKeyWithoutPlaintextRequest{Request: req, Input: input, Copy: c.GenerateDataKeyWithoutPlaintextRequest}
}

// GenerateDataKeyWithoutPlaintextRequest is the request type for the
// GenerateDataKeyWithoutPlaintext API operation.
type GenerateDataKeyWithoutPlaintextRequest struct {
	*aws.Request
	Input *GenerateDataKeyWithoutPlaintextInput
	Copy  func(*GenerateDataKeyWithoutPlaintextInput) GenerateDataKeyWithoutPlaintextRequest
}

// Send marshals and sends the GenerateDataKeyWithoutPlaintext API request.
func (r GenerateDataKeyWithoutPlaintextRequest) Send(ctx context.Context) (*GenerateDataKeyWithoutPlaintextResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateDataKeyWithoutPlaintextResponse{
		GenerateDataKeyWithoutPlaintextOutput: r.Request.Data.(*GenerateDataKeyWithoutPlaintextOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateDataKeyWithoutPlaintextResponse is the response type for the
// GenerateDataKeyWithoutPlaintext API operation.
type GenerateDataKeyWithoutPlaintextResponse struct {
	*GenerateDataKeyWithoutPlaintextOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateDataKeyWithoutPlaintext request.
func (r *GenerateDataKeyWithoutPlaintextResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
