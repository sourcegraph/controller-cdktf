package cloudassetorganizationfeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cloudassetorganizationfeed/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloud_asset_organization_feed google_cloud_asset_organization_feed}.
type CloudAssetOrganizationFeed interface {
	cdktf.TerraformResource
	AssetNames() *[]*string
	SetAssetNames(val *[]*string)
	AssetNamesInput() *[]*string
	AssetTypes() *[]*string
	SetAssetTypes(val *[]*string)
	AssetTypesInput() *[]*string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() CloudAssetOrganizationFeedConditionOutputReference
	ConditionInput() *CloudAssetOrganizationFeedCondition
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FeedId() *string
	SetFeedId(val *string)
	FeedIdInput() *string
	FeedOutputConfig() CloudAssetOrganizationFeedFeedOutputConfigOutputReference
	FeedOutputConfigInput() *CloudAssetOrganizationFeedFeedOutputConfig
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	OrgId() *string
	SetOrgId(val *string)
	OrgIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CloudAssetOrganizationFeedTimeoutsOutputReference
	TimeoutsInput() any
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
	PutCondition(value *CloudAssetOrganizationFeedCondition)
	PutFeedOutputConfig(value *CloudAssetOrganizationFeedFeedOutputConfig)
	PutTimeouts(value *CloudAssetOrganizationFeedTimeouts)
	ResetAssetNames()
	ResetAssetTypes()
	ResetCondition()
	ResetContentType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
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

// The jsii proxy struct for CloudAssetOrganizationFeed
type jsiiProxy_CloudAssetOrganizationFeed struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) AssetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) AssetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) AssetTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) AssetTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Condition() CloudAssetOrganizationFeedConditionOutputReference {
	var returns CloudAssetOrganizationFeedConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) ConditionInput() *CloudAssetOrganizationFeedCondition {
	var returns *CloudAssetOrganizationFeedCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) FeedId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) FeedIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) FeedOutputConfig() CloudAssetOrganizationFeedFeedOutputConfigOutputReference {
	var returns CloudAssetOrganizationFeedFeedOutputConfigOutputReference
	_jsii_.Get(
		j,
		"feedOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) FeedOutputConfigInput() *CloudAssetOrganizationFeedFeedOutputConfig {
	var returns *CloudAssetOrganizationFeedFeedOutputConfig
	_jsii_.Get(
		j,
		"feedOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) Timeouts() CloudAssetOrganizationFeedTimeoutsOutputReference {
	var returns CloudAssetOrganizationFeedTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloud_asset_organization_feed google_cloud_asset_organization_feed} Resource.
func NewCloudAssetOrganizationFeed(scope constructs.Construct, id *string, config *CloudAssetOrganizationFeedConfig) CloudAssetOrganizationFeed {
	_init_.Initialize()

	if err := validateNewCloudAssetOrganizationFeedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAssetOrganizationFeed{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/cloud_asset_organization_feed google_cloud_asset_organization_feed} Resource.
func NewCloudAssetOrganizationFeed_Override(c CloudAssetOrganizationFeed, scope constructs.Construct, id *string, config *CloudAssetOrganizationFeedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		[]any{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetAssetNames(val *[]*string) {
	if err := j.validateSetAssetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNames",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetAssetTypes(val *[]*string) {
	if err := j.validateSetAssetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetTypes",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetBillingProject(val *string) {
	if err := j.validateSetBillingProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetFeedId(val *string) {
	if err := j.validateSetFeedIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedId",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAssetOrganizationFeed) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CloudAssetOrganizationFeed resource upon running "cdktf plan <stack-name>".
func CloudAssetOrganizationFeed_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudAssetOrganizationFeed_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
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
func CloudAssetOrganizationFeed_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateCloudAssetOrganizationFeed_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAssetOrganizationFeed_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateCloudAssetOrganizationFeed_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAssetOrganizationFeed_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateCloudAssetOrganizationFeed_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func CloudAssetOrganizationFeed_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cloudAssetOrganizationFeed.CloudAssetOrganizationFeed",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) AddOverride(path *string, value any) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]any{path, value},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]any{id, provider},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAssetOrganizationFeed) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]any{id},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) MoveTo(moveTarget *string, index any) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]any{id},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) PutCondition(value *CloudAssetOrganizationFeedCondition) {
	if err := c.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCondition",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) PutFeedOutputConfig(value *CloudAssetOrganizationFeedFeedOutputConfig) {
	if err := c.validatePutFeedOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFeedOutputConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) PutTimeouts(value *CloudAssetOrganizationFeedTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetAssetNames() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetAssetTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetAssetTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetContentType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssetOrganizationFeed) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
