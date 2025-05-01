package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

type GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
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
	ContainerLogMaxFiles() *float64
	SetContainerLogMaxFiles(val *float64)
	ContainerLogMaxFilesInput() *float64
	ContainerLogMaxSize() *string
	SetContainerLogMaxSize(val *string)
	ContainerLogMaxSizeInput() *string
	CpuCfsQuota() interface{}
	SetCpuCfsQuota(val interface{})
	CpuCfsQuotaInput() interface{}
	CpuCfsQuotaPeriod() *string
	SetCpuCfsQuotaPeriod(val *string)
	CpuCfsQuotaPeriodInput() *string
	CpuManagerPolicy() *string
	SetCpuManagerPolicy(val *string)
	CpuManagerPolicyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImageGcHighThresholdPercent() *float64
	SetImageGcHighThresholdPercent(val *float64)
	ImageGcHighThresholdPercentInput() *float64
	ImageGcLowThresholdPercent() *float64
	SetImageGcLowThresholdPercent(val *float64)
	ImageGcLowThresholdPercentInput() *float64
	ImageMaximumGcAge() *string
	SetImageMaximumGcAge(val *string)
	ImageMaximumGcAgeInput() *string
	ImageMinimumGcAge() *string
	SetImageMinimumGcAge(val *string)
	ImageMinimumGcAgeInput() *string
	InsecureKubeletReadonlyPortEnabled() *string
	SetInsecureKubeletReadonlyPortEnabled(val *string)
	InsecureKubeletReadonlyPortEnabledInput() *string
	InternalValue() *GoogleContainerNodePoolNodeConfigKubeletConfig
	SetInternalValue(val *GoogleContainerNodePoolNodeConfigKubeletConfig)
	PodPidsLimit() *float64
	SetPodPidsLimit(val *float64)
	PodPidsLimitInput() *float64
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
	ResetAllowedUnsafeSysctls()
	ResetContainerLogMaxFiles()
	ResetContainerLogMaxSize()
	ResetCpuCfsQuota()
	ResetCpuCfsQuotaPeriod()
	ResetCpuManagerPolicy()
	ResetImageGcHighThresholdPercent()
	ResetImageGcLowThresholdPercent()
	ResetImageMaximumGcAge()
	ResetImageMinimumGcAge()
	ResetInsecureKubeletReadonlyPortEnabled()
	ResetPodPidsLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference
type jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerLogMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerLogMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuota() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CpuManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageGcHighThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageGcHighThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageGcLowThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageGcLowThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageMaximumGcAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMaximumGcAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageMaximumGcAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMaximumGcAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageMinimumGcAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMinimumGcAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ImageMinimumGcAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMinimumGcAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) InsecureKubeletReadonlyPortEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) InsecureKubeletReadonlyPortEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) InternalValue() *GoogleContainerNodePoolNodeConfigKubeletConfig {
	var returns *GoogleContainerNodePoolNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) PodPidsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podPidsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) PodPidsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podPidsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleContainerNodePoolNodeConfigKubeletConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolNodeConfigKubeletConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleContainerNodePoolNodeConfigKubeletConfigOutputReference_Override(g GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerNodePool.GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetAllowedUnsafeSysctls(val *[]*string) {
	if err := j.validateSetAllowedUnsafeSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetContainerLogMaxFiles(val *float64) {
	if err := j.validateSetContainerLogMaxFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxFiles",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetContainerLogMaxSize(val *string) {
	if err := j.validateSetContainerLogMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxSize",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetCpuCfsQuota(val interface{}) {
	if err := j.validateSetCpuCfsQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuota",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetCpuCfsQuotaPeriod(val *string) {
	if err := j.validateSetCpuCfsQuotaPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuotaPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetCpuManagerPolicy(val *string) {
	if err := j.validateSetCpuManagerPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManagerPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetImageGcHighThresholdPercent(val *float64) {
	if err := j.validateSetImageGcHighThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcHighThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetImageGcLowThresholdPercent(val *float64) {
	if err := j.validateSetImageGcLowThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcLowThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetImageMaximumGcAge(val *string) {
	if err := j.validateSetImageMaximumGcAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageMaximumGcAge",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetImageMinimumGcAge(val *string) {
	if err := j.validateSetImageMinimumGcAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageMinimumGcAge",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetInsecureKubeletReadonlyPortEnabled(val *string) {
	if err := j.validateSetInsecureKubeletReadonlyPortEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureKubeletReadonlyPortEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetInternalValue(val *GoogleContainerNodePoolNodeConfigKubeletConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetPodPidsLimit(val *float64) {
	if err := j.validateSetPodPidsLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podPidsLimit",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetContainerLogMaxFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerLogMaxFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetContainerLogMaxSize() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerLogMaxSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuCfsQuota() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuCfsQuota",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuCfsQuotaPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuCfsQuotaPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuManagerPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCpuManagerPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetImageGcHighThresholdPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetImageGcHighThresholdPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetImageGcLowThresholdPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetImageGcLowThresholdPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetImageMaximumGcAge() {
	_jsii_.InvokeVoid(
		g,
		"resetImageMaximumGcAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetImageMinimumGcAge() {
	_jsii_.InvokeVoid(
		g,
		"resetImageMinimumGcAge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetInsecureKubeletReadonlyPortEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetInsecureKubeletReadonlyPortEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ResetPodPidsLimit() {
	_jsii_.InvokeVoid(
		g,
		"resetPodPidsLimit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerNodePoolNodeConfigKubeletConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

