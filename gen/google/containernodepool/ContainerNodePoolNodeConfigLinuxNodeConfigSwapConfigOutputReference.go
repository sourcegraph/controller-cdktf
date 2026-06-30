package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/containernodepool/internal"
)

type ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	BootDiskProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
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
	DedicatedLocalSsdProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	DedicatedLocalSsdProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	EncryptionConfig() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	EphemeralLocalSsdProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	EphemeralLocalSsdProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	SetInternalValue(val *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig)
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
	PutBootDiskProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile)
	PutDedicatedLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile)
	PutEncryptionConfig(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig)
	PutEphemeralLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile)
	ResetBootDiskProfile()
	ResetDedicatedLocalSsdProfile()
	ResetEnabled()
	ResetEncryptionConfig()
	ResetEphemeralLocalSsdProfile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
type jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference {
	var returns ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	_jsii_.Get(
		j,
		"bootDiskProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	_jsii_.Get(
		j,
		"bootDiskProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference {
	var returns ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfig() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference {
	var returns ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfigInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference {
	var returns ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InternalValue() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference_Override(c ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetInternalValue(val *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutBootDiskProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile) {
	if err := c.validatePutBootDiskProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBootDiskProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutDedicatedLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile) {
	if err := c.validatePutDedicatedLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDedicatedLocalSsdProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEncryptionConfig(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig) {
	if err := c.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEphemeralLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile) {
	if err := c.validatePutEphemeralLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEphemeralLocalSsdProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetBootDiskProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetDedicatedLocalSsdProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetDedicatedLocalSsdProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEphemeralLocalSsdProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetEphemeralLocalSsdProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
