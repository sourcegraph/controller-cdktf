package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference interface {
	cdktf.ComplexObject
	BitbucketServerConfigResource() *string
	SetBitbucketServerConfigResource(val *string)
	BitbucketServerConfigResourceInput() *string
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
	InternalValue() *GoogleCloudbuildTriggerBitbucketServerTriggerConfig
	SetInternalValue(val *GoogleCloudbuildTriggerBitbucketServerTriggerConfig)
	ProjectKey() *string
	SetProjectKey(val *string)
	ProjectKeyInput() *string
	PullRequest() GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference
	PullRequestInput() *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequest
	Push() GoogleCloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference
	PushInput() *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPush
	RepoSlug() *string
	SetRepoSlug(val *string)
	RepoSlugInput() *string
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
	PutPullRequest(value *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequest)
	PutPush(value *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPush)
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

// The jsii proxy struct for GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) BitbucketServerConfigResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitbucketServerConfigResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) BitbucketServerConfigResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitbucketServerConfigResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InternalValue() *GoogleCloudbuildTriggerBitbucketServerTriggerConfig {
	var returns *GoogleCloudbuildTriggerBitbucketServerTriggerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ProjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ProjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PullRequest() GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference {
	var returns GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PullRequestInput() *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequest {
	var returns *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Push() GoogleCloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference {
	var returns GoogleCloudbuildTriggerBitbucketServerTriggerConfigPushOutputReference
	_jsii_.Get(
		j,
		"push",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PushInput() *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPush {
	var returns *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPush
	_jsii_.Get(
		j,
		"pushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) RepoSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) RepoSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repoSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference_Override(g GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetBitbucketServerConfigResource(val *string) {
	if err := j.validateSetBitbucketServerConfigResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitbucketServerConfigResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetInternalValue(val *GoogleCloudbuildTriggerBitbucketServerTriggerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetProjectKey(val *string) {
	if err := j.validateSetProjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectKey",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetRepoSlug(val *string) {
	if err := j.validateSetRepoSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repoSlug",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PutPullRequest(value *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPullRequest) {
	if err := g.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) PutPush(value *GoogleCloudbuildTriggerBitbucketServerTriggerConfigPush) {
	if err := g.validatePutPushParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPush",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ResetPush() {
	_jsii_.InvokeVoid(
		g,
		"resetPush",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBitbucketServerTriggerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

