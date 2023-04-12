package googlenotebooksruntime

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlenotebooksruntime/internal"
)

type GoogleNotebooksRuntimeSoftwareConfigOutputReference interface {
	cdktf.ComplexObject
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
	CustomGpuDriverPath() *string
	SetCustomGpuDriverPath(val *string)
	CustomGpuDriverPathInput() *string
	EnableHealthMonitoring() interface{}
	SetEnableHealthMonitoring(val interface{})
	EnableHealthMonitoringInput() interface{}
	// Experimental.
	Fqn() *string
	IdleShutdown() interface{}
	SetIdleShutdown(val interface{})
	IdleShutdownInput() interface{}
	IdleShutdownTimeout() *float64
	SetIdleShutdownTimeout(val *float64)
	IdleShutdownTimeoutInput() *float64
	InstallGpuDriver() interface{}
	SetInstallGpuDriver(val interface{})
	InstallGpuDriverInput() interface{}
	InternalValue() *GoogleNotebooksRuntimeSoftwareConfig
	SetInternalValue(val *GoogleNotebooksRuntimeSoftwareConfig)
	Kernels() GoogleNotebooksRuntimeSoftwareConfigKernelsList
	KernelsInput() interface{}
	NotebookUpgradeSchedule() *string
	SetNotebookUpgradeSchedule(val *string)
	NotebookUpgradeScheduleInput() *string
	PostStartupScript() *string
	SetPostStartupScript(val *string)
	PostStartupScriptBehavior() *string
	SetPostStartupScriptBehavior(val *string)
	PostStartupScriptBehaviorInput() *string
	PostStartupScriptInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Upgradeable() cdktf.IResolvable
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
	PutKernels(value interface{})
	ResetCustomGpuDriverPath()
	ResetEnableHealthMonitoring()
	ResetIdleShutdown()
	ResetIdleShutdownTimeout()
	ResetInstallGpuDriver()
	ResetKernels()
	ResetNotebookUpgradeSchedule()
	ResetPostStartupScript()
	ResetPostStartupScriptBehavior()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNotebooksRuntimeSoftwareConfigOutputReference
type jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) CustomGpuDriverPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) CustomGpuDriverPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) EnableHealthMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) EnableHealthMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) IdleShutdown() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idleShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleShutdownTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) IdleShutdownTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleShutdownTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) InstallGpuDriver() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) InstallGpuDriverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) InternalValue() *GoogleNotebooksRuntimeSoftwareConfig {
	var returns *GoogleNotebooksRuntimeSoftwareConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) Kernels() GoogleNotebooksRuntimeSoftwareConfigKernelsList {
	var returns GoogleNotebooksRuntimeSoftwareConfigKernelsList
	_jsii_.Get(
		j,
		"kernels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) KernelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kernelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) NotebookUpgradeSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookUpgradeSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) NotebookUpgradeScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookUpgradeScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) PostStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) PostStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) Upgradeable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"upgradeable",
		&returns,
	)
	return returns
}


func NewGoogleNotebooksRuntimeSoftwareConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNotebooksRuntimeSoftwareConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNotebooksRuntimeSoftwareConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleNotebooksRuntime.GoogleNotebooksRuntimeSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNotebooksRuntimeSoftwareConfigOutputReference_Override(g GoogleNotebooksRuntimeSoftwareConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleNotebooksRuntime.GoogleNotebooksRuntimeSoftwareConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetCustomGpuDriverPath(val *string) {
	if err := j.validateSetCustomGpuDriverPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customGpuDriverPath",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetEnableHealthMonitoring(val interface{}) {
	if err := j.validateSetEnableHealthMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHealthMonitoring",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetIdleShutdown(val interface{}) {
	if err := j.validateSetIdleShutdownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleShutdown",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetIdleShutdownTimeout(val *float64) {
	if err := j.validateSetIdleShutdownTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleShutdownTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetInstallGpuDriver(val interface{}) {
	if err := j.validateSetInstallGpuDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installGpuDriver",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetInternalValue(val *GoogleNotebooksRuntimeSoftwareConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetNotebookUpgradeSchedule(val *string) {
	if err := j.validateSetNotebookUpgradeScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookUpgradeSchedule",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetPostStartupScript(val *string) {
	if err := j.validateSetPostStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetPostStartupScriptBehavior(val *string) {
	if err := j.validateSetPostStartupScriptBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScriptBehavior",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) PutKernels(value interface{}) {
	if err := g.validatePutKernelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putKernels",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetCustomGpuDriverPath() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomGpuDriverPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetEnableHealthMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableHealthMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetIdleShutdown() {
	_jsii_.InvokeVoid(
		g,
		"resetIdleShutdown",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetIdleShutdownTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetIdleShutdownTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetInstallGpuDriver() {
	_jsii_.InvokeVoid(
		g,
		"resetInstallGpuDriver",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetKernels() {
	_jsii_.InvokeVoid(
		g,
		"resetKernels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetNotebookUpgradeSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebookUpgradeSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetPostStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetPostStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ResetPostStartupScriptBehavior() {
	_jsii_.InvokeVoid(
		g,
		"resetPostStartupScriptBehavior",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNotebooksRuntimeSoftwareConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

