package googledatafusioninstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatafusioninstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_fusion_instance google_data_fusion_instance}.
type GoogleDataFusionInstance interface {
	cdktf.TerraformResource
	Accelerators() GoogleDataFusionInstanceAcceleratorsList
	AcceleratorsInput() interface{}
	ApiEndpoint() *string
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
	CreateTime() *string
	CryptoKeyConfig() GoogleDataFusionInstanceCryptoKeyConfigOutputReference
	CryptoKeyConfigInput() *GoogleDataFusionInstanceCryptoKeyConfig
	DataprocServiceAccount() *string
	SetDataprocServiceAccount(val *string)
	DataprocServiceAccountInput() *string
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
	EnableRbac() interface{}
	SetEnableRbac(val interface{})
	EnableRbacInput() interface{}
	EnableStackdriverLogging() interface{}
	SetEnableStackdriverLogging(val interface{})
	EnableStackdriverLoggingInput() interface{}
	EnableStackdriverMonitoring() interface{}
	SetEnableStackdriverMonitoring(val interface{})
	EnableStackdriverMonitoringInput() interface{}
	EventPublishConfig() GoogleDataFusionInstanceEventPublishConfigOutputReference
	EventPublishConfigInput() *GoogleDataFusionInstanceEventPublishConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsBucket() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GoogleDataFusionInstanceNetworkConfigOutputReference
	NetworkConfigInput() *GoogleDataFusionInstanceNetworkConfig
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	P4ServiceAccount() *string
	PrivateInstance() interface{}
	SetPrivateInstance(val interface{})
	PrivateInstanceInput() interface{}
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ServiceAccount() *string
	ServiceEndpoint() *string
	State() *string
	StateMessage() *string
	TenantProjectId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDataFusionInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdateTime() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutAccelerators(value interface{})
	PutCryptoKeyConfig(value *GoogleDataFusionInstanceCryptoKeyConfig)
	PutEventPublishConfig(value *GoogleDataFusionInstanceEventPublishConfig)
	PutNetworkConfig(value *GoogleDataFusionInstanceNetworkConfig)
	PutTimeouts(value *GoogleDataFusionInstanceTimeouts)
	ResetAccelerators()
	ResetCryptoKeyConfig()
	ResetDataprocServiceAccount()
	ResetDescription()
	ResetDisplayName()
	ResetEnableRbac()
	ResetEnableStackdriverLogging()
	ResetEnableStackdriverMonitoring()
	ResetEventPublishConfig()
	ResetId()
	ResetLabels()
	ResetNetworkConfig()
	ResetOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateInstance()
	ResetProject()
	ResetRegion()
	ResetTimeouts()
	ResetVersion()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleDataFusionInstance
type jsiiProxy_GoogleDataFusionInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDataFusionInstance) Accelerators() GoogleDataFusionInstanceAcceleratorsList {
	var returns GoogleDataFusionInstanceAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) CryptoKeyConfig() GoogleDataFusionInstanceCryptoKeyConfigOutputReference {
	var returns GoogleDataFusionInstanceCryptoKeyConfigOutputReference
	_jsii_.Get(
		j,
		"cryptoKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) CryptoKeyConfigInput() *GoogleDataFusionInstanceCryptoKeyConfig {
	var returns *GoogleDataFusionInstanceCryptoKeyConfig
	_jsii_.Get(
		j,
		"cryptoKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DataprocServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DataprocServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableRbac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableRbacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRbacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableStackdriverLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableStackdriverLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableStackdriverMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EnableStackdriverMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStackdriverMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EventPublishConfig() GoogleDataFusionInstanceEventPublishConfigOutputReference {
	var returns GoogleDataFusionInstanceEventPublishConfigOutputReference
	_jsii_.Get(
		j,
		"eventPublishConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) EventPublishConfigInput() *GoogleDataFusionInstanceEventPublishConfig {
	var returns *GoogleDataFusionInstanceEventPublishConfig
	_jsii_.Get(
		j,
		"eventPublishConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) GcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) NetworkConfig() GoogleDataFusionInstanceNetworkConfigOutputReference {
	var returns GoogleDataFusionInstanceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) NetworkConfigInput() *GoogleDataFusionInstanceNetworkConfig {
	var returns *GoogleDataFusionInstanceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) P4ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"p4ServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) PrivateInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) PrivateInstanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ServiceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TenantProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Timeouts() GoogleDataFusionInstanceTimeoutsOutputReference {
	var returns GoogleDataFusionInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataFusionInstance) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_fusion_instance google_data_fusion_instance} Resource.
func NewGoogleDataFusionInstance(scope constructs.Construct, id *string, config *GoogleDataFusionInstanceConfig) GoogleDataFusionInstance {
	_init_.Initialize()

	if err := validateNewGoogleDataFusionInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataFusionInstance{}

	_jsii_.Create(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_data_fusion_instance google_data_fusion_instance} Resource.
func NewGoogleDataFusionInstance_Override(g GoogleDataFusionInstance, scope constructs.Construct, id *string, config *GoogleDataFusionInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetDataprocServiceAccount(val *string) {
	if err := j.validateSetDataprocServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataprocServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetEnableRbac(val interface{}) {
	if err := j.validateSetEnableRbacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRbac",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetEnableStackdriverLogging(val interface{}) {
	if err := j.validateSetEnableStackdriverLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverLogging",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetEnableStackdriverMonitoring(val interface{}) {
	if err := j.validateSetEnableStackdriverMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStackdriverMonitoring",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetOptions(val *map[string]*string) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetPrivateInstance(val interface{}) {
	if err := j.validateSetPrivateInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateInstance",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_GoogleDataFusionInstance)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
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
func GoogleDataFusionInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataFusionInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataFusionInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataFusionInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDataFusionInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDataFusionInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDataFusionInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleDataFusionInstance.GoogleDataFusionInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDataFusionInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataFusionInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataFusionInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataFusionInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) PutAccelerators(value interface{}) {
	if err := g.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) PutCryptoKeyConfig(value *GoogleDataFusionInstanceCryptoKeyConfig) {
	if err := g.validatePutCryptoKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCryptoKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) PutEventPublishConfig(value *GoogleDataFusionInstanceEventPublishConfig) {
	if err := g.validatePutEventPublishConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEventPublishConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) PutNetworkConfig(value *GoogleDataFusionInstanceNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) PutTimeouts(value *GoogleDataFusionInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetAccelerators() {
	_jsii_.InvokeVoid(
		g,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetCryptoKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCryptoKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetDataprocServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetEnableRbac() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableRbac",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetEnableStackdriverLogging() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableStackdriverLogging",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetEnableStackdriverMonitoring() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableStackdriverMonitoring",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetEventPublishConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEventPublishConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetPrivateInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataFusionInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataFusionInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataFusionInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataFusionInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

