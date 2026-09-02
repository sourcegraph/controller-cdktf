package googlebeyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebeyondcorpsecuritygatewayapplication/internal"
)

type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference interface {
	cdktn.ComplexObject
	AllowedClientHeaders() *[]*string
	SetAllowedClientHeaders(val *[]*string)
	AllowedClientHeadersInput() *[]*string
	ClientIp() interface{}
	SetClientIp(val interface{})
	ClientIpInput() interface{}
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
	ContextualHeaders() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
	ContextualHeadersInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GatewayIdentity() *string
	SetGatewayIdentity(val *string)
	GatewayIdentityInput() *string
	InternalValue() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
	SetInternalValue(val *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol)
	MetadataHeaders() *map[string]*string
	SetMetadataHeaders(val *map[string]*string)
	MetadataHeadersInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutContextualHeaders(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders)
	ResetAllowedClientHeaders()
	ResetClientIp()
	ResetContextualHeaders()
	ResetGatewayIdentity()
	ResetMetadataHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference
type jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) AllowedClientHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) AllowedClientHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ClientIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ClientIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ContextualHeaders() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
	_jsii_.Get(
		j,
		"contextualHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ContextualHeadersInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	_jsii_.Get(
		j,
		"contextualHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GatewayIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GatewayIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InternalValue() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) MetadataHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) MetadataHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBeyondcorpSecurityGatewayApplication.GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference_Override(g GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBeyondcorpSecurityGatewayApplication.GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetAllowedClientHeaders(val *[]*string) {
	if err := j.validateSetAllowedClientHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClientHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetClientIp(val interface{}) {
	if err := j.validateSetClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIp",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetGatewayIdentity(val *string) {
	if err := j.validateSetGatewayIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIdentity",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetInternalValue(val *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetMetadataHeaders(val *map[string]*string) {
	if err := j.validateSetMetadataHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) PutContextualHeaders(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders) {
	if err := g.validatePutContextualHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContextualHeaders",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetAllowedClientHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedClientHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetClientIp() {
	_jsii_.InvokeVoid(
		g,
		"resetClientIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetContextualHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetContextualHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetGatewayIdentity() {
	_jsii_.InvokeVoid(
		g,
		"resetGatewayIdentity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetMetadataHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

