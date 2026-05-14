package googlecomputeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregioninstancegroupmanager/internal"
)

type GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference interface {
	cdktf.ComplexObject
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	AutoDelete() interface{}
	SetAutoDelete(val interface{})
	AutoDeleteInput() interface{}
	Boot() interface{}
	SetBoot(val interface{})
	BootInput() interface{}
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	DiskEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey
	DiskName() *string
	SetDiskName(val *string)
	DiskNameInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	GuestOsFeatures() *[]*string
	SetGuestOsFeatures(val *[]*string)
	GuestOsFeaturesInput() *[]*string
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList
	LabelsInput() interface{}
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
	ResourceManagerTags() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList
	ResourceManagerTagsInput() interface{}
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Source() *string
	SetSource(val *string)
	SourceImage() *string
	SetSourceImage(val *string)
	SourceImageEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference
	SourceImageEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey
	SourceImageInput() *string
	SourceInput() *string
	SourceSnapshot() *string
	SetSourceSnapshot(val *string)
	SourceSnapshotEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference
	SourceSnapshotEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	SourceSnapshotInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutDiskEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey)
	PutLabels(value interface{})
	PutResourceManagerTags(value interface{})
	PutSourceImageEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey)
	PutSourceSnapshotEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey)
	ResetArchitecture()
	ResetAutoDelete()
	ResetBoot()
	ResetDeviceName()
	ResetDiskEncryptionKey()
	ResetDiskName()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetGuestOsFeatures()
	ResetInterface()
	ResetLabels()
	ResetMode()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	ResetResourceManagerTags()
	ResetResourcePolicies()
	ResetSource()
	ResetSourceImage()
	ResetSourceImageEncryptionKey()
	ResetSourceSnapshot()
	ResetSourceSnapshotEncryptionKey()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference
type jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) AutoDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) AutoDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Boot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) BootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey {
	var returns *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GuestOsFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GuestOsFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Labels() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourceManagerTags() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTagsList
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourceManagerTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceImageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey {
	var returns *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey
	_jsii_.Get(
		j,
		"sourceImageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotEncryptionKey() GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference {
	var returns GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotEncryptionKeyInput() *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey {
	var returns *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) SourceSnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference_Override(g GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionInstanceGroupManager.GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetAutoDelete(val interface{}) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetBoot(val interface{}) {
	if err := j.validateSetBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boot",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskName(val *string) {
	if err := j.validateSetDiskNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetGuestOsFeatures(val *[]*string) {
	if err := j.validateSetGuestOsFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestOsFeatures",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSourceImage(val *string) {
	if err := j.validateSetSourceImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImage",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetSourceSnapshot(val *string) {
	if err := j.validateSetSourceSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSnapshot",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutDiskEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksDiskEncryptionKey) {
	if err := g.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutLabels(value interface{}) {
	if err := g.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLabels",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutResourceManagerTags(value interface{}) {
	if err := g.validatePutResourceManagerTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceManagerTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutSourceImageEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceImageEncryptionKey) {
	if err := g.validatePutSourceImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceImageEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) PutSourceSnapshotEncryptionKey(value *GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksSourceSnapshotEncryptionKey) {
	if err := g.validatePutSourceSnapshotEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceSnapshotEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		g,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetBoot() {
	_jsii_.InvokeVoid(
		g,
		"resetBoot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		g,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskName() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetGuestOsFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestOsFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		g,
		"resetSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceImage() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceImageEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceImageEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceSnapshot() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceSnapshot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetSourceSnapshotEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceSnapshotEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

