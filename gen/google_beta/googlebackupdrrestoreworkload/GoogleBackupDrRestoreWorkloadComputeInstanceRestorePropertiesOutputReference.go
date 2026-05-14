package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	AllocationAffinity() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference
	AllocationAffinityInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
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
	ConfidentialInstanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig
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
	Disks() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList
	DisksInput() interface{}
	DisplayDevice() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference
	DisplayDeviceInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	// Experimental.
	Fqn() *string
	GuestAccelerators() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList
	GuestAcceleratorsInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	InstanceEncryptionKey() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference
	InstanceEncryptionKeyInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey
	InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties
	SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties)
	KeyRevocationActionType() *string
	SetKeyRevocationActionType(val *string)
	KeyRevocationActionTypeInput() *string
	Labels() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList
	LabelsInput() interface{}
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference
	MetadataInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterfaces() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList
	NetworkInterfacesInput() interface{}
	NetworkPerformanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig
	Params() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference
	ParamsInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams
	PrivateIpv6GoogleAccess() *string
	SetPrivateIpv6GoogleAccess(val *string)
	PrivateIpv6GoogleAccessInput() *string
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Scheduling() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference
	SchedulingInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling
	ServiceAccounts() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList
	ServiceAccountsInput() interface{}
	ShieldedInstanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	Tags() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference
	TagsInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags
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
	PutAdvancedMachineFeatures(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures)
	PutAllocationAffinity(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity)
	PutConfidentialInstanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig)
	PutDisks(value interface{})
	PutDisplayDevice(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice)
	PutGuestAccelerators(value interface{})
	PutInstanceEncryptionKey(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey)
	PutLabels(value interface{})
	PutMetadata(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata)
	PutNetworkInterfaces(value interface{})
	PutNetworkPerformanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig)
	PutParams(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams)
	PutScheduling(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling)
	PutServiceAccounts(value interface{})
	PutShieldedInstanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig)
	PutTags(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags)
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

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AdvancedMachineFeatures() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AdvancedMachineFeaturesInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AllocationAffinity() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference
	_jsii_.Get(
		j,
		"allocationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) AllocationAffinityInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
	_jsii_.Get(
		j,
		"allocationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ConfidentialInstanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ConfidentialInstanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Disks() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisksList
	_jsii_.Get(
		j,
		"disks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisplayDevice() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDeviceOutputReference
	_jsii_.Get(
		j,
		"displayDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) DisplayDeviceInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice
	_jsii_.Get(
		j,
		"displayDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GuestAccelerators() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesGuestAcceleratorsList
	_jsii_.Get(
		j,
		"guestAccelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GuestAcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InstanceEncryptionKey() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"instanceEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InstanceEncryptionKeyInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey
	_jsii_.Get(
		j,
		"instanceEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) KeyRevocationActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) KeyRevocationActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Labels() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Metadata() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MetadataInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkInterfaces() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkInterfacesList
	_jsii_.Get(
		j,
		"networkInterfaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkInterfacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkPerformanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) NetworkPerformanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Params() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ParamsInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PrivateIpv6GoogleAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Scheduling() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) SchedulingInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ServiceAccounts() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesServiceAccountsList
	_jsii_.Get(
		j,
		"serviceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ServiceAccountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ShieldedInstanceConfig() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ShieldedInstanceConfigInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Tags() GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference {
	var returns GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTagsOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TagsInput() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference_Override(g GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestoreProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetKeyRevocationActionType(val *string) {
	if err := j.validateSetKeyRevocationActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRevocationActionType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetPrivateIpv6GoogleAccess(val *string) {
	if err := j.validateSetPrivateIpv6GoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpv6GoogleAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutAdvancedMachineFeatures(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutAllocationAffinity(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity) {
	if err := g.validatePutAllocationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllocationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutConfidentialInstanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesConfidentialInstanceConfig) {
	if err := g.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutDisks(value interface{}) {
	if err := g.validatePutDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDisks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutDisplayDevice(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesDisplayDevice) {
	if err := g.validatePutDisplayDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDisplayDevice",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutGuestAccelerators(value interface{}) {
	if err := g.validatePutGuestAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerators",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutInstanceEncryptionKey(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesInstanceEncryptionKey) {
	if err := g.validatePutInstanceEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutLabels(value interface{}) {
	if err := g.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLabels",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutMetadata(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesMetadata) {
	if err := g.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetadata",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutNetworkInterfaces(value interface{}) {
	if err := g.validatePutNetworkInterfacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterfaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutNetworkPerformanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesNetworkPerformanceConfig) {
	if err := g.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutParams(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutScheduling(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutServiceAccounts(value interface{}) {
	if err := g.validatePutServiceAccountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccounts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutShieldedInstanceConfig(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) PutTags(value *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesTags) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetAllocationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetAllocationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		g,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDisks() {
	_jsii_.InvokeVoid(
		g,
		"resetDisks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetDisplayDevice() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayDevice",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetGuestAccelerators() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		g,
		"resetHostname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetInstanceEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetKeyRevocationActionType() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyRevocationActionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetNetworkInterfaces() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterfaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetPrivateIpv6GoogleAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIpv6GoogleAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetServiceAccounts() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccounts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

