package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/containercluster/internal"
)

type ContainerClusterNodeConfigLinuxNodeConfigOutputReference interface {
	cdktn.ComplexObject
	AccurateTimeConfig() ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	AccurateTimeConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig
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
	HugepagesConfig() ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	HugepagesConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig
	InternalValue() *ContainerClusterNodeConfigLinuxNodeConfig
	SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfig)
	NodeKernelModuleLoading() ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	NodeKernelModuleLoadingInput() *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	SwapConfig() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference
	SwapConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig
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
	PutAccurateTimeConfig(value *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig)
	PutHugepagesConfig(value *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig)
	PutNodeKernelModuleLoading(value *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading)
	PutSwapConfig(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig)
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

// The jsii proxy struct for ContainerClusterNodeConfigLinuxNodeConfigOutputReference
type jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfig() ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfigOutputReference
	_jsii_.Get(
		j,
		"accurateTimeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) AccurateTimeConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig
	_jsii_.Get(
		j,
		"accurateTimeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) CgroupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) CgroupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cgroupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) HugepagesConfig() ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
	_jsii_.Get(
		j,
		"hugepagesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) HugepagesConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig
	_jsii_.Get(
		j,
		"hugepagesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) InternalValue() *ContainerClusterNodeConfigLinuxNodeConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoading() ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoadingOutputReference
	_jsii_.Get(
		j,
		"nodeKernelModuleLoading",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) NodeKernelModuleLoadingInput() *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading
	_jsii_.Get(
		j,
		"nodeKernelModuleLoadingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) SwapConfig() ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigSwapConfigOutputReference
	_jsii_.Get(
		j,
		"swapConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) SwapConfigInput() *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig
	_jsii_.Get(
		j,
		"swapConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) Sysctls() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) SysctlsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefrag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefrag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageDefragInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageDefragInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) TransparentHugepageEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transparentHugepageEnabledInput",
		&returns,
	)
	return returns
}


func NewContainerClusterNodeConfigLinuxNodeConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodeConfigLinuxNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodeConfigLinuxNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodeConfigLinuxNodeConfigOutputReference_Override(c ContainerClusterNodeConfigLinuxNodeConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.containerCluster.ContainerClusterNodeConfigLinuxNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetCgroupMode(val *string) {
	if err := j.validateSetCgroupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cgroupMode",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetInternalValue(val *ContainerClusterNodeConfigLinuxNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetSysctls(val *map[string]*string) {
	if err := j.validateSetSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sysctls",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageDefrag(val *string) {
	if err := j.validateSetTransparentHugepageDefragParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageDefrag",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference)SetTransparentHugepageEnabled(val *string) {
	if err := j.validateSetTransparentHugepageEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transparentHugepageEnabled",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutAccurateTimeConfig(value *ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig) {
	if err := c.validatePutAccurateTimeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAccurateTimeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutHugepagesConfig(value *ContainerClusterNodeConfigLinuxNodeConfigHugepagesConfig) {
	if err := c.validatePutHugepagesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHugepagesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutNodeKernelModuleLoading(value *ContainerClusterNodeConfigLinuxNodeConfigNodeKernelModuleLoading) {
	if err := c.validatePutNodeKernelModuleLoadingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeKernelModuleLoading",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) PutSwapConfig(value *ContainerClusterNodeConfigLinuxNodeConfigSwapConfig) {
	if err := c.validatePutSwapConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSwapConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetAccurateTimeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAccurateTimeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetCgroupMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCgroupMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetHugepagesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetHugepagesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetNodeKernelModuleLoading() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeKernelModuleLoading",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetSwapConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSwapConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetSysctls() {
	_jsii_.InvokeVoid(
		c,
		"resetSysctls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageDefrag() {
	_jsii_.InvokeVoid(
		c,
		"resetTransparentHugepageDefrag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ResetTransparentHugepageEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetTransparentHugepageEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodeConfigLinuxNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

