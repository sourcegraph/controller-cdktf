package backupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/backupdrrestoreworkload/internal"
)

type BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	AllocationAffinity() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference
	AllocationAffinityInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
	CanIpForward() interface{}
	SetCanIpForward(val interface{})
	CanIpForwardInput() interface{}
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
	ConfidentialInstanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disks() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList
	DisksInput() interface{}
	DisplayDevice() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference
	DisplayDeviceInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	// Experimental.
	Fqn() *string
	GuestAccelerators() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList
	GuestAcceleratorsInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InstanceEncryptionKey() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference
	InstanceEncryptionKeyInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey
	InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestoreProperties
	SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestoreProperties)
	KeyRevocationActionType() *string
	SetKeyRevocationActionType(val *string)
	KeyRevocationActionTypeInput() *string
	Labels() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList
	LabelsInput() interface{}
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference
	MetadataInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterfaces() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	NetworkPerformanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig
	Params() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference
	ParamsInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams
	PrivateIpv6GoogleAccess() *string
	SetPrivateIpv6GoogleAccess(val *string)
	PrivateIpv6GoogleAccessInput() *string
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Scheduling() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference
	SchedulingInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling
	ServiceAccounts() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList
	ServiceAccountsInput() interface{}
	ShieldedInstanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	Tags() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference
	TagsInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags
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
	PutAdvancedMachineFeatures(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures)
	PutAllocationAffinity(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity)
	PutConfidentialInstanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig)
	PutDisks(value interface{})
	PutDisplayDevice(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice)
	PutGuestAccelerators(value interface{})
	PutInstanceEncryptionKey(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey)
	PutLabels(value interface{})
	PutMetadata(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata)
	PutNetworkInterfaces(value interface{})
	PutNetworkPerformanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig)
	PutParams(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams)
	PutScheduling(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling)
	PutServiceAccounts(value interface{})
	PutShieldedInstanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig)
	PutTags(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags)
	ResetAdvancedMachineFeatures()
	ResetAllocationAffinity()
	ResetCanIpForward()
	ResetConfidentialInstanceConfig()
	ResetDeletionProtection()
	ResetDescription()
	ResetDisks()
	ResetDisplayDevice()
	ResetGuestAccelerators()
	ResetHostname()
	ResetInstanceEncryptionKey()
	ResetKeyRevocationActionType()
	ResetLabels()
	ResetMachineType()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetNetworkInterfaces()
	ResetNetworkPerformanceConfig()
	ResetParams()
	ResetPrivateIpv6GoogleAccess()
	ResetResourcePolicies()
	ResetScheduling()
	ResetServiceAccounts()
	ResetShieldedInstanceConfig()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
type jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AdvancedMachineFeatures() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AdvancedMachineFeaturesInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AllocationAffinity() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference
	_jsii_.Get(
		j,
		"allocationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AllocationAffinityInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
	_jsii_.Get(
		j,
		"allocationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ConfidentialInstanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ConfidentialInstanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Disks() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList
	_jsii_.Get(
		j,
		"disks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisplayDevice() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference
	_jsii_.Get(
		j,
		"displayDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisplayDeviceInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	_jsii_.Get(
		j,
		"displayDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GuestAccelerators() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList
	_jsii_.Get(
		j,
		"guestAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GuestAcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InstanceEncryptionKey() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"instanceEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InstanceEncryptionKeyInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey
	_jsii_.Get(
		j,
		"instanceEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InternalValue() *BackupDrRestoreWorkloadComputeInstanceRestoreProperties {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestoreProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) KeyRevocationActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) KeyRevocationActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Labels() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Metadata() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MetadataInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkInterfaces() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkPerformanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkPerformanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Params() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ParamsInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PrivateIpv6GoogleAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Scheduling() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) SchedulingInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ServiceAccounts() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList
	_jsii_.Get(
		j,
		"serviceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ServiceAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ShieldedInstanceConfig() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ShieldedInstanceConfigInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Tags() BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference {
	var returns BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TagsInput() *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags {
	var returns *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference_Override(b BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.backupDrRestoreWorkload.BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetInternalValue(val *BackupDrRestoreWorkloadComputeInstanceRestoreProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetKeyRevocationActionType(val *string) {
	if err := j.validateSetKeyRevocationActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRevocationActionType",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetPrivateIpv6GoogleAccess(val *string) {
	if err := j.validateSetPrivateIpv6GoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpv6GoogleAccess",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutAdvancedMachineFeatures(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures) {
	if err := b.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutAllocationAffinity(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity) {
	if err := b.validatePutAllocationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAllocationAffinity",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutConfidentialInstanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig) {
	if err := b.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutDisks(value interface{}) {
	if err := b.validatePutDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDisks",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutDisplayDevice(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice) {
	if err := b.validatePutDisplayDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDisplayDevice",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutGuestAccelerators(value interface{}) {
	if err := b.validatePutGuestAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGuestAccelerators",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutInstanceEncryptionKey(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey) {
	if err := b.validatePutInstanceEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInstanceEncryptionKey",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutLabels(value interface{}) {
	if err := b.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLabels",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutMetadata(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata) {
	if err := b.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMetadata",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := b.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutNetworkPerformanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig) {
	if err := b.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutParams(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams) {
	if err := b.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putParams",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutScheduling(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling) {
	if err := b.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putScheduling",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutServiceAccounts(value interface{}) {
	if err := b.validatePutServiceAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putServiceAccounts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutShieldedInstanceConfig(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig) {
	if err := b.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutTags(value *BackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags) {
	if err := b.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTags",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		b,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetAllocationAffinity() {
	_jsii_.InvokeVoid(
		b,
		"resetAllocationAffinity",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		b,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		b,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDisks() {
	_jsii_.InvokeVoid(
		b,
		"resetDisks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDisplayDevice() {
	_jsii_.InvokeVoid(
		b,
		"resetDisplayDevice",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetGuestAccelerators() {
	_jsii_.InvokeVoid(
		b,
		"resetGuestAccelerators",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		b,
		"resetHostname",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetInstanceEncryptionKey() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceEncryptionKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetKeyRevocationActionType() {
	_jsii_.InvokeVoid(
		b,
		"resetKeyRevocationActionType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		b,
		"resetLabels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		b,
		"resetMachineType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadata",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		b,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetParams() {
	_jsii_.InvokeVoid(
		b,
		"resetParams",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetPrivateIpv6GoogleAccess() {
	_jsii_.InvokeVoid(
		b,
		"resetPrivateIpv6GoogleAccess",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		b,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		b,
		"resetScheduling",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetServiceAccounts() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceAccounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

