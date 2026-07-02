package discoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginewidgetconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_widget_config google_discovery_engine_widget_config}.
type DiscoveryEngineWidgetConfig interface {
	cdktf.TerraformResource
	AccessSettings() DiscoveryEngineWidgetConfigAccessSettingsOutputReference
	AccessSettingsInput() *DiscoveryEngineWidgetConfigAccessSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollectionId() *string
	SetCollectionId(val *string)
	CollectionIdInput() *string
	ConfigId() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HomepageSetting() DiscoveryEngineWidgetConfigHomepageSettingOutputReference
	HomepageSettingInput() *DiscoveryEngineWidgetConfigHomepageSetting
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
	Timeouts() DiscoveryEngineWidgetConfigTimeoutsOutputReference
	TimeoutsInput() any
	UiBranding() DiscoveryEngineWidgetConfigUiBrandingOutputReference
	UiBrandingInput() *DiscoveryEngineWidgetConfigUiBranding
	UiSettings() DiscoveryEngineWidgetConfigUiSettingsOutputReference
	UiSettingsInput() *DiscoveryEngineWidgetConfigUiSettings
	WidgetConfigId() *string
	SetWidgetConfigId(val *string)
	WidgetConfigIdInput() *string
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
	PutAccessSettings(value *DiscoveryEngineWidgetConfigAccessSettings)
	PutHomepageSetting(value *DiscoveryEngineWidgetConfigHomepageSetting)
	PutTimeouts(value *DiscoveryEngineWidgetConfigTimeouts)
	PutUiBranding(value *DiscoveryEngineWidgetConfigUiBranding)
	PutUiSettings(value *DiscoveryEngineWidgetConfigUiSettings)
	ResetAccessSettings()
	ResetCollectionId()
	ResetHomepageSetting()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	ResetUiBranding()
	ResetUiSettings()
	ResetWidgetConfigId()
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

// The jsii proxy struct for DiscoveryEngineWidgetConfig
type jsiiProxy_DiscoveryEngineWidgetConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) AccessSettings() DiscoveryEngineWidgetConfigAccessSettingsOutputReference {
	var returns DiscoveryEngineWidgetConfigAccessSettingsOutputReference
	_jsii_.Get(
		j,
		"accessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) AccessSettingsInput() *DiscoveryEngineWidgetConfigAccessSettings {
	var returns *DiscoveryEngineWidgetConfigAccessSettings
	_jsii_.Get(
		j,
		"accessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) ConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) EngineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) EngineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) HomepageSetting() DiscoveryEngineWidgetConfigHomepageSettingOutputReference {
	var returns DiscoveryEngineWidgetConfigHomepageSettingOutputReference
	_jsii_.Get(
		j,
		"homepageSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) HomepageSettingInput() *DiscoveryEngineWidgetConfigHomepageSetting {
	var returns *DiscoveryEngineWidgetConfigHomepageSetting
	_jsii_.Get(
		j,
		"homepageSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) Timeouts() DiscoveryEngineWidgetConfigTimeoutsOutputReference {
	var returns DiscoveryEngineWidgetConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) UiBranding() DiscoveryEngineWidgetConfigUiBrandingOutputReference {
	var returns DiscoveryEngineWidgetConfigUiBrandingOutputReference
	_jsii_.Get(
		j,
		"uiBranding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) UiBrandingInput() *DiscoveryEngineWidgetConfigUiBranding {
	var returns *DiscoveryEngineWidgetConfigUiBranding
	_jsii_.Get(
		j,
		"uiBrandingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) UiSettings() DiscoveryEngineWidgetConfigUiSettingsOutputReference {
	var returns DiscoveryEngineWidgetConfigUiSettingsOutputReference
	_jsii_.Get(
		j,
		"uiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) UiSettingsInput() *DiscoveryEngineWidgetConfigUiSettings {
	var returns *DiscoveryEngineWidgetConfigUiSettings
	_jsii_.Get(
		j,
		"uiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) WidgetConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) WidgetConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetConfigIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_widget_config google_discovery_engine_widget_config} Resource.
func NewDiscoveryEngineWidgetConfig(scope constructs.Construct, id *string, config *DiscoveryEngineWidgetConfigConfig) DiscoveryEngineWidgetConfig {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineWidgetConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineWidgetConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/discovery_engine_widget_config google_discovery_engine_widget_config} Resource.
func NewDiscoveryEngineWidgetConfig_Override(d DiscoveryEngineWidgetConfig, scope constructs.Construct, id *string, config *DiscoveryEngineWidgetConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		[]any{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetEngineId(val *string) {
	if err := j.validateSetEngineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineWidgetConfig) SetWidgetConfigId(val *string) {
	if err := j.validateSetWidgetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetConfigId",
		val,
	)
}

// Generates CDKTF code for importing a DiscoveryEngineWidgetConfig resource upon running "cdktf plan <stack-name>".
func DiscoveryEngineWidgetConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineWidgetConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
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
func DiscoveryEngineWidgetConfig_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineWidgetConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineWidgetConfig_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineWidgetConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineWidgetConfig_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineWidgetConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineWidgetConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) AddOverride(path *string, value any) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]any{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]any{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) MoveTo(moveTarget *string, index any) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]any{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) PutAccessSettings(value *DiscoveryEngineWidgetConfigAccessSettings) {
	if err := d.validatePutAccessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAccessSettings",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) PutHomepageSetting(value *DiscoveryEngineWidgetConfigHomepageSetting) {
	if err := d.validatePutHomepageSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHomepageSetting",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) PutTimeouts(value *DiscoveryEngineWidgetConfigTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) PutUiBranding(value *DiscoveryEngineWidgetConfigUiBranding) {
	if err := d.validatePutUiBrandingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUiBranding",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) PutUiSettings(value *DiscoveryEngineWidgetConfigUiSettings) {
	if err := d.validatePutUiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUiSettings",
		[]any{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetAccessSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAccessSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetCollectionId() {
	_jsii_.InvokeVoid(
		d,
		"resetCollectionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetHomepageSetting() {
	_jsii_.InvokeVoid(
		d,
		"resetHomepageSetting",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetUiBranding() {
	_jsii_.InvokeVoid(
		d,
		"resetUiBranding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetUiSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetUiSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ResetWidgetConfigId() {
	_jsii_.InvokeVoid(
		d,
		"resetWidgetConfigId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineWidgetConfig) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
