package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference interface {
	cdktn.ComplexObject
	BootDiskProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	BootDiskProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
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
	DedicatedLocalSsdProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	DedicatedLocalSsdProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncryptionConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	EphemeralLocalSsdProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	EphemeralLocalSsdProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	SetInternalValue(val *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig)
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
	PutBootDiskProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile)
	PutDedicatedLocalSsdProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile)
	PutEncryptionConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig)
	PutEphemeralLocalSsdProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile)
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

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	_jsii_.Get(
		j,
		"bootDiskProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	_jsii_.Get(
		j,
		"bootDiskProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfile() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfileInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InternalValue() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference_Override(g GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetInternalValue(val *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutBootDiskProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile) {
	if err := g.validatePutBootDiskProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDiskProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutDedicatedLocalSsdProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile) {
	if err := g.validatePutDedicatedLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDedicatedLocalSsdProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEncryptionConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEphemeralLocalSsdProfile(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile) {
	if err := g.validatePutEphemeralLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralLocalSsdProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetBootDiskProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetDedicatedLocalSsdProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedLocalSsdProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEphemeralLocalSsdProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralLocalSsdProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

