package connectinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/connectinstance/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/connect_instance aws_connect_instance}.
type ConnectInstance interface {
	cdktf.TerraformResource
	Arn() *string
	AutoResolveBestVoicesEnabled() interface{}
	SetAutoResolveBestVoicesEnabled(val interface{})
	AutoResolveBestVoicesEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactFlowLogsEnabled() interface{}
	SetContactFlowLogsEnabled(val interface{})
	ContactFlowLogsEnabledInput() interface{}
	ContactLensEnabled() interface{}
	SetContactLensEnabled(val interface{})
	ContactLensEnabledInput() interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreatedTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	EarlyMediaEnabled() interface{}
	SetEarlyMediaEnabled(val interface{})
	EarlyMediaEnabledInput() interface{}
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
	IdentityManagementType() *string
	SetIdentityManagementType(val *string)
	IdentityManagementTypeInput() *string
	IdInput() *string
	InboundCallsEnabled() interface{}
	SetInboundCallsEnabled(val interface{})
	InboundCallsEnabledInput() interface{}
	InstanceAlias() *string
	SetInstanceAlias(val *string)
	InstanceAliasInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiPartyConferenceEnabled() interface{}
	SetMultiPartyConferenceEnabled(val interface{})
	MultiPartyConferenceEnabledInput() interface{}
	// The tree node.
	Node() constructs.Node
	OutboundCallsEnabled() interface{}
	SetOutboundCallsEnabled(val interface{})
	OutboundCallsEnabledInput() interface{}
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
	ServiceRole() *string
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ConnectInstanceTimeoutsOutputReference
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
	PutTimeouts(value *ConnectInstanceTimeouts)
	ResetAutoResolveBestVoicesEnabled()
	ResetContactFlowLogsEnabled()
	ResetContactLensEnabled()
	ResetDirectoryId()
	ResetEarlyMediaEnabled()
	ResetId()
	ResetInstanceAlias()
	ResetMultiPartyConferenceEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ConnectInstance
type jsiiProxy_ConnectInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ConnectInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) AutoResolveBestVoicesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoicesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) AutoResolveBestVoicesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoResolveBestVoicesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ContactFlowLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactFlowLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ContactFlowLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactFlowLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ContactLensEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLensEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ContactLensEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactLensEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) EarlyMediaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMediaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) EarlyMediaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"earlyMediaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) IdentityManagementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityManagementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) IdentityManagementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityManagementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) InboundCallsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCallsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) InboundCallsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inboundCallsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) InstanceAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) InstanceAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) MultiPartyConferenceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiPartyConferenceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) MultiPartyConferenceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiPartyConferenceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) OutboundCallsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCallsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) OutboundCallsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundCallsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) Timeouts() ConnectInstanceTimeoutsOutputReference {
	var returns ConnectInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConnectInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/connect_instance aws_connect_instance} Resource.
func NewConnectInstance(scope constructs.Construct, id *string, config *ConnectInstanceConfig) ConnectInstance {
	_init_.Initialize()

	if err := validateNewConnectInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectInstance{}

	_jsii_.Create(
		"aws.connectInstance.ConnectInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/connect_instance aws_connect_instance} Resource.
func NewConnectInstance_Override(c ConnectInstance, scope constructs.Construct, id *string, config *ConnectInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.connectInstance.ConnectInstance",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ConnectInstance)SetAutoResolveBestVoicesEnabled(val interface{}) {
	if err := j.validateSetAutoResolveBestVoicesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoResolveBestVoicesEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetContactFlowLogsEnabled(val interface{}) {
	if err := j.validateSetContactFlowLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactFlowLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetContactLensEnabled(val interface{}) {
	if err := j.validateSetContactLensEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactLensEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetEarlyMediaEnabled(val interface{}) {
	if err := j.validateSetEarlyMediaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"earlyMediaEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetIdentityManagementType(val *string) {
	if err := j.validateSetIdentityManagementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityManagementType",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetInboundCallsEnabled(val interface{}) {
	if err := j.validateSetInboundCallsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundCallsEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetInstanceAlias(val *string) {
	if err := j.validateSetInstanceAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceAlias",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetMultiPartyConferenceEnabled(val interface{}) {
	if err := j.validateSetMultiPartyConferenceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiPartyConferenceEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetOutboundCallsEnabled(val interface{}) {
	if err := j.validateSetOutboundCallsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCallsEnabled",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ConnectInstance)SetProvisioners(val *[]interface{}) {
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
func ConnectInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConnectInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws.connectInstance.ConnectInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ConnectInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws.connectInstance.ConnectInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ConnectInstance) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ConnectInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ConnectInstance) PutTimeouts(value *ConnectInstanceTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConnectInstance) ResetAutoResolveBestVoicesEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoResolveBestVoicesEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetContactFlowLogsEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetContactFlowLogsEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetContactLensEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetContactLensEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		c,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetEarlyMediaEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEarlyMediaEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetInstanceAlias() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceAlias",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetMultiPartyConferenceEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetMultiPartyConferenceEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConnectInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

