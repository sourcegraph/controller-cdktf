package googlesecuresourcemanagerbranchrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesecuresourcemanagerbranchrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secure_source_manager_branch_rule google_secure_source_manager_branch_rule}.
type GoogleSecureSourceManagerBranchRule interface {
	cdktf.TerraformResource
	AllowStaleReviews() interface{}
	SetAllowStaleReviews(val interface{})
	AllowStaleReviewsInput() interface{}
	BranchRuleId() *string
	SetBranchRuleId(val *string)
	BranchRuleIdInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
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
	IncludePattern() *string
	SetIncludePattern(val *string)
	IncludePatternInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MinimumApprovalsCount() *float64
	SetMinimumApprovalsCount(val *float64)
	MinimumApprovalsCountInput() *float64
	MinimumReviewsCount() *float64
	SetMinimumReviewsCount(val *float64)
	MinimumReviewsCountInput() *float64
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
	RepositoryId() *string
	SetRepositoryId(val *string)
	RepositoryIdInput() *string
	RequireCommentsResolved() interface{}
	SetRequireCommentsResolved(val interface{})
	RequireCommentsResolvedInput() interface{}
	RequireLinearHistory() interface{}
	SetRequireLinearHistory(val interface{})
	RequireLinearHistoryInput() interface{}
	RequirePullRequest() interface{}
	SetRequirePullRequest(val interface{})
	RequirePullRequestInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleSecureSourceManagerBranchRuleTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	PutTimeouts(value *GoogleSecureSourceManagerBranchRuleTimeouts)
	ResetAllowStaleReviews()
	ResetDisabled()
	ResetId()
	ResetMinimumApprovalsCount()
	ResetMinimumReviewsCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRequireCommentsResolved()
	ResetRequireLinearHistory()
	ResetRequirePullRequest()
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

// The jsii proxy struct for GoogleSecureSourceManagerBranchRule
type jsiiProxy_GoogleSecureSourceManagerBranchRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) AllowStaleReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) AllowStaleReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowStaleReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) BranchRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) BranchRuleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchRuleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) IncludePattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) IncludePatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) MinimumApprovalsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) MinimumApprovalsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) MinimumReviewsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumReviewsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) MinimumReviewsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumReviewsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RepositoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RepositoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequireCommentsResolved() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommentsResolved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequireCommentsResolvedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCommentsResolvedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequireLinearHistory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLinearHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequireLinearHistoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLinearHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequirePullRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) RequirePullRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Timeouts() GoogleSecureSourceManagerBranchRuleTimeoutsOutputReference {
	var returns GoogleSecureSourceManagerBranchRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secure_source_manager_branch_rule google_secure_source_manager_branch_rule} Resource.
func NewGoogleSecureSourceManagerBranchRule(scope constructs.Construct, id *string, config *GoogleSecureSourceManagerBranchRuleConfig) GoogleSecureSourceManagerBranchRule {
	_init_.Initialize()

	if err := validateNewGoogleSecureSourceManagerBranchRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSecureSourceManagerBranchRule{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_secure_source_manager_branch_rule google_secure_source_manager_branch_rule} Resource.
func NewGoogleSecureSourceManagerBranchRule_Override(g GoogleSecureSourceManagerBranchRule, scope constructs.Construct, id *string, config *GoogleSecureSourceManagerBranchRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetAllowStaleReviews(val interface{}) {
	if err := j.validateSetAllowStaleReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowStaleReviews",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetBranchRuleId(val *string) {
	if err := j.validateSetBranchRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branchRuleId",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetIncludePattern(val *string) {
	if err := j.validateSetIncludePatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePattern",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetMinimumApprovalsCount(val *float64) {
	if err := j.validateSetMinimumApprovalsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumApprovalsCount",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetMinimumReviewsCount(val *float64) {
	if err := j.validateSetMinimumReviewsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumReviewsCount",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetRepositoryId(val *string) {
	if err := j.validateSetRepositoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryId",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetRequireCommentsResolved(val interface{}) {
	if err := j.validateSetRequireCommentsResolvedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCommentsResolved",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetRequireLinearHistory(val interface{}) {
	if err := j.validateSetRequireLinearHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLinearHistory",
		val,
	)
}

func (j *jsiiProxy_GoogleSecureSourceManagerBranchRule)SetRequirePullRequest(val interface{}) {
	if err := j.validateSetRequirePullRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePullRequest",
		val,
	)
}

// Generates CDKTF code for importing a GoogleSecureSourceManagerBranchRule resource upon running "cdktf plan <stack-name>".
func GoogleSecureSourceManagerBranchRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleSecureSourceManagerBranchRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
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
func GoogleSecureSourceManagerBranchRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecureSourceManagerBranchRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSecureSourceManagerBranchRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecureSourceManagerBranchRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSecureSourceManagerBranchRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSecureSourceManagerBranchRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleSecureSourceManagerBranchRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleSecureSourceManagerBranchRule.GoogleSecureSourceManagerBranchRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) PutTimeouts(value *GoogleSecureSourceManagerBranchRuleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetAllowStaleReviews() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowStaleReviews",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetMinimumApprovalsCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumApprovalsCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetMinimumReviewsCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumReviewsCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetRequireCommentsResolved() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireCommentsResolved",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetRequireLinearHistory() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireLinearHistory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetRequirePullRequest() {
	_jsii_.InvokeVoid(
		g,
		"resetRequirePullRequest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSecureSourceManagerBranchRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

