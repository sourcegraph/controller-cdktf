package computenetworkpeeringroutesconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/computenetworkpeeringroutesconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_network_peering_routes_config google_compute_network_peering_routes_config}.
type ComputeNetworkPeeringRoutesConfig interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExportCustomRoutes() any
	SetExportCustomRoutes(val any)
	ExportCustomRoutesInput() any
	ExportSubnetRoutesWithPublicIp() any
	SetExportSubnetRoutesWithPublicIp(val any)
	ExportSubnetRoutesWithPublicIpInput() any
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
	ImportSubnetRoutesWithPublicIp() any
	SetImportSubnetRoutesWithPublicIp(val any)
	ImportSubnetRoutesWithPublicIpInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	Peering() *string
	SetPeering(val *string)
	PeeringInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeNetworkPeeringRoutesConfigTimeoutsOutputReference
	TimeoutsInput() any
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
	PutTimeouts(value *ComputeNetworkPeeringRoutesConfigTimeouts)
	ResetExportSubnetRoutesWithPublicIp()
	ResetId()
	ResetImportSubnetRoutesWithPublicIp()
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

// The jsii proxy struct for ComputeNetworkPeeringRoutesConfig
type jsiiProxy_ComputeNetworkPeeringRoutesConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ExportCustomRoutes() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ExportCustomRoutesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"exportCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ExportSubnetRoutesWithPublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"exportSubnetRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ExportSubnetRoutesWithPublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"exportSubnetRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ImportCustomRoutes() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ImportCustomRoutesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"importCustomRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ImportSubnetRoutesWithPublicIp() any {
	var returns any
	_jsii_.Get(
		j,
		"importSubnetRoutesWithPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ImportSubnetRoutesWithPublicIpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"importSubnetRoutesWithPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Peering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) PeeringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) Timeouts() ComputeNetworkPeeringRoutesConfigTimeoutsOutputReference {
	var returns ComputeNetworkPeeringRoutesConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_network_peering_routes_config google_compute_network_peering_routes_config} Resource.
func NewComputeNetworkPeeringRoutesConfig(scope constructs.Construct, id *string, config *ComputeNetworkPeeringRoutesConfigConfig) ComputeNetworkPeeringRoutesConfig {
	_init_.Initialize()

	if err := validateNewComputeNetworkPeeringRoutesConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeNetworkPeeringRoutesConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/compute_network_peering_routes_config google_compute_network_peering_routes_config} Resource.
func NewComputeNetworkPeeringRoutesConfig_Override(c ComputeNetworkPeeringRoutesConfig, scope constructs.Construct, id *string, config *ComputeNetworkPeeringRoutesConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetExportCustomRoutes(val any) {
	if err := j.validateSetExportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetExportSubnetRoutesWithPublicIp(val any) {
	if err := j.validateSetExportSubnetRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportSubnetRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetImportCustomRoutes(val any) {
	if err := j.validateSetImportCustomRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importCustomRoutes",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetImportSubnetRoutesWithPublicIp(val any) {
	if err := j.validateSetImportSubnetRoutesWithPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importSubnetRoutesWithPublicIp",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetPeering(val *string) {
	if err := j.validateSetPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peering",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ComputeNetworkPeeringRoutesConfig resource upon running "cdktf plan <stack-name>".
func ComputeNetworkPeeringRoutesConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeNetworkPeeringRoutesConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
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
func ComputeNetworkPeeringRoutesConfig_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateComputeNetworkPeeringRoutesConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeNetworkPeeringRoutesConfig_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateComputeNetworkPeeringRoutesConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeNetworkPeeringRoutesConfig_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateComputeNetworkPeeringRoutesConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ComputeNetworkPeeringRoutesConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeNetworkPeeringRoutesConfig.ComputeNetworkPeeringRoutesConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) PutTimeouts(value *ComputeNetworkPeeringRoutesConfigTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetExportSubnetRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		c,
		"resetExportSubnetRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetImportSubnetRoutesWithPublicIp() {
	_jsii_.InvokeVoid(
		c,
		"resetImportSubnetRoutesWithPublicIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkPeeringRoutesConfig) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
