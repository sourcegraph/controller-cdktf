package googleiamworkforcepoolproviderscimtenant

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleiamworkforcepoolproviderscimtenant/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_tenant google_iam_workforce_pool_provider_scim_tenant}.
type GoogleIamWorkforcePoolProviderScimTenant interface {
	cdktf.TerraformResource
	BaseUri() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClaimMapping() *map[string]*string
	SetClaimMapping(val *map[string]*string)
	ClaimMappingInput() *map[string]*string
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HardDelete() interface{}
	SetHardDelete(val interface{})
	HardDeleteInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderId() *string
	SetProviderId(val *string)
	ProviderIdInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurgeTime() *string
	// Experimental.
	RawOverrides() interface{}
	ScimTenantId() *string
	SetScimTenantId(val *string)
	ScimTenantIdInput() *string
	ServiceAgent() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleIamWorkforcePoolProviderScimTenantTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkforcePoolId() *string
	SetWorkforcePoolId(val *string)
	WorkforcePoolIdInput() *string
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
	PutTimeouts(value *GoogleIamWorkforcePoolProviderScimTenantTimeouts)
	ResetClaimMapping()
	ResetDescription()
	ResetDisplayName()
	ResetHardDelete()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for GoogleIamWorkforcePoolProviderScimTenant
type jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) BaseUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ClaimMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ClaimMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"claimMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) HardDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) HardDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) PurgeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purgeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ScimTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ScimTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scimTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ServiceAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) Timeouts() GoogleIamWorkforcePoolProviderScimTenantTimeoutsOutputReference {
	var returns GoogleIamWorkforcePoolProviderScimTenantTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) WorkforcePoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforcePoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) WorkforcePoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workforcePoolIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_tenant google_iam_workforce_pool_provider_scim_tenant} Resource.
func NewGoogleIamWorkforcePoolProviderScimTenant(scope constructs.Construct, id *string, config *GoogleIamWorkforcePoolProviderScimTenantConfig) GoogleIamWorkforcePoolProviderScimTenant {
	_init_.Initialize()

	if err := validateNewGoogleIamWorkforcePoolProviderScimTenantParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_provider_scim_tenant google_iam_workforce_pool_provider_scim_tenant} Resource.
func NewGoogleIamWorkforcePoolProviderScimTenant_Override(g GoogleIamWorkforcePoolProviderScimTenant, scope constructs.Construct, id *string, config *GoogleIamWorkforcePoolProviderScimTenantConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetClaimMapping(val *map[string]*string) {
	if err := j.validateSetClaimMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"claimMapping",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetHardDelete(val interface{}) {
	if err := j.validateSetHardDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetProviderId(val *string) {
	if err := j.validateSetProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerId",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetScimTenantId(val *string) {
	if err := j.validateSetScimTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scimTenantId",
		val,
	)
}

func (j *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant)SetWorkforcePoolId(val *string) {
	if err := j.validateSetWorkforcePoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workforcePoolId",
		val,
	)
}

// Generates CDKTF code for importing a GoogleIamWorkforcePoolProviderScimTenant resource upon running "cdktf plan <stack-name>".
func GoogleIamWorkforcePoolProviderScimTenant_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProviderScimTenant_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
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
func GoogleIamWorkforcePoolProviderScimTenant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProviderScimTenant_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIamWorkforcePoolProviderScimTenant_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProviderScimTenant_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleIamWorkforcePoolProviderScimTenant_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleIamWorkforcePoolProviderScimTenant_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleIamWorkforcePoolProviderScimTenant_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleIamWorkforcePoolProviderScimTenant.GoogleIamWorkforcePoolProviderScimTenant",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) PutTimeouts(value *GoogleIamWorkforcePoolProviderScimTenantTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetClaimMapping() {
	_jsii_.InvokeVoid(
		g,
		"resetClaimMapping",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetHardDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetHardDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIamWorkforcePoolProviderScimTenant) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

