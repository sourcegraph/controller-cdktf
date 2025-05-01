package googlecomputeinstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancetemplate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template google_compute_instance_template}.
type GoogleComputeInstanceTemplate interface {
	cdktf.TerraformResource
	AdvancedMachineFeatures() GoogleComputeInstanceTemplateAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleComputeInstanceTemplateAdvancedMachineFeatures
	CanIpForward() interface{}
	SetCanIpForward(val interface{})
	CanIpForwardInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidentialInstanceConfig() GoogleComputeInstanceTemplateConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *GoogleComputeInstanceTemplateConfidentialInstanceConfig
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disk() GoogleComputeInstanceTemplateDiskList
	DiskInput() interface{}
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
	GuestAccelerator() GoogleComputeInstanceTemplateGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceDescription() *string
	SetInstanceDescription(val *string)
	InstanceDescriptionInput() *string
	KeyRevocationActionType() *string
	SetKeyRevocationActionType(val *string)
	KeyRevocationActionTypeInput() *string
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
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	NetworkInterface() GoogleComputeInstanceTemplateNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NetworkPerformanceConfig() GoogleComputeInstanceTemplateNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *GoogleComputeInstanceTemplateNetworkPerformanceConfig
	// The tree node.
	Node() constructs.Node
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReservationAffinity() GoogleComputeInstanceTemplateReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleComputeInstanceTemplateReservationAffinity
	ResourceManagerTags() *map[string]*string
	SetResourceManagerTags(val *map[string]*string)
	ResourceManagerTagsInput() *map[string]*string
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Scheduling() GoogleComputeInstanceTemplateSchedulingOutputReference
	SchedulingInput() *GoogleComputeInstanceTemplateScheduling
	SelfLink() *string
	SelfLinkUnique() *string
	ServiceAccount() GoogleComputeInstanceTemplateServiceAccountOutputReference
	ServiceAccountInput() *GoogleComputeInstanceTemplateServiceAccount
	ShieldedInstanceConfig() GoogleComputeInstanceTemplateShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleComputeInstanceTemplateShieldedInstanceConfig
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
	Timeouts() GoogleComputeInstanceTemplateTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAdvancedMachineFeatures(value *GoogleComputeInstanceTemplateAdvancedMachineFeatures)
	PutConfidentialInstanceConfig(value *GoogleComputeInstanceTemplateConfidentialInstanceConfig)
	PutDisk(value interface{})
	PutGuestAccelerator(value interface{})
	PutNetworkInterface(value interface{})
	PutNetworkPerformanceConfig(value *GoogleComputeInstanceTemplateNetworkPerformanceConfig)
	PutReservationAffinity(value *GoogleComputeInstanceTemplateReservationAffinity)
	PutScheduling(value *GoogleComputeInstanceTemplateScheduling)
	PutServiceAccount(value *GoogleComputeInstanceTemplateServiceAccount)
	PutShieldedInstanceConfig(value *GoogleComputeInstanceTemplateShieldedInstanceConfig)
	PutTimeouts(value *GoogleComputeInstanceTemplateTimeouts)
	ResetAdvancedMachineFeatures()
	ResetCanIpForward()
	ResetConfidentialInstanceConfig()
	ResetDescription()
	ResetEnableDisplay()
	ResetGuestAccelerator()
	ResetId()
	ResetInstanceDescription()
	ResetKeyRevocationActionType()
	ResetLabels()
	ResetMetadata()
	ResetMetadataStartupScript()
	ResetMinCpuPlatform()
	ResetName()
	ResetNamePrefix()
	ResetNetworkInterface()
	ResetNetworkPerformanceConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerMetadata()
	ResetProject()
	ResetRegion()
	ResetReservationAffinity()
	ResetResourceManagerTags()
	ResetResourcePolicies()
	ResetScheduling()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleComputeInstanceTemplate
type jsiiProxy_GoogleComputeInstanceTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) AdvancedMachineFeatures() GoogleComputeInstanceTemplateAdvancedMachineFeaturesOutputReference {
	var returns GoogleComputeInstanceTemplateAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) AdvancedMachineFeaturesInput() *GoogleComputeInstanceTemplateAdvancedMachineFeatures {
	var returns *GoogleComputeInstanceTemplateAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ConfidentialInstanceConfig() GoogleComputeInstanceTemplateConfidentialInstanceConfigOutputReference {
	var returns GoogleComputeInstanceTemplateConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ConfidentialInstanceConfigInput() *GoogleComputeInstanceTemplateConfidentialInstanceConfig {
	var returns *GoogleComputeInstanceTemplateConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Disk() GoogleComputeInstanceTemplateDiskList {
	var returns GoogleComputeInstanceTemplateDiskList
	_jsii_.Get(
		j,
		"disk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) DiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) EnableDisplay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) EnableDisplayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) GuestAccelerator() GoogleComputeInstanceTemplateGuestAcceleratorList {
	var returns GoogleComputeInstanceTemplateGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) InstanceDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) InstanceDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) KeyRevocationActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) KeyRevocationActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyRevocationActionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MetadataFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MetadataStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MetadataStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NetworkInterface() GoogleComputeInstanceTemplateNetworkInterfaceList {
	var returns GoogleComputeInstanceTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NetworkPerformanceConfig() GoogleComputeInstanceTemplateNetworkPerformanceConfigOutputReference {
	var returns GoogleComputeInstanceTemplateNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) NetworkPerformanceConfigInput() *GoogleComputeInstanceTemplateNetworkPerformanceConfig {
	var returns *GoogleComputeInstanceTemplateNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) PartnerMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"partnerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) PartnerMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"partnerMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ReservationAffinity() GoogleComputeInstanceTemplateReservationAffinityOutputReference {
	var returns GoogleComputeInstanceTemplateReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ReservationAffinityInput() *GoogleComputeInstanceTemplateReservationAffinity {
	var returns *GoogleComputeInstanceTemplateReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ResourceManagerTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ResourceManagerTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Scheduling() GoogleComputeInstanceTemplateSchedulingOutputReference {
	var returns GoogleComputeInstanceTemplateSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) SchedulingInput() *GoogleComputeInstanceTemplateScheduling {
	var returns *GoogleComputeInstanceTemplateScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) SelfLinkUnique() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLinkUnique",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ServiceAccount() GoogleComputeInstanceTemplateServiceAccountOutputReference {
	var returns GoogleComputeInstanceTemplateServiceAccountOutputReference
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ServiceAccountInput() *GoogleComputeInstanceTemplateServiceAccount {
	var returns *GoogleComputeInstanceTemplateServiceAccount
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ShieldedInstanceConfig() GoogleComputeInstanceTemplateShieldedInstanceConfigOutputReference {
	var returns GoogleComputeInstanceTemplateShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) ShieldedInstanceConfigInput() *GoogleComputeInstanceTemplateShieldedInstanceConfig {
	var returns *GoogleComputeInstanceTemplateShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TagsFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) Timeouts() GoogleComputeInstanceTemplateTimeoutsOutputReference {
	var returns GoogleComputeInstanceTemplateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template google_compute_instance_template} Resource.
func NewGoogleComputeInstanceTemplate(scope constructs.Construct, id *string, config *GoogleComputeInstanceTemplateConfig) GoogleComputeInstanceTemplate {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceTemplate{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_instance_template google_compute_instance_template} Resource.
func NewGoogleComputeInstanceTemplate_Override(g GoogleComputeInstanceTemplate, scope constructs.Construct, id *string, config *GoogleComputeInstanceTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetEnableDisplay(val interface{}) {
	if err := j.validateSetEnableDisplayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDisplay",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetInstanceDescription(val *string) {
	if err := j.validateSetInstanceDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceDescription",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetKeyRevocationActionType(val *string) {
	if err := j.validateSetKeyRevocationActionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyRevocationActionType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetMetadataStartupScript(val *string) {
	if err := j.validateSetMetadataStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataStartupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetPartnerMetadata(val *map[string]*string) {
	if err := j.validateSetPartnerMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerMetadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetResourceManagerTags(val *map[string]*string) {
	if err := j.validateSetResourceManagerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceManagerTags",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceTemplate)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeInstanceTemplate resource upon running "cdktf plan <stack-name>".
func GoogleComputeInstanceTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
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
func GoogleComputeInstanceTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInstanceTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeInstanceTemplate.GoogleComputeInstanceTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceTemplate) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutAdvancedMachineFeatures(value *GoogleComputeInstanceTemplateAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutConfidentialInstanceConfig(value *GoogleComputeInstanceTemplateConfidentialInstanceConfig) {
	if err := g.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutDisk(value interface{}) {
	if err := g.validatePutDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutNetworkInterface(value interface{}) {
	if err := g.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutNetworkPerformanceConfig(value *GoogleComputeInstanceTemplateNetworkPerformanceConfig) {
	if err := g.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutReservationAffinity(value *GoogleComputeInstanceTemplateReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutScheduling(value *GoogleComputeInstanceTemplateScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutServiceAccount(value *GoogleComputeInstanceTemplateServiceAccount) {
	if err := g.validatePutServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutShieldedInstanceConfig(value *GoogleComputeInstanceTemplateShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) PutTimeouts(value *GoogleComputeInstanceTemplateTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		g,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetEnableDisplay() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDisplay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetInstanceDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetKeyRevocationActionType() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyRevocationActionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetMetadataStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetPartnerMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetPartnerMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

