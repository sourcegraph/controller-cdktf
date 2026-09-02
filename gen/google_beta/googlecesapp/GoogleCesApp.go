package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_app google_ces_app}.
type GoogleCesApp interface {
	cdktn.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AudioProcessingConfig() GoogleCesAppAudioProcessingConfigOutputReference
	AudioProcessingConfigInput() *GoogleCesAppAudioProcessingConfig
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCertificateSettings() GoogleCesAppClientCertificateSettingsOutputReference
	ClientCertificateSettingsInput() *GoogleCesAppClientCertificateSettings
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
	CreateTime() *string
	DataStoreSettings() GoogleCesAppDataStoreSettingsOutputReference
	DataStoreSettingsInput() *GoogleCesAppDataStoreSettings
	DefaultChannelProfile() GoogleCesAppDefaultChannelProfileOutputReference
	DefaultChannelProfileInput() *GoogleCesAppDefaultChannelProfile
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentCount() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Etag() *string
	EvaluationMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsOutputReference
	EvaluationMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholds
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalInstruction() *string
	SetGlobalInstruction(val *string)
	GlobalInstructionInput() *string
	Guardrails() *[]*string
	SetGuardrails(val *[]*string)
	GuardrailsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LanguageSettings() GoogleCesAppLanguageSettingsOutputReference
	LanguageSettingsInput() *GoogleCesAppLanguageSettings
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingSettings() GoogleCesAppLoggingSettingsOutputReference
	LoggingSettingsInput() *GoogleCesAppLoggingSettings
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	ModelSettings() GoogleCesAppModelSettingsOutputReference
	ModelSettingsInput() *GoogleCesAppModelSettings
	Name() *string
	// The tree node.
	Node() constructs.Node
	Pinned() interface{}
	SetPinned(val interface{})
	PinnedInput() interface{}
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RootAgent() *string
	SetRootAgent(val *string)
	RootAgentInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCesAppTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZoneSettings() GoogleCesAppTimeZoneSettingsOutputReference
	TimeZoneSettingsInput() *GoogleCesAppTimeZoneSettings
	UpdateTime() *string
	VariableDeclarations() GoogleCesAppVariableDeclarationsList
	VariableDeclarationsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutAudioProcessingConfig(value *GoogleCesAppAudioProcessingConfig)
	PutClientCertificateSettings(value *GoogleCesAppClientCertificateSettings)
	PutDataStoreSettings(value *GoogleCesAppDataStoreSettings)
	PutDefaultChannelProfile(value *GoogleCesAppDefaultChannelProfile)
	PutEvaluationMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholds)
	PutLanguageSettings(value *GoogleCesAppLanguageSettings)
	PutLoggingSettings(value *GoogleCesAppLoggingSettings)
	PutModelSettings(value *GoogleCesAppModelSettings)
	PutTimeouts(value *GoogleCesAppTimeouts)
	PutTimeZoneSettings(value *GoogleCesAppTimeZoneSettings)
	PutVariableDeclarations(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAudioProcessingConfig()
	ResetClientCertificateSettings()
	ResetDataStoreSettings()
	ResetDefaultChannelProfile()
	ResetDescription()
	ResetEvaluationMetricsThresholds()
	ResetGlobalInstruction()
	ResetGuardrails()
	ResetId()
	ResetLanguageSettings()
	ResetLoggingSettings()
	ResetMetadata()
	ResetModelSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPinned()
	ResetProject()
	ResetRootAgent()
	ResetTimeouts()
	ResetTimeZoneSettings()
	ResetVariableDeclarations()
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for GoogleCesApp
type jsiiProxy_GoogleCesApp struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleCesApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) AudioProcessingConfig() GoogleCesAppAudioProcessingConfigOutputReference {
	var returns GoogleCesAppAudioProcessingConfigOutputReference
	_jsii_.Get(
		j,
		"audioProcessingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) AudioProcessingConfigInput() *GoogleCesAppAudioProcessingConfig {
	var returns *GoogleCesAppAudioProcessingConfig
	_jsii_.Get(
		j,
		"audioProcessingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ClientCertificateSettings() GoogleCesAppClientCertificateSettingsOutputReference {
	var returns GoogleCesAppClientCertificateSettingsOutputReference
	_jsii_.Get(
		j,
		"clientCertificateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ClientCertificateSettingsInput() *GoogleCesAppClientCertificateSettings {
	var returns *GoogleCesAppClientCertificateSettings
	_jsii_.Get(
		j,
		"clientCertificateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DataStoreSettings() GoogleCesAppDataStoreSettingsOutputReference {
	var returns GoogleCesAppDataStoreSettingsOutputReference
	_jsii_.Get(
		j,
		"dataStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DataStoreSettingsInput() *GoogleCesAppDataStoreSettings {
	var returns *GoogleCesAppDataStoreSettings
	_jsii_.Get(
		j,
		"dataStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DefaultChannelProfile() GoogleCesAppDefaultChannelProfileOutputReference {
	var returns GoogleCesAppDefaultChannelProfileOutputReference
	_jsii_.Get(
		j,
		"defaultChannelProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DefaultChannelProfileInput() *GoogleCesAppDefaultChannelProfile {
	var returns *GoogleCesAppDefaultChannelProfile
	_jsii_.Get(
		j,
		"defaultChannelProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DeploymentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) EvaluationMetricsThresholds() GoogleCesAppEvaluationMetricsThresholdsOutputReference {
	var returns GoogleCesAppEvaluationMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"evaluationMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) EvaluationMetricsThresholdsInput() *GoogleCesAppEvaluationMetricsThresholds {
	var returns *GoogleCesAppEvaluationMetricsThresholds
	_jsii_.Get(
		j,
		"evaluationMetricsThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) GlobalInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) GlobalInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) GuardrailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) LanguageSettings() GoogleCesAppLanguageSettingsOutputReference {
	var returns GoogleCesAppLanguageSettingsOutputReference
	_jsii_.Get(
		j,
		"languageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) LanguageSettingsInput() *GoogleCesAppLanguageSettings {
	var returns *GoogleCesAppLanguageSettings
	_jsii_.Get(
		j,
		"languageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) LoggingSettings() GoogleCesAppLoggingSettingsOutputReference {
	var returns GoogleCesAppLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"loggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) LoggingSettingsInput() *GoogleCesAppLoggingSettings {
	var returns *GoogleCesAppLoggingSettings
	_jsii_.Get(
		j,
		"loggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ModelSettings() GoogleCesAppModelSettingsOutputReference {
	var returns GoogleCesAppModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ModelSettingsInput() *GoogleCesAppModelSettings {
	var returns *GoogleCesAppModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Pinned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pinned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) PinnedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pinnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) RootAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) RootAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) Timeouts() GoogleCesAppTimeoutsOutputReference {
	var returns GoogleCesAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TimeZoneSettings() GoogleCesAppTimeZoneSettingsOutputReference {
	var returns GoogleCesAppTimeZoneSettingsOutputReference
	_jsii_.Get(
		j,
		"timeZoneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) TimeZoneSettingsInput() *GoogleCesAppTimeZoneSettings {
	var returns *GoogleCesAppTimeZoneSettings
	_jsii_.Get(
		j,
		"timeZoneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) VariableDeclarations() GoogleCesAppVariableDeclarationsList {
	var returns GoogleCesAppVariableDeclarationsList
	_jsii_.Get(
		j,
		"variableDeclarations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesApp) VariableDeclarationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variableDeclarationsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_app google_ces_app} Resource.
func NewGoogleCesApp(scope constructs.Construct, id *string, config *GoogleCesAppConfig) GoogleCesApp {
	_init_.Initialize()

	if err := validateNewGoogleCesAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesApp{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_ces_app google_ces_app} Resource.
func NewGoogleCesApp_Override(g GoogleCesApp, scope constructs.Construct, id *string, config *GoogleCesAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetGlobalInstruction(val *string) {
	if err := j.validateSetGlobalInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalInstruction",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetGuardrails(val *[]*string) {
	if err := j.validateSetGuardrailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrails",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetPinned(val interface{}) {
	if err := j.validateSetPinnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinned",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCesApp)SetRootAgent(val *string) {
	if err := j.validateSetRootAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootAgent",
		val,
	)
}

// Generates CDKTN code for importing a GoogleCesApp resource upon running "cdktn plan <stack-name>".
func GoogleCesApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCesApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
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
func GoogleCesApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCesApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCesApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCesApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCesApp) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCesApp) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCesApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesApp) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCesApp) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesApp) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCesApp) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCesApp) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutAudioProcessingConfig(value *GoogleCesAppAudioProcessingConfig) {
	if err := g.validatePutAudioProcessingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAudioProcessingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutClientCertificateSettings(value *GoogleCesAppClientCertificateSettings) {
	if err := g.validatePutClientCertificateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientCertificateSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutDataStoreSettings(value *GoogleCesAppDataStoreSettings) {
	if err := g.validatePutDataStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataStoreSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutDefaultChannelProfile(value *GoogleCesAppDefaultChannelProfile) {
	if err := g.validatePutDefaultChannelProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultChannelProfile",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutEvaluationMetricsThresholds(value *GoogleCesAppEvaluationMetricsThresholds) {
	if err := g.validatePutEvaluationMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEvaluationMetricsThresholds",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutLanguageSettings(value *GoogleCesAppLanguageSettings) {
	if err := g.validatePutLanguageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLanguageSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutLoggingSettings(value *GoogleCesAppLoggingSettings) {
	if err := g.validatePutLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutModelSettings(value *GoogleCesAppModelSettings) {
	if err := g.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putModelSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutTimeouts(value *GoogleCesAppTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutTimeZoneSettings(value *GoogleCesAppTimeZoneSettings) {
	if err := g.validatePutTimeZoneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeZoneSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) PutVariableDeclarations(value interface{}) {
	if err := g.validatePutVariableDeclarationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVariableDeclarations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCesApp) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetAudioProcessingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioProcessingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetClientCertificateSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertificateSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetDataStoreSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStoreSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetDefaultChannelProfile() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultChannelProfile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetEvaluationMetricsThresholds() {
	_jsii_.InvokeVoid(
		g,
		"resetEvaluationMetricsThresholds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetGlobalInstruction() {
	_jsii_.InvokeVoid(
		g,
		"resetGlobalInstruction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetGuardrails() {
	_jsii_.InvokeVoid(
		g,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetLanguageSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetLanguageSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetLoggingSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetModelSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetPinned() {
	_jsii_.InvokeVoid(
		g,
		"resetPinned",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetRootAgent() {
	_jsii_.InvokeVoid(
		g,
		"resetRootAgent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetTimeZoneSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZoneSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) ResetVariableDeclarations() {
	_jsii_.InvokeVoid(
		g,
		"resetVariableDeclarations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesApp) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

