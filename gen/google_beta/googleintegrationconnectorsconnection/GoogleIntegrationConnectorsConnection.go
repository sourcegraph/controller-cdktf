package googleintegrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationconnectorsconnection/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_integration_connectors_connection google_integration_connectors_connection}.
type GoogleIntegrationConnectorsConnection interface {
	cdktf.TerraformResource
	AuthConfig() GoogleIntegrationConnectorsConnectionAuthConfigOutputReference
	AuthConfigInput() *GoogleIntegrationConnectorsConnectionAuthConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigVariable() GoogleIntegrationConnectorsConnectionConfigVariableList
	ConfigVariableInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionRevision() *string
	ConnectorVersion() *string
	SetConnectorVersion(val *string)
	ConnectorVersionInfraConfig() GoogleIntegrationConnectorsConnectionConnectorVersionInfraConfigList
	ConnectorVersionInput() *string
	ConnectorVersionLaunchStage() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationConfig() GoogleIntegrationConnectorsConnectionDestinationConfigList
	DestinationConfigInput() interface{}
	EffectiveLabels() cdktf.StringMap
	EventingConfig() GoogleIntegrationConnectorsConnectionEventingConfigOutputReference
	EventingConfigInput() *GoogleIntegrationConnectorsConnectionEventingConfig
	EventingEnablementType() *string
	SetEventingEnablementType(val *string)
	EventingEnablementTypeInput() *string
	EventingRuntimeData() GoogleIntegrationConnectorsConnectionEventingRuntimeDataList
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LockConfig() GoogleIntegrationConnectorsConnectionLockConfigOutputReference
	LockConfigInput() *GoogleIntegrationConnectorsConnectionLockConfig
	LogConfig() GoogleIntegrationConnectorsConnectionLogConfigOutputReference
	LogConfigInput() *GoogleIntegrationConnectorsConnectionLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeConfig() GoogleIntegrationConnectorsConnectionNodeConfigOutputReference
	NodeConfigInput() *GoogleIntegrationConnectorsConnectionNodeConfig
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
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceDirectory() *string
	SslConfig() GoogleIntegrationConnectorsConnectionSslConfigOutputReference
	SslConfigInput() *GoogleIntegrationConnectorsConnectionSslConfig
	Status() GoogleIntegrationConnectorsConnectionStatusList
	SubscriptionType() *string
	Suspended() interface{}
	SetSuspended(val interface{})
	SuspendedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleIntegrationConnectorsConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAuthConfig(value *GoogleIntegrationConnectorsConnectionAuthConfig)
	PutConfigVariable(value interface{})
	PutDestinationConfig(value interface{})
	PutEventingConfig(value *GoogleIntegrationConnectorsConnectionEventingConfig)
	PutLockConfig(value *GoogleIntegrationConnectorsConnectionLockConfig)
	PutLogConfig(value *GoogleIntegrationConnectorsConnectionLogConfig)
	PutNodeConfig(value *GoogleIntegrationConnectorsConnectionNodeConfig)
	PutSslConfig(value *GoogleIntegrationConnectorsConnectionSslConfig)
	PutTimeouts(value *GoogleIntegrationConnectorsConnectionTimeouts)
	ResetAuthConfig()
	ResetConfigVariable()
	ResetDescription()
	ResetDestinationConfig()
	ResetEventingConfig()
	ResetEventingEnablementType()
	ResetId()
	ResetLabels()
	ResetLockConfig()
	ResetLogConfig()
	ResetNodeConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetServiceAccount()
	ResetSslConfig()
	ResetSuspended()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleIntegrationConnectorsConnection
type jsiiProxy_GoogleIntegrationConnectorsConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) AuthConfig() GoogleIntegrationConnectorsConnectionAuthConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) AuthConfigInput() *GoogleIntegrationConnectorsConnectionAuthConfig {
	var returns *GoogleIntegrationConnectorsConnectionAuthConfig
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConfigVariable() GoogleIntegrationConnectorsConnectionConfigVariableList {
	var returns GoogleIntegrationConnectorsConnectionConfigVariableList
	_jsii_.Get(
		j,
		"configVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConfigVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConnectionRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConnectorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConnectorVersionInfraConfig() GoogleIntegrationConnectorsConnectionConnectorVersionInfraConfigList {
	var returns GoogleIntegrationConnectorsConnectionConnectorVersionInfraConfigList
	_jsii_.Get(
		j,
		"connectorVersionInfraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConnectorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConnectorVersionLaunchStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersionLaunchStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) DestinationConfig() GoogleIntegrationConnectorsConnectionDestinationConfigList {
	var returns GoogleIntegrationConnectorsConnectionDestinationConfigList
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) DestinationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EventingConfig() GoogleIntegrationConnectorsConnectionEventingConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionEventingConfigOutputReference
	_jsii_.Get(
		j,
		"eventingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EventingConfigInput() *GoogleIntegrationConnectorsConnectionEventingConfig {
	var returns *GoogleIntegrationConnectorsConnectionEventingConfig
	_jsii_.Get(
		j,
		"eventingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EventingEnablementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventingEnablementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EventingEnablementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventingEnablementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) EventingRuntimeData() GoogleIntegrationConnectorsConnectionEventingRuntimeDataList {
	var returns GoogleIntegrationConnectorsConnectionEventingRuntimeDataList
	_jsii_.Get(
		j,
		"eventingRuntimeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LockConfig() GoogleIntegrationConnectorsConnectionLockConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionLockConfigOutputReference
	_jsii_.Get(
		j,
		"lockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LockConfigInput() *GoogleIntegrationConnectorsConnectionLockConfig {
	var returns *GoogleIntegrationConnectorsConnectionLockConfig
	_jsii_.Get(
		j,
		"lockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LogConfig() GoogleIntegrationConnectorsConnectionLogConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) LogConfigInput() *GoogleIntegrationConnectorsConnectionLogConfig {
	var returns *GoogleIntegrationConnectorsConnectionLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) NodeConfig() GoogleIntegrationConnectorsConnectionNodeConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) NodeConfigInput() *GoogleIntegrationConnectorsConnectionNodeConfig {
	var returns *GoogleIntegrationConnectorsConnectionNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) ServiceDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) SslConfig() GoogleIntegrationConnectorsConnectionSslConfigOutputReference {
	var returns GoogleIntegrationConnectorsConnectionSslConfigOutputReference
	_jsii_.Get(
		j,
		"sslConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) SslConfigInput() *GoogleIntegrationConnectorsConnectionSslConfig {
	var returns *GoogleIntegrationConnectorsConnectionSslConfig
	_jsii_.Get(
		j,
		"sslConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Status() GoogleIntegrationConnectorsConnectionStatusList {
	var returns GoogleIntegrationConnectorsConnectionStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Suspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) SuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) Timeouts() GoogleIntegrationConnectorsConnectionTimeoutsOutputReference {
	var returns GoogleIntegrationConnectorsConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_integration_connectors_connection google_integration_connectors_connection} Resource.
func NewGoogleIntegrationConnectorsConnection(scope constructs.Construct, id *string, config *GoogleIntegrationConnectorsConnectionConfig) GoogleIntegrationConnectorsConnection {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationConnectorsConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationConnectorsConnection{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_integration_connectors_connection google_integration_connectors_connection} Resource.
func NewGoogleIntegrationConnectorsConnection_Override(g GoogleIntegrationConnectorsConnection, scope constructs.Construct, id *string, config *GoogleIntegrationConnectorsConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetConnectorVersion(val *string) {
	if err := j.validateSetConnectorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetEventingEnablementType(val *string) {
	if err := j.validateSetEventingEnablementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventingEnablementType",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationConnectorsConnection)SetSuspended(val interface{}) {
	if err := j.validateSetSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspended",
		val,
	)
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
func GoogleIntegrationConnectorsConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIntegrationConnectorsConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIntegrationConnectorsConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIntegrationConnectorsConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIntegrationConnectorsConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIntegrationConnectorsConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleIntegrationConnectorsConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleIntegrationConnectorsConnection.GoogleIntegrationConnectorsConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutAuthConfig(value *GoogleIntegrationConnectorsConnectionAuthConfig) {
	if err := g.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutConfigVariable(value interface{}) {
	if err := g.validatePutConfigVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigVariable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutDestinationConfig(value interface{}) {
	if err := g.validatePutDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutEventingConfig(value *GoogleIntegrationConnectorsConnectionEventingConfig) {
	if err := g.validatePutEventingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEventingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutLockConfig(value *GoogleIntegrationConnectorsConnectionLockConfig) {
	if err := g.validatePutLockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLockConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutLogConfig(value *GoogleIntegrationConnectorsConnectionLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutNodeConfig(value *GoogleIntegrationConnectorsConnectionNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutSslConfig(value *GoogleIntegrationConnectorsConnectionSslConfig) {
	if err := g.validatePutSslConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSslConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) PutTimeouts(value *GoogleIntegrationConnectorsConnectionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetConfigVariable() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigVariable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetEventingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEventingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetEventingEnablementType() {
	_jsii_.InvokeVoid(
		g,
		"resetEventingEnablementType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetLockConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLockConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetSslConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSslConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetSuspended() {
	_jsii_.InvokeVoid(
		g,
		"resetSuspended",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationConnectorsConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

