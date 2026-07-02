package beyondcorpsecuritygateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/beyondcorpsecuritygateway/internal"
)

type BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference interface {
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
	DeviceInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference
	DeviceInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo
	// Experimental.
	Fqn() *string
	GroupInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference
	GroupInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo
	InternalValue() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders
	SetInternalValue(val *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders)
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
	UserInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference
	UserInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo
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
	PutDeviceInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo)
	PutGroupInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo)
	PutUserInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo)
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

// The jsii proxy struct for BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference
type jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) DeviceInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference {
	var returns BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference
	_jsii_.Get(
		j,
		"deviceInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) DeviceInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo {
	var returns *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo
	_jsii_.Get(
		j,
		"deviceInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GroupInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference {
	var returns BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference
	_jsii_.Get(
		j,
		"groupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GroupInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo {
	var returns *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo
	_jsii_.Get(
		j,
		"groupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InternalValue() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders {
	var returns *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) UserInfo() BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference {
	var returns BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference
	_jsii_.Get(
		j,
		"userInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) UserInfoInput() *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo {
	var returns *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo
	_jsii_.Get(
		j,
		"userInfoInput",
		&returns,
	)
	return returns
}

func NewBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGateway.BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference_Override(b BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.beyondcorpSecurityGateway.BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetInternalValue(val *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutDeviceInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo) {
	if err := b.validatePutDeviceInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDeviceInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutGroupInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo) {
	if err := b.validatePutGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGroupInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutUserInfo(value *BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo) {
	if err := b.validatePutUserInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserInfo",
		[]any{value},
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetDeviceInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetDeviceInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetGroupInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetGroupInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetUserInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetUserInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (b *jsiiProxy_BeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
