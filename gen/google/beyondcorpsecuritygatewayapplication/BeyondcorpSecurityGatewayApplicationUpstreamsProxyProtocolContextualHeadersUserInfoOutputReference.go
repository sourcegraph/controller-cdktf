package beyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/beyondcorpsecuritygatewayapplication/internal"
)

type BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
	SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo)
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
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
	ResetOutputType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference
type jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference {
	_init_.Initialize()

	if err := validateNewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference_Override(b BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

