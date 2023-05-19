package googleorganizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleorganizationpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_organization_policy google_organization_policy}.
type GoogleOrganizationPolicy interface {
	cdktf.TerraformResource
	BooleanPolicy() GoogleOrganizationPolicyBooleanPolicyOutputReference
	BooleanPolicyInput() *GoogleOrganizationPolicyBooleanPolicy
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Constraint() *string
	SetConstraint(val *string)
	ConstraintInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListPolicy() GoogleOrganizationPolicyListPolicyOutputReference
	ListPolicyInput() *GoogleOrganizationPolicyListPolicy
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RestorePolicy() GoogleOrganizationPolicyRestorePolicyOutputReference
	RestorePolicyInput() *GoogleOrganizationPolicyRestorePolicy
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleOrganizationPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	Version() *float64
	SetVersion(val *float64)
	VersionInput() *float64
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
	PutBooleanPolicy(value *GoogleOrganizationPolicyBooleanPolicy)
	PutListPolicy(value *GoogleOrganizationPolicyListPolicy)
	PutRestorePolicy(value *GoogleOrganizationPolicyRestorePolicy)
	PutTimeouts(value *GoogleOrganizationPolicyTimeouts)
	ResetBooleanPolicy()
	ResetId()
	ResetListPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRestorePolicy()
	ResetTimeouts()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleOrganizationPolicy
type jsiiProxy_GoogleOrganizationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleOrganizationPolicy) BooleanPolicy() GoogleOrganizationPolicyBooleanPolicyOutputReference {
	var returns GoogleOrganizationPolicyBooleanPolicyOutputReference
	_jsii_.Get(
		j,
		"booleanPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) BooleanPolicyInput() *GoogleOrganizationPolicyBooleanPolicy {
	var returns *GoogleOrganizationPolicyBooleanPolicy
	_jsii_.Get(
		j,
		"booleanPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Constraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) ConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) ListPolicy() GoogleOrganizationPolicyListPolicyOutputReference {
	var returns GoogleOrganizationPolicyListPolicyOutputReference
	_jsii_.Get(
		j,
		"listPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) ListPolicyInput() *GoogleOrganizationPolicyListPolicy {
	var returns *GoogleOrganizationPolicyListPolicy
	_jsii_.Get(
		j,
		"listPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) RestorePolicy() GoogleOrganizationPolicyRestorePolicyOutputReference {
	var returns GoogleOrganizationPolicyRestorePolicyOutputReference
	_jsii_.Get(
		j,
		"restorePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) RestorePolicyInput() *GoogleOrganizationPolicyRestorePolicy {
	var returns *GoogleOrganizationPolicyRestorePolicy
	_jsii_.Get(
		j,
		"restorePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Timeouts() GoogleOrganizationPolicyTimeoutsOutputReference {
	var returns GoogleOrganizationPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationPolicy) VersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_organization_policy google_organization_policy} Resource.
func NewGoogleOrganizationPolicy(scope constructs.Construct, id *string, config *GoogleOrganizationPolicyConfig) GoogleOrganizationPolicy {
	_init_.Initialize()

	if err := validateNewGoogleOrganizationPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOrganizationPolicy{}

	_jsii_.Create(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_organization_policy google_organization_policy} Resource.
func NewGoogleOrganizationPolicy_Override(g GoogleOrganizationPolicy, scope constructs.Construct, id *string, config *GoogleOrganizationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetConstraint(val *string) {
	if err := j.validateSetConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"constraint",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationPolicy)SetVersion(val *float64) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
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
func GoogleOrganizationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOrganizationPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOrganizationPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleOrganizationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleOrganizationPolicy.GoogleOrganizationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleOrganizationPolicy) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrganizationPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) PutBooleanPolicy(value *GoogleOrganizationPolicyBooleanPolicy) {
	if err := g.validatePutBooleanPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBooleanPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) PutListPolicy(value *GoogleOrganizationPolicyListPolicy) {
	if err := g.validatePutListPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putListPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) PutRestorePolicy(value *GoogleOrganizationPolicyRestorePolicy) {
	if err := g.validatePutRestorePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestorePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) PutTimeouts(value *GoogleOrganizationPolicyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetBooleanPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetBooleanPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetListPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetListPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetRestorePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRestorePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

