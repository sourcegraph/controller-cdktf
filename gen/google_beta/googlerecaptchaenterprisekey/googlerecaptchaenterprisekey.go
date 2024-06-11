package googlerecaptchaenterprisekey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlerecaptchaenterprisekey/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_recaptcha_enterprise_key google_recaptcha_enterprise_key}.
type GoogleRecaptchaEnterpriseKey interface {
	cdktf.TerraformResource
	AndroidSettings() GoogleRecaptchaEnterpriseKeyAndroidSettingsOutputReference
	AndroidSettingsInput() *GoogleRecaptchaEnterpriseKeyAndroidSettings
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
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
	IosSettings() GoogleRecaptchaEnterpriseKeyIosSettingsOutputReference
	IosSettingsInput() *GoogleRecaptchaEnterpriseKeyIosSettings
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
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
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestingOptions() GoogleRecaptchaEnterpriseKeyTestingOptionsOutputReference
	TestingOptionsInput() *GoogleRecaptchaEnterpriseKeyTestingOptions
	Timeouts() GoogleRecaptchaEnterpriseKeyTimeoutsOutputReference
	TimeoutsInput() interface{}
	WafSettings() GoogleRecaptchaEnterpriseKeyWafSettingsOutputReference
	WafSettingsInput() *GoogleRecaptchaEnterpriseKeyWafSettings
	WebSettings() GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference
	WebSettingsInput() *GoogleRecaptchaEnterpriseKeyWebSettings
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
	PutAndroidSettings(value *GoogleRecaptchaEnterpriseKeyAndroidSettings)
	PutIosSettings(value *GoogleRecaptchaEnterpriseKeyIosSettings)
	PutTestingOptions(value *GoogleRecaptchaEnterpriseKeyTestingOptions)
	PutTimeouts(value *GoogleRecaptchaEnterpriseKeyTimeouts)
	PutWafSettings(value *GoogleRecaptchaEnterpriseKeyWafSettings)
	PutWebSettings(value *GoogleRecaptchaEnterpriseKeyWebSettings)
	ResetAndroidSettings()
	ResetId()
	ResetIosSettings()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTestingOptions()
	ResetTimeouts()
	ResetWafSettings()
	ResetWebSettings()
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

// The jsii proxy struct for GoogleRecaptchaEnterpriseKey
type jsiiProxy_GoogleRecaptchaEnterpriseKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) AndroidSettings() GoogleRecaptchaEnterpriseKeyAndroidSettingsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyAndroidSettingsOutputReference
	_jsii_.Get(
		j,
		"androidSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) AndroidSettingsInput() *GoogleRecaptchaEnterpriseKeyAndroidSettings {
	var returns *GoogleRecaptchaEnterpriseKeyAndroidSettings
	_jsii_.Get(
		j,
		"androidSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) IosSettings() GoogleRecaptchaEnterpriseKeyIosSettingsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyIosSettingsOutputReference
	_jsii_.Get(
		j,
		"iosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) IosSettingsInput() *GoogleRecaptchaEnterpriseKeyIosSettings {
	var returns *GoogleRecaptchaEnterpriseKeyIosSettings
	_jsii_.Get(
		j,
		"iosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TestingOptions() GoogleRecaptchaEnterpriseKeyTestingOptionsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyTestingOptionsOutputReference
	_jsii_.Get(
		j,
		"testingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TestingOptionsInput() *GoogleRecaptchaEnterpriseKeyTestingOptions {
	var returns *GoogleRecaptchaEnterpriseKeyTestingOptions
	_jsii_.Get(
		j,
		"testingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) Timeouts() GoogleRecaptchaEnterpriseKeyTimeoutsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) WafSettings() GoogleRecaptchaEnterpriseKeyWafSettingsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyWafSettingsOutputReference
	_jsii_.Get(
		j,
		"wafSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) WafSettingsInput() *GoogleRecaptchaEnterpriseKeyWafSettings {
	var returns *GoogleRecaptchaEnterpriseKeyWafSettings
	_jsii_.Get(
		j,
		"wafSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) WebSettings() GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference {
	var returns GoogleRecaptchaEnterpriseKeyWebSettingsOutputReference
	_jsii_.Get(
		j,
		"webSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey) WebSettingsInput() *GoogleRecaptchaEnterpriseKeyWebSettings {
	var returns *GoogleRecaptchaEnterpriseKeyWebSettings
	_jsii_.Get(
		j,
		"webSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_recaptcha_enterprise_key google_recaptcha_enterprise_key} Resource.
func NewGoogleRecaptchaEnterpriseKey(scope constructs.Construct, id *string, config *GoogleRecaptchaEnterpriseKeyConfig) GoogleRecaptchaEnterpriseKey {
	_init_.Initialize()

	if err := validateNewGoogleRecaptchaEnterpriseKeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRecaptchaEnterpriseKey{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_recaptcha_enterprise_key google_recaptcha_enterprise_key} Resource.
func NewGoogleRecaptchaEnterpriseKey_Override(g GoogleRecaptchaEnterpriseKey, scope constructs.Construct, id *string, config *GoogleRecaptchaEnterpriseKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleRecaptchaEnterpriseKey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleRecaptchaEnterpriseKey resource upon running "cdktf plan <stack-name>".
func GoogleRecaptchaEnterpriseKey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleRecaptchaEnterpriseKey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
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
func GoogleRecaptchaEnterpriseKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRecaptchaEnterpriseKey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleRecaptchaEnterpriseKey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRecaptchaEnterpriseKey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleRecaptchaEnterpriseKey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleRecaptchaEnterpriseKey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleRecaptchaEnterpriseKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleRecaptchaEnterpriseKey.GoogleRecaptchaEnterpriseKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutAndroidSettings(value *GoogleRecaptchaEnterpriseKeyAndroidSettings) {
	if err := g.validatePutAndroidSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAndroidSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutIosSettings(value *GoogleRecaptchaEnterpriseKeyIosSettings) {
	if err := g.validatePutIosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIosSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutTestingOptions(value *GoogleRecaptchaEnterpriseKeyTestingOptions) {
	if err := g.validatePutTestingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTestingOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutTimeouts(value *GoogleRecaptchaEnterpriseKeyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutWafSettings(value *GoogleRecaptchaEnterpriseKeyWafSettings) {
	if err := g.validatePutWafSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWafSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) PutWebSettings(value *GoogleRecaptchaEnterpriseKeyWebSettings) {
	if err := g.validatePutWebSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetAndroidSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetAndroidSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetIosSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetIosSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetTestingOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetTestingOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetWafSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWafSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ResetWebSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetWebSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRecaptchaEnterpriseKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

