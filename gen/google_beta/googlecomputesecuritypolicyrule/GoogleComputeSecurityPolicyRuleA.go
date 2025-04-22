package googlecomputesecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputesecuritypolicyrule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_security_policy_rule google_compute_security_policy_rule}.
type GoogleComputeSecurityPolicyRuleA interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeaderAction() GoogleComputeSecurityPolicyRuleHeaderActionAOutputReference
	HeaderActionInput() *GoogleComputeSecurityPolicyRuleHeaderActionA
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() GoogleComputeSecurityPolicyRuleMatchAOutputReference
	MatchInput() *GoogleComputeSecurityPolicyRuleMatchA
	// The tree node.
	Node() constructs.Node
	PreconfiguredWafConfig() GoogleComputeSecurityPolicyRulePreconfiguredWafConfigAOutputReference
	PreconfiguredWafConfigInput() *GoogleComputeSecurityPolicyRulePreconfiguredWafConfigA
	Preview() interface{}
	SetPreview(val interface{})
	PreviewInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	RateLimitOptions() GoogleComputeSecurityPolicyRuleRateLimitOptionsAOutputReference
	RateLimitOptionsInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsA
	// Experimental.
	RawOverrides() interface{}
	RedirectOptions() GoogleComputeSecurityPolicyRuleRedirectOptionsAOutputReference
	RedirectOptionsInput() *GoogleComputeSecurityPolicyRuleRedirectOptionsA
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeSecurityPolicyRuleTimeoutsOutputReference
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
	PutHeaderAction(value *GoogleComputeSecurityPolicyRuleHeaderActionA)
	PutMatch(value *GoogleComputeSecurityPolicyRuleMatchA)
	PutPreconfiguredWafConfig(value *GoogleComputeSecurityPolicyRulePreconfiguredWafConfigA)
	PutRateLimitOptions(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsA)
	PutRedirectOptions(value *GoogleComputeSecurityPolicyRuleRedirectOptionsA)
	PutTimeouts(value *GoogleComputeSecurityPolicyRuleTimeouts)
	ResetDescription()
	ResetHeaderAction()
	ResetId()
	ResetMatch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreconfiguredWafConfig()
	ResetPreview()
	ResetProject()
	ResetRateLimitOptions()
	ResetRedirectOptions()
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

// The jsii proxy struct for GoogleComputeSecurityPolicyRuleA
type jsiiProxy_GoogleComputeSecurityPolicyRuleA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) HeaderAction() GoogleComputeSecurityPolicyRuleHeaderActionAOutputReference {
	var returns GoogleComputeSecurityPolicyRuleHeaderActionAOutputReference
	_jsii_.Get(
		j,
		"headerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) HeaderActionInput() *GoogleComputeSecurityPolicyRuleHeaderActionA {
	var returns *GoogleComputeSecurityPolicyRuleHeaderActionA
	_jsii_.Get(
		j,
		"headerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Match() GoogleComputeSecurityPolicyRuleMatchAOutputReference {
	var returns GoogleComputeSecurityPolicyRuleMatchAOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) MatchInput() *GoogleComputeSecurityPolicyRuleMatchA {
	var returns *GoogleComputeSecurityPolicyRuleMatchA
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PreconfiguredWafConfig() GoogleComputeSecurityPolicyRulePreconfiguredWafConfigAOutputReference {
	var returns GoogleComputeSecurityPolicyRulePreconfiguredWafConfigAOutputReference
	_jsii_.Get(
		j,
		"preconfiguredWafConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PreconfiguredWafConfigInput() *GoogleComputeSecurityPolicyRulePreconfiguredWafConfigA {
	var returns *GoogleComputeSecurityPolicyRulePreconfiguredWafConfigA
	_jsii_.Get(
		j,
		"preconfiguredWafConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Preview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"previewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) RateLimitOptions() GoogleComputeSecurityPolicyRuleRateLimitOptionsAOutputReference {
	var returns GoogleComputeSecurityPolicyRuleRateLimitOptionsAOutputReference
	_jsii_.Get(
		j,
		"rateLimitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) RateLimitOptionsInput() *GoogleComputeSecurityPolicyRuleRateLimitOptionsA {
	var returns *GoogleComputeSecurityPolicyRuleRateLimitOptionsA
	_jsii_.Get(
		j,
		"rateLimitOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) RedirectOptions() GoogleComputeSecurityPolicyRuleRedirectOptionsAOutputReference {
	var returns GoogleComputeSecurityPolicyRuleRedirectOptionsAOutputReference
	_jsii_.Get(
		j,
		"redirectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) RedirectOptionsInput() *GoogleComputeSecurityPolicyRuleRedirectOptionsA {
	var returns *GoogleComputeSecurityPolicyRuleRedirectOptionsA
	_jsii_.Get(
		j,
		"redirectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) Timeouts() GoogleComputeSecurityPolicyRuleTimeoutsOutputReference {
	var returns GoogleComputeSecurityPolicyRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_security_policy_rule google_compute_security_policy_rule} Resource.
func NewGoogleComputeSecurityPolicyRuleA(scope constructs.Construct, id *string, config *GoogleComputeSecurityPolicyRuleAConfig) GoogleComputeSecurityPolicyRuleA {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyRuleAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyRuleA{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_compute_security_policy_rule google_compute_security_policy_rule} Resource.
func NewGoogleComputeSecurityPolicyRuleA_Override(g GoogleComputeSecurityPolicyRuleA, scope constructs.Construct, id *string, config *GoogleComputeSecurityPolicyRuleAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetPreview(val interface{}) {
	if err := j.validateSetPreviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preview",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyRuleA)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeSecurityPolicyRuleA resource upon running "cdktf plan <stack-name>".
func GoogleComputeSecurityPolicyRuleA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeSecurityPolicyRuleA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
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
func GoogleComputeSecurityPolicyRuleA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSecurityPolicyRuleA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeSecurityPolicyRuleA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSecurityPolicyRuleA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeSecurityPolicyRuleA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeSecurityPolicyRuleA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeSecurityPolicyRuleA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleComputeSecurityPolicyRule.GoogleComputeSecurityPolicyRuleA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutHeaderAction(value *GoogleComputeSecurityPolicyRuleHeaderActionA) {
	if err := g.validatePutHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaderAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutMatch(value *GoogleComputeSecurityPolicyRuleMatchA) {
	if err := g.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMatch",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutPreconfiguredWafConfig(value *GoogleComputeSecurityPolicyRulePreconfiguredWafConfigA) {
	if err := g.validatePutPreconfiguredWafConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreconfiguredWafConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutRateLimitOptions(value *GoogleComputeSecurityPolicyRuleRateLimitOptionsA) {
	if err := g.validatePutRateLimitOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRateLimitOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutRedirectOptions(value *GoogleComputeSecurityPolicyRuleRedirectOptionsA) {
	if err := g.validatePutRedirectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRedirectOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) PutTimeouts(value *GoogleComputeSecurityPolicyRuleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetHeaderAction() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderAction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetPreconfiguredWafConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPreconfiguredWafConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetPreview() {
	_jsii_.InvokeVoid(
		g,
		"resetPreview",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetRateLimitOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetRateLimitOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetRedirectOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetRedirectOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyRuleA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

