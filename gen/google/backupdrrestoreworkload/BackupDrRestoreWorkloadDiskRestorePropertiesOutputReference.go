package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference interface {
	cdktn.ComplexObject
	AccessMode() *string
	SetAccessMode(val *string)
	AccessModeInput() *string
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskEncryptionKey() BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey
	EnableConfidentialCompute() interface{}
	SetEnableConfidentialCompute(val interface{})
	EnableConfidentialComputeInput() interface{}
	// Experimental.
	Fqn() *string
	GuestOsFeature() BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList
	GuestOsFeatureInput() interface{}
	InternalValue() *BackupDrRestoreWorkloadDiskRestoreProperties
	SetInternalValue(val *BackupDrRestoreWorkloadDiskRestoreProperties)
	Labels() BackupDrRestoreWorkloadDiskRestorePropertiesLabelsList
	LabelsInput() interface{}
	Licenses() *[]*string
	SetLicenses(val *[]*string)
	LicensesInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PhysicalBlockSizeBytes() *float64
	SetPhysicalBlockSizeBytes(val *float64)
	PhysicalBlockSizeBytesInput() *float64
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
	ResourceManagerTags() BackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList
	ResourceManagerTagsInput() interface{}
	ResourcePolicy() *[]*string
	SetResourcePolicy(val *[]*string)
	ResourcePolicyInput() *[]*string
	SizeGb() *float64
	SetSizeGb(val *float64)
	SizeGbInput() *float64
	StoragePool() *string
	SetStoragePool(val *string)
	StoragePoolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDiskEncryptionKey(value *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey)
	PutGuestOsFeature(value interface{})
	PutLabels(value interface{})
	PutResourceManagerTags(value interface{})
	ResetAccessMode()
	ResetArchitecture()
	ResetDescription()
	ResetDiskEncryptionKey()
	ResetEnableConfidentialCompute()
	ResetGuestOsFeature()
	ResetLabels()
	ResetLicenses()
	ResetPhysicalBlockSizeBytes()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	ResetResourceManagerTags()
	ResetResourcePolicy()
	ResetStoragePool()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
type jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DiskEncryptionKey() BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference {
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DiskEncryptionKeyInput() *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey {
	var returns *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) EnableConfidentialCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) EnableConfidentialComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GuestOsFeature() BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList {
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList
	_jsii_.Get(
		j,
		"guestOsFeature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GuestOsFeatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestOsFeatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InternalValue() *BackupDrRestoreWorkloadDiskRestoreProperties {
	var returns *BackupDrRestoreWorkloadDiskRestoreProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Labels() BackupDrRestoreWorkloadDiskRestorePropertiesLabelsList {
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) LicensesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licensesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PhysicalBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PhysicalBlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourceManagerTags() BackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList {
	var returns BackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourceManagerTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourcePolicy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourcePolicyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadDiskRestorePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference_Override(b BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetEnableConfidentialCompute(val interface{}) {
	if err := j.validateSetEnableConfidentialComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidentialCompute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetInternalValue(val *BackupDrRestoreWorkloadDiskRestoreProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetLicenses(val *[]*string) {
	if err := j.validateSetLicensesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenses",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetPhysicalBlockSizeBytes(val *float64) {
	if err := j.validateSetPhysicalBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalBlockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetResourcePolicy(val *[]*string) {
	if err := j.validateSetResourcePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicy",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetSizeGb(val *float64) {
	if err := j.validateSetSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutDiskEncryptionKey(value *BackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey) {
	if err := b.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutGuestOsFeature(value interface{}) {
	if err := b.validatePutGuestOsFeatureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGuestOsFeature",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutLabels(value interface{}) {
	if err := b.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLabels",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutResourceManagerTags(value interface{}) {
	if err := b.validatePutResourceManagerTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResourceManagerTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		b,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		b,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetEnableConfidentialCompute() {
	_jsii_.InvokeVoid(
		b,
		"resetEnableConfidentialCompute",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetGuestOsFeature() {
	_jsii_.InvokeVoid(
		b,
		"resetGuestOsFeature",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetLicenses() {
	_jsii_.InvokeVoid(
		b,
		"resetLicenses",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetPhysicalBlockSizeBytes() {
	_jsii_.InvokeVoid(
		b,
		"resetPhysicalBlockSizeBytes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		b,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		b,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		b,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetResourcePolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetResourcePolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetStoragePool() {
	_jsii_.InvokeVoid(
		b,
		"resetStoragePool",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

