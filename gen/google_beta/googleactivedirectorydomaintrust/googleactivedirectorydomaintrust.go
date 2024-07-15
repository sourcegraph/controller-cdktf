package googleactivedirectorydomaintrust

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleactivedirectorydomaintrust/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain_trust google_active_directory_domain_trust}.
type GoogleActiveDirectoryDomainTrust interface {
	cdktf.TerraformResource
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
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
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
	SelectiveAuthentication() interface{}
	SetSelectiveAuthentication(val interface{})
	SelectiveAuthenticationInput() interface{}
	TargetDnsIpAddresses() *[]*string
	SetTargetDnsIpAddresses(val *[]*string)
	TargetDnsIpAddressesInput() *[]*string
	TargetDomainName() *string
	SetTargetDomainName(val *string)
	TargetDomainNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleActiveDirectoryDomainTrustTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustDirection() *string
	SetTrustDirection(val *string)
	TrustDirectionInput() *string
	TrustHandshakeSecret() *string
	SetTrustHandshakeSecret(val *string)
	TrustHandshakeSecretInput() *string
	TrustType() *string
	SetTrustType(val *string)
	TrustTypeInput() *string
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
	PutTimeouts(value *GoogleActiveDirectoryDomainTrustTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSelectiveAuthentication()
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

// The jsii proxy struct for GoogleActiveDirectoryDomainTrust
type jsiiProxy_GoogleActiveDirectoryDomainTrust struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) SelectiveAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectiveAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) SelectiveAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectiveAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TargetDnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetDnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TargetDnsIpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetDnsIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TargetDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TargetDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) Timeouts() GoogleActiveDirectoryDomainTrustTimeoutsOutputReference {
	var returns GoogleActiveDirectoryDomainTrustTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustHandshakeSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustHandshakeSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustHandshakeSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustHandshakeSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust) TrustTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain_trust google_active_directory_domain_trust} Resource.
func NewGoogleActiveDirectoryDomainTrust(scope constructs.Construct, id *string, config *GoogleActiveDirectoryDomainTrustConfig) GoogleActiveDirectoryDomainTrust {
	_init_.Initialize()

	if err := validateNewGoogleActiveDirectoryDomainTrustParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleActiveDirectoryDomainTrust{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_active_directory_domain_trust google_active_directory_domain_trust} Resource.
func NewGoogleActiveDirectoryDomainTrust_Override(g GoogleActiveDirectoryDomainTrust, scope constructs.Construct, id *string, config *GoogleActiveDirectoryDomainTrustConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetSelectiveAuthentication(val interface{}) {
	if err := j.validateSetSelectiveAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectiveAuthentication",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetTargetDnsIpAddresses(val *[]*string) {
	if err := j.validateSetTargetDnsIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDnsIpAddresses",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetTargetDomainName(val *string) {
	if err := j.validateSetTargetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDomainName",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetTrustDirection(val *string) {
	if err := j.validateSetTrustDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustDirection",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetTrustHandshakeSecret(val *string) {
	if err := j.validateSetTrustHandshakeSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustHandshakeSecret",
		val,
	)
}

func (j *jsiiProxy_GoogleActiveDirectoryDomainTrust)SetTrustType(val *string) {
	if err := j.validateSetTrustTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustType",
		val,
	)
}

// Generates CDKTF code for importing a GoogleActiveDirectoryDomainTrust resource upon running "cdktf plan <stack-name>".
func GoogleActiveDirectoryDomainTrust_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleActiveDirectoryDomainTrust_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
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
func GoogleActiveDirectoryDomainTrust_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleActiveDirectoryDomainTrust_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleActiveDirectoryDomainTrust_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleActiveDirectoryDomainTrust_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleActiveDirectoryDomainTrust_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleActiveDirectoryDomainTrust_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleActiveDirectoryDomainTrust_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleActiveDirectoryDomainTrust.GoogleActiveDirectoryDomainTrust",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) PutTimeouts(value *GoogleActiveDirectoryDomainTrustTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ResetSelectiveAuthentication() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectiveAuthentication",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleActiveDirectoryDomainTrust) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

