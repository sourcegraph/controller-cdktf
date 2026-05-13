package googleapigeedeveloperapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeedeveloperapp/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app google_apigee_developer_app}.
type GoogleApigeeDeveloperApp interface {
	cdktf.TerraformResource
	ApiProducts() *[]*string
	SetApiProducts(val *[]*string)
	ApiProductsInput() *[]*string
	AppFamily() *string
	SetAppFamily(val *string)
	AppFamilyInput() *string
	AppId() *string
	Attributes() GoogleApigeeDeveloperAppAttributesList
	AttributesInput() interface{}
	CallbackUrl() *string
	SetCallbackUrl(val *string)
	CallbackUrlInput() *string
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
	CreatedAt() *string
	Credentials() GoogleApigeeDeveloperAppCredentialsList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeveloperEmail() *string
	SetDeveloperEmail(val *string)
	DeveloperEmailInput() *string
	DeveloperId() *string
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
	KeyExpiresIn() *string
	SetKeyExpiresIn(val *string)
	KeyExpiresInInput() *string
	LastModifiedAt() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleApigeeDeveloperAppTimeoutsOutputReference
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
	PutAttributes(value interface{})
	PutTimeouts(value *GoogleApigeeDeveloperAppTimeouts)
	ResetApiProducts()
	ResetAppFamily()
	ResetAttributes()
	ResetId()
	ResetKeyExpiresIn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScopes()
	ResetStatus()
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

// The jsii proxy struct for GoogleApigeeDeveloperApp
type jsiiProxy_GoogleApigeeDeveloperApp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) ApiProducts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProducts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) ApiProductsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProductsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) AppFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) AppFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Attributes() GoogleApigeeDeveloperAppAttributesList {
	var returns GoogleApigeeDeveloperAppAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) CallbackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Credentials() GoogleApigeeDeveloperAppCredentialsList {
	var returns GoogleApigeeDeveloperAppCredentialsList
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) DeveloperEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) DeveloperEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) DeveloperId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) KeyExpiresIn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyExpiresIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) KeyExpiresInInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyExpiresInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) LastModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) OrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) OrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) Timeouts() GoogleApigeeDeveloperAppTimeoutsOutputReference {
	var returns GoogleApigeeDeveloperAppTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app google_apigee_developer_app} Resource.
func NewGoogleApigeeDeveloperApp(scope constructs.Construct, id *string, config *GoogleApigeeDeveloperAppConfig) GoogleApigeeDeveloperApp {
	_init_.Initialize()

	if err := validateNewGoogleApigeeDeveloperAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeDeveloperApp{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app google_apigee_developer_app} Resource.
func NewGoogleApigeeDeveloperApp_Override(g GoogleApigeeDeveloperApp, scope constructs.Construct, id *string, config *GoogleApigeeDeveloperAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetApiProducts(val *[]*string) {
	if err := j.validateSetApiProductsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiProducts",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetAppFamily(val *string) {
	if err := j.validateSetAppFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appFamily",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetCallbackUrl(val *string) {
	if err := j.validateSetCallbackUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"callbackUrl",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetDeveloperEmail(val *string) {
	if err := j.validateSetDeveloperEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetKeyExpiresIn(val *string) {
	if err := j.validateSetKeyExpiresInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyExpiresIn",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetOrgId(val *string) {
	if err := j.validateSetOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgId",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeDeveloperApp)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a GoogleApigeeDeveloperApp resource upon running "cdktf plan <stack-name>".
func GoogleApigeeDeveloperApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleApigeeDeveloperApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
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
func GoogleApigeeDeveloperApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeDeveloperApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeDeveloperApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeDeveloperApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApigeeDeveloperApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApigeeDeveloperApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleApigeeDeveloperApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleApigeeDeveloperApp.GoogleApigeeDeveloperApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeDeveloperApp) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) PutAttributes(value interface{}) {
	if err := g.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) PutTimeouts(value *GoogleApigeeDeveloperAppTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetApiProducts() {
	_jsii_.InvokeVoid(
		g,
		"resetApiProducts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetAppFamily() {
	_jsii_.InvokeVoid(
		g,
		"resetAppFamily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetKeyExpiresIn() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyExpiresIn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetStatus() {
	_jsii_.InvokeVoid(
		g,
		"resetStatus",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeDeveloperApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

