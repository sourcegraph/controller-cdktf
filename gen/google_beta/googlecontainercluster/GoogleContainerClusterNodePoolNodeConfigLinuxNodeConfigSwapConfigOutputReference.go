package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	BootDiskProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
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
	DedicatedLocalSsdProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	DedicatedLocalSsdProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncryptionConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	EphemeralLocalSsdProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	EphemeralLocalSsdProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig
	SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig)
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
	PutBootDiskProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile)
	PutDedicatedLocalSsdProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile)
	PutEncryptionConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig)
	PutEphemeralLocalSsdProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile)
	ResetBootDiskProfile()
	ResetDedicatedLocalSsdProfile()
	ResetEnabled()
	ResetEncryptionConfig()
	ResetEphemeralLocalSsdProfile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
type jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfileOutputReference
	_jsii_.Get(
		j,
		"bootDiskProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) BootDiskProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile
	_jsii_.Get(
		j,
		"bootDiskProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) DedicatedLocalSsdProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile
	_jsii_.Get(
		j,
		"dedicatedLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfig() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EncryptionConfigInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfile() GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference {
	var returns GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfileOutputReference
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) EphemeralLocalSsdProfileInput() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile
	_jsii_.Get(
		j,
		"ephemeralLocalSsdProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InternalValue() *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig {
	var returns *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference_Override(g GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetInternalValue(val *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutBootDiskProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigBootDiskProfile) {
	if err := g.validatePutBootDiskProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDiskProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutDedicatedLocalSsdProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigDedicatedLocalSsdProfile) {
	if err := g.validatePutDedicatedLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDedicatedLocalSsdProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEncryptionConfig(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) PutEphemeralLocalSsdProfile(value *GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigEphemeralLocalSsdProfile) {
	if err := g.validatePutEphemeralLocalSsdProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEphemeralLocalSsdProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetBootDiskProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetDedicatedLocalSsdProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetDedicatedLocalSsdProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ResetEphemeralLocalSsdProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetEphemeralLocalSsdProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

