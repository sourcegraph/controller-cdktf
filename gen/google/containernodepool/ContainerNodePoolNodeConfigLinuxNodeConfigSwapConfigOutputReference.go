package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/containernodepool/internal"
)

type ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference interface {
	cdktn.ComplexObject
	BootDiskProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	BootDiskProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
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
	DedicatedLocalSsdProfile() ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	DedicatedLocalSsdProfileInput() *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
type jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference struct {
	internal.Type__cdktnComplexObject
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

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference_Override(c ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetInternalValue(val *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutDedicatedLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile) {
	if err := c.validatePutDedicatedLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDedicatedLocalSsdProfile",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEncryptionConfig(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig) {
	if err := c.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEphemeralLocalSsdProfile(value *ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile) {
	if err := c.validatePutEphemeralLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEphemeralLocalSsdProfile",
		[]interface{}{value},
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
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

