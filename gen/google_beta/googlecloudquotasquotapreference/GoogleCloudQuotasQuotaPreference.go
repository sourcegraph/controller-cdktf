package googlecloudquotasquotapreference

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudquotasquotapreference/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_cloud_quotas_quota_preference google_cloud_quotas_quota_preference}.
type GoogleCloudQuotasQuotaPreference interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactEmail() *string
	SetContactEmail(val *string)
	ContactEmailInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Dimensions() *map[string]*string
	SetDimensions(val *map[string]*string)
	DimensionsInput() *map[string]*string
	Etag() *string
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
	IgnoreSafetyChecks() *string
	SetIgnoreSafetyChecks(val *string)
	IgnoreSafetyChecksInput() *string
	Justification() *string
	SetJustification(val *string)
	JustificationInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parent() *string
	SetParent(val *string)
	ParentInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuotaConfig() GoogleCloudQuotasQuotaPreferenceQuotaConfigOutputReference
	QuotaConfigInput() *GoogleCloudQuotasQuotaPreferenceQuotaConfig
	QuotaId() *string
	SetQuotaId(val *string)
	QuotaIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudQuotasQuotaPreferenceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
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
	PutQuotaConfig(value *GoogleCloudQuotasQuotaPreferenceQuotaConfig)
	PutTimeouts(value *GoogleCloudQuotasQuotaPreferenceTimeouts)
	ResetContactEmail()
	ResetDimensions()
	ResetId()
	ResetIgnoreSafetyChecks()
	ResetJustification()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParent()
	ResetQuotaId()
	ResetService()
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

// The jsii proxy struct for GoogleCloudQuotasQuotaPreference
type jsiiProxy_GoogleCloudQuotasQuotaPreference struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ContactEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Dimensions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) DimensionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) IgnoreSafetyChecks() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreSafetyChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) IgnoreSafetyChecksInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreSafetyChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) JustificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Parent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ParentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) QuotaConfig() GoogleCloudQuotasQuotaPreferenceQuotaConfigOutputReference {
	var returns GoogleCloudQuotasQuotaPreferenceQuotaConfigOutputReference
	_jsii_.Get(
		j,
		"quotaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) QuotaConfigInput() *GoogleCloudQuotasQuotaPreferenceQuotaConfig {
	var returns *GoogleCloudQuotasQuotaPreferenceQuotaConfig
	_jsii_.Get(
		j,
		"quotaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) QuotaId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) QuotaIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quotaIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) Timeouts() GoogleCloudQuotasQuotaPreferenceTimeoutsOutputReference {
	var returns GoogleCloudQuotasQuotaPreferenceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_cloud_quotas_quota_preference google_cloud_quotas_quota_preference} Resource.
func NewGoogleCloudQuotasQuotaPreference(scope constructs.Construct, id *string, config *GoogleCloudQuotasQuotaPreferenceConfig) GoogleCloudQuotasQuotaPreference {
	_init_.Initialize()

	if err := validateNewGoogleCloudQuotasQuotaPreferenceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudQuotasQuotaPreference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_cloud_quotas_quota_preference google_cloud_quotas_quota_preference} Resource.
func NewGoogleCloudQuotasQuotaPreference_Override(g GoogleCloudQuotasQuotaPreference, scope constructs.Construct, id *string, config *GoogleCloudQuotasQuotaPreferenceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetContactEmail(val *string) {
	if err := j.validateSetContactEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetDimensions(val *map[string]*string) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetIgnoreSafetyChecks(val *string) {
	if err := j.validateSetIgnoreSafetyChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreSafetyChecks",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetJustification(val *string) {
	if err := j.validateSetJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"justification",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetParent(val *string) {
	if err := j.validateSetParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parent",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetQuotaId(val *string) {
	if err := j.validateSetQuotaIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotaId",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudQuotasQuotaPreference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

// Generates CDKTF code for importing a GoogleCloudQuotasQuotaPreference resource upon running "cdktf plan <stack-name>".
func GoogleCloudQuotasQuotaPreference_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCloudQuotasQuotaPreference_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
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
func GoogleCloudQuotasQuotaPreference_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudQuotasQuotaPreference_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudQuotasQuotaPreference_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudQuotasQuotaPreference_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudQuotasQuotaPreference_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudQuotasQuotaPreference_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudQuotasQuotaPreference_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleCloudQuotasQuotaPreference.GoogleCloudQuotasQuotaPreference",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) PutQuotaConfig(value *GoogleCloudQuotasQuotaPreferenceQuotaConfig) {
	if err := g.validatePutQuotaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQuotaConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) PutTimeouts(value *GoogleCloudQuotasQuotaPreferenceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetContactEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetContactEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetDimensions() {
	_jsii_.InvokeVoid(
		g,
		"resetDimensions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetIgnoreSafetyChecks() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreSafetyChecks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetJustification() {
	_jsii_.InvokeVoid(
		g,
		"resetJustification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetParent() {
	_jsii_.InvokeVoid(
		g,
		"resetParent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetQuotaId() {
	_jsii_.InvokeVoid(
		g,
		"resetQuotaId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetService() {
	_jsii_.InvokeVoid(
		g,
		"resetService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudQuotasQuotaPreference) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

