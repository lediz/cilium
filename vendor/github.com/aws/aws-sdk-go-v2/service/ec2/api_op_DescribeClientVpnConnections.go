// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClientVpnConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. Filter names and values are case-sensitive.
	//
	//    * connection-id - The ID of the connection.
	//
	//    * username - For Active Directory client authentication, the user name
	//    of the client who established the client connection.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeClientVpnConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeClientVpnConnectionsInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeClientVpnConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the active and terminated client connections.
	Connections []VpnConnection `locationName:"connections" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeClientVpnConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClientVpnConnections = "DescribeClientVpnConnections"

// DescribeClientVpnConnectionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes active client connections and connections that have been terminated
// within the last 60 minutes for the specified Client VPN endpoint.
//
//    // Example sending a request using DescribeClientVpnConnectionsRequest.
//    req := client.DescribeClientVpnConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeClientVpnConnections
func (c *Client) DescribeClientVpnConnectionsRequest(input *DescribeClientVpnConnectionsInput) DescribeClientVpnConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeClientVpnConnections,
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
		input = &DescribeClientVpnConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeClientVpnConnectionsOutput{})

	return DescribeClientVpnConnectionsRequest{Request: req, Input: input, Copy: c.DescribeClientVpnConnectionsRequest}
}

// DescribeClientVpnConnectionsRequest is the request type for the
// DescribeClientVpnConnections API operation.
type DescribeClientVpnConnectionsRequest struct {
	*aws.Request
	Input *DescribeClientVpnConnectionsInput
	Copy  func(*DescribeClientVpnConnectionsInput) DescribeClientVpnConnectionsRequest
}

// Send marshals and sends the DescribeClientVpnConnections API request.
func (r DescribeClientVpnConnectionsRequest) Send(ctx context.Context) (*DescribeClientVpnConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClientVpnConnectionsResponse{
		DescribeClientVpnConnectionsOutput: r.Request.Data.(*DescribeClientVpnConnectionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClientVpnConnectionsRequestPaginator returns a paginator for DescribeClientVpnConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClientVpnConnectionsRequest(input)
//   p := ec2.NewDescribeClientVpnConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClientVpnConnectionsPaginator(req DescribeClientVpnConnectionsRequest) DescribeClientVpnConnectionsPaginator {
	return DescribeClientVpnConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClientVpnConnectionsInput
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

// DescribeClientVpnConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClientVpnConnectionsPaginator struct {
	aws.Pager
}

func (p *DescribeClientVpnConnectionsPaginator) CurrentPage() *DescribeClientVpnConnectionsOutput {
	return p.Pager.CurrentPage().(*DescribeClientVpnConnectionsOutput)
}

// DescribeClientVpnConnectionsResponse is the response type for the
// DescribeClientVpnConnections API operation.
type DescribeClientVpnConnectionsResponse struct {
	*DescribeClientVpnConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClientVpnConnections request.
func (r *DescribeClientVpnConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
