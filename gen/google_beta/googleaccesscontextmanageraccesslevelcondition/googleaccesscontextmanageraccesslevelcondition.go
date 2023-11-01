package googleaccesscontextmanageraccesslevelcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleaccesscontextmanageraccesslevelcondition/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_access_context_manager_access_level_condition google_access_context_manager_access_level_condition}.
type GoogleAccessContextManagerAccessLevelCondition interface {
	cdktf.TerraformResource
	AccessLevel() *string
	SetAccessLevel(val *string)
	AccessLevelInput() *string
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
	DevicePolicy() GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference
	DevicePolicyInput() *GoogleAccessContextManagerAccessLevelConditionDevicePolicy
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
	IpSubnetworks() *[]*string
	SetIpSubnetworks(val *[]*string)
	IpSubnetworksInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Members() *[]*string
	SetMembers(val *[]*string)
	MembersInput() *[]*string
	Negate() interface{}
	SetNegate(val interface{})
	NegateInput() interface{}
	// The tree node.
	Node() constructs.Node
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
	Regions() *[]*string
	SetRegions(val *[]*string)
	RegionsInput() *[]*string
	RequiredAccessLevels() *[]*string
	SetRequiredAccessLevels(val *[]*string)
	RequiredAccessLevelsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAccessContextManagerAccessLevelConditionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcNetworkSources() GoogleAccessContextManagerAccessLevelConditionVpcNetworkSourcesList
	VpcNetworkSourcesInput() interface{}
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
	PutDevicePolicy(value *GoogleAccessContextManagerAccessLevelConditionDevicePolicy)
	PutTimeouts(value *GoogleAccessContextManagerAccessLevelConditionTimeouts)
	PutVpcNetworkSources(value interface{})
	ResetDevicePolicy()
	ResetId()
	ResetIpSubnetworks()
	ResetMembers()
	ResetNegate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegions()
	ResetRequiredAccessLevels()
	ResetTimeouts()
	ResetVpcNetworkSources()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleAccessContextManagerAccessLevelCondition
type jsiiProxy_GoogleAccessContextManagerAccessLevelCondition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) AccessLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) AccessLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) DevicePolicy() GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference {
	var returns GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference
	_jsii_.Get(
		j,
		"devicePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) DevicePolicyInput() *GoogleAccessContextManagerAccessLevelConditionDevicePolicy {
	var returns *GoogleAccessContextManagerAccessLevelConditionDevicePolicy
	_jsii_.Get(
		j,
		"devicePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) IpSubnetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSubnetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) IpSubnetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSubnetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Members() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) MembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Negate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) NegateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) RequiredAccessLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredAccessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) RequiredAccessLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiredAccessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) Timeouts() GoogleAccessContextManagerAccessLevelConditionTimeoutsOutputReference {
	var returns GoogleAccessContextManagerAccessLevelConditionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) VpcNetworkSources() GoogleAccessContextManagerAccessLevelConditionVpcNetworkSourcesList {
	var returns GoogleAccessContextManagerAccessLevelConditionVpcNetworkSourcesList
	_jsii_.Get(
		j,
		"vpcNetworkSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) VpcNetworkSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcNetworkSourcesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_access_context_manager_access_level_condition google_access_context_manager_access_level_condition} Resource.
func NewGoogleAccessContextManagerAccessLevelCondition(scope constructs.Construct, id *string, config *GoogleAccessContextManagerAccessLevelConditionConfig) GoogleAccessContextManagerAccessLevelCondition {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerAccessLevelConditionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerAccessLevelCondition{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.4.0/docs/resources/google_access_context_manager_access_level_condition google_access_context_manager_access_level_condition} Resource.
func NewGoogleAccessContextManagerAccessLevelCondition_Override(g GoogleAccessContextManagerAccessLevelCondition, scope constructs.Construct, id *string, config *GoogleAccessContextManagerAccessLevelConditionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetAccessLevel(val *string) {
	if err := j.validateSetAccessLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetIpSubnetworks(val *[]*string) {
	if err := j.validateSetIpSubnetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipSubnetworks",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetMembers(val *[]*string) {
	if err := j.validateSetMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"members",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetNegate(val interface{}) {
	if err := j.validateSetNegateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negate",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition)SetRequiredAccessLevels(val *[]*string) {
	if err := j.validateSetRequiredAccessLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredAccessLevels",
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
func GoogleAccessContextManagerAccessLevelCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAccessContextManagerAccessLevelCondition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAccessContextManagerAccessLevelCondition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAccessContextManagerAccessLevelCondition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAccessContextManagerAccessLevelCondition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAccessContextManagerAccessLevelCondition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAccessContextManagerAccessLevelCondition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelCondition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) PutDevicePolicy(value *GoogleAccessContextManagerAccessLevelConditionDevicePolicy) {
	if err := g.validatePutDevicePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDevicePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) PutTimeouts(value *GoogleAccessContextManagerAccessLevelConditionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) PutVpcNetworkSources(value interface{}) {
	if err := g.validatePutVpcNetworkSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcNetworkSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetDevicePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDevicePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetIpSubnetworks() {
	_jsii_.InvokeVoid(
		g,
		"resetIpSubnetworks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetMembers() {
	_jsii_.InvokeVoid(
		g,
		"resetMembers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetNegate() {
	_jsii_.InvokeVoid(
		g,
		"resetNegate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetRequiredAccessLevels() {
	_jsii_.InvokeVoid(
		g,
		"resetRequiredAccessLevels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ResetVpcNetworkSources() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcNetworkSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelCondition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

