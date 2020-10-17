// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcPeeringConnectionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * accepter-vpc-info.cidr-block - The IPv4 CIDR block of the accepter VPC.
	//
	//    * accepter-vpc-info.owner-id - The AWS account ID of the owner of the
	//    accepter VPC.
	//
	//    * accepter-vpc-info.vpc-id - The ID of the accepter VPC.
	//
	//    * expiration-time - The expiration date and time for the VPC peering connection.
	//
	//    * requester-vpc-info.cidr-block - The IPv4 CIDR block of the requester's
	//    VPC.
	//
	//    * requester-vpc-info.owner-id - The AWS account ID of the owner of the
	//    requester VPC.
	//
	//    * requester-vpc-info.vpc-id - The ID of the requester VPC.
	//
	//    * status-code - The status of the VPC peering connection (pending-acceptance
	//    | failed | expired | provisioning | active | deleting | deleted | rejected).
	//
	//    * status-message - A message that provides more information about the
	//    status of the VPC peering connection, if applicable.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * vpc-peering-connection-id - The ID of the VPC peering connection.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// One or more VPC peering connection IDs.
	//
	// Default: Describes all your VPC peering connections.
	VpcPeeringConnectionIds []string `locationName:"VpcPeeringConnectionId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcPeeringConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcPeeringConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVpcPeeringConnectionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeVpcPeeringConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the VPC peering connections.
	VpcPeeringConnections []VpcPeeringConnection `locationName:"vpcPeeringConnectionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcPeeringConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeVpcPeeringConnections = "DescribeVpcPeeringConnections"

// DescribeVpcPeeringConnectionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your VPC peering connections.
//
//    // Example sending a request using DescribeVpcPeeringConnectionsRequest.
//    req := client.DescribeVpcPeeringConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeVpcPeeringConnections
func (c *Client) DescribeVpcPeeringConnectionsRequest(input *DescribeVpcPeeringConnectionsInput) DescribeVpcPeeringConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpcPeeringConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeVpcPeeringConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpcPeeringConnectionsOutput{})

	return DescribeVpcPeeringConnectionsRequest{Request: req, Input: input, Copy: c.DescribeVpcPeeringConnectionsRequest}
}

// DescribeVpcPeeringConnectionsRequest is the request type for the
// DescribeVpcPeeringConnections API operation.
type DescribeVpcPeeringConnectionsRequest struct {
	*aws.Request
	Input *DescribeVpcPeeringConnectionsInput
	Copy  func(*DescribeVpcPeeringConnectionsInput) DescribeVpcPeeringConnectionsRequest
}

// Send marshals and sends the DescribeVpcPeeringConnections API request.
func (r DescribeVpcPeeringConnectionsRequest) Send(ctx context.Context) (*DescribeVpcPeeringConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpcPeeringConnectionsResponse{
		DescribeVpcPeeringConnectionsOutput: r.Request.Data.(*DescribeVpcPeeringConnectionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeVpcPeeringConnectionsRequestPaginator returns a paginator for DescribeVpcPeeringConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeVpcPeeringConnectionsRequest(input)
//   p := ec2.NewDescribeVpcPeeringConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeVpcPeeringConnectionsPaginator(req DescribeVpcPeeringConnectionsRequest) DescribeVpcPeeringConnectionsPaginator {
	return DescribeVpcPeeringConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeVpcPeeringConnectionsInput
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

// DescribeVpcPeeringConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeVpcPeeringConnectionsPaginator struct {
	aws.Pager
}

func (p *DescribeVpcPeeringConnectionsPaginator) CurrentPage() *DescribeVpcPeeringConnectionsOutput {
	return p.Pager.CurrentPage().(*DescribeVpcPeeringConnectionsOutput)
}

// DescribeVpcPeeringConnectionsResponse is the response type for the
// DescribeVpcPeeringConnections API operation.
type DescribeVpcPeeringConnectionsResponse struct {
	*DescribeVpcPeeringConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpcPeeringConnections request.
func (r *DescribeVpcPeeringConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
