package networkzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/networkzone/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone okta_network_zone}.
type NetworkZone interface {
	cdktf.TerraformResource
	Asns() *[]*string
	SetAsns(val *[]*string)
	AsnsInput() *[]*string
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
	DynamicLocations() *[]*string
	SetDynamicLocations(val *[]*string)
	DynamicLocationsExclude() *[]*string
	SetDynamicLocationsExclude(val *[]*string)
	DynamicLocationsExcludeInput() *[]*string
	DynamicLocationsInput() *[]*string
	DynamicProxyType() *string
	SetDynamicProxyType(val *string)
	DynamicProxyTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gateways() *[]*string
	SetGateways(val *[]*string)
	GatewaysInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpServiceCategoriesExclude() *[]*string
	SetIpServiceCategoriesExclude(val *[]*string)
	IpServiceCategoriesExcludeInput() *[]*string
	IpServiceCategoriesInclude() *[]*string
	SetIpServiceCategoriesInclude(val *[]*string)
	IpServiceCategoriesIncludeInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	Proxies() *[]*string
	SetProxies(val *[]*string)
	ProxiesInput() *[]*string
	// Experimental.
	RawOverrides() any
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Usage() *string
	SetUsage(val *string)
	UsageInput() *string
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
	ResetAsns()
	ResetDynamicLocations()
	ResetDynamicLocationsExclude()
	ResetDynamicProxyType()
	ResetGateways()
	ResetId()
	ResetIpServiceCategoriesExclude()
	ResetIpServiceCategoriesInclude()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProxies()
	ResetStatus()
	ResetUsage()
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

// The jsii proxy struct for NetworkZone
type jsiiProxy_NetworkZone struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkZone) Asns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) AsnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dynamicLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicLocationsExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dynamicLocationsExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicLocationsExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dynamicLocationsExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dynamicLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicProxyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicProxyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) DynamicProxyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamicProxyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Gateways() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gateways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) GatewaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatewaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) IpServiceCategoriesExclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipServiceCategoriesExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) IpServiceCategoriesExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipServiceCategoriesExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) IpServiceCategoriesInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipServiceCategoriesInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) IpServiceCategoriesIncludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipServiceCategoriesIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Proxies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) ProxiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"proxiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) Usage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkZone) UsageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone okta_network_zone} Resource.
func NewNetworkZone(scope constructs.Construct, id *string, config *NetworkZoneConfig) NetworkZone {
	_init_.Initialize()

	if err := validateNewNetworkZoneParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkZone{}

	_jsii_.Create(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/network_zone okta_network_zone} Resource.
func NewNetworkZone_Override(n NetworkZone, scope constructs.Construct, id *string, config *NetworkZoneConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		[]any{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkZone) SetAsns(val *[]*string) {
	if err := j.validateSetAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asns",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetDynamicLocations(val *[]*string) {
	if err := j.validateSetDynamicLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicLocations",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetDynamicLocationsExclude(val *[]*string) {
	if err := j.validateSetDynamicLocationsExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicLocationsExclude",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetDynamicProxyType(val *string) {
	if err := j.validateSetDynamicProxyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicProxyType",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetGateways(val *[]*string) {
	if err := j.validateSetGatewaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateways",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetIpServiceCategoriesExclude(val *[]*string) {
	if err := j.validateSetIpServiceCategoriesExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipServiceCategoriesExclude",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetIpServiceCategoriesInclude(val *[]*string) {
	if err := j.validateSetIpServiceCategoriesIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipServiceCategoriesInclude",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetProxies(val *[]*string) {
	if err := j.validateSetProxiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxies",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_NetworkZone) SetUsage(val *string) {
	if err := j.validateSetUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usage",
		val,
	)
}

// Generates CDKTF code for importing a NetworkZone resource upon running "cdktf plan <stack-name>".
func NetworkZone_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkZone_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.networkZone.NetworkZone",
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
func NetworkZone_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkZone_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkZone_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkZone_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkZone_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateNetworkZone_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func NetworkZone_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.networkZone.NetworkZone",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkZone) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (n *jsiiProxy_NetworkZone) AddOverride(path *string, value any) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]any{path, value},
	)
}

func (n *jsiiProxy_NetworkZone) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]any{id, provider},
	)
}

func (n *jsiiProxy_NetworkZone) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkZone) MoveTo(moveTarget *string, index any) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkZone) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]any{id},
	)
}

func (n *jsiiProxy_NetworkZone) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkZone) ResetAsns() {
	_jsii_.InvokeVoid(
		n,
		"resetAsns",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetDynamicLocations() {
	_jsii_.InvokeVoid(
		n,
		"resetDynamicLocations",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetDynamicLocationsExclude() {
	_jsii_.InvokeVoid(
		n,
		"resetDynamicLocationsExclude",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetDynamicProxyType() {
	_jsii_.InvokeVoid(
		n,
		"resetDynamicProxyType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetGateways() {
	_jsii_.InvokeVoid(
		n,
		"resetGateways",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetIpServiceCategoriesExclude() {
	_jsii_.InvokeVoid(
		n,
		"resetIpServiceCategoriesExclude",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetIpServiceCategoriesInclude() {
	_jsii_.InvokeVoid(
		n,
		"resetIpServiceCategoriesInclude",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetProxies() {
	_jsii_.InvokeVoid(
		n,
		"resetProxies",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetStatus() {
	_jsii_.InvokeVoid(
		n,
		"resetStatus",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) ResetUsage() {
	_jsii_.InvokeVoid(
		n,
		"resetUsage",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkZone) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkZone) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
