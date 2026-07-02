package connectquickconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/connectquickconnect/internal"
)

type ConnectQuickConnectQuickConnectConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ConnectQuickConnectQuickConnectConfig
	SetInternalValue(val *ConnectQuickConnectQuickConnectConfig)
	PhoneConfig() ConnectQuickConnectQuickConnectConfigPhoneConfigList
	PhoneConfigInput() any
	QueueConfig() ConnectQuickConnectQuickConnectConfigQueueConfigList
	QueueConfigInput() any
	QuickConnectType() *string
	SetQuickConnectType(val *string)
	QuickConnectTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserConfig() ConnectQuickConnectQuickConnectConfigUserConfigList
	UserConfigInput() any
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
	PutPhoneConfig(value any)
	PutQueueConfig(value any)
	PutUserConfig(value any)
	ResetPhoneConfig()
	ResetQueueConfig()
	ResetUserConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConnectQuickConnectQuickConnectConfigOutputReference
type jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) InternalValue() *ConnectQuickConnectQuickConnectConfig {
	var returns *ConnectQuickConnectQuickConnectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) PhoneConfig() ConnectQuickConnectQuickConnectConfigPhoneConfigList {
	var returns ConnectQuickConnectQuickConnectConfigPhoneConfigList
	_jsii_.Get(
		j,
		"phoneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) PhoneConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"phoneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) QueueConfig() ConnectQuickConnectQuickConnectConfigQueueConfigList {
	var returns ConnectQuickConnectQuickConnectConfigQueueConfigList
	_jsii_.Get(
		j,
		"queueConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) QueueConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"queueConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) QuickConnectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) QuickConnectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickConnectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) UserConfig() ConnectQuickConnectQuickConnectConfigUserConfigList {
	var returns ConnectQuickConnectQuickConnectConfigUserConfigList
	_jsii_.Get(
		j,
		"userConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) UserConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"userConfigInput",
		&returns,
	)
	return returns
}

func NewConnectQuickConnectQuickConnectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ConnectQuickConnectQuickConnectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewConnectQuickConnectQuickConnectConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.connectQuickConnect.ConnectQuickConnectQuickConnectConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConnectQuickConnectQuickConnectConfigOutputReference_Override(c ConnectQuickConnectQuickConnectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.connectQuickConnect.ConnectQuickConnectQuickConnectConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetInternalValue(val *ConnectQuickConnectQuickConnectConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetQuickConnectType(val *string) {
	if err := j.validateSetQuickConnectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quickConnectType",
		val,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) PutPhoneConfig(value any) {
	if err := c.validatePutPhoneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPhoneConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) PutQueueConfig(value any) {
	if err := c.validatePutQueueConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueueConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) PutUserConfig(value any) {
	if err := c.validatePutUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUserConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ResetPhoneConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPhoneConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ResetQueueConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetQueueConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ResetUserConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetUserConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectQuickConnectQuickConnectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
