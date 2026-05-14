package googlebeyondcorpsecuritygateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebeyondcorpsecuritygateway/internal"
)

type GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference interface {
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
	DeviceInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference
	DeviceInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo
	// Experimental.
	Fqn() *string
	GroupInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference
	GroupInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo
	InternalValue() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders
	SetInternalValue(val *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders)
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
	UserInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference
	UserInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo
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
	PutDeviceInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo)
	PutGroupInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo)
	PutUserInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo)
	ResetDeviceInfo()
	ResetGroupInfo()
	ResetOutputType()
	ResetUserInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference
type jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) DeviceInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfoOutputReference
	_jsii_.Get(
		j,
		"deviceInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) DeviceInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo {
	var returns *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo
	_jsii_.Get(
		j,
		"deviceInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GroupInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfoOutputReference
	_jsii_.Get(
		j,
		"groupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GroupInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo {
	var returns *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo
	_jsii_.Get(
		j,
		"groupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InternalValue() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders {
	var returns *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) UserInfo() GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfoOutputReference
	_jsii_.Get(
		j,
		"userInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) UserInfoInput() *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo {
	var returns *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo
	_jsii_.Get(
		j,
		"userInfoInput",
		&returns,
	)
	return returns
}


func NewGoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBeyondcorpSecurityGateway.GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference_Override(g GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBeyondcorpSecurityGateway.GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetInternalValue(val *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeaders) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutDeviceInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersDeviceInfo) {
	if err := g.validatePutDeviceInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeviceInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutGroupInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersGroupInfo) {
	if err := g.validatePutGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGroupInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) PutUserInfo(value *GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersUserInfo) {
	if err := g.validatePutUserInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetDeviceInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetDeviceInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetGroupInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ResetUserInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetUserInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayProxyProtocolConfigContextualHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

