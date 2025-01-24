package googletpuv2vm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googletpuv2vm/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_tpu_v2_vm google_tpu_v2_vm}.
type GoogleTpuV2Vm interface {
	cdktf.TerraformResource
	AcceleratorConfig() GoogleTpuV2VmAcceleratorConfigOutputReference
	AcceleratorConfigInput() *GoogleTpuV2VmAcceleratorConfig
	AcceleratorType() *string
	SetAcceleratorType(val *string)
	AcceleratorTypeInput() *string
	ApiVersion() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrBlock() *string
	SetCidrBlock(val *string)
	CidrBlockInput() *string
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
	DataDisks() GoogleTpuV2VmDataDisksList
	DataDisksInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() *string
	HealthDescription() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	MultisliceNode() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GoogleTpuV2VmNetworkConfigOutputReference
	NetworkConfigInput() *GoogleTpuV2VmNetworkConfig
	NetworkEndpoints() GoogleTpuV2VmNetworkEndpointsList
	// The tree node.
	Node() constructs.Node
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
	QueuedResource() *string
	// Experimental.
	RawOverrides() interface{}
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	SchedulingConfig() GoogleTpuV2VmSchedulingConfigOutputReference
	SchedulingConfigInput() *GoogleTpuV2VmSchedulingConfig
	ServiceAccount() GoogleTpuV2VmServiceAccountOutputReference
	ServiceAccountInput() *GoogleTpuV2VmServiceAccount
	ShieldedInstanceConfig() GoogleTpuV2VmShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *GoogleTpuV2VmShieldedInstanceConfig
	State() *string
	Symptoms() GoogleTpuV2VmSymptomsList
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
	Timeouts() GoogleTpuV2VmTimeoutsOutputReference
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
	PutAcceleratorConfig(value *GoogleTpuV2VmAcceleratorConfig)
	PutDataDisks(value interface{})
	PutNetworkConfig(value *GoogleTpuV2VmNetworkConfig)
	PutSchedulingConfig(value *GoogleTpuV2VmSchedulingConfig)
	PutServiceAccount(value *GoogleTpuV2VmServiceAccount)
	PutShieldedInstanceConfig(value *GoogleTpuV2VmShieldedInstanceConfig)
	PutTimeouts(value *GoogleTpuV2VmTimeouts)
	ResetAcceleratorConfig()
	ResetAcceleratorType()
	ResetCidrBlock()
	ResetDataDisks()
	ResetDescription()
	ResetId()
	ResetLabels()
	ResetMetadata()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSchedulingConfig()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
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

// The jsii proxy struct for GoogleTpuV2Vm
type jsiiProxy_GoogleTpuV2Vm struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleTpuV2Vm) AcceleratorConfig() GoogleTpuV2VmAcceleratorConfigOutputReference {
	var returns GoogleTpuV2VmAcceleratorConfigOutputReference
	_jsii_.Get(
		j,
		"acceleratorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) AcceleratorConfigInput() *GoogleTpuV2VmAcceleratorConfig {
	var returns *GoogleTpuV2VmAcceleratorConfig
	_jsii_.Get(
		j,
		"acceleratorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) AcceleratorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) AcceleratorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) DataDisks() GoogleTpuV2VmDataDisksList {
	var returns GoogleTpuV2VmDataDisksList
	_jsii_.Get(
		j,
		"dataDisks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) DataDisksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDisksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Health() *string {
	var returns *string
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) HealthDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) MultisliceNode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"multisliceNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) NetworkConfig() GoogleTpuV2VmNetworkConfigOutputReference {
	var returns GoogleTpuV2VmNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) NetworkConfigInput() *GoogleTpuV2VmNetworkConfig {
	var returns *GoogleTpuV2VmNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) NetworkEndpoints() GoogleTpuV2VmNetworkEndpointsList {
	var returns GoogleTpuV2VmNetworkEndpointsList
	_jsii_.Get(
		j,
		"networkEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) QueuedResource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queuedResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) SchedulingConfig() GoogleTpuV2VmSchedulingConfigOutputReference {
	var returns GoogleTpuV2VmSchedulingConfigOutputReference
	_jsii_.Get(
		j,
		"schedulingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) SchedulingConfigInput() *GoogleTpuV2VmSchedulingConfig {
	var returns *GoogleTpuV2VmSchedulingConfig
	_jsii_.Get(
		j,
		"schedulingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ServiceAccount() GoogleTpuV2VmServiceAccountOutputReference {
	var returns GoogleTpuV2VmServiceAccountOutputReference
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ServiceAccountInput() *GoogleTpuV2VmServiceAccount {
	var returns *GoogleTpuV2VmServiceAccount
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ShieldedInstanceConfig() GoogleTpuV2VmShieldedInstanceConfigOutputReference {
	var returns GoogleTpuV2VmShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ShieldedInstanceConfigInput() *GoogleTpuV2VmShieldedInstanceConfig {
	var returns *GoogleTpuV2VmShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Symptoms() GoogleTpuV2VmSymptomsList {
	var returns GoogleTpuV2VmSymptomsList
	_jsii_.Get(
		j,
		"symptoms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Timeouts() GoogleTpuV2VmTimeoutsOutputReference {
	var returns GoogleTpuV2VmTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTpuV2Vm) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_tpu_v2_vm google_tpu_v2_vm} Resource.
func NewGoogleTpuV2Vm(scope constructs.Construct, id *string, config *GoogleTpuV2VmConfig) GoogleTpuV2Vm {
	_init_.Initialize()

	if err := validateNewGoogleTpuV2VmParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTpuV2Vm{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_tpu_v2_vm google_tpu_v2_vm} Resource.
func NewGoogleTpuV2Vm_Override(g GoogleTpuV2Vm, scope constructs.Construct, id *string, config *GoogleTpuV2VmConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetAcceleratorType(val *string) {
	if err := j.validateSetAcceleratorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratorType",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetCidrBlock(val *string) {
	if err := j.validateSetCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrBlock",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_GoogleTpuV2Vm)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a GoogleTpuV2Vm resource upon running "cdktf plan <stack-name>".
func GoogleTpuV2Vm_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleTpuV2Vm_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
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
func GoogleTpuV2Vm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleTpuV2Vm_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleTpuV2Vm_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleTpuV2Vm_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleTpuV2Vm_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleTpuV2Vm_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleTpuV2Vm_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleTpuV2Vm.GoogleTpuV2Vm",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTpuV2Vm) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTpuV2Vm) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTpuV2Vm) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutAcceleratorConfig(value *GoogleTpuV2VmAcceleratorConfig) {
	if err := g.validatePutAcceleratorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAcceleratorConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutDataDisks(value interface{}) {
	if err := g.validatePutDataDisksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataDisks",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutNetworkConfig(value *GoogleTpuV2VmNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutSchedulingConfig(value *GoogleTpuV2VmSchedulingConfig) {
	if err := g.validatePutSchedulingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSchedulingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutServiceAccount(value *GoogleTpuV2VmServiceAccount) {
	if err := g.validatePutServiceAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putServiceAccount",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutShieldedInstanceConfig(value *GoogleTpuV2VmShieldedInstanceConfig) {
	if err := g.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) PutTimeouts(value *GoogleTpuV2VmTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetAcceleratorConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetAcceleratorType() {
	_jsii_.InvokeVoid(
		g,
		"resetAcceleratorType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetCidrBlock() {
	_jsii_.InvokeVoid(
		g,
		"resetCidrBlock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetDataDisks() {
	_jsii_.InvokeVoid(
		g,
		"resetDataDisks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetSchedulingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedulingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTpuV2Vm) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTpuV2Vm) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

