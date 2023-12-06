package policyruleidpdiscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/policyruleidpdiscovery/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_idp_discovery okta_policy_rule_idp_discovery}.
type PolicyRuleIdpDiscovery interface {
	cdktf.TerraformResource
	AppExclude() PolicyRuleIdpDiscoveryAppExcludeList
	AppExcludeInput() interface{}
	AppInclude() PolicyRuleIdpDiscoveryAppIncludeList
	AppIncludeInput() interface{}
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
	IdpId() *string
	SetIdpId(val *string)
	IdpIdInput() *string
	IdpType() *string
	SetIdpType(val *string)
	IdpTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConnection() *string
	SetNetworkConnection(val *string)
	NetworkConnectionInput() *string
	NetworkExcludes() *[]*string
	SetNetworkExcludes(val *[]*string)
	NetworkExcludesInput() *[]*string
	NetworkIncludes() *[]*string
	SetNetworkIncludes(val *[]*string)
	NetworkIncludesInput() *[]*string
	// The tree node.
	Node() constructs.Node
	PlatformInclude() PolicyRuleIdpDiscoveryPlatformIncludeList
	PlatformIncludeInput() interface{}
	Policyid() *string
	SetPolicyid(val *string)
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyidInput() *string
	PolicyIdInput() *string
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserIdentifierAttribute() *string
	SetUserIdentifierAttribute(val *string)
	UserIdentifierAttributeInput() *string
	UserIdentifierPatterns() PolicyRuleIdpDiscoveryUserIdentifierPatternsList
	UserIdentifierPatternsInput() interface{}
	UserIdentifierType() *string
	SetUserIdentifierType(val *string)
	UserIdentifierTypeInput() *string
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
	PutAppExclude(value interface{})
	PutAppInclude(value interface{})
	PutPlatformInclude(value interface{})
	PutUserIdentifierPatterns(value interface{})
	ResetAppExclude()
	ResetAppInclude()
	ResetId()
	ResetIdpId()
	ResetIdpType()
	ResetNetworkConnection()
	ResetNetworkExcludes()
	ResetNetworkIncludes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlatformInclude()
	ResetPolicyid()
	ResetPolicyId()
	ResetPriority()
	ResetStatus()
	ResetUserIdentifierAttribute()
	ResetUserIdentifierPatterns()
	ResetUserIdentifierType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PolicyRuleIdpDiscovery
type jsiiProxy_PolicyRuleIdpDiscovery struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) AppExclude() PolicyRuleIdpDiscoveryAppExcludeList {
	var returns PolicyRuleIdpDiscoveryAppExcludeList
	_jsii_.Get(
		j,
		"appExclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) AppExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appExcludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) AppInclude() PolicyRuleIdpDiscoveryAppIncludeList {
	var returns PolicyRuleIdpDiscoveryAppIncludeList
	_jsii_.Get(
		j,
		"appInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) AppIncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) IdpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) IdpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) IdpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) IdpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkConnection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkExcludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkExcludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkIncludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) NetworkIncludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkIncludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PlatformInclude() PolicyRuleIdpDiscoveryPlatformIncludeList {
	var returns PolicyRuleIdpDiscoveryPlatformIncludeList
	_jsii_.Get(
		j,
		"platformInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PlatformIncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"platformIncludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Policyid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PolicyidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentifierAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentifierAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierPatterns() PolicyRuleIdpDiscoveryUserIdentifierPatternsList {
	var returns PolicyRuleIdpDiscoveryUserIdentifierPatternsList
	_jsii_.Get(
		j,
		"userIdentifierPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierPatternsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userIdentifierPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentifierType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery) UserIdentifierTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdentifierTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_idp_discovery okta_policy_rule_idp_discovery} Resource.
func NewPolicyRuleIdpDiscovery(scope constructs.Construct, id *string, config *PolicyRuleIdpDiscoveryConfig) PolicyRuleIdpDiscovery {
	_init_.Initialize()

	if err := validateNewPolicyRuleIdpDiscoveryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyRuleIdpDiscovery{}

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/policy_rule_idp_discovery okta_policy_rule_idp_discovery} Resource.
func NewPolicyRuleIdpDiscovery_Override(p PolicyRuleIdpDiscovery, scope constructs.Construct, id *string, config *PolicyRuleIdpDiscoveryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetIdpId(val *string) {
	if err := j.validateSetIdpIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetIdpType(val *string) {
	if err := j.validateSetIdpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpType",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetNetworkConnection(val *string) {
	if err := j.validateSetNetworkConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConnection",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetNetworkExcludes(val *[]*string) {
	if err := j.validateSetNetworkExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkExcludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetNetworkIncludes(val *[]*string) {
	if err := j.validateSetNetworkIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIncludes",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetPolicyid(val *string) {
	if err := j.validateSetPolicyidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyid",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetPolicyId(val *string) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetUserIdentifierAttribute(val *string) {
	if err := j.validateSetUserIdentifierAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIdentifierAttribute",
		val,
	)
}

func (j *jsiiProxy_PolicyRuleIdpDiscovery)SetUserIdentifierType(val *string) {
	if err := j.validateSetUserIdentifierTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIdentifierType",
		val,
	)
}

// Generates CDKTF code for importing a PolicyRuleIdpDiscovery resource upon running "cdktf plan <stack-name>".
func PolicyRuleIdpDiscovery_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePolicyRuleIdpDiscovery_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
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
func PolicyRuleIdpDiscovery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleIdpDiscovery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleIdpDiscovery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleIdpDiscovery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PolicyRuleIdpDiscovery_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePolicyRuleIdpDiscovery_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PolicyRuleIdpDiscovery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.policyRuleIdpDiscovery.PolicyRuleIdpDiscovery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) PutAppExclude(value interface{}) {
	if err := p.validatePutAppExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAppExclude",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) PutAppInclude(value interface{}) {
	if err := p.validatePutAppIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAppInclude",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) PutPlatformInclude(value interface{}) {
	if err := p.validatePutPlatformIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPlatformInclude",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) PutUserIdentifierPatterns(value interface{}) {
	if err := p.validatePutUserIdentifierPatternsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUserIdentifierPatterns",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetAppExclude() {
	_jsii_.InvokeVoid(
		p,
		"resetAppExclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetAppInclude() {
	_jsii_.InvokeVoid(
		p,
		"resetAppInclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetIdpId() {
	_jsii_.InvokeVoid(
		p,
		"resetIdpId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetIdpType() {
	_jsii_.InvokeVoid(
		p,
		"resetIdpType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetNetworkConnection() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkConnection",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetNetworkExcludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkExcludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetNetworkIncludes() {
	_jsii_.InvokeVoid(
		p,
		"resetNetworkIncludes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetPlatformInclude() {
	_jsii_.InvokeVoid(
		p,
		"resetPlatformInclude",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetPolicyid() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyid",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetPolicyId() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetPriority() {
	_jsii_.InvokeVoid(
		p,
		"resetPriority",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetStatus() {
	_jsii_.InvokeVoid(
		p,
		"resetStatus",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetUserIdentifierAttribute() {
	_jsii_.InvokeVoid(
		p,
		"resetUserIdentifierAttribute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetUserIdentifierPatterns() {
	_jsii_.InvokeVoid(
		p,
		"resetUserIdentifierPatterns",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ResetUserIdentifierType() {
	_jsii_.InvokeVoid(
		p,
		"resetUserIdentifierType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyRuleIdpDiscovery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

