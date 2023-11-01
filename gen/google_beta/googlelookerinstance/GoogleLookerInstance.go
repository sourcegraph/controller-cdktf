package googlelookerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlelookerinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_looker_instance google_looker_instance}.
type GoogleLookerInstance interface {
	cdktf.TerraformResource
	AdminSettings() GoogleLookerInstanceAdminSettingsOutputReference
	AdminSettingsInput() *GoogleLookerInstanceAdminSettings
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumerNetwork() *string
	SetConsumerNetwork(val *string)
	ConsumerNetworkInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DenyMaintenancePeriod() GoogleLookerInstanceDenyMaintenancePeriodOutputReference
	DenyMaintenancePeriodInput() *GoogleLookerInstanceDenyMaintenancePeriod
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EgressPublicIp() *string
	EncryptionConfig() GoogleLookerInstanceEncryptionConfigOutputReference
	EncryptionConfigInput() *GoogleLookerInstanceEncryptionConfig
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
	IngressPrivateIp() *string
	IngressPublicIp() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LookerUri() *string
	LookerVersion() *string
	MaintenanceWindow() GoogleLookerInstanceMaintenanceWindowOutputReference
	MaintenanceWindowInput() *GoogleLookerInstanceMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OauthConfig() GoogleLookerInstanceOauthConfigOutputReference
	OauthConfigInput() *GoogleLookerInstanceOauthConfig
	PlatformEdition() *string
	SetPlatformEdition(val *string)
	PlatformEditionInput() *string
	PrivateIpEnabled() interface{}
	SetPrivateIpEnabled(val interface{})
	PrivateIpEnabledInput() interface{}
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
	PublicIpEnabled() interface{}
	SetPublicIpEnabled(val interface{})
	PublicIpEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReservedRange() *string
	SetReservedRange(val *string)
	ReservedRangeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleLookerInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	UserMetadata() GoogleLookerInstanceUserMetadataOutputReference
	UserMetadataInput() *GoogleLookerInstanceUserMetadata
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
	PutAdminSettings(value *GoogleLookerInstanceAdminSettings)
	PutDenyMaintenancePeriod(value *GoogleLookerInstanceDenyMaintenancePeriod)
	PutEncryptionConfig(value *GoogleLookerInstanceEncryptionConfig)
	PutMaintenanceWindow(value *GoogleLookerInstanceMaintenanceWindow)
	PutOauthConfig(value *GoogleLookerInstanceOauthConfig)
	PutTimeouts(value *GoogleLookerInstanceTimeouts)
	PutUserMetadata(value *GoogleLookerInstanceUserMetadata)
	ResetAdminSettings()
	ResetConsumerNetwork()
	ResetDenyMaintenancePeriod()
	ResetEncryptionConfig()
	ResetId()
	ResetMaintenanceWindow()
	ResetOauthConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformEdition()
	ResetPrivateIpEnabled()
	ResetProject()
	ResetPublicIpEnabled()
	ResetRegion()
	ResetReservedRange()
	ResetTimeouts()
	ResetUserMetadata()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleLookerInstance
type jsiiProxy_GoogleLookerInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleLookerInstance) AdminSettings() GoogleLookerInstanceAdminSettingsOutputReference {
	var returns GoogleLookerInstanceAdminSettingsOutputReference
	_jsii_.Get(
		j,
		"adminSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) AdminSettingsInput() *GoogleLookerInstanceAdminSettings {
	var returns *GoogleLookerInstanceAdminSettings
	_jsii_.Get(
		j,
		"adminSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ConsumerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ConsumerNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) DenyMaintenancePeriod() GoogleLookerInstanceDenyMaintenancePeriodOutputReference {
	var returns GoogleLookerInstanceDenyMaintenancePeriodOutputReference
	_jsii_.Get(
		j,
		"denyMaintenancePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) DenyMaintenancePeriodInput() *GoogleLookerInstanceDenyMaintenancePeriod {
	var returns *GoogleLookerInstanceDenyMaintenancePeriod
	_jsii_.Get(
		j,
		"denyMaintenancePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) EgressPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) EncryptionConfig() GoogleLookerInstanceEncryptionConfigOutputReference {
	var returns GoogleLookerInstanceEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) EncryptionConfigInput() *GoogleLookerInstanceEncryptionConfig {
	var returns *GoogleLookerInstanceEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) IngressPrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressPrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) IngressPublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) LookerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) LookerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) MaintenanceWindow() GoogleLookerInstanceMaintenanceWindowOutputReference {
	var returns GoogleLookerInstanceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) MaintenanceWindowInput() *GoogleLookerInstanceMaintenanceWindow {
	var returns *GoogleLookerInstanceMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) OauthConfig() GoogleLookerInstanceOauthConfigOutputReference {
	var returns GoogleLookerInstanceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) OauthConfigInput() *GoogleLookerInstanceOauthConfig {
	var returns *GoogleLookerInstanceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PlatformEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PlatformEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PrivateIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PrivateIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PublicIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) PublicIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ReservedRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) ReservedRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reservedRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) Timeouts() GoogleLookerInstanceTimeoutsOutputReference {
	var returns GoogleLookerInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) UserMetadata() GoogleLookerInstanceUserMetadataOutputReference {
	var returns GoogleLookerInstanceUserMetadataOutputReference
	_jsii_.Get(
		j,
		"userMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLookerInstance) UserMetadataInput() *GoogleLookerInstanceUserMetadata {
	var returns *GoogleLookerInstanceUserMetadata
	_jsii_.Get(
		j,
		"userMetadataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_looker_instance google_looker_instance} Resource.
func NewGoogleLookerInstance(scope constructs.Construct, id *string, config *GoogleLookerInstanceConfig) GoogleLookerInstance {
	_init_.Initialize()

	if err := validateNewGoogleLookerInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLookerInstance{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_looker_instance google_looker_instance} Resource.
func NewGoogleLookerInstance_Override(g GoogleLookerInstance, scope constructs.Construct, id *string, config *GoogleLookerInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetConsumerNetwork(val *string) {
	if err := j.validateSetConsumerNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetPlatformEdition(val *string) {
	if err := j.validateSetPlatformEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformEdition",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetPrivateIpEnabled(val interface{}) {
	if err := j.validateSetPrivateIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetPublicIpEnabled(val interface{}) {
	if err := j.validateSetPublicIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleLookerInstance)SetReservedRange(val *string) {
	if err := j.validateSetReservedRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedRange",
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
func GoogleLookerInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLookerInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleLookerInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLookerInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleLookerInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleLookerInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleLookerInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleLookerInstance.GoogleLookerInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleLookerInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLookerInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLookerInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLookerInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLookerInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLookerInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLookerInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLookerInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLookerInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLookerInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleLookerInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutAdminSettings(value *GoogleLookerInstanceAdminSettings) {
	if err := g.validatePutAdminSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdminSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutDenyMaintenancePeriod(value *GoogleLookerInstanceDenyMaintenancePeriod) {
	if err := g.validatePutDenyMaintenancePeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDenyMaintenancePeriod",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutEncryptionConfig(value *GoogleLookerInstanceEncryptionConfig) {
	if err := g.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutMaintenanceWindow(value *GoogleLookerInstanceMaintenanceWindow) {
	if err := g.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutOauthConfig(value *GoogleLookerInstanceOauthConfig) {
	if err := g.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutTimeouts(value *GoogleLookerInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) PutUserMetadata(value *GoogleLookerInstanceUserMetadata) {
	if err := g.validatePutUserMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserMetadata",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetAdminSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAdminSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetConsumerNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumerNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetDenyMaintenancePeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetDenyMaintenancePeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetPlatformEdition() {
	_jsii_.InvokeVoid(
		g,
		"resetPlatformEdition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetPrivateIpEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivateIpEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetPublicIpEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicIpEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetReservedRange() {
	_jsii_.InvokeVoid(
		g,
		"resetReservedRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) ResetUserMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetUserMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLookerInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLookerInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLookerInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLookerInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

