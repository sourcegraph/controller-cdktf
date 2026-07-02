package vmwareenginenetworkpeering

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/vmwareenginenetworkpeering/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vmwareengine_network_peering google_vmwareengine_network_peering}.
type VmwareengineNetworkPeering interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExportCustomRoutes() any
	SetExportCustomRoutes(val any)
	ExportCustomRoutesInput() any
	ExportCustomRoutesWithPublicIp() any
	SetExportCustomRoutesWithPublicIp(val any)
	ExportCustomRoutesWithPublicIpInput() any
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
	ImportCustomRoutes() any
	SetImportCustomRoutes(val any)
	ImportCustomRoutesInput() any
	ImportCustomRoutesWithPublicIp() any
	SetImportCustomRoutesWithPublicIp(val any)
	ImportCustomRoutesWithPublicIpInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerNetwork() *string
	SetPeerNetwork(val *string)
	PeerNetworkInput() *string
	PeerNetworkType() *string
	SetPeerNetworkType(val *string)
	PeerNetworkTypeInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	State() *string
	StateDetails() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VmwareengineNetworkPeeringTimeoutsOutputReference
	TimeoutsInput() any
	Uid() *string
	UpdateTime() *string
	VmwareEngineNetwork() *string
	SetVmwareEngineNetwork(val *string)
	VmwareEngineNetworkCanonical() *string
	VmwareEngineNetworkInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *VmwareengineNetworkPeeringTimeouts)
	ResetDescription()
	ResetExportCustomRoutes()
	ResetExportCustomRoutesWithPublicIp()
	ResetId()
	ResetImportCustomRoutes()
	ResetImportCustomRoutesWithPublicIp()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for VmwareengineNetworkPeering
type jsiiProxy_VmwareengineNetworkPeering struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VmwareengineNetworkPeering) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ExportCustomRoutes() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ExportCustomRoutesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ExportCustomRoutesWithPublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ExportCustomRoutesWithPublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ImportCustomRoutes() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ImportCustomRoutesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ImportCustomRoutesWithPublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ImportCustomRoutesWithPublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) PeerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) PeerNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) PeerNetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetworkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) PeerNetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerNetworkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Timeouts() VmwareengineNetworkPeeringTimeoutsOutputReference {
	var returns VmwareengineNetworkPeeringTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) VmwareEngineNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) VmwareEngineNetworkCanonical() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetworkCanonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineNetworkPeering) VmwareEngineNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetworkInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vmwareengine_network_peering google_vmwareengine_network_peering} Resource.
func NewVmwareengineNetworkPeering(scope constructs.Construct, id *string, config *VmwareengineNetworkPeeringConfig) VmwareengineNetworkPeering {
	_init_.Initialize()

	if err := validateNewVmwareengineNetworkPeeringParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareengineNetworkPeering{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/vmwareengine_network_peering google_vmwareengine_network_peering} Resource.
func NewVmwareengineNetworkPeering_Override(v VmwareengineNetworkPeering, scope constructs.Construct, id *string, config *VmwareengineNetworkPeeringConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		[]any{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetExportCustomRoutes(val any) {
	if err := j.validateSetExportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetExportCustomRoutesWithPublicIp(val any) {
	if err := j.validateSetExportCustomRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportCustomRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetImportCustomRoutes(val any) {
	if err := j.validateSetImportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetImportCustomRoutesWithPublicIp(val any) {
	if err := j.validateSetImportCustomRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importCustomRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetPeerNetwork(val *string) {
	if err := j.validateSetPeerNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerNetwork",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetPeerNetworkType(val *string) {
	if err := j.validateSetPeerNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerNetworkType",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VmwareengineNetworkPeering) SetVmwareEngineNetwork(val *string) {
	if err := j.validateSetVmwareEngineNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmwareEngineNetwork",
		val,
	)
}

// Generates CDKTF code for importing a VmwareengineNetworkPeering resource upon running "cdktf plan <stack-name>".
func VmwareengineNetworkPeering_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVmwareengineNetworkPeering_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func VmwareengineNetworkPeering_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateVmwareengineNetworkPeering_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func VmwareengineNetworkPeering_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateVmwareengineNetworkPeering_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func VmwareengineNetworkPeering_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateVmwareengineNetworkPeering_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func VmwareengineNetworkPeering_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.vmwareengineNetworkPeering.VmwareengineNetworkPeering",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) AddOverride(path *string, value any) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]any{path, value},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]any{id, provider},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]any{id},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) MoveTo(moveTarget *string, index any) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]any{id},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) PutTimeouts(value *VmwareengineNetworkPeeringTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]any{value},
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetExportCustomRoutes() {
	_jsii_.InvokeVoid(
		v,
		"resetExportCustomRoutes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetExportCustomRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		v,
		"resetExportCustomRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetImportCustomRoutes() {
	_jsii_.InvokeVoid(
		v,
		"resetImportCustomRoutes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetImportCustomRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		v,
		"resetImportCustomRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetProject() {
	_jsii_.InvokeVoid(
		v,
		"resetProject",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineNetworkPeering) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineNetworkPeering) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
