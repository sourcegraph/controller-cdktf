package googlebinaryauthorizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebinaryauthorizationpolicy/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_policy google_binary_authorization_policy}.
type GoogleBinaryAuthorizationPolicy interface {
	cdktf.TerraformResource
	AdmissionWhitelistPatterns() GoogleBinaryAuthorizationPolicyAdmissionWhitelistPatternsList
	AdmissionWhitelistPatternsInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterAdmissionRules() GoogleBinaryAuthorizationPolicyClusterAdmissionRulesList
	ClusterAdmissionRulesInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultAdmissionRule() GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference
	DefaultAdmissionRuleInput() *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule
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
	GlobalPolicyEvaluationMode() *string
	SetGlobalPolicyEvaluationMode(val *string)
	GlobalPolicyEvaluationModeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBinaryAuthorizationPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAdmissionWhitelistPatterns(value interface{})
	PutClusterAdmissionRules(value interface{})
	PutDefaultAdmissionRule(value *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule)
	PutTimeouts(value *GoogleBinaryAuthorizationPolicyTimeouts)
	ResetAdmissionWhitelistPatterns()
	ResetClusterAdmissionRules()
	ResetDescription()
	ResetGlobalPolicyEvaluationMode()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for GoogleBinaryAuthorizationPolicy
type jsiiProxy_GoogleBinaryAuthorizationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) AdmissionWhitelistPatterns() GoogleBinaryAuthorizationPolicyAdmissionWhitelistPatternsList {
	var returns GoogleBinaryAuthorizationPolicyAdmissionWhitelistPatternsList
	_jsii_.Get(
		j,
		"admissionWhitelistPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) AdmissionWhitelistPatternsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"admissionWhitelistPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) ClusterAdmissionRules() GoogleBinaryAuthorizationPolicyClusterAdmissionRulesList {
	var returns GoogleBinaryAuthorizationPolicyClusterAdmissionRulesList
	_jsii_.Get(
		j,
		"clusterAdmissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) ClusterAdmissionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterAdmissionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) DefaultAdmissionRule() GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference {
	var returns GoogleBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference
	_jsii_.Get(
		j,
		"defaultAdmissionRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) DefaultAdmissionRuleInput() *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule {
	var returns *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule
	_jsii_.Get(
		j,
		"defaultAdmissionRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) GlobalPolicyEvaluationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalPolicyEvaluationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) GlobalPolicyEvaluationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalPolicyEvaluationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) Timeouts() GoogleBinaryAuthorizationPolicyTimeoutsOutputReference {
	var returns GoogleBinaryAuthorizationPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_policy google_binary_authorization_policy} Resource.
func NewGoogleBinaryAuthorizationPolicy(scope constructs.Construct, id *string, config *GoogleBinaryAuthorizationPolicyConfig) GoogleBinaryAuthorizationPolicy {
	_init_.Initialize()

	if err := validateNewGoogleBinaryAuthorizationPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBinaryAuthorizationPolicy{}

	_jsii_.Create(
		"google-beta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_binary_authorization_policy google_binary_authorization_policy} Resource.
func NewGoogleBinaryAuthorizationPolicy_Override(g GoogleBinaryAuthorizationPolicy, scope constructs.Construct, id *string, config *GoogleBinaryAuthorizationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicy",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetGlobalPolicyEvaluationMode(val *string) {
	if err := j.validateSetGlobalPolicyEvaluationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalPolicyEvaluationMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBinaryAuthorizationPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleBinaryAuthorizationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBinaryAuthorizationPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBinaryAuthorizationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleBinaryAuthorizationPolicy.GoogleBinaryAuthorizationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) PutAdmissionWhitelistPatterns(value interface{}) {
	if err := g.validatePutAdmissionWhitelistPatternsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdmissionWhitelistPatterns",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) PutClusterAdmissionRules(value interface{}) {
	if err := g.validatePutClusterAdmissionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterAdmissionRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) PutDefaultAdmissionRule(value *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule) {
	if err := g.validatePutDefaultAdmissionRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultAdmissionRule",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) PutTimeouts(value *GoogleBinaryAuthorizationPolicyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetAdmissionWhitelistPatterns() {
	_jsii_.InvokeVoid(
		g,
		"resetAdmissionWhitelistPatterns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetClusterAdmissionRules() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterAdmissionRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetGlobalPolicyEvaluationMode() {
	_jsii_.InvokeVoid(
		g,
		"resetGlobalPolicyEvaluationMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBinaryAuthorizationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
