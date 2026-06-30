package cesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cesapp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_app google_ces_app}.
type CesApp interface {
	cdktf.TerraformResource
	AppId() *string
	SetAppId(val *string)
	AppIdInput() *string
	AudioProcessingConfig() CesAppAudioProcessingConfigOutputReference
	AudioProcessingConfigInput() *CesAppAudioProcessingConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificateSettings() CesAppClientCertificateSettingsOutputReference
	ClientCertificateSettingsInput() *CesAppClientCertificateSettings
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
	CreateTime() *string
	DataStoreSettings() CesAppDataStoreSettingsOutputReference
	DataStoreSettingsInput() *CesAppDataStoreSettings
	DefaultChannelProfile() CesAppDefaultChannelProfileOutputReference
	DefaultChannelProfileInput() *CesAppDefaultChannelProfile
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
	EvaluationMetricsThresholds() CesAppEvaluationMetricsThresholdsOutputReference
	EvaluationMetricsThresholdsInput() *CesAppEvaluationMetricsThresholds
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
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
	LanguageSettings() CesAppLanguageSettingsOutputReference
	LanguageSettingsInput() *CesAppLanguageSettings
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingSettings() CesAppLoggingSettingsOutputReference
	LoggingSettingsInput() *CesAppLoggingSettings
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	ModelSettings() CesAppModelSettingsOutputReference
	ModelSettingsInput() *CesAppModelSettings
	Name() *string
	// The tree node.
	Node() constructs.Node
	Pinned() any
	SetPinned(val any)
	PinnedInput() any
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
	RootAgent() *string
	SetRootAgent(val *string)
	RootAgentInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CesAppTimeoutsOutputReference
	TimeoutsInput() any
	TimeZoneSettings() CesAppTimeZoneSettingsOutputReference
	TimeZoneSettingsInput() *CesAppTimeZoneSettings
	UpdateTime() *string
	VariableDeclarations() CesAppVariableDeclarationsList
	VariableDeclarationsInput() any
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
	PutAudioProcessingConfig(value *CesAppAudioProcessingConfig)
	PutClientCertificateSettings(value *CesAppClientCertificateSettings)
	PutDataStoreSettings(value *CesAppDataStoreSettings)
	PutDefaultChannelProfile(value *CesAppDefaultChannelProfile)
	PutEvaluationMetricsThresholds(value *CesAppEvaluationMetricsThresholds)
	PutLanguageSettings(value *CesAppLanguageSettings)
	PutLoggingSettings(value *CesAppLoggingSettings)
	PutModelSettings(value *CesAppModelSettings)
	PutTimeouts(value *CesAppTimeouts)
	PutTimeZoneSettings(value *CesAppTimeZoneSettings)
	PutVariableDeclarations(value any)
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

// The jsii proxy struct for CesApp
type jsiiProxy_CesApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CesApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) AppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) AudioProcessingConfig() CesAppAudioProcessingConfigOutputReference {
	var returns CesAppAudioProcessingConfigOutputReference
	_jsii_.Get(
		j,
		"audioProcessingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) AudioProcessingConfigInput() *CesAppAudioProcessingConfig {
	var returns *CesAppAudioProcessingConfig
	_jsii_.Get(
		j,
		"audioProcessingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ClientCertificateSettings() CesAppClientCertificateSettingsOutputReference {
	var returns CesAppClientCertificateSettingsOutputReference
	_jsii_.Get(
		j,
		"clientCertificateSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ClientCertificateSettingsInput() *CesAppClientCertificateSettings {
	var returns *CesAppClientCertificateSettings
	_jsii_.Get(
		j,
		"clientCertificateSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DataStoreSettings() CesAppDataStoreSettingsOutputReference {
	var returns CesAppDataStoreSettingsOutputReference
	_jsii_.Get(
		j,
		"dataStoreSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DataStoreSettingsInput() *CesAppDataStoreSettings {
	var returns *CesAppDataStoreSettings
	_jsii_.Get(
		j,
		"dataStoreSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DefaultChannelProfile() CesAppDefaultChannelProfileOutputReference {
	var returns CesAppDefaultChannelProfileOutputReference
	_jsii_.Get(
		j,
		"defaultChannelProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DefaultChannelProfileInput() *CesAppDefaultChannelProfile {
	var returns *CesAppDefaultChannelProfile
	_jsii_.Get(
		j,
		"defaultChannelProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DeploymentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) EvaluationMetricsThresholds() CesAppEvaluationMetricsThresholdsOutputReference {
	var returns CesAppEvaluationMetricsThresholdsOutputReference
	_jsii_.Get(
		j,
		"evaluationMetricsThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) EvaluationMetricsThresholdsInput() *CesAppEvaluationMetricsThresholds {
	var returns *CesAppEvaluationMetricsThresholds
	_jsii_.Get(
		j,
		"evaluationMetricsThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) GlobalInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) GlobalInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Guardrails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) GuardrailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guardrailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) LanguageSettings() CesAppLanguageSettingsOutputReference {
	var returns CesAppLanguageSettingsOutputReference
	_jsii_.Get(
		j,
		"languageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) LanguageSettingsInput() *CesAppLanguageSettings {
	var returns *CesAppLanguageSettings
	_jsii_.Get(
		j,
		"languageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) LoggingSettings() CesAppLoggingSettingsOutputReference {
	var returns CesAppLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"loggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) LoggingSettingsInput() *CesAppLoggingSettings {
	var returns *CesAppLoggingSettings
	_jsii_.Get(
		j,
		"loggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ModelSettings() CesAppModelSettingsOutputReference {
	var returns CesAppModelSettingsOutputReference
	_jsii_.Get(
		j,
		"modelSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ModelSettingsInput() *CesAppModelSettings {
	var returns *CesAppModelSettings
	_jsii_.Get(
		j,
		"modelSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Pinned() any {
	var returns any
	_jsii_.Get(
		j,
		"pinned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) PinnedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pinnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) RootAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) RootAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) Timeouts() CesAppTimeoutsOutputReference {
	var returns CesAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TimeZoneSettings() CesAppTimeZoneSettingsOutputReference {
	var returns CesAppTimeZoneSettingsOutputReference
	_jsii_.Get(
		j,
		"timeZoneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) TimeZoneSettingsInput() *CesAppTimeZoneSettings {
	var returns *CesAppTimeZoneSettings
	_jsii_.Get(
		j,
		"timeZoneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) VariableDeclarations() CesAppVariableDeclarationsList {
	var returns CesAppVariableDeclarationsList
	_jsii_.Get(
		j,
		"variableDeclarations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesApp) VariableDeclarationsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"variableDeclarationsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_app google_ces_app} Resource.
func NewCesApp(scope constructs.Construct, id *string, config *CesAppConfig) CesApp {
	_init_.Initialize()

	if err := validateNewCesAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesApp{}

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesApp",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/ces_app google_ces_app} Resource.
func NewCesApp_Override(c CesApp, scope constructs.Construct, id *string, config *CesAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesApp.CesApp",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CesApp) SetAppId(val *string) {
	if err := j.validateSetAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetGlobalInstruction(val *string) {
	if err := j.validateSetGlobalInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalInstruction",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetGuardrails(val *[]*string) {
	if err := j.validateSetGuardrailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrails",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetPinned(val any) {
	if err := j.validateSetPinnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pinned",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CesApp) SetRootAgent(val *string) {
	if err := j.validateSetRootAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootAgent",
		val,
	)
}

// Generates CDKTF code for importing a CesApp resource upon running "cdktf plan <stack-name>".
func CesApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCesApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesApp.CesApp",
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
func CesApp_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateCesApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesApp.CesApp",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CesApp_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateCesApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesApp.CesApp",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CesApp_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateCesApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cesApp.CesApp",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func CesApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cesApp.CesApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CesApp) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_CesApp) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_CesApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesApp) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_CesApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesApp) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_CesApp) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_CesApp) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_CesApp) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_CesApp) PutAudioProcessingConfig(value *CesAppAudioProcessingConfig) {
	if err := c.validatePutAudioProcessingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAudioProcessingConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutClientCertificateSettings(value *CesAppClientCertificateSettings) {
	if err := c.validatePutClientCertificateSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientCertificateSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutDataStoreSettings(value *CesAppDataStoreSettings) {
	if err := c.validatePutDataStoreSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataStoreSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutDefaultChannelProfile(value *CesAppDefaultChannelProfile) {
	if err := c.validatePutDefaultChannelProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultChannelProfile",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutEvaluationMetricsThresholds(value *CesAppEvaluationMetricsThresholds) {
	if err := c.validatePutEvaluationMetricsThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEvaluationMetricsThresholds",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutLanguageSettings(value *CesAppLanguageSettings) {
	if err := c.validatePutLanguageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLanguageSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutLoggingSettings(value *CesAppLoggingSettings) {
	if err := c.validatePutLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoggingSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutModelSettings(value *CesAppModelSettings) {
	if err := c.validatePutModelSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModelSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutTimeouts(value *CesAppTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutTimeZoneSettings(value *CesAppTimeZoneSettings) {
	if err := c.validatePutTimeZoneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeZoneSettings",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) PutVariableDeclarations(value any) {
	if err := c.validatePutVariableDeclarationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVariableDeclarations",
		[]any{value},
	)
}

func (c *jsiiProxy_CesApp) ResetAudioProcessingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAudioProcessingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetClientCertificateSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetClientCertificateSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetDataStoreSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetDataStoreSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetDefaultChannelProfile() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultChannelProfile",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetEvaluationMetricsThresholds() {
	_jsii_.InvokeVoid(
		c,
		"resetEvaluationMetricsThresholds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetGlobalInstruction() {
	_jsii_.InvokeVoid(
		c,
		"resetGlobalInstruction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetGuardrails() {
	_jsii_.InvokeVoid(
		c,
		"resetGuardrails",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetLanguageSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetLoggingSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetModelSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetModelSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetPinned() {
	_jsii_.InvokeVoid(
		c,
		"resetPinned",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetRootAgent() {
	_jsii_.InvokeVoid(
		c,
		"resetRootAgent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetTimeZoneSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeZoneSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) ResetVariableDeclarations() {
	_jsii_.InvokeVoid(
		c,
		"resetVariableDeclarations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesApp) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesApp) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
