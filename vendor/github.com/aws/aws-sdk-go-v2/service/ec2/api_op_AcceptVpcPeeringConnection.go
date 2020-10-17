// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AcceptVpcPeeringConnectionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC peering connection. You must specify this parameter in
	// the request.
	VpcPeeringConnectionId *string `locationName:"vpcPeeringConnectionId" type:"string"`
}

// String returns the string representation
func (s AcceptVpcPeeringConnectionInput) String() string {
	return awsutil.Prettify(s)
}

type AcceptVpcPeeringConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the VPC peering connection.
	VpcPeeringConnection *VpcPeeringConnection `locationName:"vpcPeeringConnection" type:"structure"`
}

// String returns the string representation
func (s AcceptVpcPeeringConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptVpcPeeringConnection = "AcceptVpcPeeringConnection"

// AcceptVpcPeeringConnectionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Accept a VPC peering connection request. To accept a request, the VPC peering
// connection must be in the pending-acceptance state, and you must be the owner
// of the peer VPC. Use DescribeVpcPeeringConnections to view your outstanding
// VPC peering connection requests.
//
// For an inter-Region VPC peering connection request, you must accept the VPC
// peering connection in the Region of the accepter VPC.
//
//    // Example sending a request using AcceptVpcPeeringConnectionRequest.
//    req := client.AcceptVpcPeeringConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptVpcPeeringConnection
func (c *Client) AcceptVpcPeeringConnectionRequest(input *AcceptVpcPeeringConnectionInput) AcceptVpcPeeringConnectionRequest {
	op := &aws.Operation{
		Name:       opAcceptVpcPeeringConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptVpcPeeringConnectionInput{}
	}

	req := c.newRequest(op, input, &AcceptVpcPeeringConnectionOutput{})

	return AcceptVpcPeeringConnectionRequest{Request: req, Input: input, Copy: c.AcceptVpcPeeringConnectionRequest}
}

// AcceptVpcPeeringConnectionRequest is the request type for the
// AcceptVpcPeeringConnection API operation.
type AcceptVpcPeeringConnectionRequest struct {
	*aws.Request
	Input *AcceptVpcPeeringConnectionInput
	Copy  func(*AcceptVpcPeeringConnectionInput) AcceptVpcPeeringConnectionRequest
}

// Send marshals and sends the AcceptVpcPeeringConnection API request.
func (r AcceptVpcPeeringConnectionRequest) Send(ctx context.Context) (*AcceptVpcPeeringConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptVpcPeeringConnectionResponse{
		AcceptVpcPeeringConnectionOutput: r.Request.Data.(*AcceptVpcPeeringConnectionOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptVpcPeeringConnectionResponse is the response type for the
// AcceptVpcPeeringConnection API operation.
type AcceptVpcPeeringConnectionResponse struct {
	*AcceptVpcPeeringConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptVpcPeeringConnection request.
func (r *AcceptVpcPeeringConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
