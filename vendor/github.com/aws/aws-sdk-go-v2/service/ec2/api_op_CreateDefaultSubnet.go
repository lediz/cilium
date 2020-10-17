// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDefaultSubnetInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which to create the default subnet.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s CreateDefaultSubnetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDefaultSubnetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDefaultSubnetInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDefaultSubnetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the subnet.
	Subnet *Subnet `locationName:"subnet" type:"structure"`
}

// String returns the string representation
func (s CreateDefaultSubnetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDefaultSubnet = "CreateDefaultSubnet"

// CreateDefaultSubnetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a default subnet with a size /20 IPv4 CIDR block in the specified
// Availability Zone in your default VPC. You can have only one default subnet
// per Availability Zone. For more information, see Creating a Default Subnet
// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#create-default-subnet)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateDefaultSubnetRequest.
//    req := client.CreateDefaultSubnetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateDefaultSubnet
func (c *Client) CreateDefaultSubnetRequest(input *CreateDefaultSubnetInput) CreateDefaultSubnetRequest {
	op := &aws.Operation{
		Name:       opCreateDefaultSubnet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDefaultSubnetInput{}
	}

	req := c.newRequest(op, input, &CreateDefaultSubnetOutput{})

	return CreateDefaultSubnetRequest{Request: req, Input: input, Copy: c.CreateDefaultSubnetRequest}
}

// CreateDefaultSubnetRequest is the request type for the
// CreateDefaultSubnet API operation.
type CreateDefaultSubnetRequest struct {
	*aws.Request
	Input *CreateDefaultSubnetInput
	Copy  func(*CreateDefaultSubnetInput) CreateDefaultSubnetRequest
}

// Send marshals and sends the CreateDefaultSubnet API request.
func (r CreateDefaultSubnetRequest) Send(ctx context.Context) (*CreateDefaultSubnetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDefaultSubnetResponse{
		CreateDefaultSubnetOutput: r.Request.Data.(*CreateDefaultSubnetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDefaultSubnetResponse is the response type for the
// CreateDefaultSubnet API operation.
type CreateDefaultSubnetResponse struct {
	*CreateDefaultSubnetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDefaultSubnet request.
func (r *CreateDefaultSubnetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
