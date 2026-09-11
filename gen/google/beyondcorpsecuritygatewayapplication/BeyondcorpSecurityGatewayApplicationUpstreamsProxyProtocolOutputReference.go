package beyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/beyondcorpsecuritygatewayapplication/internal"
)

type BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference interface {
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
	ContextualHeaders() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
	ContextualHeadersInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
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
	InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
	SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol)
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
	PutContextualHeaders(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders)
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

// The jsii proxy struct for BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference
type jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) AllowedClientHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) AllowedClientHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedClientHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ClientIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ClientIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ContextualHeaders() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
	_jsii_.Get(
		j,
		"contextualHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ContextualHeadersInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	_jsii_.Get(
		j,
		"contextualHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GatewayIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GatewayIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) MetadataHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) MetadataHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference {
	_init_.Initialize()

	if err := validateNewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference_Override(b BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetAllowedClientHeaders(val *[]*string) {
	if err := j.validateSetAllowedClientHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClientHeaders",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetClientIp(val interface{}) {
	if err := j.validateSetClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientIp",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetGatewayIdentity(val *string) {
	if err := j.validateSetGatewayIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayIdentity",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocol) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetMetadataHeaders(val *map[string]*string) {
	if err := j.validateSetMetadataHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataHeaders",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) PutContextualHeaders(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders) {
	if err := b.validatePutContextualHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContextualHeaders",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetAllowedClientHeaders() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowedClientHeaders",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetClientIp() {
	_jsii_.InvokeVoid(
		b,
		"resetClientIp",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetContextualHeaders() {
	_jsii_.InvokeVoid(
		b,
		"resetContextualHeaders",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetGatewayIdentity() {
	_jsii_.InvokeVoid(
		b,
		"resetGatewayIdentity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ResetMetadataHeaders() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadataHeaders",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

