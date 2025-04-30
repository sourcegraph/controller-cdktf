package googlecloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudtasksqueue/internal"
)

type GoogleCloudTasksQueueHttpTargetOutputReference interface {
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
	HeaderOverrides() GoogleCloudTasksQueueHttpTargetHeaderOverridesList
	HeaderOverridesInput() interface{}
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() *GoogleCloudTasksQueueHttpTarget
	SetInternalValue(val *GoogleCloudTasksQueueHttpTarget)
	OauthToken() GoogleCloudTasksQueueHttpTargetOauthTokenOutputReference
	OauthTokenInput() *GoogleCloudTasksQueueHttpTargetOauthToken
	OidcToken() GoogleCloudTasksQueueHttpTargetOidcTokenOutputReference
	OidcTokenInput() *GoogleCloudTasksQueueHttpTargetOidcToken
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriOverride() GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference
	UriOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverride
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
	PutHeaderOverrides(value interface{})
	PutOauthToken(value *GoogleCloudTasksQueueHttpTargetOauthToken)
	PutOidcToken(value *GoogleCloudTasksQueueHttpTargetOidcToken)
	PutUriOverride(value *GoogleCloudTasksQueueHttpTargetUriOverride)
	ResetHeaderOverrides()
	ResetHttpMethod()
	ResetOauthToken()
	ResetOidcToken()
	ResetUriOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudTasksQueueHttpTargetOutputReference
type jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) HeaderOverrides() GoogleCloudTasksQueueHttpTargetHeaderOverridesList {
	var returns GoogleCloudTasksQueueHttpTargetHeaderOverridesList
	_jsii_.Get(
		j,
		"headerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) HeaderOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) InternalValue() *GoogleCloudTasksQueueHttpTarget {
	var returns *GoogleCloudTasksQueueHttpTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) OauthToken() GoogleCloudTasksQueueHttpTargetOauthTokenOutputReference {
	var returns GoogleCloudTasksQueueHttpTargetOauthTokenOutputReference
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) OauthTokenInput() *GoogleCloudTasksQueueHttpTargetOauthToken {
	var returns *GoogleCloudTasksQueueHttpTargetOauthToken
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) OidcToken() GoogleCloudTasksQueueHttpTargetOidcTokenOutputReference {
	var returns GoogleCloudTasksQueueHttpTargetOidcTokenOutputReference
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) OidcTokenInput() *GoogleCloudTasksQueueHttpTargetOidcToken {
	var returns *GoogleCloudTasksQueueHttpTargetOidcToken
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) UriOverride() GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference {
	var returns GoogleCloudTasksQueueHttpTargetUriOverrideOutputReference
	_jsii_.Get(
		j,
		"uriOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) UriOverrideInput() *GoogleCloudTasksQueueHttpTargetUriOverride {
	var returns *GoogleCloudTasksQueueHttpTargetUriOverride
	_jsii_.Get(
		j,
		"uriOverrideInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudTasksQueueHttpTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudTasksQueueHttpTargetOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudTasksQueueHttpTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudTasksQueueHttpTargetOutputReference_Override(g GoogleCloudTasksQueueHttpTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudTasksQueue.GoogleCloudTasksQueueHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetInternalValue(val *GoogleCloudTasksQueueHttpTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) PutHeaderOverrides(value interface{}) {
	if err := g.validatePutHeaderOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderOverrides",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) PutOauthToken(value *GoogleCloudTasksQueueHttpTargetOauthToken) {
	if err := g.validatePutOauthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) PutOidcToken(value *GoogleCloudTasksQueueHttpTargetOidcToken) {
	if err := g.validatePutOidcTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOidcToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) PutUriOverride(value *GoogleCloudTasksQueueHttpTargetUriOverride) {
	if err := g.validatePutUriOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUriOverride",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ResetHeaderOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ResetOauthToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ResetOidcToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ResetUriOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetUriOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudTasksQueueHttpTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

