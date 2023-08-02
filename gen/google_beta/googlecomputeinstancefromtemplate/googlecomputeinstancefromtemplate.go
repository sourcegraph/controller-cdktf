package googlecomputeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeinstancefromtemplate/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_instance_from_template google_compute_instance_from_template}.
type GoogleComputeInstanceFromTemplate interface {
	cdktf.TerraformResource
	AdvancedMachineFeatures() GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference
	AdvancedMachineFeaturesInput() *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures
	AllowStoppingForUpdate() interface{}
	SetAllowStoppingForUpdate(val interface{})
	AllowStoppingForUpdateInput() interface{}
	AttachedDisk() GoogleComputeInstanceFromTemplateAttachedDiskList
	AttachedDiskInput() interface{}
	BootDisk() GoogleComputeInstanceFromTemplateBootDiskOutputReference
	BootDiskInput() *GoogleComputeInstanceFromTemplateBootDisk
	CanIpForward() interface{}
	SetCanIpForward(val interface{})
	CanIpForwardInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidentialInstanceConfig() GoogleComputeInstanceFromTemplateConfidentialInstanceConfigOutputReference
	ConfidentialInstanceConfigInput() *GoogleComputeInstanceFromTemplateConfidentialInstanceConfig
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
	GuestAccelerator() GoogleComputeInstanceFromTemplateGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
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
	NetworkInterface() GoogleComputeInstanceFromTemplateNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NetworkPerformanceConfig() GoogleComputeInstanceFromTemplateNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *GoogleComputeInstanceFromTemplateNetworkPerformanceConfig
	// The tree node.
	Node() constructs.Node
	Params() GoogleComputeInstanceFromTemplateParamsOutputReference
	ParamsInput() *GoogleComputeInstanceFromTemplateParams
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
	ReservationAffinity() GoogleComputeInstanceFromTemplateReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleComputeInstanceFromTemplateReservationAffinity
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Scheduling() GoogleComputeInstanceFromTemplateSchedulingOutputReference
	SchedulingInput() *GoogleComputeInstanceFromTemplateScheduling
	ScratchDisk() GoogleComputeInstanceFromTemplateScratchDiskList
	ScratchDiskInput() interface{}
	SelfLink() *string
	ServiceAccount() GoogleComputeInstanceFromTemplateServiceAccountList
	ServiceAccountInput() interface{}
	ShieldedInstanceConfig() GoogleComputeInstanceFromTemplateShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleComputeInstanceFromTemplateShieldedInstanceConfig
	SourceInstanceTemplate() *string
	SetSourceInstanceTemplate(val *string)
	SourceInstanceTemplateInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsFingerprint() *string
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeInstanceFromTemplateTimeoutsOutputReference
	TimeoutsInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedMachineFeatures(value *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures)
	PutAttachedDisk(value interface{})
	PutBootDisk(value *GoogleComputeInstanceFromTemplateBootDisk)
	PutConfidentialInstanceConfig(value *GoogleComputeInstanceFromTemplateConfidentialInstanceConfig)
	PutGuestAccelerator(value interface{})
	PutNetworkInterface(value interface{})
	PutNetworkPerformanceConfig(value *GoogleComputeInstanceFromTemplateNetworkPerformanceConfig)
	PutParams(value *GoogleComputeInstanceFromTemplateParams)
	PutReservationAffinity(value *GoogleComputeInstanceFromTemplateReservationAffinity)
	PutScheduling(value *GoogleComputeInstanceFromTemplateScheduling)
	PutScratchDisk(value interface{})
	PutServiceAccount(value interface{})
	PutShieldedInstanceConfig(value *GoogleComputeInstanceFromTemplateShieldedInstanceConfig)
	PutTimeouts(value *GoogleComputeInstanceFromTemplateTimeouts)
	ResetAdvancedMachineFeatures()
	ResetAllowStoppingForUpdate()
	ResetAttachedDisk()
	ResetBootDisk()
	ResetCanIpForward()
	ResetConfidentialInstanceConfig()
	ResetDeletionProtection()
	ResetDescription()
	ResetDesiredStatus()
	ResetEnableDisplay()
	ResetGuestAccelerator()
	ResetHostname()
	ResetId()
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
	ResetProject()
	ResetReservationAffinity()
	ResetResourcePolicies()
	ResetScheduling()
	ResetScratchDisk()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetTags()
	ResetTimeouts()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleComputeInstanceFromTemplate
type jsiiProxy_GoogleComputeInstanceFromTemplate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AdvancedMachineFeatures() GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference {
	var returns GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AdvancedMachineFeaturesInput() *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures {
	var returns *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"advancedMachineFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AllowStoppingForUpdate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AllowStoppingForUpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStoppingForUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AttachedDisk() GoogleComputeInstanceFromTemplateAttachedDiskList {
	var returns GoogleComputeInstanceFromTemplateAttachedDiskList
	_jsii_.Get(
		j,
		"attachedDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) AttachedDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachedDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) BootDisk() GoogleComputeInstanceFromTemplateBootDiskOutputReference {
	var returns GoogleComputeInstanceFromTemplateBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) BootDiskInput() *GoogleComputeInstanceFromTemplateBootDisk {
	var returns *GoogleComputeInstanceFromTemplateBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) CanIpForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) CanIpForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canIpForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ConfidentialInstanceConfig() GoogleComputeInstanceFromTemplateConfidentialInstanceConfigOutputReference {
	var returns GoogleComputeInstanceFromTemplateConfidentialInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"confidentialInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ConfidentialInstanceConfigInput() *GoogleComputeInstanceFromTemplateConfidentialInstanceConfig {
	var returns *GoogleComputeInstanceFromTemplateConfidentialInstanceConfig
	_jsii_.Get(
		j,
		"confidentialInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) CpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) CurrentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DesiredStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) DesiredStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) EnableDisplay() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) EnableDisplayInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDisplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) GuestAccelerator() GoogleComputeInstanceFromTemplateGuestAcceleratorList {
	var returns GoogleComputeInstanceFromTemplateGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MetadataFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MetadataStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MetadataStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) NetworkInterface() GoogleComputeInstanceFromTemplateNetworkInterfaceList {
	var returns GoogleComputeInstanceFromTemplateNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) NetworkPerformanceConfig() GoogleComputeInstanceFromTemplateNetworkPerformanceConfigOutputReference {
	var returns GoogleComputeInstanceFromTemplateNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) NetworkPerformanceConfigInput() *GoogleComputeInstanceFromTemplateNetworkPerformanceConfig {
	var returns *GoogleComputeInstanceFromTemplateNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Params() GoogleComputeInstanceFromTemplateParamsOutputReference {
	var returns GoogleComputeInstanceFromTemplateParamsOutputReference
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ParamsInput() *GoogleComputeInstanceFromTemplateParams {
	var returns *GoogleComputeInstanceFromTemplateParams
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ReservationAffinity() GoogleComputeInstanceFromTemplateReservationAffinityOutputReference {
	var returns GoogleComputeInstanceFromTemplateReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ReservationAffinityInput() *GoogleComputeInstanceFromTemplateReservationAffinity {
	var returns *GoogleComputeInstanceFromTemplateReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Scheduling() GoogleComputeInstanceFromTemplateSchedulingOutputReference {
	var returns GoogleComputeInstanceFromTemplateSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) SchedulingInput() *GoogleComputeInstanceFromTemplateScheduling {
	var returns *GoogleComputeInstanceFromTemplateScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ScratchDisk() GoogleComputeInstanceFromTemplateScratchDiskList {
	var returns GoogleComputeInstanceFromTemplateScratchDiskList
	_jsii_.Get(
		j,
		"scratchDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ScratchDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scratchDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ServiceAccount() GoogleComputeInstanceFromTemplateServiceAccountList {
	var returns GoogleComputeInstanceFromTemplateServiceAccountList
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ServiceAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ShieldedInstanceConfig() GoogleComputeInstanceFromTemplateShieldedInstanceConfigOutputReference {
	var returns GoogleComputeInstanceFromTemplateShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ShieldedInstanceConfigInput() *GoogleComputeInstanceFromTemplateShieldedInstanceConfig {
	var returns *GoogleComputeInstanceFromTemplateShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) SourceInstanceTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInstanceTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) SourceInstanceTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInstanceTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TagsFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Timeouts() GoogleComputeInstanceFromTemplateTimeoutsOutputReference {
	var returns GoogleComputeInstanceFromTemplateTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_instance_from_template google_compute_instance_from_template} Resource.
func NewGoogleComputeInstanceFromTemplate(scope constructs.Construct, id *string, config *GoogleComputeInstanceFromTemplateConfig) GoogleComputeInstanceFromTemplate {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromTemplate{}

	_jsii_.Create(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.76.0/docs/resources/google_compute_instance_from_template google_compute_instance_from_template} Resource.
func NewGoogleComputeInstanceFromTemplate_Override(g GoogleComputeInstanceFromTemplate, scope constructs.Construct, id *string, config *GoogleComputeInstanceFromTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetAllowStoppingForUpdate(val interface{}) {
	if err := j.validateSetAllowStoppingForUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStoppingForUpdate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetCanIpForward(val interface{}) {
	if err := j.validateSetCanIpForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canIpForward",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetDesiredStatus(val *string) {
	if err := j.validateSetDesiredStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredStatus",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetEnableDisplay(val interface{}) {
	if err := j.validateSetEnableDisplayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDisplay",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetMetadataStartupScript(val *string) {
	if err := j.validateSetMetadataStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataStartupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetSourceInstanceTemplate(val *string) {
	if err := j.validateSetSourceInstanceTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceInstanceTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplate)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
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
func GoogleComputeInstanceFromTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceFromTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeInstanceFromTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeInstanceFromTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeInstanceFromTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutAdvancedMachineFeatures(value *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures) {
	if err := g.validatePutAdvancedMachineFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedMachineFeatures",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutAttachedDisk(value interface{}) {
	if err := g.validatePutAttachedDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttachedDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutBootDisk(value *GoogleComputeInstanceFromTemplateBootDisk) {
	if err := g.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutConfidentialInstanceConfig(value *GoogleComputeInstanceFromTemplateConfidentialInstanceConfig) {
	if err := g.validatePutConfidentialInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfidentialInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutGuestAccelerator(value interface{}) {
	if err := g.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutNetworkInterface(value interface{}) {
	if err := g.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutNetworkPerformanceConfig(value *GoogleComputeInstanceFromTemplateNetworkPerformanceConfig) {
	if err := g.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutParams(value *GoogleComputeInstanceFromTemplateParams) {
	if err := g.validatePutParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParams",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutReservationAffinity(value *GoogleComputeInstanceFromTemplateReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutScheduling(value *GoogleComputeInstanceFromTemplateScheduling) {
	if err := g.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScheduling",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutScratchDisk(value interface{}) {
	if err := g.validatePutScratchDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScratchDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutServiceAccount(value interface{}) {
	if err := g.validatePutServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutShieldedInstanceConfig(value *GoogleComputeInstanceFromTemplateShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) PutTimeouts(value *GoogleComputeInstanceFromTemplateTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetAdvancedMachineFeatures() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedMachineFeatures",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetAllowStoppingForUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowStoppingForUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetAttachedDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetAttachedDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetBootDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetCanIpForward() {
	_jsii_.InvokeVoid(
		g,
		"resetCanIpForward",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetConfidentialInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConfidentialInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetDesiredStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetEnableDisplay() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableDisplay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		g,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetHostname() {
	_jsii_.InvokeVoid(
		g,
		"resetHostname",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetMetadataStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetParams() {
	_jsii_.InvokeVoid(
		g,
		"resetParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetScheduling() {
	_jsii_.InvokeVoid(
		g,
		"resetScheduling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetScratchDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetScratchDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

