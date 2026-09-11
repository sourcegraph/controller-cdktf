package googlebeyondcorpsecuritygatewayapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebeyondcorpsecuritygatewayapplication/internal"
)

type GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference interface {
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
	DeviceInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference
	DeviceInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo
	// Experimental.
	Fqn() *string
	GroupInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference
	GroupInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo
	InternalValue() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	SetInternalValue(val *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders)
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
	UserInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference
	UserInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
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
	PutDeviceInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo)
	PutGroupInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo)
	PutUserInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo)
	ResetDeviceInfo()
	ResetGroupInfo()
	ResetOutputType()
	ResetUserInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference
type jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) DeviceInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfoOutputReference
	_jsii_.Get(
		j,
		"deviceInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) DeviceInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo
	_jsii_.Get(
		j,
		"deviceInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GroupInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfoOutputReference
	_jsii_.Get(
		j,
		"groupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GroupInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo
	_jsii_.Get(
		j,
		"groupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InternalValue() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) UserInfo() GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference {
	var returns GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfoOutputReference
	_jsii_.Get(
		j,
		"userInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) UserInfoInput() *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo {
	var returns *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo
	_jsii_.Get(
		j,
		"userInfoInput",
		&returns,
	)
	return returns
}


func NewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBeyondcorpSecurityGatewayApplication.GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference_Override(g GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBeyondcorpSecurityGatewayApplication.GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetInternalValue(val *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeaders) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutDeviceInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersDeviceInfo) {
	if err := g.validatePutDeviceInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeviceInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutGroupInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersGroupInfo) {
	if err := g.validatePutGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGroupInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) PutUserInfo(value *GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersUserInfo) {
	if err := g.validatePutUserInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserInfo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetDeviceInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetDeviceInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetGroupInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetGroupInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ResetUserInfo() {
	_jsii_.InvokeVoid(
		g,
		"resetUserInfo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBeyondcorpSecurityGatewayApplicationUpstreamsProxyProtocolContextualHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

