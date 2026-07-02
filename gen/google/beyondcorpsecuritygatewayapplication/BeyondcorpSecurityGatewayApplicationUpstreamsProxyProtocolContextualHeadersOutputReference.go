package beyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/beyondcorpsecuritygatewayapplication/internal"
)

type BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	DeviceInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference
	DeviceInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo
	// Experimental.
	Fqn() *string
	GroupInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference
	GroupInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo
	InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders)
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference
	UserInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutDeviceInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo)
	PutGroupInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo)
	PutUserInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo)
	ResetDeviceInfo()
	ResetGroupInfo()
	ResetOutputType()
	ResetUserInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
type jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) DeviceInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference
	_jsii_.Get(
		j,
		"deviceInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) DeviceInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo
	_jsii_.Get(
		j,
		"deviceInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GroupInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference
	_jsii_.Get(
		j,
		"groupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GroupInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo
	_jsii_.Get(
		j,
		"groupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InternalValue() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) UserInfo() BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference {
	var returns BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference
	_jsii_.Get(
		j,
		"userInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) UserInfoInput() *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo {
	var returns *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
	_jsii_.Get(
		j,
		"userInfoInput",
		&returns,
	)
	return returns
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference_Override(b BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGatewayApplication.BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetInternalValue(val *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutDeviceInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo) {
	if err := b.validatePutDeviceInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDeviceInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutGroupInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo) {
	if err := b.validatePutGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGroupInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutUserInfo(value *BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo) {
	if err := b.validatePutUserInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetDeviceInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetDeviceInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetGroupInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetGroupInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetUserInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetUserInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
