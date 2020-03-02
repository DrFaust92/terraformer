// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type ListMultipartUploadsInput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to which the multipart upload was initiated.
	//
	// When using this API with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation using an access point through the AWS SDKs, you
	// provide the access point ARN in place of the bucket name. For more information
	// about access point ARNs, see Using Access Points (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html)
	// in the Amazon Simple Storage Service Developer Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Character you use to group keys.
	//
	// All keys that contain the same string between the prefix, if specified, and
	// the first occurrence of the delimiter after the prefix are grouped under
	// a single result element, CommonPrefixes. If you don't specify the prefix
	// parameter, then the substring starts at the beginning of the key. The keys
	// that are grouped under CommonPrefixes result element are not returned elsewhere
	// in the response.
	Delimiter *string `location:"querystring" locationName:"delimiter" type:"string"`

	// Requests Amazon S3 to encode the object keys in the response and specifies
	// the encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters
	// with an ASCII value from 0 to 10. For characters that are not supported in
	// XML 1.0, you can add this parameter to request that Amazon S3 encode the
	// keys in the response.
	EncodingType EncodingType `location:"querystring" locationName:"encoding-type" type:"string" enum:"true"`

	// Together with upload-id-marker, this parameter specifies the multipart upload
	// after which listing should begin.
	//
	// If upload-id-marker is not specified, only the keys lexicographically greater
	// than the specified key-marker will be included in the list.
	//
	// If upload-id-marker is specified, any multipart uploads for a key equal to
	// the key-marker might also be included, provided those multipart uploads have
	// upload IDs lexicographically greater than the specified upload-id-marker.
	KeyMarker *string `location:"querystring" locationName:"key-marker" type:"string"`

	// Sets the maximum number of multipart uploads, from 1 to 1,000, to return
	// in the response body. 1,000 is the maximum number of uploads that can be
	// returned in a response.
	MaxUploads *int64 `location:"querystring" locationName:"max-uploads" type:"integer"`

	// Lists in-progress uploads only for those keys that begin with the specified
	// prefix. You can use prefixes to separate a bucket into different grouping
	// of keys. (You can think of using prefix to make groups in the same way you'd
	// use a folder in a file system.)
	Prefix *string `location:"querystring" locationName:"prefix" type:"string"`

	// Together with key-marker, specifies the multipart upload after which listing
	// should begin. If key-marker is not specified, the upload-id-marker parameter
	// is ignored. Otherwise, any multipart uploads for a key equal to the key-marker
	// might be included in the list only if they have an upload ID lexicographically
	// greater than the specified upload-id-marker.
	UploadIdMarker *string `location:"querystring" locationName:"upload-id-marker" type:"string"`
}

// String returns the string representation
func (s ListMultipartUploadsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMultipartUploadsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMultipartUploadsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListMultipartUploadsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultipartUploadsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "encoding-type", v, metadata)
	}
	if s.KeyMarker != nil {
		v := *s.KeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "key-marker", protocol.StringValue(v), metadata)
	}
	if s.MaxUploads != nil {
		v := *s.MaxUploads

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-uploads", protocol.Int64Value(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "prefix", protocol.StringValue(v), metadata)
	}
	if s.UploadIdMarker != nil {
		v := *s.UploadIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "upload-id-marker", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *ListMultipartUploadsInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *ListMultipartUploadsInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type ListMultipartUploadsOutput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to which the multipart upload was initiated.
	Bucket *string `type:"string"`

	// If you specify a delimiter in the request, then the result returns each distinct
	// key prefix containing the delimiter in a CommonPrefixes element. The distinct
	// key prefixes are returned in the Prefix child element.
	CommonPrefixes []CommonPrefix `type:"list" flattened:"true"`

	// Contains the delimiter you specified in the request. If you don't specify
	// a delimiter in your request, this element is absent from the response.
	Delimiter *string `type:"string"`

	// Encoding type used by Amazon S3 to encode object keys in the response.
	//
	// If you specify encoding-type request parameter, Amazon S3 includes this element
	// in the response, and returns encoded key name values in the following response
	// elements:
	//
	// Delimiter, KeyMarker, Prefix, NextKeyMarker, Key.
	EncodingType EncodingType `type:"string" enum:"true"`

	// Indicates whether the returned list of multipart uploads is truncated. A
	// value of true indicates that the list was truncated. The list can be truncated
	// if the number of multipart uploads exceeds the limit allowed or specified
	// by max uploads.
	IsTruncated *bool `type:"boolean"`

	// The key at or after which the listing began.
	KeyMarker *string `type:"string"`

	// Maximum number of multipart uploads that could have been included in the
	// response.
	MaxUploads *int64 `type:"integer"`

	// When a list is truncated, this element specifies the value that should be
	// used for the key-marker request parameter in a subsequent request.
	NextKeyMarker *string `type:"string"`

	// When a list is truncated, this element specifies the value that should be
	// used for the upload-id-marker request parameter in a subsequent request.
	NextUploadIdMarker *string `type:"string"`

	// When a prefix is provided in the request, this field contains the specified
	// prefix. The result contains only keys starting with the specified prefix.
	Prefix *string `type:"string"`

	// Upload ID after which listing began.
	UploadIdMarker *string `type:"string"`

	// Container for elements related to a particular multipart upload. A response
	// can contain zero or more Upload elements.
	Uploads []MultipartUpload `locationName:"Upload" type:"list" flattened:"true"`
}

