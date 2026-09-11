package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference interface {
	cdktn.ComplexObject
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
	EnableNestedVirtualization() interface{}
	SetEnableNestedVirtualization(val interface{})
	EnableNestedVirtualizationInput() interface{}
	EnableUefiNetworking() interface{}
	SetEnableUefiNetworking(val interface{})
	EnableUefiNetworkingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
	VisibleCoreCount() *float64
	SetVisibleCoreCount(val *float64)
	VisibleCoreCountInput() *float64
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
	ResetEnableNestedVirtualization()
	ResetEnableUefiNetworking()
	ResetThreadsPerCore()
	ResetVisibleCoreCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference
type jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) EnableNestedVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) EnableNestedVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) EnableUefiNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUefiNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) EnableUefiNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUefiNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) VisibleCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) VisibleCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCountInput",
		&returns,
	)
	return returns
}


func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference_Override(b BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetEnableNestedVirtualization(val interface{}) {
	if err := j.validateSetEnableNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetEnableUefiNetworking(val interface{}) {
	if err := j.validateSetEnableUefiNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUefiNetworking",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference)SetVisibleCoreCount(val *float64) {
	if err := j.validateSetVisibleCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleCoreCount",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ResetEnableNestedVirtualization() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableNestedVirtualization",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ResetEnableUefiNetworking() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableUefiNetworking",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		b,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ResetVisibleCoreCount() {
	_jsii_.InvokeVoid(
		b,
		"resetVisibleCoreCount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

