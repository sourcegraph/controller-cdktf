package discoveryenginelicenseconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginelicenseconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_license_config google_discovery_engine_license_config}.
type DiscoveryEngineLicenseConfig interface {
	cdktf.TerraformResource
	AutoRenew() interface{}
	SetAutoRenew(val interface{})
	AutoRenewInput() interface{}
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
	EndDate() DiscoveryEngineLicenseConfigEndDateOutputReference
	EndDateInput() *DiscoveryEngineLicenseConfigEndDate
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FreeTrial() interface{}
	SetFreeTrial(val interface{})
	FreeTrialInput() interface{}
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LicenseConfigId() *string
	SetLicenseConfigId(val *string)
	LicenseConfigIdInput() *string
	LicenseCount() *float64
	SetLicenseCount(val *float64)
	LicenseCountInput() *float64
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
	// Experimental.
	RawOverrides() interface{}
	StartDate() DiscoveryEngineLicenseConfigStartDateOutputReference
	StartDateInput() *DiscoveryEngineLicenseConfigStartDate
	SubscriptionTerm() *string
	SetSubscriptionTerm(val *string)
	SubscriptionTermInput() *string
	SubscriptionTier() *string
	SetSubscriptionTier(val *string)
	SubscriptionTierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DiscoveryEngineLicenseConfigTimeoutsOutputReference
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
	PutEndDate(value *DiscoveryEngineLicenseConfigEndDate)
	PutStartDate(value *DiscoveryEngineLicenseConfigStartDate)
	PutTimeouts(value *DiscoveryEngineLicenseConfigTimeouts)
	ResetAutoRenew()
	ResetEndDate()
	ResetFreeTrial()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for DiscoveryEngineLicenseConfig
type jsiiProxy_DiscoveryEngineLicenseConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) AutoRenew() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) AutoRenewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRenewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) EndDate() DiscoveryEngineLicenseConfigEndDateOutputReference {
	var returns DiscoveryEngineLicenseConfigEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) EndDateInput() *DiscoveryEngineLicenseConfigEndDate {
	var returns *DiscoveryEngineLicenseConfigEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) FreeTrial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freeTrial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) FreeTrialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freeTrialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) LicenseConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) LicenseConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) LicenseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"licenseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) LicenseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"licenseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) StartDate() DiscoveryEngineLicenseConfigStartDateOutputReference {
	var returns DiscoveryEngineLicenseConfigStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) StartDateInput() *DiscoveryEngineLicenseConfigStartDate {
	var returns *DiscoveryEngineLicenseConfigStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) SubscriptionTerm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) SubscriptionTermInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) SubscriptionTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) SubscriptionTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) Timeouts() DiscoveryEngineLicenseConfigTimeoutsOutputReference {
	var returns DiscoveryEngineLicenseConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_license_config google_discovery_engine_license_config} Resource.
func NewDiscoveryEngineLicenseConfig(scope constructs.Construct, id *string, config *DiscoveryEngineLicenseConfigConfig) DiscoveryEngineLicenseConfig {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineLicenseConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineLicenseConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_license_config google_discovery_engine_license_config} Resource.
func NewDiscoveryEngineLicenseConfig_Override(d DiscoveryEngineLicenseConfig, scope constructs.Construct, id *string, config *DiscoveryEngineLicenseConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetAutoRenew(val interface{}) {
	if err := j.validateSetAutoRenewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRenew",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetFreeTrial(val interface{}) {
	if err := j.validateSetFreeTrialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freeTrial",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetLicenseConfigId(val *string) {
	if err := j.validateSetLicenseConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseConfigId",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetLicenseCount(val *float64) {
	if err := j.validateSetLicenseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseCount",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetSubscriptionTerm(val *string) {
	if err := j.validateSetSubscriptionTermParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionTerm",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineLicenseConfig)SetSubscriptionTier(val *string) {
	if err := j.validateSetSubscriptionTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionTier",
		val,
	)
}

// Generates CDKTF code for importing a DiscoveryEngineLicenseConfig resource upon running "cdktf plan <stack-name>".
func DiscoveryEngineLicenseConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDiscoveryEngineLicenseConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
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
func DiscoveryEngineLicenseConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineLicenseConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineLicenseConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineLicenseConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DiscoveryEngineLicenseConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDiscoveryEngineLicenseConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DiscoveryEngineLicenseConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.discoveryEngineLicenseConfig.DiscoveryEngineLicenseConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) PutEndDate(value *DiscoveryEngineLicenseConfigEndDate) {
	if err := d.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEndDate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) PutStartDate(value *DiscoveryEngineLicenseConfigStartDate) {
	if err := d.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStartDate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) PutTimeouts(value *DiscoveryEngineLicenseConfigTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetAutoRenew() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoRenew",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetEndDate() {
	_jsii_.InvokeVoid(
		d,
		"resetEndDate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetFreeTrial() {
	_jsii_.InvokeVoid(
		d,
		"resetFreeTrial",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineLicenseConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