// String returns the string representation
func (s ListMultipartUploadsOutput) String() string {
	return awsutil.Prettify(s)
}

func (s *ListMultipartUploadsOutput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListMultipartUploadsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.CommonPrefixes != nil {
		v := s.CommonPrefixes

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "CommonPrefixes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EncodingType", v, metadata)
	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.KeyMarker != nil {
		v := *s.KeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KeyMarker", protocol.StringValue(v), metadata)
	}
	if s.MaxUploads != nil {
		v := *s.MaxUploads

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxUploads", protocol.Int64Value(v), metadata)
	}
	if s.NextKeyMarker != nil {
		v := *s.NextKeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextKeyMarker", protocol.StringValue(v), metadata)
	}
	if s.NextUploadIdMarker != nil {
		v := *s.NextUploadIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextUploadIdMarker", protocol.StringValue(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Prefix", protocol.StringValue(v), metadata)
	}
	if s.UploadIdMarker != nil {
		v := *s.UploadIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UploadIdMarker", protocol.StringValue(v), metadata)
	}
	if s.Uploads != nil {
		v := s.Uploads

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Upload", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListMultipartUploads = "ListMultipartUploads"

// ListMultipartUploadsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// This operation lists in-progress multipart uploads. An in-progress multipart
// upload is a multipart upload that has been initiated using the Initiate Multipart
// Upload request, but has not yet been completed or aborted.
//
// This operation returns at most 1,000 multipart uploads in the response. 1,000
// multipart uploads is the maximum number of uploads a response can include,
// which is also the default value. You can further limit the number of uploads
// in a response by specifying the max-uploads parameter in the response. If
// additional multipart uploads satisfy the list criteria, the response will
// contain an IsTruncated element with the value true. To list the additional
// multipart uploads, use the key-marker and upload-id-marker request parameters.
//
// In the response, the uploads are sorted by key. If your application has initiated
// more than one multipart upload using the same object key, then uploads in
// the response are first sorted by key. Additionally, uploads are sorted in
// ascending order within each key by the upload initiation time.
//
// For more information on multipart uploads, see Uploading Objects Using Multipart
// Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html).
//
// For information on permissions required to use the multipart upload API,
// see Multipart Upload API and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html).
//
// The following operations are related to ListMultipartUploads:
//
//    * CreateMultipartUpload
//
//    * UploadPart
//
//    * CompleteMultipartUpload
//
//    * ListParts
//
//    * AbortMultipartUpload
//
//    // Example sending a request using ListMultipartUploadsRequest.
//    req := client.ListMultipartUploadsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListMultipartUploads
func (c *Client) ListMultipartUploadsRequest(input *ListMultipartUploadsInput) ListMultipartUploadsRequest {
	op := &aws.Operation{
		Name:       opListMultipartUploads,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?uploads",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"KeyMarker", "UploadIdMarker"},
			OutputTokens:    []string{"NextKeyMarker", "NextUploadIdMarker"},
			LimitToken:      "MaxUploads",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListMultipartUploadsInput{}
	}

	req := c.newRequest(op, input, &ListMultipartUploadsOutput{})
	return ListMultipartUploadsRequest{Request: req, Input: input, Copy: c.ListMultipartUploadsRequest}
}

// ListMultipartUploadsRequest is the request type for the
// ListMultipartUploads API operation.
type ListMultipartUploadsRequest struct {
	*aws.Request
	Input *ListMultipartUploadsInput
	Copy  func(*ListMultipartUploadsInput) ListMultipartUploadsRequest
}

// Send marshals and sends the ListMultipartUploads API request.
func (r ListMultipartUploadsRequest) Send(ctx context.Context) (*ListMultipartUploadsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMultipartUploadsResponse{
		ListMultipartUploadsOutput: r.Request.Data.(*ListMultipartUploadsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMultipartUploadsRequestPaginator returns a paginator for ListMultipartUploads.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMultipartUploadsRequest(input)
//   p := s3.NewListMultipartUploadsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMultipartUploadsPaginator(req ListMultipartUploadsRequest) ListMultipartUploadsPaginator {
	return ListMultipartUploadsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListMultipartUploadsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListMultipartUploadsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMultipartUploadsPaginator struct {
	aws.Pager
}

func (p *ListMultipartUploadsPaginator) CurrentPage() *ListMultipartUploadsOutput {
	return p.Pager.CurrentPage().(*ListMultipartUploadsOutput)
}

// ListMultipartUploadsResponse is the response type for the
// ListMultipartUploads API operation.
type ListMultipartUploadsResponse struct {
	*ListMultipartUploadsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMultipartUploads request.
func (r *ListMultipartUploadsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
