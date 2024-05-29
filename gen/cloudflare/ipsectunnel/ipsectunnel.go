package ipsectunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/ipsectunnel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ipsec_tunnel cloudflare_ipsec_tunnel}.
type IpsecTunnel interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowNullCipher() interface{}
	SetAllowNullCipher(val interface{})
	AllowNullCipherInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudflareEndpoint() *string
	SetCloudflareEndpoint(val *string)
	CloudflareEndpointInput() *string
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
	CustomerEndpoint() *string
	SetCustomerEndpoint(val *string)
	CustomerEndpointInput() *string
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
	FqdnId() *string
	SetFqdnId(val *string)
	FqdnIdInput() *string
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
	HexId() *string
	SetHexId(val *string)
	HexIdInput() *string
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
	Psk() *string
	SetPsk(val *string)
	PskInput() *string
	// Experimental.
	RawOverrides() interface{}
	RemoteId() *string
	SetRemoteId(val *string)
	RemoteIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
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
	ResetAccountId()
	ResetAllowNullCipher()
	ResetDescription()
	ResetFqdnId()
	ResetHealthCheckEnabled()
	ResetHealthCheckTarget()
	ResetHealthCheckType()
	ResetHexId()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPsk()
	ResetRemoteId()
	ResetUserId()
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

// The jsii proxy struct for IpsecTunnel
type jsiiProxy_IpsecTunnel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IpsecTunnel) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) AllowNullCipher() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNullCipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) AllowNullCipherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNullCipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) CloudflareEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) CloudflareEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudflareEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) CustomerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) CustomerEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) FqdnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) FqdnIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdnIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) HexIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hexIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) InterfaceAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) InterfaceAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) Psk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) PskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) RemoteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) RemoteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpsecTunnel) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ipsec_tunnel cloudflare_ipsec_tunnel} Resource.
func NewIpsecTunnel(scope constructs.Construct, id *string, config *IpsecTunnelConfig) IpsecTunnel {
	_init_.Initialize()

	if err := validateNewIpsecTunnelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpsecTunnel{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/ipsec_tunnel cloudflare_ipsec_tunnel} Resource.
func NewIpsecTunnel_Override(i IpsecTunnel, scope constructs.Construct, id *string, config *IpsecTunnelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetAllowNullCipher(val interface{}) {
	if err := j.validateSetAllowNullCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowNullCipher",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetCloudflareEndpoint(val *string) {
	if err := j.validateSetCloudflareEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudflareEndpoint",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetCustomerEndpoint(val *string) {
	if err := j.validateSetCustomerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerEndpoint",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetFqdnId(val *string) {
	if err := j.validateSetFqdnIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fqdnId",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetHealthCheckEnabled(val interface{}) {
	if err := j.validateSetHealthCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetHealthCheckTarget(val *string) {
	if err := j.validateSetHealthCheckTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckTarget",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetHexId(val *string) {
	if err := j.validateSetHexIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hexId",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetInterfaceAddress(val *string) {
	if err := j.validateSetInterfaceAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceAddress",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetPsk(val *string) {
	if err := j.validateSetPskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"psk",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetRemoteId(val *string) {
	if err := j.validateSetRemoteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteId",
		val,
	)
}

func (j *jsiiProxy_IpsecTunnel)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Generates CDKTF code for importing a IpsecTunnel resource upon running "cdktf plan <stack-name>".
func IpsecTunnel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIpsecTunnel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
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
func IpsecTunnel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpsecTunnel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IpsecTunnel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpsecTunnel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IpsecTunnel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpsecTunnel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IpsecTunnel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.ipsecTunnel.IpsecTunnel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IpsecTunnel) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IpsecTunnel) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IpsecTunnel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IpsecTunnel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IpsecTunnel) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IpsecTunnel) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IpsecTunnel) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetAllowNullCipher() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowNullCipher",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetFqdnId() {
	_jsii_.InvokeVoid(
		i,
		"resetFqdnId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetHealthCheckEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetHealthCheckEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetHealthCheckTarget() {
	_jsii_.InvokeVoid(
		i,
		"resetHealthCheckTarget",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		i,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetHexId() {
	_jsii_.InvokeVoid(
		i,
		"resetHexId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetPsk() {
	_jsii_.InvokeVoid(
		i,
		"resetPsk",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetRemoteId() {
	_jsii_.InvokeVoid(
		i,
		"resetRemoteId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) ResetUserId() {
	_jsii_.InvokeVoid(
		i,
		"resetUserId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IpsecTunnel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpsecTunnel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

