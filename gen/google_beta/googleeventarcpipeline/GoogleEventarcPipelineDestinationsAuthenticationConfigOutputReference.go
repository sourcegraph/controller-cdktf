package googleeventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleeventarcpipeline/internal"
)

type GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference interface {
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
	GoogleOidc() GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidcOutputReference
	GoogleOidcInput() *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc
	InternalValue() *GoogleEventarcPipelineDestinationsAuthenticationConfig
	SetInternalValue(val *GoogleEventarcPipelineDestinationsAuthenticationConfig)
	OauthToken() GoogleEventarcPipelineDestinationsAuthenticationConfigOauthTokenOutputReference
	OauthTokenInput() *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken
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
	PutGoogleOidc(value *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc)
	PutOauthToken(value *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken)
	ResetGoogleOidc()
	ResetOauthToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference
type jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GoogleOidc() GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidcOutputReference {
	var returns GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidcOutputReference
	_jsii_.Get(
		j,
		"googleOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GoogleOidcInput() *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc {
	var returns *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc
	_jsii_.Get(
		j,
		"googleOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) InternalValue() *GoogleEventarcPipelineDestinationsAuthenticationConfig {
	var returns *GoogleEventarcPipelineDestinationsAuthenticationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) OauthToken() GoogleEventarcPipelineDestinationsAuthenticationConfigOauthTokenOutputReference {
	var returns GoogleEventarcPipelineDestinationsAuthenticationConfigOauthTokenOutputReference
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) OauthTokenInput() *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken {
	var returns *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEventarcPipelineDestinationsAuthenticationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference_Override(g GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcPipeline.GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference)SetInternalValue(val *GoogleEventarcPipelineDestinationsAuthenticationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) PutGoogleOidc(value *GoogleEventarcPipelineDestinationsAuthenticationConfigGoogleOidc) {
	if err := g.validatePutGoogleOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleOidc",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) PutOauthToken(value *GoogleEventarcPipelineDestinationsAuthenticationConfigOauthToken) {
	if err := g.validatePutOauthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthToken",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ResetGoogleOidc() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleOidc",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ResetOauthToken() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEventarcPipelineDestinationsAuthenticationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

