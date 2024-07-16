package googledialogflowcxsecuritysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxsecuritysettings/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_security_settings google_dialogflow_cx_security_settings}.
type GoogleDialogflowCxSecuritySettings interface {
	cdktf.TerraformResource
	AudioExportSettings() GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference
	AudioExportSettingsInput() *GoogleDialogflowCxSecuritySettingsAudioExportSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DeidentifyTemplate() *string
	SetDeidentifyTemplate(val *string)
	DeidentifyTemplateInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	InsightsExportSettings() GoogleDialogflowCxSecuritySettingsInsightsExportSettingsOutputReference
	InsightsExportSettingsInput() *GoogleDialogflowCxSecuritySettingsInsightsExportSettings
	InspectTemplate() *string
	SetInspectTemplate(val *string)
	InspectTemplateInput() *string
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
	PurgeDataTypes() *[]*string
	SetPurgeDataTypes(val *[]*string)
	PurgeDataTypesInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RedactionScope() *string
	SetRedactionScope(val *string)
	RedactionScopeInput() *string
	RedactionStrategy() *string
	SetRedactionStrategy(val *string)
	RedactionStrategyInput() *string
	RetentionStrategy() *string
	SetRetentionStrategy(val *string)
	RetentionStrategyInput() *string
	RetentionWindowDays() *float64
	SetRetentionWindowDays(val *float64)
	RetentionWindowDaysInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDialogflowCxSecuritySettingsTimeoutsOutputReference
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
	PutAudioExportSettings(value *GoogleDialogflowCxSecuritySettingsAudioExportSettings)
	PutInsightsExportSettings(value *GoogleDialogflowCxSecuritySettingsInsightsExportSettings)
	PutTimeouts(value *GoogleDialogflowCxSecuritySettingsTimeouts)
	ResetAudioExportSettings()
	ResetDeidentifyTemplate()
	ResetId()
	ResetInsightsExportSettings()
	ResetInspectTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPurgeDataTypes()
	ResetRedactionScope()
	ResetRedactionStrategy()
	ResetRetentionStrategy()
	ResetRetentionWindowDays()
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

// The jsii proxy struct for GoogleDialogflowCxSecuritySettings
type jsiiProxy_GoogleDialogflowCxSecuritySettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) AudioExportSettings() GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference {
	var returns GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference
	_jsii_.Get(
		j,
		"audioExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) AudioExportSettingsInput() *GoogleDialogflowCxSecuritySettingsAudioExportSettings {
	var returns *GoogleDialogflowCxSecuritySettingsAudioExportSettings
	_jsii_.Get(
		j,
		"audioExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) DeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) DeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) InsightsExportSettings() GoogleDialogflowCxSecuritySettingsInsightsExportSettingsOutputReference {
	var returns GoogleDialogflowCxSecuritySettingsInsightsExportSettingsOutputReference
	_jsii_.Get(
		j,
		"insightsExportSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) InsightsExportSettingsInput() *GoogleDialogflowCxSecuritySettingsInsightsExportSettings {
	var returns *GoogleDialogflowCxSecuritySettingsInsightsExportSettings
	_jsii_.Get(
		j,
		"insightsExportSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) InspectTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) InspectTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) PurgeDataTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"purgeDataTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) PurgeDataTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"purgeDataTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RedactionScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RedactionScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RedactionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RedactionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redactionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RetentionStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RetentionStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RetentionWindowDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) RetentionWindowDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionWindowDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) Timeouts() GoogleDialogflowCxSecuritySettingsTimeoutsOutputReference {
	var returns GoogleDialogflowCxSecuritySettingsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_security_settings google_dialogflow_cx_security_settings} Resource.
func NewGoogleDialogflowCxSecuritySettings(scope constructs.Construct, id *string, config *GoogleDialogflowCxSecuritySettingsConfig) GoogleDialogflowCxSecuritySettings {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxSecuritySettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxSecuritySettings{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_security_settings google_dialogflow_cx_security_settings} Resource.
func NewGoogleDialogflowCxSecuritySettings_Override(g GoogleDialogflowCxSecuritySettings, scope constructs.Construct, id *string, config *GoogleDialogflowCxSecuritySettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetDeidentifyTemplate(val *string) {
	if err := j.validateSetDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetInspectTemplate(val *string) {
	if err := j.validateSetInspectTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetPurgeDataTypes(val *[]*string) {
	if err := j.validateSetPurgeDataTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeDataTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetRedactionScope(val *string) {
	if err := j.validateSetRedactionScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactionScope",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetRedactionStrategy(val *string) {
	if err := j.validateSetRedactionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redactionStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetRetentionStrategy(val *string) {
	if err := j.validateSetRetentionStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionStrategy",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettings)SetRetentionWindowDays(val *float64) {
	if err := j.validateSetRetentionWindowDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionWindowDays",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDialogflowCxSecuritySettings resource upon running "cdktf plan <stack-name>".
func GoogleDialogflowCxSecuritySettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxSecuritySettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
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
func GoogleDialogflowCxSecuritySettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxSecuritySettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowCxSecuritySettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxSecuritySettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowCxSecuritySettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxSecuritySettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDialogflowCxSecuritySettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) PutAudioExportSettings(value *GoogleDialogflowCxSecuritySettingsAudioExportSettings) {
	if err := g.validatePutAudioExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAudioExportSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) PutInsightsExportSettings(value *GoogleDialogflowCxSecuritySettingsInsightsExportSettings) {
	if err := g.validatePutInsightsExportSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInsightsExportSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) PutTimeouts(value *GoogleDialogflowCxSecuritySettingsTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetAudioExportSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioExportSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetDeidentifyTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetInsightsExportSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetInsightsExportSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetInspectTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetPurgeDataTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetPurgeDataTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetRedactionScope() {
	_jsii_.InvokeVoid(
		g,
		"resetRedactionScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetRedactionStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetRedactionStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetRetentionStrategy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionStrategy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetRetentionWindowDays() {
	_jsii_.InvokeVoid(
		g,
		"resetRetentionWindowDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

