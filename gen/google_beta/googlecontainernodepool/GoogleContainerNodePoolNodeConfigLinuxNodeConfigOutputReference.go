package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference interface {
	cdktn.ComplexObject
	AccurateTimeConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	AccurateTimeConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig
	CgroupMode() *string
	SetCgroupMode(val *string)
	CgroupModeInput() *string
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
	// Experimental.
	Fqn() *string
	HugepagesConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	HugepagesConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	InternalValue() *GoogleContainerNodePoolNodeConfigLinuxNodeConfig
	SetInternalValue(val *GoogleContainerNodePoolNodeConfigLinuxNodeConfig)
	NodeKernelModuleLoading() GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	NodeKernelModuleLoadingInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	SwapConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
	SwapConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	Sysctls() *map[string]*string
	SetSysctls(val *map[string]*string)
	SysctlsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TransparentHugepageDefrag() *string
	SetTransparentHugepageDefrag(val *string)
	TransparentHugepageDefragInput() *string
	TransparentHugepageEnabled() *string
	SetTransparentHugepageEnabled(val *string)
	TransparentHugepageEnabledInput() *string
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
	PutAccurateTimeConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig)
	PutHugepagesConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig)
	PutNodeKernelModuleLoading(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading)
	PutSwapConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig)
	ResetAccurateTimeConfig()
	ResetCgroupMode()
	ResetHugepagesConfig()
	ResetNodeKernelModuleLoading()
	ResetSwapConfig()
	ResetSysctls()
	ResetTransparentHugepageDefrag()
	ResetTransparentHugepageEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	_jsii_.Get(
		j,
		"accurateTimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig
	_jsii_.Get(
		j,
		"accurateTimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) CgroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) CgroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) HugepagesConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	_jsii_.Get(
		j,
		"hugepagesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) HugepagesConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	_jsii_.Get(
		j,
		"hugepagesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) InternalValue() *GoogleContainerNodePoolNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoading() GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	_jsii_.Get(
		j,
		"nodeKernelModuleLoading",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoadingInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	_jsii_.Get(
		j,
		"nodeKernelModuleLoadingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) SwapConfig() GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfigOutputReference
	_jsii_.Get(
		j,
		"swapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) SwapConfigInput() *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig {
	var returns *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"swapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) Sysctls() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) SysctlsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabledInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference_Override(g GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetCgroupMode(val *string) {
	if err := j.validateSetCgroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetInternalValue(val *GoogleContainerNodePoolNodeConfigLinuxNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetSysctls(val *map[string]*string) {
	if err := j.validateSetSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sysctls",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageDefrag(val *string) {
	if err := j.validateSetTransparentHugepageDefragParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageDefrag",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageEnabled(val *string) {
	if err := j.validateSetTransparentHugepageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageEnabled",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) PutAccurateTimeConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig) {
	if err := g.validatePutAccurateTimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccurateTimeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) PutHugepagesConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig) {
	if err := g.validatePutHugepagesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHugepagesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) PutNodeKernelModuleLoading(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigNodeKernelModuleLoading) {
	if err := g.validatePutNodeKernelModuleLoadingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeKernelModuleLoading",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) PutSwapConfig(value *GoogleContainerNodePoolNodeConfigLinuxNodeConfigSwapConfig) {
	if err := g.validatePutSwapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSwapConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetAccurateTimeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccurateTimeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetCgroupMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCgroupMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetHugepagesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHugepagesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetNodeKernelModuleLoading() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeKernelModuleLoading",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetSwapConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSwapConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetSysctls() {
	_jsii_.InvokeVoid(
		g,
		"resetSysctls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageDefrag() {
	_jsii_.InvokeVoid(
		g,
		"resetTransparentHugepageDefrag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetTransparentHugepageEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigLinuxNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

