package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudbuildtrigger/internal"
)

type CloudbuildTriggerDeveloperConnectEventConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GitRepositoryLink() *string
	SetGitRepositoryLink(val *string)
	GitRepositoryLinkInput() *string
	GitRepositoryLinkType() *string
	InternalValue() *CloudbuildTriggerDeveloperConnectEventConfig
	SetInternalValue(val *CloudbuildTriggerDeveloperConnectEventConfig)
	PullRequest() CloudbuildTriggerDeveloperConnectEventConfigPullRequestOutputReference
	PullRequestInput() *CloudbuildTriggerDeveloperConnectEventConfigPullRequest
	Push() CloudbuildTriggerDeveloperConnectEventConfigPushOutputReference
	PushInput() *CloudbuildTriggerDeveloperConnectEventConfigPush
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutPullRequest(value *CloudbuildTriggerDeveloperConnectEventConfigPullRequest)
	PutPush(value *CloudbuildTriggerDeveloperConnectEventConfigPush)
	ResetPullRequest()
	ResetPush()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerDeveloperConnectEventConfigOutputReference
type jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GitRepositoryLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GitRepositoryLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GitRepositoryLinkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitRepositoryLinkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) InternalValue() *CloudbuildTriggerDeveloperConnectEventConfig {
	var returns *CloudbuildTriggerDeveloperConnectEventConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) PullRequest() CloudbuildTriggerDeveloperConnectEventConfigPullRequestOutputReference {
	var returns CloudbuildTriggerDeveloperConnectEventConfigPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) PullRequestInput() *CloudbuildTriggerDeveloperConnectEventConfigPullRequest {
	var returns *CloudbuildTriggerDeveloperConnectEventConfigPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) Push() CloudbuildTriggerDeveloperConnectEventConfigPushOutputReference {
	var returns CloudbuildTriggerDeveloperConnectEventConfigPushOutputReference
	_jsii_.Get(
		j,
		"push",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) PushInput() *CloudbuildTriggerDeveloperConnectEventConfigPush {
	var returns *CloudbuildTriggerDeveloperConnectEventConfigPush
	_jsii_.Get(
		j,
		"pushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerDeveloperConnectEventConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerDeveloperConnectEventConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerDeveloperConnectEventConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerDeveloperConnectEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerDeveloperConnectEventConfigOutputReference_Override(c CloudbuildTriggerDeveloperConnectEventConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerDeveloperConnectEventConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetGitRepositoryLink(val *string) {
	if err := j.validateSetGitRepositoryLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitRepositoryLink",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetInternalValue(val *CloudbuildTriggerDeveloperConnectEventConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) PutPullRequest(value *CloudbuildTriggerDeveloperConnectEventConfigPullRequest) {
	if err := c.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) PutPush(value *CloudbuildTriggerDeveloperConnectEventConfigPush) {
	if err := c.validatePutPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPush",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ResetPush() {
	_jsii_.InvokeVoid(
		c,
		"resetPush",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerDeveloperConnectEventConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

