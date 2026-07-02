package hypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/hypercomputeclustercluster/internal"
)

type HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference interface {
	cdktf.ComplexObject
	BootDisk() HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference
	BootDiskInput() *HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk
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
	Count() *string
	SetCount(val *string)
	CountInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableOsLogin() any
	SetEnableOsLogin(val any)
	EnableOsLoginInput() any
	EnablePublicIps() any
	SetEnablePublicIps(val any)
	EnablePublicIpsInput() any
	// Experimental.
	Fqn() *string
	Instances() HypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList
	InternalValue() *HypercomputeclusterClusterOrchestratorSlurmLoginNodes
	SetInternalValue(val *HypercomputeclusterClusterOrchestratorSlurmLoginNodes)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	StartupScript() *string
	SetStartupScript(val *string)
	StartupScriptInput() *string
	StorageConfigs() HypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList
	StorageConfigsInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutBootDisk(value *HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk)
	PutStorageConfigs(value any)
	ResetBootDisk()
	ResetEnableOsLogin()
	ResetEnablePublicIps()
	ResetLabels()
	ResetStartupScript()
	ResetStorageConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
type jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) BootDisk() HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference {
	var returns HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) BootDiskInput() *HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk {
	var returns *HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Count() *string {
	var returns *string
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) CountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnableOsLogin() any {
	var returns any
	_jsii_.Get(
		j,
		"enableOsLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnableOsLoginInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableOsLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnablePublicIps() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnablePublicIpsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enablePublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Instances() HypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList {
	var returns HypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InternalValue() *HypercomputeclusterClusterOrchestratorSlurmLoginNodes {
	var returns *HypercomputeclusterClusterOrchestratorSlurmLoginNodes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StorageConfigs() HypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList {
	var returns HypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList
	_jsii_.Get(
		j,
		"storageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StorageConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"storageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}

func NewHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference {
	_init_.Initialize()

	if err := validateNewHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference_Override(h HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.hypercomputeclusterCluster.HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference",
		[]any{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetCount(val *string) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetEnableOsLogin(val any) {
	if err := j.validateSetEnableOsLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOsLogin",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetEnablePublicIps(val any) {
	if err := j.validateSetEnablePublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicIps",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetInternalValue(val *HypercomputeclusterClusterOrchestratorSlurmLoginNodes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetStartupScript(val *string) {
	if err := j.validateSetStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScript",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) PutBootDisk(value *HypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk) {
	if err := h.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putBootDisk",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) PutStorageConfigs(value any) {
	if err := h.validatePutStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStorageConfigs",
		[]any{value},
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		h,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetEnableOsLogin() {
	_jsii_.InvokeVoid(
		h,
		"resetEnableOsLogin",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetEnablePublicIps() {
	_jsii_.InvokeVoid(
		h,
		"resetEnablePublicIps",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		h,
		"resetLabels",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetStartupScript() {
	_jsii_.InvokeVoid(
		h,
		"resetStartupScript",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetStorageConfigs() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageConfigs",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		h,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
