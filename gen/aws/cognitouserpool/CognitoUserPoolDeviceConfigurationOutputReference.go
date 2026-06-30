package cognitouserpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/cognitouserpool/internal"
)

type CognitoUserPoolDeviceConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChallengeRequiredOnNewDevice() any
	SetChallengeRequiredOnNewDevice(val any)
	ChallengeRequiredOnNewDeviceInput() any
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
	DeviceOnlyRememberedOnUserPrompt() any
	SetDeviceOnlyRememberedOnUserPrompt(val any)
	DeviceOnlyRememberedOnUserPromptInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *CognitoUserPoolDeviceConfiguration
	SetInternalValue(val *CognitoUserPoolDeviceConfiguration)
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
	ResetChallengeRequiredOnNewDevice()
	ResetDeviceOnlyRememberedOnUserPrompt()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CognitoUserPoolDeviceConfigurationOutputReference
type jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ChallengeRequiredOnNewDevice() any {
	var returns any
	_jsii_.Get(
		j,
		"challengeRequiredOnNewDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ChallengeRequiredOnNewDeviceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"challengeRequiredOnNewDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) DeviceOnlyRememberedOnUserPrompt() any {
	var returns any
	_jsii_.Get(
		j,
		"deviceOnlyRememberedOnUserPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) DeviceOnlyRememberedOnUserPromptInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deviceOnlyRememberedOnUserPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) InternalValue() *CognitoUserPoolDeviceConfiguration {
	var returns *CognitoUserPoolDeviceConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCognitoUserPoolDeviceConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CognitoUserPoolDeviceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewCognitoUserPoolDeviceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolDeviceConfigurationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCognitoUserPoolDeviceConfigurationOutputReference_Override(c CognitoUserPoolDeviceConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUserPool.CognitoUserPoolDeviceConfigurationOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetChallengeRequiredOnNewDevice(val any) {
	if err := j.validateSetChallengeRequiredOnNewDeviceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"challengeRequiredOnNewDevice",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetDeviceOnlyRememberedOnUserPrompt(val any) {
	if err := j.validateSetDeviceOnlyRememberedOnUserPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOnlyRememberedOnUserPrompt",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetInternalValue(val *CognitoUserPoolDeviceConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ResetChallengeRequiredOnNewDevice() {
	_jsii_.InvokeVoid(
		c,
		"resetChallengeRequiredOnNewDevice",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ResetDeviceOnlyRememberedOnUserPrompt() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceOnlyRememberedOnUserPrompt",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
