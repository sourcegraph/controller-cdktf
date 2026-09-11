package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference interface {
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
	DiskEncryptionKey() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey
	EnableConfidentialCompute() interface{}
	SetEnableConfidentialCompute(val interface{})
	EnableConfidentialComputeInput() interface{}
	// Experimental.
	Fqn() *string
	GuestOsFeature() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList
	GuestOsFeatureInput() interface{}
	InternalValue() *GoogleBackupDrRestoreWorkloadDiskRestoreProperties
	SetInternalValue(val *GoogleBackupDrRestoreWorkloadDiskRestoreProperties)
	Labels() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesLabelsList
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
	ResourceManagerTags() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList
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
	PutDiskEncryptionKey(value *GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey)
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

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) AccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) AccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DiskEncryptionKey() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference {
	var returns GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) DiskEncryptionKeyInput() *GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey {
	var returns *GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) EnableConfidentialCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) EnableConfidentialComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GuestOsFeature() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList {
	var returns GoogleBackupDrRestoreWorkloadDiskRestorePropertiesGuestOsFeatureList
	_jsii_.Get(
		j,
		"guestOsFeature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GuestOsFeatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestOsFeatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InternalValue() *GoogleBackupDrRestoreWorkloadDiskRestoreProperties {
	var returns *GoogleBackupDrRestoreWorkloadDiskRestoreProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Labels() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesLabelsList {
	var returns GoogleBackupDrRestoreWorkloadDiskRestorePropertiesLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) LicensesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licensesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PhysicalBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PhysicalBlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourceManagerTags() GoogleBackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList {
	var returns GoogleBackupDrRestoreWorkloadDiskRestorePropertiesResourceManagerTagsList
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourceManagerTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourcePolicy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResourcePolicyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) SizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference_Override(g GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetAccessMode(val *string) {
	if err := j.validateSetAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetEnableConfidentialCompute(val interface{}) {
	if err := j.validateSetEnableConfidentialComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidentialCompute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetInternalValue(val *GoogleBackupDrRestoreWorkloadDiskRestoreProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetLicenses(val *[]*string) {
	if err := j.validateSetLicensesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenses",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetPhysicalBlockSizeBytes(val *float64) {
	if err := j.validateSetPhysicalBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalBlockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetResourcePolicy(val *[]*string) {
	if err := j.validateSetResourcePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetSizeGb(val *float64) {
	if err := j.validateSetSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutDiskEncryptionKey(value *GoogleBackupDrRestoreWorkloadDiskRestorePropertiesDiskEncryptionKey) {
	if err := g.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutGuestOsFeature(value interface{}) {
	if err := g.validatePutGuestOsFeatureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestOsFeature",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutLabels(value interface{}) {
	if err := g.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLabels",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) PutResourceManagerTags(value interface{}) {
	if err := g.validatePutResourceManagerTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceManagerTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetAccessMode() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		g,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetEnableConfidentialCompute() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableConfidentialCompute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetGuestOsFeature() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestOsFeature",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetLicenses() {
	_jsii_.InvokeVoid(
		g,
		"resetLicenses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetPhysicalBlockSizeBytes() {
	_jsii_.InvokeVoid(
		g,
		"resetPhysicalBlockSizeBytes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetResourcePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ResetStoragePool() {
	_jsii_.InvokeVoid(
		g,
		"resetStoragePool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadDiskRestorePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

