package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainercluster/internal"
)

type GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference interface {
	cdktn.ComplexObject
	AccurateTimeConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	AccurateTimeConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig
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
	HugepagesConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	HugepagesConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig
	InternalValue() *GoogleContainerClusterNodeConfigLinuxNodeConfig
	SetInternalValue(val *GoogleContainerClusterNodeConfigLinuxNodeConfig)
	NodeKernelModuleLoading() GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	NodeKernelModuleLoadingInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	SwapConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference
	SwapConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfig
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
	PutAccurateTimeConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig)
	PutHugepagesConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig)
	PutNodeKernelModuleLoading(value *GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading)
	PutSwapConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfig)
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

// The jsii proxy struct for GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference
type jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	_jsii_.Get(
		j,
		"accurateTimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig
	_jsii_.Get(
		j,
		"accurateTimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) CgroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) CgroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) HugepagesConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	_jsii_.Get(
		j,
		"hugepagesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) HugepagesConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig
	_jsii_.Get(
		j,
		"hugepagesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) InternalValue() *GoogleContainerClusterNodeConfigLinuxNodeConfig {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoading() GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference {
	var returns GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	_jsii_.Get(
		j,
		"nodeKernelModuleLoading",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoadingInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	_jsii_.Get(
		j,
		"nodeKernelModuleLoadingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) SwapConfig() GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference
	_jsii_.Get(
		j,
		"swapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) SwapConfigInput() *GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfig {
	var returns *GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"swapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) Sysctls() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) SysctlsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabledInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodeConfigLinuxNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference_Override(g GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetCgroupMode(val *string) {
	if err := j.validateSetCgroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupMode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetInternalValue(val *GoogleContainerClusterNodeConfigLinuxNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetSysctls(val *map[string]*string) {
	if err := j.validateSetSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sysctls",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageDefrag(val *string) {
	if err := j.validateSetTransparentHugepageDefragParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageDefrag",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageEnabled(val *string) {
	if err := j.validateSetTransparentHugepageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageEnabled",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutAccurateTimeConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig) {
	if err := g.validatePutAccurateTimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccurateTimeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutHugepagesConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig) {
	if err := g.validatePutHugepagesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHugepagesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutNodeKernelModuleLoading(value *GoogleContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading) {
	if err := g.validatePutNodeKernelModuleLoadingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeKernelModuleLoading",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutSwapConfig(value *GoogleContainerClusterNodeConfigLinuxNodeConfigSwapConfig) {
	if err := g.validatePutSwapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSwapConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetAccurateTimeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccurateTimeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetCgroupMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCgroupMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetHugepagesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHugepagesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetNodeKernelModuleLoading() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeKernelModuleLoading",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetSwapConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSwapConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetSysctls() {
	_jsii_.InvokeVoid(
		g,
		"resetSysctls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageDefrag() {
	_jsii_.InvokeVoid(
		g,
		"resetTransparentHugepageDefrag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetTransparentHugepageEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigLinuxNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

