package googledialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledialogflowconversationprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_conversation_profile google_dialogflow_conversation_profile}.
type GoogleDialogflowConversationProfile interface {
	cdktf.TerraformResource
	AutomatedAgentConfig() GoogleDialogflowConversationProfileAutomatedAgentConfigOutputReference
	AutomatedAgentConfigInput() *GoogleDialogflowConversationProfileAutomatedAgentConfig
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
	HumanAgentAssistantConfig() GoogleDialogflowConversationProfileHumanAgentAssistantConfigOutputReference
	HumanAgentAssistantConfigInput() *GoogleDialogflowConversationProfileHumanAgentAssistantConfig
	HumanAgentHandoffConfig() GoogleDialogflowConversationProfileHumanAgentHandoffConfigOutputReference
	HumanAgentHandoffConfigInput() *GoogleDialogflowConversationProfileHumanAgentHandoffConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() GoogleDialogflowConversationProfileLoggingConfigOutputReference
	LoggingConfigInput() *GoogleDialogflowConversationProfileLoggingConfig
	Name() *string
	NewMessageEventNotificationConfig() GoogleDialogflowConversationProfileNewMessageEventNotificationConfigOutputReference
	NewMessageEventNotificationConfigInput() *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig
	NewRecognitionResultNotificationConfig() GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfigOutputReference
	NewRecognitionResultNotificationConfigInput() *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig
	// The tree node.
	Node() constructs.Node
	NotificationConfig() GoogleDialogflowConversationProfileNotificationConfigOutputReference
	NotificationConfigInput() *GoogleDialogflowConversationProfileNotificationConfig
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
	SecuritySettings() *string
	SetSecuritySettings(val *string)
	SecuritySettingsInput() *string
	SttConfig() GoogleDialogflowConversationProfileSttConfigOutputReference
	SttConfigInput() *GoogleDialogflowConversationProfileSttConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDialogflowConversationProfileTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	TtsConfig() GoogleDialogflowConversationProfileTtsConfigOutputReference
	TtsConfigInput() *GoogleDialogflowConversationProfileTtsConfig
	UseBidiStreaming() interface{}
	SetUseBidiStreaming(val interface{})
	UseBidiStreamingInput() interface{}
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
	PutAutomatedAgentConfig(value *GoogleDialogflowConversationProfileAutomatedAgentConfig)
	PutHumanAgentAssistantConfig(value *GoogleDialogflowConversationProfileHumanAgentAssistantConfig)
	PutHumanAgentHandoffConfig(value *GoogleDialogflowConversationProfileHumanAgentHandoffConfig)
	PutLoggingConfig(value *GoogleDialogflowConversationProfileLoggingConfig)
	PutNewMessageEventNotificationConfig(value *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig)
	PutNewRecognitionResultNotificationConfig(value *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig)
	PutNotificationConfig(value *GoogleDialogflowConversationProfileNotificationConfig)
	PutSttConfig(value *GoogleDialogflowConversationProfileSttConfig)
	PutTimeouts(value *GoogleDialogflowConversationProfileTimeouts)
	PutTtsConfig(value *GoogleDialogflowConversationProfileTtsConfig)
	ResetAutomatedAgentConfig()
	ResetHumanAgentAssistantConfig()
	ResetHumanAgentHandoffConfig()
	ResetId()
	ResetLanguageCode()
	ResetLoggingConfig()
	ResetNewMessageEventNotificationConfig()
	ResetNewRecognitionResultNotificationConfig()
	ResetNotificationConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSecuritySettings()
	ResetSttConfig()
	ResetTimeouts()
	ResetTimeZone()
	ResetTtsConfig()
	ResetUseBidiStreaming()
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

// The jsii proxy struct for GoogleDialogflowConversationProfile
type jsiiProxy_GoogleDialogflowConversationProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) AutomatedAgentConfig() GoogleDialogflowConversationProfileAutomatedAgentConfigOutputReference {
	var returns GoogleDialogflowConversationProfileAutomatedAgentConfigOutputReference
	_jsii_.Get(
		j,
		"automatedAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) AutomatedAgentConfigInput() *GoogleDialogflowConversationProfileAutomatedAgentConfig {
	var returns *GoogleDialogflowConversationProfileAutomatedAgentConfig
	_jsii_.Get(
		j,
		"automatedAgentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) HumanAgentAssistantConfig() GoogleDialogflowConversationProfileHumanAgentAssistantConfigOutputReference {
	var returns GoogleDialogflowConversationProfileHumanAgentAssistantConfigOutputReference
	_jsii_.Get(
		j,
		"humanAgentAssistantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) HumanAgentAssistantConfigInput() *GoogleDialogflowConversationProfileHumanAgentAssistantConfig {
	var returns *GoogleDialogflowConversationProfileHumanAgentAssistantConfig
	_jsii_.Get(
		j,
		"humanAgentAssistantConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) HumanAgentHandoffConfig() GoogleDialogflowConversationProfileHumanAgentHandoffConfigOutputReference {
	var returns GoogleDialogflowConversationProfileHumanAgentHandoffConfigOutputReference
	_jsii_.Get(
		j,
		"humanAgentHandoffConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) HumanAgentHandoffConfigInput() *GoogleDialogflowConversationProfileHumanAgentHandoffConfig {
	var returns *GoogleDialogflowConversationProfileHumanAgentHandoffConfig
	_jsii_.Get(
		j,
		"humanAgentHandoffConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) LoggingConfig() GoogleDialogflowConversationProfileLoggingConfigOutputReference {
	var returns GoogleDialogflowConversationProfileLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) LoggingConfigInput() *GoogleDialogflowConversationProfileLoggingConfig {
	var returns *GoogleDialogflowConversationProfileLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NewMessageEventNotificationConfig() GoogleDialogflowConversationProfileNewMessageEventNotificationConfigOutputReference {
	var returns GoogleDialogflowConversationProfileNewMessageEventNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"newMessageEventNotificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NewMessageEventNotificationConfigInput() *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig {
	var returns *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig
	_jsii_.Get(
		j,
		"newMessageEventNotificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NewRecognitionResultNotificationConfig() GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfigOutputReference {
	var returns GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"newRecognitionResultNotificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NewRecognitionResultNotificationConfigInput() *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig {
	var returns *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig
	_jsii_.Get(
		j,
		"newRecognitionResultNotificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NotificationConfig() GoogleDialogflowConversationProfileNotificationConfigOutputReference {
	var returns GoogleDialogflowConversationProfileNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) NotificationConfigInput() *GoogleDialogflowConversationProfileNotificationConfig {
	var returns *GoogleDialogflowConversationProfileNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) SecuritySettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) SecuritySettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) SttConfig() GoogleDialogflowConversationProfileSttConfigOutputReference {
	var returns GoogleDialogflowConversationProfileSttConfigOutputReference
	_jsii_.Get(
		j,
		"sttConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) SttConfigInput() *GoogleDialogflowConversationProfileSttConfig {
	var returns *GoogleDialogflowConversationProfileSttConfig
	_jsii_.Get(
		j,
		"sttConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) Timeouts() GoogleDialogflowConversationProfileTimeoutsOutputReference {
	var returns GoogleDialogflowConversationProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TtsConfig() GoogleDialogflowConversationProfileTtsConfigOutputReference {
	var returns GoogleDialogflowConversationProfileTtsConfigOutputReference
	_jsii_.Get(
		j,
		"ttsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) TtsConfigInput() *GoogleDialogflowConversationProfileTtsConfig {
	var returns *GoogleDialogflowConversationProfileTtsConfig
	_jsii_.Get(
		j,
		"ttsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) UseBidiStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBidiStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile) UseBidiStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBidiStreamingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_conversation_profile google_dialogflow_conversation_profile} Resource.
func NewGoogleDialogflowConversationProfile(scope constructs.Construct, id *string, config *GoogleDialogflowConversationProfileConfig) GoogleDialogflowConversationProfile {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowConversationProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowConversationProfile{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/resources/google_dialogflow_conversation_profile google_dialogflow_conversation_profile} Resource.
func NewGoogleDialogflowConversationProfile_Override(g GoogleDialogflowConversationProfile, scope constructs.Construct, id *string, config *GoogleDialogflowConversationProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetSecuritySettings(val *string) {
	if err := j.validateSetSecuritySettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securitySettings",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowConversationProfile)SetUseBidiStreaming(val interface{}) {
	if err := j.validateSetUseBidiStreamingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBidiStreaming",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDialogflowConversationProfile resource upon running "cdktf plan <stack-name>".
func GoogleDialogflowConversationProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDialogflowConversationProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
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
func GoogleDialogflowConversationProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowConversationProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowConversationProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowConversationProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDialogflowConversationProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDialogflowConversationProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDialogflowConversationProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDialogflowConversationProfile.GoogleDialogflowConversationProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowConversationProfile) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutAutomatedAgentConfig(value *GoogleDialogflowConversationProfileAutomatedAgentConfig) {
	if err := g.validatePutAutomatedAgentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutomatedAgentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutHumanAgentAssistantConfig(value *GoogleDialogflowConversationProfileHumanAgentAssistantConfig) {
	if err := g.validatePutHumanAgentAssistantConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHumanAgentAssistantConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutHumanAgentHandoffConfig(value *GoogleDialogflowConversationProfileHumanAgentHandoffConfig) {
	if err := g.validatePutHumanAgentHandoffConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHumanAgentHandoffConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutLoggingConfig(value *GoogleDialogflowConversationProfileLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutNewMessageEventNotificationConfig(value *GoogleDialogflowConversationProfileNewMessageEventNotificationConfig) {
	if err := g.validatePutNewMessageEventNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewMessageEventNotificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutNewRecognitionResultNotificationConfig(value *GoogleDialogflowConversationProfileNewRecognitionResultNotificationConfig) {
	if err := g.validatePutNewRecognitionResultNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewRecognitionResultNotificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutNotificationConfig(value *GoogleDialogflowConversationProfileNotificationConfig) {
	if err := g.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutSttConfig(value *GoogleDialogflowConversationProfileSttConfig) {
	if err := g.validatePutSttConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSttConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutTimeouts(value *GoogleDialogflowConversationProfileTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) PutTtsConfig(value *GoogleDialogflowConversationProfileTtsConfig) {
	if err := g.validatePutTtsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTtsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetAutomatedAgentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutomatedAgentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetHumanAgentAssistantConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHumanAgentAssistantConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetHumanAgentHandoffConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHumanAgentHandoffConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		g,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetNewMessageEventNotificationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNewMessageEventNotificationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetNewRecognitionResultNotificationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNewRecognitionResultNotificationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetSttConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSttConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetTtsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTtsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ResetUseBidiStreaming() {
	_jsii_.InvokeVoid(
		g,
		"resetUseBidiStreaming",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowConversationProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

