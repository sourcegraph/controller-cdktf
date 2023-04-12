package gretunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/gretunnel/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/cloudflare/r/gre_tunnel cloudflare_gre_tunnel}.
type GreTunnel interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudflareGreEndpoint() *string
	SetCloudflareGreEndpoint(val *string)
	CloudflareGreEndpointInput() *string
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
	CustomerGreEndpoint() *string
	SetCustomerGreEndpoint(val *string)
	CustomerGreEndpointInput() *string
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
	HealthCheckEnabled() interface{}
	SetHealthCheckEnabled(val interface{})
	HealthCheckEnabledInput() interface{}
	HealthCheckTarget() *string
	SetHealthCheckTarget(val *string)
	HealthCheckTargetInput() *string
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterfaceAddress() *string
	SetInterfaceAddress(val *string)
	InterfaceAddressInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	ResetAccountId()
	ResetDescription()
	ResetHealthCheckEnabled()
	ResetHealthCheckTarget()
	ResetHealthCheckType()
	ResetId()
	ResetMtu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GreTunnel
type jsiiProxy_GreTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GreTunnel) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) CloudflareGreEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareGreEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) CloudflareGreEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareGreEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) CustomerGreEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGreEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) CustomerGreEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGreEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) InterfaceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) InterfaceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GreTunnel) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/gre_tunnel cloudflare_gre_tunnel} Resource.
func NewGreTunnel(scope constructs.Construct, id *string, config *GreTunnelConfig) GreTunnel {
	_init_.Initialize()

	if err := validateNewGreTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GreTunnel{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/cloudflare/r/gre_tunnel cloudflare_gre_tunnel} Resource.
func NewGreTunnel_Override(g GreTunnel, scope constructs.Construct, id *string, config *GreTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GreTunnel)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetCloudflareGreEndpoint(val *string) {
	if err := j.validateSetCloudflareGreEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudflareGreEndpoint",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetCustomerGreEndpoint(val *string) {
	if err := j.validateSetCustomerGreEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGreEndpoint",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetHealthCheckEnabled(val interface{}) {
	if err := j.validateSetHealthCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetHealthCheckTarget(val *string) {
	if err := j.validateSetHealthCheckTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckTarget",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetInterfaceAddress(val *string) {
	if err := j.validateSetInterfaceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceAddress",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GreTunnel)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
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
func GreTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GreTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GreTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGreTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GreTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.greTunnel.GreTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GreTunnel) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GreTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GreTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GreTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GreTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GreTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GreTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GreTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GreTunnel) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GreTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GreTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GreTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GreTunnel) ResetAccountId() {
	_jsii_.InvokeVoid(
		g,
		"resetAccountId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetHealthCheckEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetHealthCheckTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetMtu() {
	_jsii_.InvokeVoid(
		g,
		"resetMtu",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) ResetTtl() {
	_jsii_.InvokeVoid(
		g,
		"resetTtl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GreTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GreTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GreTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GreTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

