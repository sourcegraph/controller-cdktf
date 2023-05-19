package accesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accesspolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy cloudflare_access_policy}.
type AccessPolicy interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	ApprovalGroup() AccessPolicyApprovalGroupList
	ApprovalGroupInput() interface{}
	ApprovalRequired() interface{}
	SetApprovalRequired(val interface{})
	ApprovalRequiredInput() interface{}
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
	Decision() *string
	SetDecision(val *string)
	DecisionInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Exclude() AccessPolicyExcludeList
	ExcludeInput() interface{}
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
	Include() AccessPolicyIncludeList
	IncludeInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurposeJustificationPrompt() *string
	SetPurposeJustificationPrompt(val *string)
	PurposeJustificationPromptInput() *string
	PurposeJustificationRequired() interface{}
	SetPurposeJustificationRequired(val interface{})
	PurposeJustificationRequiredInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	Require() AccessPolicyRequireList
	RequireInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutApprovalGroup(value interface{})
	PutExclude(value interface{})
	PutInclude(value interface{})
	PutRequire(value interface{})
	ResetAccountId()
	ResetApprovalGroup()
	ResetApprovalRequired()
	ResetExclude()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurposeJustificationPrompt()
	ResetPurposeJustificationRequired()
	ResetRequire()
	ResetZoneId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AccessPolicy
type jsiiProxy_AccessPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccessPolicy) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApprovalGroup() AccessPolicyApprovalGroupList {
	var returns AccessPolicyApprovalGroupList
	_jsii_.Get(
		j,
		"approvalGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApprovalGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApprovalRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ApprovalRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"approvalRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Decision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) DecisionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Exclude() AccessPolicyExcludeList {
	var returns AccessPolicyExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Include() AccessPolicyIncludeList {
	var returns AccessPolicyIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) IncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) PurposeJustificationPrompt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) PurposeJustificationPromptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purposeJustificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) PurposeJustificationRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purposeJustificationRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) PurposeJustificationRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purposeJustificationRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) Require() AccessPolicyRequireList {
	var returns AccessPolicyRequireList
	_jsii_.Get(
		j,
		"require",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) RequireInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPolicy) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy cloudflare_access_policy} Resource.
func NewAccessPolicy(scope constructs.Construct, id *string, config *AccessPolicyConfig) AccessPolicy {
	_init_.Initialize()

	if err := validateNewAccessPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessPolicy{}

	_jsii_.Create(
		"cloudflare.accessPolicy.AccessPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_policy cloudflare_access_policy} Resource.
func NewAccessPolicy_Override(a AccessPolicy, scope constructs.Construct, id *string, config *AccessPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.accessPolicy.AccessPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccessPolicy)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetApprovalRequired(val interface{}) {
	if err := j.validateSetApprovalRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetDecision(val *string) {
	if err := j.validateSetDecisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decision",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetPurposeJustificationPrompt(val *string) {
	if err := j.validateSetPurposeJustificationPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationPrompt",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetPurposeJustificationRequired(val interface{}) {
	if err := j.validateSetPurposeJustificationRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purposeJustificationRequired",
		val,
	)
}

func (j *jsiiProxy_AccessPolicy)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
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
func AccessPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessPolicy.AccessPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessPolicy.AccessPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessPolicy.AccessPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccessPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.accessPolicy.AccessPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccessPolicy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccessPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccessPolicy) PutApprovalGroup(value interface{}) {
	if err := a.validatePutApprovalGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApprovalGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicy) PutExclude(value interface{}) {
	if err := a.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicy) PutInclude(value interface{}) {
	if err := a.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInclude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicy) PutRequire(value interface{}) {
	if err := a.validatePutRequireParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequire",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessPolicy) ResetAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetApprovalGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetApprovalRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetApprovalRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetExclude() {
	_jsii_.InvokeVoid(
		a,
		"resetExclude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetPurposeJustificationPrompt() {
	_jsii_.InvokeVoid(
		a,
		"resetPurposeJustificationPrompt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetPurposeJustificationRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetPurposeJustificationRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetRequire() {
	_jsii_.InvokeVoid(
		a,
		"resetRequire",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) ResetZoneId() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

