package migrationcenterpreferenceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/migrationcenterpreferenceset/internal"
)

type MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences)
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	MachinePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference
	MachinePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences
	PersistentDiskType() *string
	SetPersistentDiskType(val *string)
	PersistentDiskTypeInput() *string
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
	PutMachinePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences)
	ResetLicenseType()
	ResetMachinePreferences()
	ResetPersistentDiskType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference
type jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) MachinePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference {
	var returns MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesOutputReference
	_jsii_.Get(
		j,
		"machinePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) MachinePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences
	_jsii_.Get(
		j,
		"machinePreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) PersistentDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) PersistentDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"persistentDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference_Override(m MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetPersistentDiskType(val *string) {
	if err := j.validateSetPersistentDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"persistentDiskType",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) PutMachinePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferences) {
	if err := m.validatePutMachinePreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMachinePreferences",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ResetLicenseType() {
	_jsii_.InvokeVoid(
		m,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ResetMachinePreferences() {
	_jsii_.InvokeVoid(
		m,
		"resetMachinePreferences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ResetPersistentDiskType() {
	_jsii_.InvokeVoid(
		m,
		"resetPersistentDiskType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

