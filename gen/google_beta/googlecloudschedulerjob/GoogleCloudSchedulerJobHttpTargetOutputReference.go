package googlecloudschedulerjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudschedulerjob/internal"
)

type GoogleCloudSchedulerJobHttpTargetOutputReference interface {
	cdktf.ComplexObject
	Body() *string
	SetBody(val *string)
	BodyInput() *string
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
	Headers() *map[string]*string
	SetHeaders(val *map[string]*string)
	HeadersInput() *map[string]*string
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() *GoogleCloudSchedulerJobHttpTarget
	SetInternalValue(val *GoogleCloudSchedulerJobHttpTarget)
	OauthToken() GoogleCloudSchedulerJobHttpTargetOauthTokenOutputReference
	OauthTokenInput() *GoogleCloudSchedulerJobHttpTargetOauthToken
	OidcToken() GoogleCloudSchedulerJobHttpTargetOidcTokenOutputReference
	OidcTokenInput() *GoogleCloudSchedulerJobHttpTargetOidcToken
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	PutOauthToken(value *GoogleCloudSchedulerJobHttpTargetOauthToken)
	PutOidcToken(value *GoogleCloudSchedulerJobHttpTargetOidcToken)
	ResetBody()
	ResetHeaders()
	ResetHttpMethod()
	ResetOauthToken()
	ResetOidcToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudSchedulerJobHttpTargetOutputReference
type jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) BodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) Headers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) HeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) InternalValue() *GoogleCloudSchedulerJobHttpTarget {
	var returns *GoogleCloudSchedulerJobHttpTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) OauthToken() GoogleCloudSchedulerJobHttpTargetOauthTokenOutputReference {
	var returns GoogleCloudSchedulerJobHttpTargetOauthTokenOutputReference
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) OauthTokenInput() *GoogleCloudSchedulerJobHttpTargetOauthToken {
	var returns *GoogleCloudSchedulerJobHttpTargetOauthToken
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) OidcToken() GoogleCloudSchedulerJobHttpTargetOidcTokenOutputReference {
	var returns GoogleCloudSchedulerJobHttpTargetOidcTokenOutputReference
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) OidcTokenInput() *GoogleCloudSchedulerJobHttpTargetOidcToken {
	var returns *GoogleCloudSchedulerJobHttpTargetOidcToken
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudSchedulerJobHttpTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudSchedulerJobHttpTargetOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudSchedulerJobHttpTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference{}

	_jsii_.Create(
		"google-beta.googleCloudSchedulerJob.GoogleCloudSchedulerJobHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudSchedulerJobHttpTargetOutputReference_Override(g GoogleCloudSchedulerJobHttpTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudSchedulerJob.GoogleCloudSchedulerJobHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetBody(val *string) {
	if err := j.validateSetBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetHeaders(val *map[string]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetInternalValue(val *GoogleCloudSchedulerJobHttpTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) PutOauthToken(value *GoogleCloudSchedulerJobHttpTargetOauthToken) {
	if err := g.validatePutOauthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) PutOidcToken(value *GoogleCloudSchedulerJobHttpTargetOidcToken) {
	if err := g.validatePutOidcTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOidcToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ResetBody() {
	_jsii_.InvokeVoid(
		g,
		"resetBody",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ResetOauthToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ResetOidcToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCloudSchedulerJobHttpTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

