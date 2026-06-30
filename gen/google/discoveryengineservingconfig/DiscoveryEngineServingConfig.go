package discoveryengineservingconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryengineservingconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_serving_config google_discovery_engine_serving_config}.
type DiscoveryEngineServingConfig interface {
	cdktf.TerraformResource
	BoostControlIds() *[]*string
	SetBoostControlIds(val *[]*string)
	BoostControlIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollectionId() *string
	SetCollectionId(val *string)
	CollectionIdInput() *string
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
	EngineId() *string
	SetEngineId(val *string)
	EngineIdInput() *string
	FilterControlIds() *[]*string
	SetFilterControlIds(val *[]*string)
	FilterControlIdsInput() *[]*string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PromoteControlIds() *[]*string
	SetPromoteControlIds(val *[]*string)
	PromoteControlIdsInput() *[]*string
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
	RedirectControlIds() *[]*string
	SetRedirectControlIds(val *[]*string)
	RedirectControlIdsInput() *[]*string
	ServingConfigId() *string
	SetServingConfigId(val *string)
	ServingConfigIdInput() *string
	SynonymsControlIds() *[]*string
	SetSynonymsControlIds(val *[]*string)
	SynonymsControlIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DiscoveryEngineServingConfigTimeoutsOutputReference
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
	PutTimeouts(value *DiscoveryEngineServingConfigTimeouts)
	ResetBoostControlIds()
	ResetCollectionId()
	ResetFilterControlIds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPromoteControlIds()
	ResetRedirectControlIds()
	ResetServingConfigId()
	ResetSynonymsControlIds()
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

// The jsii proxy struct for DiscoveryEngineServingConfig
type jsiiProxy_DiscoveryEngineServingConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) BoostControlIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boostControlIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) BoostControlIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boostControlIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) EngineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) EngineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) FilterControlIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterControlIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) FilterControlIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterControlIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) PromoteControlIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"promoteControlIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) PromoteControlIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"promoteControlIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) RedirectControlIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectControlIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) RedirectControlIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"redirectControlIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) ServingConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) ServingConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SynonymsControlIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"synonymsControlIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SynonymsControlIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"synonymsControlIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) Timeouts() DiscoveryEngineServingConfigTimeoutsOutputReference {
	var returns DiscoveryEngineServingConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_serving_config google_discovery_engine_serving_config} Resource.
func NewDiscoveryEngineServingConfig(scope constructs.Construct, id *string, config *DiscoveryEngineServingConfigConfig) DiscoveryEngineServingConfig {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineServingConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineServingConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_serving_config google_discovery_engine_serving_config} Resource.
func NewDiscoveryEngineServingConfig_Override(d DiscoveryEngineServingConfig, scope constructs.Construct, id *string, config *DiscoveryEngineServingConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetBoostControlIds(val *[]*string) {
	if err := j.validateSetBoostControlIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boostControlIds",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetEngineId(val *string) {
	if err := j.validateSetEngineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetFilterControlIds(val *[]*string) {
	if err := j.validateSetFilterControlIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterControlIds",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetPromoteControlIds(val *[]*string) {
	if err := j.validateSetPromoteControlIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promoteControlIds",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetRedirectControlIds(val *[]*string) {
	if err := j.validateSetRedirectControlIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectControlIds",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetServingConfigId(val *string) {
	if err := j.validateSetServingConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servingConfigId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineServingConfig) SetSynonymsControlIds(val *[]*string) {
	if err := j.validateSetSynonymsControlIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synonymsControlIds",
		val,
	)
}

// Generates CDKTF code for importing a DiscoveryEngineServingConfig resource upon running "cdktf plan <stack-name>".
func DiscoveryEngineServingConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineServingConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
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
func DiscoveryEngineServingConfig_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineServingConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineServingConfig_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineServingConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineServingConfig_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineServingConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineServingConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.discoveryEngineServingConfig.DiscoveryEngineServingConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) PutTimeouts(value *DiscoveryEngineServingConfigTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetBoostControlIds() {
	_jsii_.InvokeVoid(
		d,
		"resetBoostControlIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetCollectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetCollectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetFilterControlIds() {
	_jsii_.InvokeVoid(
		d,
		"resetFilterControlIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetPromoteControlIds() {
	_jsii_.InvokeVoid(
		d,
		"resetPromoteControlIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetRedirectControlIds() {
	_jsii_.InvokeVoid(
		d,
		"resetRedirectControlIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetServingConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetServingConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetSynonymsControlIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSynonymsControlIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineServingConfig) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
