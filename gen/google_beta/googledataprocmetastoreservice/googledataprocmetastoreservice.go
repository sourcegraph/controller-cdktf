package googledataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledataprocmetastoreservice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_dataproc_metastore_service google_dataproc_metastore_service}.
type GoogleDataprocMetastoreService interface {
	cdktf.TerraformResource
	ArtifactGcsUri() *string
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
	DatabaseType() *string
	SetDatabaseType(val *string)
	DatabaseTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktf.StringMap
	EncryptionConfig() GoogleDataprocMetastoreServiceEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleDataprocMetastoreServiceEncryptionConfig
	EndpointUri() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HiveMetastoreConfig() GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference
	HiveMetastoreConfigInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfig
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
	MaintenanceWindow() GoogleDataprocMetastoreServiceMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleDataprocMetastoreServiceMaintenanceWindow
	MetadataIntegration() GoogleDataprocMetastoreServiceMetadataIntegrationOutputReference
	MetadataIntegrationInput() *GoogleDataprocMetastoreServiceMetadataIntegration
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkConfig() GoogleDataprocMetastoreServiceNetworkConfigOutputReference
	NetworkConfigInput() *GoogleDataprocMetastoreServiceNetworkConfig
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelInput() *string
	ScalingConfig() GoogleDataprocMetastoreServiceScalingConfigOutputReference
	ScalingConfigInput() *GoogleDataprocMetastoreServiceScalingConfig
	ServiceId() *string
	SetServiceId(val *string)
	ServiceIdInput() *string
	State() *string
	StateMessage() *string
	TelemetryConfig() GoogleDataprocMetastoreServiceTelemetryConfigOutputReference
	TelemetryConfigInput() *GoogleDataprocMetastoreServiceTelemetryConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() GoogleDataprocMetastoreServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	PutEncryptionConfig(value *GoogleDataprocMetastoreServiceEncryptionConfig)
	PutHiveMetastoreConfig(value *GoogleDataprocMetastoreServiceHiveMetastoreConfig)
	PutMaintenanceWindow(value *GoogleDataprocMetastoreServiceMaintenanceWindow)
	PutMetadataIntegration(value *GoogleDataprocMetastoreServiceMetadataIntegration)
	PutNetworkConfig(value *GoogleDataprocMetastoreServiceNetworkConfig)
	PutScalingConfig(value *GoogleDataprocMetastoreServiceScalingConfig)
	PutTelemetryConfig(value *GoogleDataprocMetastoreServiceTelemetryConfig)
	PutTimeouts(value *GoogleDataprocMetastoreServiceTimeouts)
	ResetDatabaseType()
	ResetEncryptionConfig()
	ResetHiveMetastoreConfig()
	ResetId()
	ResetLabels()
	ResetLocation()
	ResetMaintenanceWindow()
	ResetMetadataIntegration()
	ResetNetwork()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProject()
	ResetReleaseChannel()
	ResetScalingConfig()
	ResetTelemetryConfig()
	ResetTier()
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

// The jsii proxy struct for GoogleDataprocMetastoreService
type jsiiProxy_GoogleDataprocMetastoreService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ArtifactGcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactGcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) DatabaseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) EncryptionConfig() GoogleDataprocMetastoreServiceEncryptionConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) EncryptionConfigInput() *GoogleDataprocMetastoreServiceEncryptionConfig {
	var returns *GoogleDataprocMetastoreServiceEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) HiveMetastoreConfig() GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceHiveMetastoreConfigOutputReference
	_jsii_.Get(
		j,
		"hiveMetastoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) HiveMetastoreConfigInput() *GoogleDataprocMetastoreServiceHiveMetastoreConfig {
	var returns *GoogleDataprocMetastoreServiceHiveMetastoreConfig
	_jsii_.Get(
		j,
		"hiveMetastoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) MaintenanceWindow() GoogleDataprocMetastoreServiceMaintenanceWindowOutputReference {
	var returns GoogleDataprocMetastoreServiceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) MaintenanceWindowInput() *GoogleDataprocMetastoreServiceMaintenanceWindow {
	var returns *GoogleDataprocMetastoreServiceMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) MetadataIntegration() GoogleDataprocMetastoreServiceMetadataIntegrationOutputReference {
	var returns GoogleDataprocMetastoreServiceMetadataIntegrationOutputReference
	_jsii_.Get(
		j,
		"metadataIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) MetadataIntegrationInput() *GoogleDataprocMetastoreServiceMetadataIntegration {
	var returns *GoogleDataprocMetastoreServiceMetadataIntegration
	_jsii_.Get(
		j,
		"metadataIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) NetworkConfig() GoogleDataprocMetastoreServiceNetworkConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) NetworkConfigInput() *GoogleDataprocMetastoreServiceNetworkConfig {
	var returns *GoogleDataprocMetastoreServiceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ScalingConfig() GoogleDataprocMetastoreServiceScalingConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceScalingConfigOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ScalingConfigInput() *GoogleDataprocMetastoreServiceScalingConfig {
	var returns *GoogleDataprocMetastoreServiceScalingConfig
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) ServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TelemetryConfig() GoogleDataprocMetastoreServiceTelemetryConfigOutputReference {
	var returns GoogleDataprocMetastoreServiceTelemetryConfigOutputReference
	_jsii_.Get(
		j,
		"telemetryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TelemetryConfigInput() *GoogleDataprocMetastoreServiceTelemetryConfig {
	var returns *GoogleDataprocMetastoreServiceTelemetryConfig
	_jsii_.Get(
		j,
		"telemetryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Timeouts() GoogleDataprocMetastoreServiceTimeoutsOutputReference {
	var returns GoogleDataprocMetastoreServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocMetastoreService) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_dataproc_metastore_service google_dataproc_metastore_service} Resource.
func NewGoogleDataprocMetastoreService(scope constructs.Construct, id *string, config *GoogleDataprocMetastoreServiceConfig) GoogleDataprocMetastoreService {
	_init_.Initialize()

	if err := validateNewGoogleDataprocMetastoreServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocMetastoreService{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_dataproc_metastore_service google_dataproc_metastore_service} Resource.
func NewGoogleDataprocMetastoreService_Override(g GoogleDataprocMetastoreService, scope constructs.Construct, id *string, config *GoogleDataprocMetastoreServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetDatabaseType(val *string) {
	if err := j.validateSetDatabaseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetReleaseChannel(val *string) {
	if err := j.validateSetReleaseChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetServiceId(val *string) {
	if err := j.validateSetServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceId",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocMetastoreService)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
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
func GoogleDataprocMetastoreService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocMetastoreService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocMetastoreService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocMetastoreService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataprocMetastoreService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataprocMetastoreService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataprocMetastoreService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDataprocMetastoreService.GoogleDataprocMetastoreService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocMetastoreService) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutEncryptionConfig(value *GoogleDataprocMetastoreServiceEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutHiveMetastoreConfig(value *GoogleDataprocMetastoreServiceHiveMetastoreConfig) {
	if err := g.validatePutHiveMetastoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHiveMetastoreConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutMaintenanceWindow(value *GoogleDataprocMetastoreServiceMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutMetadataIntegration(value *GoogleDataprocMetastoreServiceMetadataIntegration) {
	if err := g.validatePutMetadataIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetadataIntegration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutNetworkConfig(value *GoogleDataprocMetastoreServiceNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutScalingConfig(value *GoogleDataprocMetastoreServiceScalingConfig) {
	if err := g.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutTelemetryConfig(value *GoogleDataprocMetastoreServiceTelemetryConfig) {
	if err := g.validatePutTelemetryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTelemetryConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) PutTimeouts(value *GoogleDataprocMetastoreServiceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetDatabaseType() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetHiveMetastoreConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHiveMetastoreConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetMetadataIntegration() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataIntegration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetPort() {
	_jsii_.InvokeVoid(
		g,
		"resetPort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetTelemetryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTelemetryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetTier() {
	_jsii_.InvokeVoid(
		g,
		"resetTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocMetastoreService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

