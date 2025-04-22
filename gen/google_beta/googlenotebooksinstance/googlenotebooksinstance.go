package googlenotebooksinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenotebooksinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_notebooks_instance google_notebooks_instance}.
type GoogleNotebooksInstance interface {
	cdktf.TerraformResource
	AcceleratorConfig() GoogleNotebooksInstanceAcceleratorConfigOutputReference
	AcceleratorConfigInput() *GoogleNotebooksInstanceAcceleratorConfig
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerImage() GoogleNotebooksInstanceContainerImageOutputReference
	ContainerImageInput() *GoogleNotebooksInstanceContainerImage
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	SetCreateTime(val *string)
	CreateTimeInput() *string
	CustomGpuDriverPath() *string
	SetCustomGpuDriverPath(val *string)
	CustomGpuDriverPathInput() *string
	DataDiskSizeGb() *float64
	SetDataDiskSizeGb(val *float64)
	DataDiskSizeGbInput() *float64
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	DiskEncryption() *string
	SetDiskEncryption(val *string)
	DiskEncryptionInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallGpuDriver() interface{}
	SetInstallGpuDriver(val interface{})
	InstallGpuDriverInput() interface{}
	InstanceOwners() *[]*string
	SetInstanceOwners(val *[]*string)
	InstanceOwnersInput() *[]*string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	// The tree node.
	Node() constructs.Node
	NoProxyAccess() interface{}
	SetNoProxyAccess(val interface{})
	NoProxyAccessInput() interface{}
	NoPublicIp() interface{}
	SetNoPublicIp(val interface{})
	NoPublicIpInput() interface{}
	NoRemoveDataDisk() interface{}
	SetNoRemoveDataDisk(val interface{})
	NoRemoveDataDiskInput() interface{}
	PostStartupScript() *string
	SetPostStartupScript(val *string)
	PostStartupScriptInput() *string
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
	ProxyUri() *string
	// Experimental.
	RawOverrides() interface{}
	ReservationAffinity() GoogleNotebooksInstanceReservationAffinityOutputReference
	ReservationAffinityInput() *GoogleNotebooksInstanceReservationAffinity
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceAccountScopes() *[]*string
	SetServiceAccountScopes(val *[]*string)
	ServiceAccountScopesInput() *[]*string
	ShieldedInstanceConfig() GoogleNotebooksInstanceShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleNotebooksInstanceShieldedInstanceConfig
	State() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleNotebooksInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	SetUpdateTime(val *string)
	UpdateTimeInput() *string
	VmImage() GoogleNotebooksInstanceVmImageOutputReference
	VmImageInput() *GoogleNotebooksInstanceVmImage
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
	PutAcceleratorConfig(value *GoogleNotebooksInstanceAcceleratorConfig)
	PutContainerImage(value *GoogleNotebooksInstanceContainerImage)
	PutReservationAffinity(value *GoogleNotebooksInstanceReservationAffinity)
	PutShieldedInstanceConfig(value *GoogleNotebooksInstanceShieldedInstanceConfig)
	PutTimeouts(value *GoogleNotebooksInstanceTimeouts)
	PutVmImage(value *GoogleNotebooksInstanceVmImage)
	ResetAcceleratorConfig()
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetContainerImage()
	ResetCreateTime()
	ResetCustomGpuDriverPath()
	ResetDataDiskSizeGb()
	ResetDataDiskType()
	ResetDesiredState()
	ResetDiskEncryption()
	ResetId()
	ResetInstallGpuDriver()
	ResetInstanceOwners()
	ResetKmsKey()
	ResetLabels()
	ResetMetadata()
	ResetNetwork()
	ResetNicType()
	ResetNoProxyAccess()
	ResetNoPublicIp()
	ResetNoRemoveDataDisk()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostStartupScript()
	ResetProject()
	ResetReservationAffinity()
	ResetServiceAccount()
	ResetServiceAccountScopes()
	ResetShieldedInstanceConfig()
	ResetSubnet()
	ResetTags()
	ResetTimeouts()
	ResetUpdateTime()
	ResetVmImage()
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

// The jsii proxy struct for GoogleNotebooksInstance
type jsiiProxy_GoogleNotebooksInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNotebooksInstance) AcceleratorConfig() GoogleNotebooksInstanceAcceleratorConfigOutputReference {
	var returns GoogleNotebooksInstanceAcceleratorConfigOutputReference
	_jsii_.Get(
		j,
		"acceleratorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) AcceleratorConfigInput() *GoogleNotebooksInstanceAcceleratorConfig {
	var returns *GoogleNotebooksInstanceAcceleratorConfig
	_jsii_.Get(
		j,
		"acceleratorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ContainerImage() GoogleNotebooksInstanceContainerImageOutputReference {
	var returns GoogleNotebooksInstanceContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ContainerImageInput() *GoogleNotebooksInstanceContainerImage {
	var returns *GoogleNotebooksInstanceContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) CreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) CustomGpuDriverPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) CustomGpuDriverPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DataDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DataDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DiskEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) DiskEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) InstallGpuDriver() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) InstallGpuDriverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) InstanceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) InstanceOwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoProxyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noProxyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoProxyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noProxyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoRemoveDataDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) NoRemoveDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) PostStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) PostStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ProxyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ReservationAffinity() GoogleNotebooksInstanceReservationAffinityOutputReference {
	var returns GoogleNotebooksInstanceReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ReservationAffinityInput() *GoogleNotebooksInstanceReservationAffinity {
	var returns *GoogleNotebooksInstanceReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ServiceAccountScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ServiceAccountScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ShieldedInstanceConfig() GoogleNotebooksInstanceShieldedInstanceConfigOutputReference {
	var returns GoogleNotebooksInstanceShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) ShieldedInstanceConfigInput() *GoogleNotebooksInstanceShieldedInstanceConfig {
	var returns *GoogleNotebooksInstanceShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) Timeouts() GoogleNotebooksInstanceTimeoutsOutputReference {
	var returns GoogleNotebooksInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) UpdateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) VmImage() GoogleNotebooksInstanceVmImageOutputReference {
	var returns GoogleNotebooksInstanceVmImageOutputReference
	_jsii_.Get(
		j,
		"vmImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNotebooksInstance) VmImageInput() *GoogleNotebooksInstanceVmImage {
	var returns *GoogleNotebooksInstanceVmImage
	_jsii_.Get(
		j,
		"vmImageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_notebooks_instance google_notebooks_instance} Resource.
func NewGoogleNotebooksInstance(scope constructs.Construct, id *string, config *GoogleNotebooksInstanceConfig) GoogleNotebooksInstance {
	_init_.Initialize()

	if err := validateNewGoogleNotebooksInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNotebooksInstance{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_notebooks_instance google_notebooks_instance} Resource.
func NewGoogleNotebooksInstance_Override(g GoogleNotebooksInstance, scope constructs.Construct, id *string, config *GoogleNotebooksInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetCreateTime(val *string) {
	if err := j.validateSetCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTime",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetCustomGpuDriverPath(val *string) {
	if err := j.validateSetCustomGpuDriverPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customGpuDriverPath",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetDataDiskSizeGb(val *float64) {
	if err := j.validateSetDataDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetDataDiskType(val *string) {
	if err := j.validateSetDataDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetDiskEncryption(val *string) {
	if err := j.validateSetDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryption",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetInstallGpuDriver(val interface{}) {
	if err := j.validateSetInstallGpuDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installGpuDriver",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetInstanceOwners(val *[]*string) {
	if err := j.validateSetInstanceOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceOwners",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetNoProxyAccess(val interface{}) {
	if err := j.validateSetNoProxyAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noProxyAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetNoPublicIp(val interface{}) {
	if err := j.validateSetNoPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPublicIp",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetNoRemoveDataDisk(val interface{}) {
	if err := j.validateSetNoRemoveDataDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRemoveDataDisk",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetPostStartupScript(val *string) {
	if err := j.validateSetPostStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetServiceAccountScopes(val *[]*string) {
	if err := j.validateSetServiceAccountScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountScopes",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleNotebooksInstance)SetUpdateTime(val *string) {
	if err := j.validateSetUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTime",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNotebooksInstance resource upon running "cdktf plan <stack-name>".
func GoogleNotebooksInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNotebooksInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
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
func GoogleNotebooksInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNotebooksInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNotebooksInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNotebooksInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNotebooksInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNotebooksInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNotebooksInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleNotebooksInstance.GoogleNotebooksInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNotebooksInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNotebooksInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNotebooksInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutAcceleratorConfig(value *GoogleNotebooksInstanceAcceleratorConfig) {
	if err := g.validatePutAcceleratorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAcceleratorConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutContainerImage(value *GoogleNotebooksInstanceContainerImage) {
	if err := g.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutReservationAffinity(value *GoogleNotebooksInstanceReservationAffinity) {
	if err := g.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutShieldedInstanceConfig(value *GoogleNotebooksInstanceShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutTimeouts(value *GoogleNotebooksInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) PutVmImage(value *GoogleNotebooksInstanceVmImage) {
	if err := g.validatePutVmImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVmImage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetAcceleratorConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetContainerImage() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetCreateTime() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetCustomGpuDriverPath() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomGpuDriverPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetDataDiskSizeGb() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDiskSizeGb",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetDesiredState() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetInstallGpuDriver() {
	_jsii_.InvokeVoid(
		g,
		"resetInstallGpuDriver",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetInstanceOwners() {
	_jsii_.InvokeVoid(
		g,
		"resetInstanceOwners",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetKmsKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetNoProxyAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetNoProxyAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetNoPublicIp() {
	_jsii_.InvokeVoid(
		g,
		"resetNoPublicIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetNoRemoveDataDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetNoRemoveDataDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetPostStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetPostStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetServiceAccountScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetSubnet() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnet",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetUpdateTime() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdateTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) ResetVmImage() {
	_jsii_.InvokeVoid(
		g,
		"resetVmImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNotebooksInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNotebooksInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

