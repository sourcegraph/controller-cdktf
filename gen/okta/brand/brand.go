package brand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/brand/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand okta_brand}.
type Brand interface {
	cdktf.TerraformResource
	AgreeToCustomPrivacyPolicy() interface{}
	SetAgreeToCustomPrivacyPolicy(val interface{})
	AgreeToCustomPrivacyPolicyInput() interface{}
	BrandId() *string
	SetBrandId(val *string)
	BrandIdInput() *string
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
	CustomPrivacyPolicyUrl() *string
	SetCustomPrivacyPolicyUrl(val *string)
	CustomPrivacyPolicyUrlInput() *string
	DefaultAppAppInstanceId() *string
	SetDefaultAppAppInstanceId(val *string)
	DefaultAppAppInstanceIdInput() *string
	DefaultAppAppLinkName() *string
	SetDefaultAppAppLinkName(val *string)
	DefaultAppAppLinkNameInput() *string
	DefaultAppClassicApplicationUri() *string
	SetDefaultAppClassicApplicationUri(val *string)
	DefaultAppClassicApplicationUriInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmailDomainId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsDefault() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Links() *string
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
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
	RemovePoweredByOkta() interface{}
	SetRemovePoweredByOkta(val interface{})
	RemovePoweredByOktaInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAgreeToCustomPrivacyPolicy()
	ResetBrandId()
	ResetCustomPrivacyPolicyUrl()
	ResetDefaultAppAppInstanceId()
	ResetDefaultAppAppLinkName()
	ResetDefaultAppClassicApplicationUri()
	ResetLocale()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemovePoweredByOkta()
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

// The jsii proxy struct for Brand
type jsiiProxy_Brand struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Brand) AgreeToCustomPrivacyPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agreeToCustomPrivacyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) AgreeToCustomPrivacyPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agreeToCustomPrivacyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) BrandId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brandId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) BrandIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brandIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) CustomPrivacyPolicyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPrivacyPolicyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) CustomPrivacyPolicyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPrivacyPolicyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppAppInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppAppInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppAppInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppAppInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppAppLinkName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppAppLinkName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppAppLinkNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppAppLinkNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppClassicApplicationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppClassicApplicationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DefaultAppClassicApplicationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAppClassicApplicationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) EmailDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) IsDefault() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Links() *string {
	var returns *string
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) RemovePoweredByOkta() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removePoweredByOkta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) RemovePoweredByOktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removePoweredByOktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Brand) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand okta_brand} Resource.
func NewBrand(scope constructs.Construct, id *string, config *BrandConfig) Brand {
	_init_.Initialize()

	if err := validateNewBrandParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Brand{}

	_jsii_.Create(
		"@cdktf/provider-okta.brand.Brand",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/brand okta_brand} Resource.
func NewBrand_Override(b Brand, scope constructs.Construct, id *string, config *BrandConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.brand.Brand",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_Brand)SetAgreeToCustomPrivacyPolicy(val interface{}) {
	if err := j.validateSetAgreeToCustomPrivacyPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agreeToCustomPrivacyPolicy",
		val,
	)
}

func (j *jsiiProxy_Brand)SetBrandId(val *string) {
	if err := j.validateSetBrandIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brandId",
		val,
	)
}

func (j *jsiiProxy_Brand)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Brand)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Brand)SetCustomPrivacyPolicyUrl(val *string) {
	if err := j.validateSetCustomPrivacyPolicyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPrivacyPolicyUrl",
		val,
	)
}

func (j *jsiiProxy_Brand)SetDefaultAppAppInstanceId(val *string) {
	if err := j.validateSetDefaultAppAppInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAppAppInstanceId",
		val,
	)
}

func (j *jsiiProxy_Brand)SetDefaultAppAppLinkName(val *string) {
	if err := j.validateSetDefaultAppAppLinkNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAppAppLinkName",
		val,
	)
}

func (j *jsiiProxy_Brand)SetDefaultAppClassicApplicationUri(val *string) {
	if err := j.validateSetDefaultAppClassicApplicationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultAppClassicApplicationUri",
		val,
	)
}

func (j *jsiiProxy_Brand)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Brand)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Brand)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Brand)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_Brand)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Brand)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Brand)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Brand)SetRemovePoweredByOkta(val interface{}) {
	if err := j.validateSetRemovePoweredByOktaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removePoweredByOkta",
		val,
	)
}

// Generates CDKTF code for importing a Brand resource upon running "cdktf plan <stack-name>".
func Brand_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBrand_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.brand.Brand",
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
func Brand_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBrand_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.brand.Brand",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Brand_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBrand_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.brand.Brand",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Brand_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBrand_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.brand.Brand",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Brand_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.brand.Brand",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_Brand) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_Brand) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_Brand) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_Brand) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_Brand) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_Brand) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_Brand) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_Brand) ResetAgreeToCustomPrivacyPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetAgreeToCustomPrivacyPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetBrandId() {
	_jsii_.InvokeVoid(
		b,
		"resetBrandId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetCustomPrivacyPolicyUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomPrivacyPolicyUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetDefaultAppAppInstanceId() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultAppAppInstanceId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetDefaultAppAppLinkName() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultAppAppLinkName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetDefaultAppClassicApplicationUri() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultAppClassicApplicationUri",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetLocale() {
	_jsii_.InvokeVoid(
		b,
		"resetLocale",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) ResetRemovePoweredByOkta() {
	_jsii_.InvokeVoid(
		b,
		"resetRemovePoweredByOkta",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Brand) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Brand) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

