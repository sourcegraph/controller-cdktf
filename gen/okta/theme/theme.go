package theme

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/theme/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/theme okta_theme}.
type Theme interface {
	cdktf.TerraformResource
	BackgroundImage() *string
	SetBackgroundImage(val *string)
	BackgroundImageInput() *string
	BackgroundImageUrl() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EmailTemplateTouchPointVariant() *string
	SetEmailTemplateTouchPointVariant(val *string)
	EmailTemplateTouchPointVariantInput() *string
	EndUserDashboardTouchPointVariant() *string
	SetEndUserDashboardTouchPointVariant(val *string)
	EndUserDashboardTouchPointVariantInput() *string
	ErrorPageTouchPointVariant() *string
	SetErrorPageTouchPointVariant(val *string)
	ErrorPageTouchPointVariantInput() *string
	Favicon() *string
	SetFavicon(val *string)
	FaviconInput() *string
	FaviconUrl() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Links() *string
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	LogoUrl() *string
	// The tree node.
	Node() constructs.Node
	PrimaryColorContrastHex() *string
	SetPrimaryColorContrastHex(val *string)
	PrimaryColorContrastHexInput() *string
	PrimaryColorHex() *string
	SetPrimaryColorHex(val *string)
	PrimaryColorHexInput() *string
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
	SecondaryColorContrastHex() *string
	SetSecondaryColorContrastHex(val *string)
	SecondaryColorContrastHexInput() *string
	SecondaryColorHex() *string
	SetSecondaryColorHex(val *string)
	SecondaryColorHexInput() *string
	SignInPageTouchPointVariant() *string
	SetSignInPageTouchPointVariant(val *string)
	SignInPageTouchPointVariantInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThemeId() *string
	SetThemeId(val *string)
	ThemeIdInput() *string
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
	ResetBackgroundImage()
	ResetEmailTemplateTouchPointVariant()
	ResetEndUserDashboardTouchPointVariant()
	ResetErrorPageTouchPointVariant()
	ResetFavicon()
	ResetLogo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryColorContrastHex()
	ResetPrimaryColorHex()
	ResetSecondaryColorContrastHex()
	ResetSecondaryColorHex()
	ResetSignInPageTouchPointVariant()
	ResetThemeId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Theme
type jsiiProxy_Theme struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Theme) BackgroundImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) BackgroundImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) BackgroundImageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backgroundImageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) BrandId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brandId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) BrandIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brandIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) EmailTemplateTouchPointVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailTemplateTouchPointVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) EmailTemplateTouchPointVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailTemplateTouchPointVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) EndUserDashboardTouchPointVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endUserDashboardTouchPointVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) EndUserDashboardTouchPointVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endUserDashboardTouchPointVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ErrorPageTouchPointVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorPageTouchPointVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ErrorPageTouchPointVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorPageTouchPointVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Favicon() *string {
	var returns *string
	_jsii_.Get(
		j,
		"favicon",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) FaviconInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faviconInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) FaviconUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faviconUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Links() *string {
	var returns *string
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) PrimaryColorContrastHex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColorContrastHex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) PrimaryColorContrastHexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColorContrastHexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) PrimaryColorHex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColorHex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) PrimaryColorHexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryColorHexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SecondaryColorContrastHex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColorContrastHex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SecondaryColorContrastHexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColorContrastHexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SecondaryColorHex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColorHex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SecondaryColorHexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryColorHexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SignInPageTouchPointVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInPageTouchPointVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) SignInPageTouchPointVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signInPageTouchPointVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ThemeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Theme) ThemeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/theme okta_theme} Resource.
func NewTheme(scope constructs.Construct, id *string, config *ThemeConfig) Theme {
	_init_.Initialize()

	if err := validateNewThemeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Theme{}

	_jsii_.Create(
		"@cdktf/provider-okta.theme.Theme",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/resources/theme okta_theme} Resource.
func NewTheme_Override(t Theme, scope constructs.Construct, id *string, config *ThemeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.theme.Theme",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_Theme)SetBackgroundImage(val *string) {
	if err := j.validateSetBackgroundImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backgroundImage",
		val,
	)
}

func (j *jsiiProxy_Theme)SetBrandId(val *string) {
	if err := j.validateSetBrandIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brandId",
		val,
	)
}

func (j *jsiiProxy_Theme)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Theme)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Theme)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Theme)SetEmailTemplateTouchPointVariant(val *string) {
	if err := j.validateSetEmailTemplateTouchPointVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailTemplateTouchPointVariant",
		val,
	)
}

func (j *jsiiProxy_Theme)SetEndUserDashboardTouchPointVariant(val *string) {
	if err := j.validateSetEndUserDashboardTouchPointVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endUserDashboardTouchPointVariant",
		val,
	)
}

func (j *jsiiProxy_Theme)SetErrorPageTouchPointVariant(val *string) {
	if err := j.validateSetErrorPageTouchPointVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorPageTouchPointVariant",
		val,
	)
}

func (j *jsiiProxy_Theme)SetFavicon(val *string) {
	if err := j.validateSetFaviconParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"favicon",
		val,
	)
}

func (j *jsiiProxy_Theme)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Theme)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Theme)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_Theme)SetPrimaryColorContrastHex(val *string) {
	if err := j.validateSetPrimaryColorContrastHexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryColorContrastHex",
		val,
	)
}

func (j *jsiiProxy_Theme)SetPrimaryColorHex(val *string) {
	if err := j.validateSetPrimaryColorHexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryColorHex",
		val,
	)
}

func (j *jsiiProxy_Theme)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Theme)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Theme)SetSecondaryColorContrastHex(val *string) {
	if err := j.validateSetSecondaryColorContrastHexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryColorContrastHex",
		val,
	)
}

func (j *jsiiProxy_Theme)SetSecondaryColorHex(val *string) {
	if err := j.validateSetSecondaryColorHexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryColorHex",
		val,
	)
}

func (j *jsiiProxy_Theme)SetSignInPageTouchPointVariant(val *string) {
	if err := j.validateSetSignInPageTouchPointVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signInPageTouchPointVariant",
		val,
	)
}

func (j *jsiiProxy_Theme)SetThemeId(val *string) {
	if err := j.validateSetThemeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"themeId",
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
func Theme_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTheme_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.theme.Theme",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Theme_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTheme_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.theme.Theme",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Theme_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTheme_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.theme.Theme",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Theme_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.theme.Theme",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_Theme) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_Theme) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_Theme) ResetBackgroundImage() {
	_jsii_.InvokeVoid(
		t,
		"resetBackgroundImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetEmailTemplateTouchPointVariant() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailTemplateTouchPointVariant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetEndUserDashboardTouchPointVariant() {
	_jsii_.InvokeVoid(
		t,
		"resetEndUserDashboardTouchPointVariant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetErrorPageTouchPointVariant() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorPageTouchPointVariant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetFavicon() {
	_jsii_.InvokeVoid(
		t,
		"resetFavicon",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetLogo() {
	_jsii_.InvokeVoid(
		t,
		"resetLogo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetPrimaryColorContrastHex() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryColorContrastHex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetPrimaryColorHex() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryColorHex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetSecondaryColorContrastHex() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryColorContrastHex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetSecondaryColorHex() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryColorHex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetSignInPageTouchPointVariant() {
	_jsii_.InvokeVoid(
		t,
		"resetSignInPageTouchPointVariant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) ResetThemeId() {
	_jsii_.InvokeVoid(
		t,
		"resetThemeId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Theme) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Theme) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

