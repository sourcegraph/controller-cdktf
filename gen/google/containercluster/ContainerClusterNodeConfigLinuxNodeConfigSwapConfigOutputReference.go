package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/containercluster/internal"
)

type ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	BootDiskProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
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
	DedicatedLocalSsdProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	DedicatedLocalSsdProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	EncryptionConfig() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	EphemeralLocalSsdProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	EphemeralLocalSsdProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig
	SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig)
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
	PutBootDiskProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile)
	PutDedicatedLocalSsdProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile)
	PutEncryptionConfig(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig)
	PutEphemeralLocalSsdProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile)
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

// The jsii proxy struct for ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference
type jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	_jsii_.Get(
		j,
		"bootDiskProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	_jsii_.Get(
		j,
		"bootDiskProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfig() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfile() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfileInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) InternalValue() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference_Override(c ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutBootDiskProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile) {
	if err := c.validatePutBootDiskProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBootDiskProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutDedicatedLocalSsdProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile) {
	if err := c.validatePutDedicatedLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDedicatedLocalSsdProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEncryptionConfig(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig) {
	if err := c.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEphemeralLocalSsdProfile(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile) {
	if err := c.validatePutEphemeralLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEphemeralLocalSsdProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetBootDiskProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetDedicatedLocalSsdProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetDedicatedLocalSsdProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEphemeralLocalSsdProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetEphemeralLocalSsdProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
