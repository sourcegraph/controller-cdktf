package googledialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowcxagent/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_cx_agent google_dialogflow_cx_agent}.
type GoogleDialogflowCxAgent interface {
	cdktn.TerraformResource
	AdvancedSettings() GoogleDialogflowCxAgentAdvancedSettingsOutputReference
	AdvancedSettingsInput() *GoogleDialogflowCxAgentAdvancedSettings
	AnswerFeedbackSettings() GoogleDialogflowCxAgentAnswerFeedbackSettingsOutputReference
	AnswerFeedbackSettingsInput() *GoogleDialogflowCxAgentAnswerFeedbackSettings
	AvatarUri() *string
	SetAvatarUri(val *string)
	AvatarUriInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCertificateSettings() GoogleDialogflowCxAgentClientCertificateSettingsOutputReference
	ClientCertificateSettingsInput() *GoogleDialogflowCxAgentClientCertificateSettings
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
	GenAppBuilderSettings() GoogleDialogflowCxAgentGenAppBuilderSettingsOutputReference
	GenAppBuilderSettingsInput() *GoogleDialogflowCxAgentGenAppBuilderSettings
	GitIntegrationSettings() GoogleDialogflowCxAgentGitIntegrationSettingsOutputReference
	GitIntegrationSettingsInput() *GoogleDialogflowCxAgentGitIntegrationSettings
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
	PersonalizationSettings() GoogleDialogflowCxAgentPersonalizationSettingsOutputReference
	PersonalizationSettingsInput() *GoogleDialogflowCxAgentPersonalizationSettings
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
	SpeechToTextSettings() GoogleDialogflowCxAgentSpeechToTextSettingsOutputReference
	SpeechToTextSettingsInput() *GoogleDialogflowCxAgentSpeechToTextSettings
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
	TextToSpeechSettings() GoogleDialogflowCxAgentTextToSpeechSettingsOutputReference
	TextToSpeechSettingsInput() *GoogleDialogflowCxAgentTextToSpeechSettings
	Timeouts() GoogleDialogflowCxAgentTimeoutsOutputReference
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
	PutAdvancedSettings(value *GoogleDialogflowCxAgentAdvancedSettings)
	PutAnswerFeedbackSettings(value *GoogleDialogflowCxAgentAnswerFeedbackSettings)
	PutClientCertificateSettings(value *GoogleDialogflowCxAgentClientCertificateSettings)
	PutGenAppBuilderSettings(value *GoogleDialogflowCxAgentGenAppBuilderSettings)
	PutGitIntegrationSettings(value *GoogleDialogflowCxAgentGitIntegrationSettings)
	PutPersonalizationSettings(value *GoogleDialogflowCxAgentPersonalizationSettings)
	PutSpeechToTextSettings(value *GoogleDialogflowCxAgentSpeechToTextSettings)
	PutTextToSpeechSettings(value *GoogleDialogflowCxAgentTextToSpeechSettings)
	PutTimeouts(value *GoogleDialogflowCxAgentTimeouts)
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

// The jsii proxy struct for GoogleDialogflowCxAgent
type jsiiProxy_GoogleDialogflowCxAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AdvancedSettings() GoogleDialogflowCxAgentAdvancedSettingsOutputReference {
	var returns GoogleDialogflowCxAgentAdvancedSettingsOutputReference
	_jsii_.Get(
		j,
		"advancedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AdvancedSettingsInput() *GoogleDialogflowCxAgentAdvancedSettings {
	var returns *GoogleDialogflowCxAgentAdvancedSettings
	_jsii_.Get(
		j,
		"advancedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AnswerFeedbackSettings() GoogleDialogflowCxAgentAnswerFeedbackSettingsOutputReference {
	var returns GoogleDialogflowCxAgentAnswerFeedbackSettingsOutputReference
	_jsii_.Get(
		j,
		"answerFeedbackSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AnswerFeedbackSettingsInput() *GoogleDialogflowCxAgentAnswerFeedbackSettings {
	var returns *GoogleDialogflowCxAgentAnswerFeedbackSettings
	_jsii_.Get(
		j,
		"answerFeedbackSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AvatarUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) AvatarUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"avatarUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) ClientCertificateSettings() GoogleDialogflowCxAgentClientCertificateSettingsOutputReference {
	var returns GoogleDialogflowCxAgentClientCertificateSettingsOutputReference
	_jsii_.Get(
		j,
		"clientCertificateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) ClientCertificateSettingsInput() *GoogleDialogflowCxAgentClientCertificateSettings {
	var returns *GoogleDialogflowCxAgentClientCertificateSettings
	_jsii_.Get(
		j,
		"clientCertificateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DefaultLanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DefaultLanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLanguageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DeleteChatEngineOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteChatEngineOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DeleteChatEngineOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteChatEngineOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableMultiLanguageTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiLanguageTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableMultiLanguageTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMultiLanguageTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableSpellCorrection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableSpellCorrectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSpellCorrectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableStackdriverLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) EnableStackdriverLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) GenAppBuilderSettings() GoogleDialogflowCxAgentGenAppBuilderSettingsOutputReference {
	var returns GoogleDialogflowCxAgentGenAppBuilderSettingsOutputReference
	_jsii_.Get(
		j,
		"genAppBuilderSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) GenAppBuilderSettingsInput() *GoogleDialogflowCxAgentGenAppBuilderSettings {
	var returns *GoogleDialogflowCxAgentGenAppBuilderSettings
	_jsii_.Get(
		j,
		"genAppBuilderSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) GitIntegrationSettings() GoogleDialogflowCxAgentGitIntegrationSettingsOutputReference {
	var returns GoogleDialogflowCxAgentGitIntegrationSettingsOutputReference
	_jsii_.Get(
		j,
		"gitIntegrationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) GitIntegrationSettingsInput() *GoogleDialogflowCxAgentGitIntegrationSettings {
	var returns *GoogleDialogflowCxAgentGitIntegrationSettings
	_jsii_.Get(
		j,
		"gitIntegrationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Locked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) LockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) PersonalizationSettings() GoogleDialogflowCxAgentPersonalizationSettingsOutputReference {
	var returns GoogleDialogflowCxAgentPersonalizationSettingsOutputReference
	_jsii_.Get(
		j,
		"personalizationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) PersonalizationSettingsInput() *GoogleDialogflowCxAgentPersonalizationSettings {
	var returns *GoogleDialogflowCxAgentPersonalizationSettings
	_jsii_.Get(
		j,
		"personalizationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SatisfiesPzi() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SatisfiesPzs() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SecuritySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SecuritySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SpeechToTextSettings() GoogleDialogflowCxAgentSpeechToTextSettingsOutputReference {
	var returns GoogleDialogflowCxAgentSpeechToTextSettingsOutputReference
	_jsii_.Get(
		j,
		"speechToTextSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SpeechToTextSettingsInput() *GoogleDialogflowCxAgentSpeechToTextSettings {
	var returns *GoogleDialogflowCxAgentSpeechToTextSettings
	_jsii_.Get(
		j,
		"speechToTextSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) StartFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) StartPlaybook() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startPlaybook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) StartPlaybookInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startPlaybookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SupportedLanguageCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) SupportedLanguageCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedLanguageCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TextToSpeechSettings() GoogleDialogflowCxAgentTextToSpeechSettingsOutputReference {
	var returns GoogleDialogflowCxAgentTextToSpeechSettingsOutputReference
	_jsii_.Get(
		j,
		"textToSpeechSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TextToSpeechSettingsInput() *GoogleDialogflowCxAgentTextToSpeechSettings {
	var returns *GoogleDialogflowCxAgentTextToSpeechSettings
	_jsii_.Get(
		j,
		"textToSpeechSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) Timeouts() GoogleDialogflowCxAgentTimeoutsOutputReference {
	var returns GoogleDialogflowCxAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxAgent) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewGoogleDialogflowCxAgent(scope constructs.Construct, id *string, config *GoogleDialogflowCxAgentConfig) GoogleDialogflowCxAgent {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxAgent{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_cx_agent google_dialogflow_cx_agent} Resource.
func NewGoogleDialogflowCxAgent_Override(g GoogleDialogflowCxAgent, scope constructs.Construct, id *string, config *GoogleDialogflowCxAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetAvatarUri(val *string) {
	if err := j.validateSetAvatarUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avatarUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetDefaultLanguageCode(val *string) {
	if err := j.validateSetDefaultLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLanguageCode",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetDeleteChatEngineOnDestroy(val interface{}) {
	if err := j.validateSetDeleteChatEngineOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteChatEngineOnDestroy",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetEnableMultiLanguageTraining(val interface{}) {
	if err := j.validateSetEnableMultiLanguageTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMultiLanguageTraining",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetEnableSpellCorrection(val interface{}) {
	if err := j.validateSetEnableSpellCorrectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSpellCorrection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetEnableStackdriverLogging(val interface{}) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetLocked(val interface{}) {
	if err := j.validateSetLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locked",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetSecuritySettings(val *string) {
	if err := j.validateSetSecuritySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securitySettings",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetStartPlaybook(val *string) {
	if err := j.validateSetStartPlaybookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startPlaybook",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetSupportedLanguageCodes(val *[]*string) {
	if err := j.validateSetSupportedLanguageCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedLanguageCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxAgent)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Generates CDKTN code for importing a GoogleDialogflowCxAgent resource upon running "cdktn plan <stack-name>".
func GoogleDialogflowCxAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
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
func GoogleDialogflowCxAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowCxAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowCxAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowCxAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDialogflowCxAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-google-beta.googleDialogflowCxAgent.GoogleDialogflowCxAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxAgent) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutAdvancedSettings(value *GoogleDialogflowCxAgentAdvancedSettings) {
	if err := g.validatePutAdvancedSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutAnswerFeedbackSettings(value *GoogleDialogflowCxAgentAnswerFeedbackSettings) {
	if err := g.validatePutAnswerFeedbackSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAnswerFeedbackSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutClientCertificateSettings(value *GoogleDialogflowCxAgentClientCertificateSettings) {
	if err := g.validatePutClientCertificateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientCertificateSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutGenAppBuilderSettings(value *GoogleDialogflowCxAgentGenAppBuilderSettings) {
	if err := g.validatePutGenAppBuilderSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGenAppBuilderSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutGitIntegrationSettings(value *GoogleDialogflowCxAgentGitIntegrationSettings) {
	if err := g.validatePutGitIntegrationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitIntegrationSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutPersonalizationSettings(value *GoogleDialogflowCxAgentPersonalizationSettings) {
	if err := g.validatePutPersonalizationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPersonalizationSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutSpeechToTextSettings(value *GoogleDialogflowCxAgentSpeechToTextSettings) {
	if err := g.validatePutSpeechToTextSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSpeechToTextSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutTextToSpeechSettings(value *GoogleDialogflowCxAgentTextToSpeechSettings) {
	if err := g.validatePutTextToSpeechSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextToSpeechSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) PutTimeouts(value *GoogleDialogflowCxAgentTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetAdvancedSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetAnswerFeedbackSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAnswerFeedbackSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetAvatarUri() {
	_jsii_.InvokeVoid(
		g,
		"resetAvatarUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetClientCertificateSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetClientCertificateSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetDeleteChatEngineOnDestroy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteChatEngineOnDestroy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetEnableMultiLanguageTraining() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableMultiLanguageTraining",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetEnableSpellCorrection() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableSpellCorrection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetGenAppBuilderSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGenAppBuilderSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetGitIntegrationSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetGitIntegrationSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetLocked() {
	_jsii_.InvokeVoid(
		g,
		"resetLocked",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetPersonalizationSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetPersonalizationSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetSpeechToTextSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSpeechToTextSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetStartPlaybook() {
	_jsii_.InvokeVoid(
		g,
		"resetStartPlaybook",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetSupportedLanguageCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetSupportedLanguageCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetTextToSpeechSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetTextToSpeechSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

