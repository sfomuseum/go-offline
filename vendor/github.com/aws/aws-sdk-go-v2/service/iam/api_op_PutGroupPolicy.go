// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an inline policy document that is embedded in the specified IAM
// group.
//
// A user can also have managed policies attached to it. To attach a managed
// policy to a group, use [AttachGroupPolicy]AttachGroupPolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// in a group, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutGroupPolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [AttachGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachGroupPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func (c *Client) PutGroupPolicy(ctx context.Context, params *PutGroupPolicyInput, optFns ...func(*Options)) (*PutGroupPolicyOutput, error) {
	if params == nil {
		params = &PutGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGroupPolicy", params, optFns, c.addOperationPutGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutGroupPolicyInput struct {

	// The name of the group to associate the policy with.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	GroupName *string

	// The policy document.
	//
	// You must provide policies in JSON format in IAM. However, for CloudFormation
	// templates formatted in YAML, you can provide the policy in JSON or YAML format.
	// CloudFormation always converts a YAML policy to JSON format before submitting it
	// to IAM.
	//
	// The [regex pattern] used to validate this parameter is a string of characters consisting of
	// the following:
	//
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	PolicyDocument *string

	// The name of the policy document.
	//
	// This parameter allows (through its [regex pattern]) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

type PutGroupPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutGroupPolicy"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpPutGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGroupPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutGroupPolicy",
	}
}
