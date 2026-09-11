package dialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/dialogflowcxagent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent}.
type DialogflowCxAgent interface {
	cdktn.TerraformResource
	AdvancedSettings() DialogflowCxAgentAdvancedSettingsOutputReference
	AdvancedSettingsInput() *DialogflowCxAgentAdvancedSettings
	AnswerFeedbackSettings() DialogflowCxAgentAnswerFeedbackSettingsOutputReference
	AnswerFeedbackSettingsInput() *DialogflowCxAgentAnswerFeedbackSettings
	AvatarUri() *string
	SetAvatarUri(val *string)
	AvatarUriInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCertificateSettings() DialogflowCxAgentClientCertificateSettingsOutputReference
	ClientCertificateSettingsInput() *DialogflowCxAgentClientCertificateSettings
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
	DefaultLanguageCode() *string
	SetDefaultLanguageCode(val *string)
	DefaultLanguageCodeInput() *string
	DeleteChatEngineOnDestroy() interface{}
	SetDeleteChatEngineOnDestroy(val interface{})
	DeleteChatEngineOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnableMultiLanguageTraining() interface{}
	SetEnableMultiLanguageTraining(val interface{})
	EnableMultiLanguageTrainingInput() interface{}
	EnableSpellCorrection() interface{}
	SetEnableSpellCorrection(val interface{})
	EnableSpellCorrectionInput() interface{}
	EnableStackdriverLogging() interface{}
	SetEnableStackdriverLogging(val interface{})
	EnableStackdriverLoggingInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GenAppBuilderSettings() DialogflowCxAgentGenAppBuilderSettingsOutputReference
	GenAppBuilderSettingsInput() *DialogflowCxAgentGenAppBuilderSettings
	GitIntegrationSettings() DialogflowCxAgentGitIntegrationSettingsOutputReference
	GitIntegrationSettingsInput() *DialogflowCxAgentGitIntegrationSettings
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Locked() interface{}
	SetLocked(val interface{})
	LockedInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	PersonalizationSettings() DialogflowCxAgentPersonalizationSettingsOutputReference
	PersonalizationSettingsInput() *DialogflowCxAgentPersonalizationSettings
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
	SatisfiesPzi() cdktn.IResolvable
	SatisfiesPzs() cdktn.IResolvable
	SecuritySettings() *string
	SetSecuritySettings(val *string)
	SecuritySettingsInput() *string
	SpeechToTextSettings() DialogflowCxAgentSpeechToTextSettingsOutputReference
	SpeechToTextSettingsInput() *DialogflowCxAgentSpeechToTextSettings
	StartFlow() *string
	StartPlaybook() *string
	SetStartPlaybook(val *string)
	StartPlaybookInput() *string
	SupportedLanguageCodes() *[]*string
	SetSupportedLanguageCodes(val *[]*string)
	SupportedLanguageCodesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TextToSpeechSettings() DialogflowCxAgentTextToSpeechSettingsOutputReference
	TextToSpeechSettingsInput() *DialogflowCxAgentTextToSpeechSettings
	Timeouts() DialogflowCxAgentTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutAdvancedSettings(value *DialogflowCxAgentAdvancedSettings)
	PutAnswerFeedbackSettings(value *DialogflowCxAgentAnswerFeedbackSettings)
	PutClientCertificateSettings(value *DialogflowCxAgentClientCertificateSettings)
	PutGenAppBuilderSettings(value *DialogflowCxAgentGenAppBuilderSettings)
	PutGitIntegrationSettings(value *DialogflowCxAgentGitIntegrationSettings)
	PutPersonalizationSettings(value *DialogflowCxAgentPersonalizationSettings)
	PutSpeechToTextSettings(value *DialogflowCxAgentSpeechToTextSettings)
	PutTextToSpeechSettings(value *DialogflowCxAgentTextToSpeechSettings)
	PutTimeouts(value *DialogflowCxAgentTimeouts)
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
	ResetAdvancedSettings()
	ResetAnswerFeedbackSettings()
	ResetAvatarUri()
	ResetClientCertificateSettings()
	ResetDeleteChatEngineOnDestroy()
	ResetDescription()
	ResetEnableMultiLanguageTraining()
	ResetEnableSpellCorrection()
	ResetEnableStackdriverLogging()
	ResetGenAppBuilderSettings()
	ResetGitIntegrationSettings()
	ResetId()
	ResetLocked()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersonalizationSettings()
	ResetProject()
	ResetSecuritySettings()
	ResetSpeechToTextSettings()
	ResetStartPlaybook()
	ResetSupportedLanguageCodes()
	ResetTextToSpeechSettings()
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

// The jsii proxy struct for DialogflowCxAgent
type jsiiProxy_DialogflowCxAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DialogflowCxAgent) AdvancedSettings() DialogflowCxAgentAdvancedSettingsOutputReference {
	var returns DialogflowCxAgentAdvancedSettingsOutputReference
	_jsii_.Get(
		j,
		"advancedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AdvancedSettingsInput() *DialogflowCxAgentAdvancedSettings {
	var returns *DialogflowCxAgentAdvancedSettings
	_jsii_.Get(
		j,
		"advancedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AnswerFeedbackSettings() DialogflowCxAgentAnswerFeedbackSettingsOutputReference {
	var returns DialogflowCxAgentAnswerFeedbackSettingsOutputReference
	_jsii_.Get(
		j,
		"answerFeedbackSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AnswerFeedbackSettingsInput() *DialogflowCxAgentAnswerFeedbackSettings {
	var returns *DialogflowCxAgentAnswerFeedbackSettings
	_jsii_.Get(
		j,
		"answerFeedbackSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AvatarUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) AvatarUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ClientCertificateSettings() DialogflowCxAgentClientCertificateSettingsOutputReference {
	var returns DialogflowCxAgentClientCertificateSettingsOutputReference
	_jsii_.Get(
		j,
		"clientCertificateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ClientCertificateSettingsInput() *DialogflowCxAgentClientCertificateSettings {
	var returns *DialogflowCxAgentClientCertificateSettings
	_jsii_.Get(
		j,
		"clientCertificateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DefaultLanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DefaultLanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DeleteChatEngineOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteChatEngineOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DeleteChatEngineOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteChatEngineOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableMultiLanguageTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiLanguageTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableMultiLanguageTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiLanguageTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableSpellCorrection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableSpellCorrectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableStackdriverLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) EnableStackdriverLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) GenAppBuilderSettings() DialogflowCxAgentGenAppBuilderSettingsOutputReference {
	var returns DialogflowCxAgentGenAppBuilderSettingsOutputReference
	_jsii_.Get(
		j,
		"genAppBuilderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) GenAppBuilderSettingsInput() *DialogflowCxAgentGenAppBuilderSettings {
	var returns *DialogflowCxAgentGenAppBuilderSettings
	_jsii_.Get(
		j,
		"genAppBuilderSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) GitIntegrationSettings() DialogflowCxAgentGitIntegrationSettingsOutputReference {
	var returns DialogflowCxAgentGitIntegrationSettingsOutputReference
	_jsii_.Get(
		j,
		"gitIntegrationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) GitIntegrationSettingsInput() *DialogflowCxAgentGitIntegrationSettings {
	var returns *DialogflowCxAgentGitIntegrationSettings
	_jsii_.Get(
		j,
		"gitIntegrationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Locked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) LockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) PersonalizationSettings() DialogflowCxAgentPersonalizationSettingsOutputReference {
	var returns DialogflowCxAgentPersonalizationSettingsOutputReference
	_jsii_.Get(
		j,
		"personalizationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) PersonalizationSettingsInput() *DialogflowCxAgentPersonalizationSettings {
	var returns *DialogflowCxAgentPersonalizationSettings
	_jsii_.Get(
		j,
		"personalizationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SatisfiesPzi() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SatisfiesPzs() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SecuritySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SecuritySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SpeechToTextSettings() DialogflowCxAgentSpeechToTextSettingsOutputReference {
	var returns DialogflowCxAgentSpeechToTextSettingsOutputReference
	_jsii_.Get(
		j,
		"speechToTextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SpeechToTextSettingsInput() *DialogflowCxAgentSpeechToTextSettings {
	var returns *DialogflowCxAgentSpeechToTextSettings
	_jsii_.Get(
		j,
		"speechToTextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) StartFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) StartPlaybook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startPlaybook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) StartPlaybookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startPlaybookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SupportedLanguageCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) SupportedLanguageCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TextToSpeechSettings() DialogflowCxAgentTextToSpeechSettingsOutputReference {
	var returns DialogflowCxAgentTextToSpeechSettingsOutputReference
	_jsii_.Get(
		j,
		"textToSpeechSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TextToSpeechSettingsInput() *DialogflowCxAgentTextToSpeechSettings {
	var returns *DialogflowCxAgentTextToSpeechSettings
	_jsii_.Get(
		j,
		"textToSpeechSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) Timeouts() DialogflowCxAgentTimeoutsOutputReference {
	var returns DialogflowCxAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgent) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewDialogflowCxAgent(scope constructs.Construct, id *string, config *DialogflowCxAgentConfig) DialogflowCxAgent {
	_init_.Initialize()

	if err := validateNewDialogflowCxAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxAgent{}

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewDialogflowCxAgent_Override(d DialogflowCxAgent, scope constructs.Construct, id *string, config *DialogflowCxAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetAvatarUri(val *string) {
	if err := j.validateSetAvatarUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatarUri",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDefaultLanguageCode(val *string) {
	if err := j.validateSetDefaultLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLanguageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDeleteChatEngineOnDestroy(val interface{}) {
	if err := j.validateSetDeleteChatEngineOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteChatEngineOnDestroy",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetEnableMultiLanguageTraining(val interface{}) {
	if err := j.validateSetEnableMultiLanguageTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultiLanguageTraining",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetEnableSpellCorrection(val interface{}) {
	if err := j.validateSetEnableSpellCorrectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSpellCorrection",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetEnableStackdriverLogging(val interface{}) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetLocked(val interface{}) {
	if err := j.validateSetLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locked",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetSecuritySettings(val *string) {
	if err := j.validateSetSecuritySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securitySettings",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetStartPlaybook(val *string) {
	if err := j.validateSetStartPlaybookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startPlaybook",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetSupportedLanguageCodes(val *[]*string) {
	if err := j.validateSetSupportedLanguageCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedLanguageCodes",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgent)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Generates CDKTN code for importing a DialogflowCxAgent resource upon running "cdktn plan <stack-name>".
func DialogflowCxAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
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
func DialogflowCxAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DialogflowCxAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDialogflowCxAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DialogflowCxAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google.dialogflowCxAgent.DialogflowCxAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutAdvancedSettings(value *DialogflowCxAgentAdvancedSettings) {
	if err := d.validatePutAdvancedSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAdvancedSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutAnswerFeedbackSettings(value *DialogflowCxAgentAnswerFeedbackSettings) {
	if err := d.validatePutAnswerFeedbackSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAnswerFeedbackSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutClientCertificateSettings(value *DialogflowCxAgentClientCertificateSettings) {
	if err := d.validatePutClientCertificateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putClientCertificateSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutGenAppBuilderSettings(value *DialogflowCxAgentGenAppBuilderSettings) {
	if err := d.validatePutGenAppBuilderSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGenAppBuilderSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutGitIntegrationSettings(value *DialogflowCxAgentGitIntegrationSettings) {
	if err := d.validatePutGitIntegrationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitIntegrationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutPersonalizationSettings(value *DialogflowCxAgentPersonalizationSettings) {
	if err := d.validatePutPersonalizationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersonalizationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutSpeechToTextSettings(value *DialogflowCxAgentSpeechToTextSettings) {
	if err := d.validatePutSpeechToTextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSpeechToTextSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutTextToSpeechSettings(value *DialogflowCxAgentTextToSpeechSettings) {
	if err := d.validatePutTextToSpeechSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTextToSpeechSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) PutTimeouts(value *DialogflowCxAgentTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetAdvancedSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAdvancedSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetAnswerFeedbackSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetAnswerFeedbackSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetAvatarUri() {
	_jsii_.InvokeVoid(
		d,
		"resetAvatarUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetClientCertificateSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetClientCertificateSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetDeleteChatEngineOnDestroy() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteChatEngineOnDestroy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetEnableMultiLanguageTraining() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableMultiLanguageTraining",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetEnableSpellCorrection() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSpellCorrection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetGenAppBuilderSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetGenAppBuilderSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetGitIntegrationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetGitIntegrationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetLocked() {
	_jsii_.InvokeVoid(
		d,
		"resetLocked",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetPersonalizationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPersonalizationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSpeechToTextSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSpeechToTextSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetStartPlaybook() {
	_jsii_.InvokeVoid(
		d,
		"resetStartPlaybook",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetSupportedLanguageCodes() {
	_jsii_.InvokeVoid(
		d,
		"resetSupportedLanguageCodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetTextToSpeechSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetTextToSpeechSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

