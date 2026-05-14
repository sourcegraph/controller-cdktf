package workstationsworkstationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/workstationsworkstationconfig/internal"
)

type WorkstationsWorkstationConfigHostGceInstanceOutputReference interface {
	cdktf.ComplexObject
	Accelerators() WorkstationsWorkstationConfigHostGceInstanceAcceleratorsList
	AcceleratorsInput() interface{}
	BoostConfigs() WorkstationsWorkstationConfigHostGceInstanceBoostConfigsList
	BoostConfigsInput() interface{}
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
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
	ConfidentialInstanceConfig() WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisablePublicIpAddresses() interface{}
	SetDisablePublicIpAddresses(val interface{})
	DisablePublicIpAddressesInput() interface{}
	DisableSsh() interface{}
	SetDisableSsh(val interface{})
	DisableSshInput() interface{}
	EnableNestedVirtualization() interface{}
	SetEnableNestedVirtualization(val interface{})
	EnableNestedVirtualizationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *WorkstationsWorkstationConfigHostGceInstance
	SetInternalValue(val *WorkstationsWorkstationConfigHostGceInstance)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	PoolSize() *float64
	SetPoolSize(val *float64)
	PoolSizeInput() *float64
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceAccountScopes() *[]*string
	SetServiceAccountScopes(val *[]*string)
	ServiceAccountScopesInput() *[]*string
	ShieldedInstanceConfig() WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmTags() *map[string]*string
	SetVmTags(val *map[string]*string)
	VmTagsInput() *map[string]*string
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
	PutAccelerators(value interface{})
	PutBoostConfigs(value interface{})
	PutConfidentialInstanceConfig(value *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig)
	PutShieldedInstanceConfig(value *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig)
	ResetAccelerators()
	ResetBoostConfigs()
	ResetBootDiskSizeGb()
	ResetConfidentialInstanceConfig()
	ResetDisablePublicIpAddresses()
	ResetDisableSsh()
	ResetEnableNestedVirtualization()
	ResetMachineType()
	ResetPoolSize()
	ResetServiceAccount()
	ResetServiceAccountScopes()
	ResetShieldedInstanceConfig()
	ResetTags()
	ResetVmTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkstationsWorkstationConfigHostGceInstanceOutputReference
type jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) Accelerators() WorkstationsWorkstationConfigHostGceInstanceAcceleratorsList {
	var returns WorkstationsWorkstationConfigHostGceInstanceAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) BoostConfigs() WorkstationsWorkstationConfigHostGceInstanceBoostConfigsList {
	var returns WorkstationsWorkstationConfigHostGceInstanceBoostConfigsList
	_jsii_.Get(
		j,
		"boostConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) BoostConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boostConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ConfidentialInstanceConfig() WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfigOutputReference {
	var returns WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ConfidentialInstanceConfigInput() *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig {
	var returns *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) DisablePublicIpAddresses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) DisablePublicIpAddressesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePublicIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) DisableSsh() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSsh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) DisableSshInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) EnableNestedVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) EnableNestedVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) InternalValue() *WorkstationsWorkstationConfigHostGceInstance {
	var returns *WorkstationsWorkstationConfigHostGceInstance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"poolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ServiceAccountScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ServiceAccountScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ShieldedInstanceConfig() WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference {
	var returns WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ShieldedInstanceConfigInput() *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig {
	var returns *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) VmTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vmTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) VmTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vmTagsInput",
		&returns,
	)
	return returns
}


func NewWorkstationsWorkstationConfigHostGceInstanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkstationsWorkstationConfigHostGceInstanceOutputReference {
	_init_.Initialize()

	if err := validateNewWorkstationsWorkstationConfigHostGceInstanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigHostGceInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkstationsWorkstationConfigHostGceInstanceOutputReference_Override(w WorkstationsWorkstationConfigHostGceInstanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.workstationsWorkstationConfig.WorkstationsWorkstationConfigHostGceInstanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetDisablePublicIpAddresses(val interface{}) {
	if err := j.validateSetDisablePublicIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePublicIpAddresses",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetDisableSsh(val interface{}) {
	if err := j.validateSetDisableSshParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSsh",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetEnableNestedVirtualization(val interface{}) {
	if err := j.validateSetEnableNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetInternalValue(val *WorkstationsWorkstationConfigHostGceInstance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetPoolSize(val *float64) {
	if err := j.validateSetPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolSize",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetServiceAccountScopes(val *[]*string) {
	if err := j.validateSetServiceAccountScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountScopes",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference)SetVmTags(val *map[string]*string) {
	if err := j.validateSetVmTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmTags",
		val,
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PutAccelerators(value interface{}) {
	if err := w.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PutBoostConfigs(value interface{}) {
	if err := w.validatePutBoostConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBoostConfigs",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PutConfidentialInstanceConfig(value *WorkstationsWorkstationConfigHostGceInstanceConfidentialInstanceConfig) {
	if err := w.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) PutShieldedInstanceConfig(value *WorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig) {
	if err := w.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		w,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetBoostConfigs() {
	_jsii_.InvokeVoid(
		w,
		"resetBoostConfigs",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		w,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetDisablePublicIpAddresses() {
	_jsii_.InvokeVoid(
		w,
		"resetDisablePublicIpAddresses",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetDisableSsh() {
	_jsii_.InvokeVoid(
		w,
		"resetDisableSsh",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetEnableNestedVirtualization() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableNestedVirtualization",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		w,
		"resetMachineType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetPoolSize() {
	_jsii_.InvokeVoid(
		w,
		"resetPoolSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetServiceAccountScopes() {
	_jsii_.InvokeVoid(
		w,
		"resetServiceAccountScopes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ResetVmTags() {
	_jsii_.InvokeVoid(
		w,
		"resetVmTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkstationsWorkstationConfigHostGceInstanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

