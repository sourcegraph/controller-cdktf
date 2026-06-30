package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/dataproccluster/internal"
)

type DataprocClusterClusterConfigMasterConfigOutputReference interface {
	cdktf.ComplexObject
	Accelerators() DataprocClusterClusterConfigMasterConfigAcceleratorsList
	AcceleratorsInput() any
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
	DiskConfig() DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference
	DiskConfigInput() *DataprocClusterClusterConfigMasterConfigDiskConfig
	// Experimental.
	Fqn() *string
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InstanceFlexibilityPolicy() DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyOutputReference
	InstanceFlexibilityPolicyInput() *DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicy
	InstanceNames() *[]*string
	InternalValue() *DataprocClusterClusterConfigMasterConfig
	SetInternalValue(val *DataprocClusterClusterConfigMasterConfig)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	NumInstances() *float64
	SetNumInstances(val *float64)
	NumInstancesInput() *float64
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
	PutAccelerators(value any)
	PutDiskConfig(value *DataprocClusterClusterConfigMasterConfigDiskConfig)
	PutInstanceFlexibilityPolicy(value *DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicy)
	ResetAccelerators()
	ResetDiskConfig()
	ResetImageUri()
	ResetInstanceFlexibilityPolicy()
	ResetMachineType()
	ResetMinCpuPlatform()
	ResetNumInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigMasterConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) Accelerators() DataprocClusterClusterConfigMasterConfigAcceleratorsList {
	var returns DataprocClusterClusterConfigMasterConfigAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) AcceleratorsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) DiskConfig() DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference {
	var returns DataprocClusterClusterConfigMasterConfigDiskConfigOutputReference
	_jsii_.Get(
		j,
		"diskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) DiskConfigInput() *DataprocClusterClusterConfigMasterConfigDiskConfig {
	var returns *DataprocClusterClusterConfigMasterConfigDiskConfig
	_jsii_.Get(
		j,
		"diskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InstanceFlexibilityPolicy() DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyOutputReference {
	var returns DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicyOutputReference
	_jsii_.Get(
		j,
		"instanceFlexibilityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InstanceFlexibilityPolicyInput() *DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicy {
	var returns *DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicy
	_jsii_.Get(
		j,
		"instanceFlexibilityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InstanceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InternalValue() *DataprocClusterClusterConfigMasterConfig {
	var returns *DataprocClusterClusterConfigMasterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) NumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) NumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataprocClusterClusterConfigMasterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigMasterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigMasterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigMasterConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigMasterConfigOutputReference_Override(d DataprocClusterClusterConfigMasterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigMasterConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetInternalValue(val *DataprocClusterClusterConfigMasterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetNumInstances(val *float64) {
	if err := j.validateSetNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numInstances",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) PutAccelerators(value any) {
	if err := d.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccelerators",
		[]any{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) PutDiskConfig(value *DataprocClusterClusterConfigMasterConfigDiskConfig) {
	if err := d.validatePutDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDiskConfig",
		[]any{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) PutInstanceFlexibilityPolicy(value *DataprocClusterClusterConfigMasterConfigInstanceFlexibilityPolicy) {
	if err := d.validatePutInstanceFlexibilityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInstanceFlexibilityPolicy",
		[]any{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		d,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetDiskConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetImageUri() {
	_jsii_.InvokeVoid(
		d,
		"resetImageUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetInstanceFlexibilityPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceFlexibilityPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		d,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ResetNumInstances() {
	_jsii_.InvokeVoid(
		d,
		"resetNumInstances",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigMasterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
