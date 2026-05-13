package googlediscoveryenginewidgetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginewidgetconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config google_discovery_engine_widget_config}.
type GoogleDiscoveryEngineWidgetConfig interface {
	cdktf.TerraformResource
	AccessSettings() GoogleDiscoveryEngineWidgetConfigAccessSettingsOutputReference
	AccessSettingsInput() *GoogleDiscoveryEngineWidgetConfigAccessSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollectionId() *string
	SetCollectionId(val *string)
	CollectionIdInput() *string
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
	HomepageSetting() GoogleDiscoveryEngineWidgetConfigHomepageSettingOutputReference
	HomepageSettingInput() *GoogleDiscoveryEngineWidgetConfigHomepageSetting
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDiscoveryEngineWidgetConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UiBranding() GoogleDiscoveryEngineWidgetConfigUiBrandingOutputReference
	UiBrandingInput() *GoogleDiscoveryEngineWidgetConfigUiBranding
	UiSettings() GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference
	UiSettingsInput() *GoogleDiscoveryEngineWidgetConfigUiSettings
	WidgetConfigId() *string
	SetWidgetConfigId(val *string)
	WidgetConfigIdInput() *string
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
	PutAccessSettings(value *GoogleDiscoveryEngineWidgetConfigAccessSettings)
	PutHomepageSetting(value *GoogleDiscoveryEngineWidgetConfigHomepageSetting)
	PutTimeouts(value *GoogleDiscoveryEngineWidgetConfigTimeouts)
	PutUiBranding(value *GoogleDiscoveryEngineWidgetConfigUiBranding)
	PutUiSettings(value *GoogleDiscoveryEngineWidgetConfigUiSettings)
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

// The jsii proxy struct for GoogleDiscoveryEngineWidgetConfig
type jsiiProxy_GoogleDiscoveryEngineWidgetConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) AccessSettings() GoogleDiscoveryEngineWidgetConfigAccessSettingsOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigAccessSettingsOutputReference
	_jsii_.Get(
		j,
		"accessSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) AccessSettingsInput() *GoogleDiscoveryEngineWidgetConfigAccessSettings {
	var returns *GoogleDiscoveryEngineWidgetConfigAccessSettings
	_jsii_.Get(
		j,
		"accessSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) CollectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) EngineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) EngineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) HomepageSetting() GoogleDiscoveryEngineWidgetConfigHomepageSettingOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigHomepageSettingOutputReference
	_jsii_.Get(
		j,
		"homepageSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) HomepageSettingInput() *GoogleDiscoveryEngineWidgetConfigHomepageSetting {
	var returns *GoogleDiscoveryEngineWidgetConfigHomepageSetting
	_jsii_.Get(
		j,
		"homepageSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) Timeouts() GoogleDiscoveryEngineWidgetConfigTimeoutsOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) UiBranding() GoogleDiscoveryEngineWidgetConfigUiBrandingOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigUiBrandingOutputReference
	_jsii_.Get(
		j,
		"uiBranding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) UiBrandingInput() *GoogleDiscoveryEngineWidgetConfigUiBranding {
	var returns *GoogleDiscoveryEngineWidgetConfigUiBranding
	_jsii_.Get(
		j,
		"uiBrandingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) UiSettings() GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference {
	var returns GoogleDiscoveryEngineWidgetConfigUiSettingsOutputReference
	_jsii_.Get(
		j,
		"uiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) UiSettingsInput() *GoogleDiscoveryEngineWidgetConfigUiSettings {
	var returns *GoogleDiscoveryEngineWidgetConfigUiSettings
	_jsii_.Get(
		j,
		"uiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) WidgetConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) WidgetConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widgetConfigIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config google_discovery_engine_widget_config} Resource.
func NewGoogleDiscoveryEngineWidgetConfig(scope constructs.Construct, id *string, config *GoogleDiscoveryEngineWidgetConfigConfig) GoogleDiscoveryEngineWidgetConfig {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineWidgetConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineWidgetConfig{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_widget_config google_discovery_engine_widget_config} Resource.
func NewGoogleDiscoveryEngineWidgetConfig_Override(g GoogleDiscoveryEngineWidgetConfig, scope constructs.Construct, id *string, config *GoogleDiscoveryEngineWidgetConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetEngineId(val *string) {
	if err := j.validateSetEngineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineId",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineWidgetConfig)SetWidgetConfigId(val *string) {
	if err := j.validateSetWidgetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"widgetConfigId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDiscoveryEngineWidgetConfig resource upon running "cdktf plan <stack-name>".
func GoogleDiscoveryEngineWidgetConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineWidgetConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
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
func GoogleDiscoveryEngineWidgetConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineWidgetConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDiscoveryEngineWidgetConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineWidgetConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDiscoveryEngineWidgetConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDiscoveryEngineWidgetConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDiscoveryEngineWidgetConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDiscoveryEngineWidgetConfig.GoogleDiscoveryEngineWidgetConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) PutAccessSettings(value *GoogleDiscoveryEngineWidgetConfigAccessSettings) {
	if err := g.validatePutAccessSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) PutHomepageSetting(value *GoogleDiscoveryEngineWidgetConfigHomepageSetting) {
	if err := g.validatePutHomepageSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHomepageSetting",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) PutTimeouts(value *GoogleDiscoveryEngineWidgetConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) PutUiBranding(value *GoogleDiscoveryEngineWidgetConfigUiBranding) {
	if err := g.validatePutUiBrandingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUiBranding",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) PutUiSettings(value *GoogleDiscoveryEngineWidgetConfigUiSettings) {
	if err := g.validatePutUiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUiSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetAccessSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetCollectionId() {
	_jsii_.InvokeVoid(
		g,
		"resetCollectionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetHomepageSetting() {
	_jsii_.InvokeVoid(
		g,
		"resetHomepageSetting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetUiBranding() {
	_jsii_.InvokeVoid(
		g,
		"resetUiBranding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetUiSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetUiSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ResetWidgetConfigId() {
	_jsii_.InvokeVoid(
		g,
		"resetWidgetConfigId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineWidgetConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

