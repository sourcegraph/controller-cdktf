package googlecomputeinstancefrommachineimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancefrommachineimage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image google_compute_instance_from_machine_image}.
type GoogleComputeInstanceFromMachineImage interface {
	cdktf.TerraformResource
	AdvancedMachineFeatures() GoogleComputeInstanceFromMachineImageAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleComputeInstanceFromMachineImageAdvancedMachineFeatures
	AllowStoppingForUpdate() interface{}
	SetAllowStoppingForUpdate(val interface{})
	AllowStoppingForUpdateInput() interface{}
	AttachedDisk() GoogleComputeInstanceFromMachineImageAttachedDiskList
	BootDisk() GoogleComputeInstanceFromMachineImageBootDiskList
	CanIpForward() interface{}
	SetCanIpForward(val interface{})
	CanIpForwardInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidentialInstanceConfig() GoogleComputeInstanceFromMachineImageConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuPlatform() *string
	CreationTimestamp() *string
	CurrentStatus() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DesiredStatus() *string
	SetDesiredStatus(val *string)
	DesiredStatusInput() *string
	EffectiveLabels() cdktf.StringMap
	EnableDisplay() interface{}
	SetEnableDisplay(val interface{})
	EnableDisplayInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestAccelerator() GoogleComputeInstanceFromMachineImageGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceEncryptionKey() GoogleComputeInstanceFromMachineImageInstanceEncryptionKeyOutputReference
	InstanceEncryptionKeyInput() *GoogleComputeInstanceFromMachineImageInstanceEncryptionKey
	InstanceId() *string
	KeyRevocationActionType() *string
	SetKeyRevocationActionType(val *string)
	KeyRevocationActionTypeInput() *string
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataFingerprint() *string
	MetadataInput() *map[string]*string
	MetadataStartupScript() *string
	SetMetadataStartupScript(val *string)
	MetadataStartupScriptInput() *string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() GoogleComputeInstanceFromMachineImageNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NetworkPerformanceConfig() GoogleComputeInstanceFromMachineImageNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *GoogleComputeInstanceFromMachineImageNetworkPerformanceConfig
	// The tree node.
	Node() constructs.Node
	Params() GoogleComputeInstanceFromMachineImageParamsOutputReference
	ParamsInput() *GoogleComputeInstanceFromMachineImageParams
	PartnerMetadata() *map[string]*string
	SetPartnerMetadata(val *map[string]*string)
	PartnerMetadataInput() *map[string]*string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReservationAffinity() GoogleComputeInstanceFromMachineImageReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleComputeInstanceFromMachineImageReservationAffinity
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Scheduling() GoogleComputeInstanceFromMachineImageSchedulingOutputReference
	SchedulingInput() *GoogleComputeInstanceFromMachineImageScheduling
	ScratchDisk() GoogleComputeInstanceFromMachineImageScratchDiskList
	SelfLink() *string
	ServiceAccount() GoogleComputeInstanceFromMachineImageServiceAccountOutputReference
	ServiceAccountInput() *GoogleComputeInstanceFromMachineImageServiceAccount
	ShieldedInstanceConfig() GoogleComputeInstanceFromMachineImageShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleComputeInstanceFromMachineImageShieldedInstanceConfig
	SourceMachineImage() *string
	SetSourceMachineImage(val *string)
	SourceMachineImageEncryptionKey() GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKeyOutputReference
	SourceMachineImageEncryptionKeyInput() *GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey
	SourceMachineImageInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsFingerprint() *string
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeInstanceFromMachineImageTimeoutsOutputReference
	TimeoutsInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedMachineFeatures(value *GoogleComputeInstanceFromMachineImageAdvancedMachineFeatures)
	PutConfidentialInstanceConfig(value *GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig)
	PutGuestAccelerator(value interface{})
	PutInstanceEncryptionKey(value *GoogleComputeInstanceFromMachineImageInstanceEncryptionKey)
	PutNetworkInterface(value interface{})
	PutNetworkPerformanceConfig(value *GoogleComputeInstanceFromMachineImageNetworkPerformanceConfig)
	PutParams(value *GoogleComputeInstanceFromMachineImageParams)
	PutReservationAffinity(value *GoogleComputeInstanceFromMachineImageReservationAffinity)
	PutScheduling(value *GoogleComputeInstanceFromMachineImageScheduling)
	PutServiceAccount(value *GoogleComputeInstanceFromMachineImageServiceAccount)
	PutShieldedInstanceConfig(value *GoogleComputeInstanceFromMachineImageShieldedInstanceConfig)
	PutSourceMachineImageEncryptionKey(value *GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey)
	PutTimeouts(value *GoogleComputeInstanceFromMachineImageTimeouts)
	ResetAdvancedMachineFeatures()
	ResetAllowStoppingForUpdate()
	ResetCanIpForward()
	ResetConfidentialInstanceConfig()
	ResetDeletionProtection()
	ResetDescription()
	ResetDesiredStatus()
	ResetEnableDisplay()
	ResetGuestAccelerator()
	ResetHostname()
	ResetId()
	ResetInstanceEncryptionKey()
	ResetKeyRevocationActionType()
	ResetLabels()
	ResetMachineType()
	ResetMetadata()
	ResetMetadataStartupScript()
	ResetMinCpuPlatform()
	ResetNetworkInterface()
	ResetNetworkPerformanceConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParams()
	ResetPartnerMetadata()
	ResetProject()
	ResetReservationAffinity()
	ResetResourcePolicies()
	ResetScheduling()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetSourceMachineImageEncryptionKey()
	ResetTags()
	ResetTimeouts()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeInstanceFromMachineImage
type jsiiProxy_GoogleComputeInstanceFromMachineImage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) AdvancedMachineFeatures() GoogleComputeInstanceFromMachineImageAdvancedMachineFeaturesOutputReference {
	var returns GoogleComputeInstanceFromMachineImageAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) AdvancedMachineFeaturesInput() *GoogleComputeInstanceFromMachineImageAdvancedMachineFeatures {
	var returns *GoogleComputeInstanceFromMachineImageAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) AllowStoppingForUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) AllowStoppingForUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) AttachedDisk() GoogleComputeInstanceFromMachineImageAttachedDiskList {
	var returns GoogleComputeInstanceFromMachineImageAttachedDiskList
	_jsii_.Get(
		j,
		"attachedDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) BootDisk() GoogleComputeInstanceFromMachineImageBootDiskList {
	var returns GoogleComputeInstanceFromMachineImageBootDiskList
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ConfidentialInstanceConfig() GoogleComputeInstanceFromMachineImageConfidentialInstanceConfigOutputReference {
	var returns GoogleComputeInstanceFromMachineImageConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ConfidentialInstanceConfigInput() *GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig {
	var returns *GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) CurrentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DesiredStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) DesiredStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) EnableDisplay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) EnableDisplayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) GuestAccelerator() GoogleComputeInstanceFromMachineImageGuestAcceleratorList {
	var returns GoogleComputeInstanceFromMachineImageGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) InstanceEncryptionKey() GoogleComputeInstanceFromMachineImageInstanceEncryptionKeyOutputReference {
	var returns GoogleComputeInstanceFromMachineImageInstanceEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"instanceEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) InstanceEncryptionKeyInput() *GoogleComputeInstanceFromMachineImageInstanceEncryptionKey {
	var returns *GoogleComputeInstanceFromMachineImageInstanceEncryptionKey
	_jsii_.Get(
		j,
		"instanceEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) KeyRevocationActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) KeyRevocationActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MetadataFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MetadataStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MetadataStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) NetworkInterface() GoogleComputeInstanceFromMachineImageNetworkInterfaceList {
	var returns GoogleComputeInstanceFromMachineImageNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) NetworkPerformanceConfig() GoogleComputeInstanceFromMachineImageNetworkPerformanceConfigOutputReference {
	var returns GoogleComputeInstanceFromMachineImageNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) NetworkPerformanceConfigInput() *GoogleComputeInstanceFromMachineImageNetworkPerformanceConfig {
	var returns *GoogleComputeInstanceFromMachineImageNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Params() GoogleComputeInstanceFromMachineImageParamsOutputReference {
	var returns GoogleComputeInstanceFromMachineImageParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ParamsInput() *GoogleComputeInstanceFromMachineImageParams {
	var returns *GoogleComputeInstanceFromMachineImageParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) PartnerMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"partnerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) PartnerMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"partnerMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ReservationAffinity() GoogleComputeInstanceFromMachineImageReservationAffinityOutputReference {
	var returns GoogleComputeInstanceFromMachineImageReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ReservationAffinityInput() *GoogleComputeInstanceFromMachineImageReservationAffinity {
	var returns *GoogleComputeInstanceFromMachineImageReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Scheduling() GoogleComputeInstanceFromMachineImageSchedulingOutputReference {
	var returns GoogleComputeInstanceFromMachineImageSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SchedulingInput() *GoogleComputeInstanceFromMachineImageScheduling {
	var returns *GoogleComputeInstanceFromMachineImageScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ScratchDisk() GoogleComputeInstanceFromMachineImageScratchDiskList {
	var returns GoogleComputeInstanceFromMachineImageScratchDiskList
	_jsii_.Get(
		j,
		"scratchDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ServiceAccount() GoogleComputeInstanceFromMachineImageServiceAccountOutputReference {
	var returns GoogleComputeInstanceFromMachineImageServiceAccountOutputReference
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ServiceAccountInput() *GoogleComputeInstanceFromMachineImageServiceAccount {
	var returns *GoogleComputeInstanceFromMachineImageServiceAccount
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ShieldedInstanceConfig() GoogleComputeInstanceFromMachineImageShieldedInstanceConfigOutputReference {
	var returns GoogleComputeInstanceFromMachineImageShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ShieldedInstanceConfigInput() *GoogleComputeInstanceFromMachineImageShieldedInstanceConfig {
	var returns *GoogleComputeInstanceFromMachineImageShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SourceMachineImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMachineImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SourceMachineImageEncryptionKey() GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKeyOutputReference {
	var returns GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceMachineImageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SourceMachineImageEncryptionKeyInput() *GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey {
	var returns *GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey
	_jsii_.Get(
		j,
		"sourceMachineImageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) SourceMachineImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceMachineImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TagsFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Timeouts() GoogleComputeInstanceFromMachineImageTimeoutsOutputReference {
	var returns GoogleComputeInstanceFromMachineImageTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image google_compute_instance_from_machine_image} Resource.
func NewGoogleComputeInstanceFromMachineImage(scope constructs.Construct, id *string, config *GoogleComputeInstanceFromMachineImageConfig) GoogleComputeInstanceFromMachineImage {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromMachineImageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromMachineImage{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_instance_from_machine_image google_compute_instance_from_machine_image} Resource.
func NewGoogleComputeInstanceFromMachineImage_Override(g GoogleComputeInstanceFromMachineImage, scope constructs.Construct, id *string, config *GoogleComputeInstanceFromMachineImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetAllowStoppingForUpdate(val interface{}) {
	if err := j.validateSetAllowStoppingForUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoppingForUpdate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetDesiredStatus(val *string) {
	if err := j.validateSetDesiredStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredStatus",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetEnableDisplay(val interface{}) {
	if err := j.validateSetEnableDisplayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDisplay",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetKeyRevocationActionType(val *string) {
	if err := j.validateSetKeyRevocationActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRevocationActionType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetMetadataStartupScript(val *string) {
	if err := j.validateSetMetadataStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataStartupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetPartnerMetadata(val *map[string]*string) {
	if err := j.validateSetPartnerMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerMetadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetSourceMachineImage(val *string) {
	if err := j.validateSetSourceMachineImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceMachineImage",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromMachineImage)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeInstanceFromMachineImage resource upon running "cdktf plan <stack-name>".
func GoogleComputeInstanceFromMachineImage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromMachineImage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleComputeInstanceFromMachineImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromMachineImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceFromMachineImage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromMachineImage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceFromMachineImage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromMachineImage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInstanceFromMachineImage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeInstanceFromMachineImage.GoogleComputeInstanceFromMachineImage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutAdvancedMachineFeatures(value *GoogleComputeInstanceFromMachineImageAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutConfidentialInstanceConfig(value *GoogleComputeInstanceFromMachineImageConfidentialInstanceConfig) {
	if err := g.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutInstanceEncryptionKey(value *GoogleComputeInstanceFromMachineImageInstanceEncryptionKey) {
	if err := g.validatePutInstanceEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstanceEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutNetworkInterface(value interface{}) {
	if err := g.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutNetworkPerformanceConfig(value *GoogleComputeInstanceFromMachineImageNetworkPerformanceConfig) {
	if err := g.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutParams(value *GoogleComputeInstanceFromMachineImageParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutReservationAffinity(value *GoogleComputeInstanceFromMachineImageReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutScheduling(value *GoogleComputeInstanceFromMachineImageScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutServiceAccount(value *GoogleComputeInstanceFromMachineImageServiceAccount) {
	if err := g.validatePutServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutShieldedInstanceConfig(value *GoogleComputeInstanceFromMachineImageShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutSourceMachineImageEncryptionKey(value *GoogleComputeInstanceFromMachineImageSourceMachineImageEncryptionKey) {
	if err := g.validatePutSourceMachineImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceMachineImageEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) PutTimeouts(value *GoogleComputeInstanceFromMachineImageTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetAllowStoppingForUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowStoppingForUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		g,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetDesiredStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetEnableDisplay() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDisplay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetHostname() {
	_jsii_.InvokeVoid(
		g,
		"resetHostname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetInstanceEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetKeyRevocationActionType() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyRevocationActionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetMetadataStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetPartnerMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetPartnerMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetSourceMachineImageEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceMachineImageEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromMachineImage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

