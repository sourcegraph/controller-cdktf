package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference interface {
	cdktf.ComplexObject
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
	EnableIntegrityMonitoring() any
	SetEnableIntegrityMonitoring(val any)
	EnableIntegrityMonitoringInput() any
	EnableSecureBoot() any
	SetEnableSecureBoot(val any)
	EnableSecureBootInput() any
	EnableVtpm() any
	SetEnableVtpm(val any)
	EnableVtpmInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig)
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
	ResetEnableIntegrityMonitoring()
	ResetEnableSecureBoot()
	ResetEnableVtpm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference
type jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableIntegrityMonitoring() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableIntegrityMonitoringInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableIntegrityMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableSecureBoot() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableSecureBootInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableSecureBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableVtpm() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) EnableVtpmInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableVtpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference_Override(b BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetEnableIntegrityMonitoring(val any) {
	if err := j.validateSetEnableIntegrityMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntegrityMonitoring",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetEnableSecureBoot(val any) {
	if err := j.validateSetEnableSecureBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecureBoot",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetEnableVtpm(val any) {
	if err := j.validateSetEnableVtpmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableVtpm",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ResetEnableIntegrityMonitoring() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableIntegrityMonitoring",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ResetEnableSecureBoot() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableSecureBoot",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ResetEnableVtpm() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableVtpm",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		b,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
