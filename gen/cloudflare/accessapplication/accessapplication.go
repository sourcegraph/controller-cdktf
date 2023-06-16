package accessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/accessapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application cloudflare_access_application}.
type AccessApplication interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedIdps() *[]*string
	SetAllowedIdps(val *[]*string)
	AllowedIdpsInput() *[]*string
	AppLauncherVisible() interface{}
	SetAppLauncherVisible(val interface{})
	AppLauncherVisibleInput() interface{}
	Aud() *string
	AutoRedirectToIdentity() interface{}
	SetAutoRedirectToIdentity(val interface{})
	AutoRedirectToIdentityInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsHeaders() AccessApplicationCorsHeadersList
	CorsHeadersInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomDenyMessage() *string
	SetCustomDenyMessage(val *string)
	CustomDenyMessageInput() *string
	CustomDenyUrl() *string
	SetCustomDenyUrl(val *string)
	CustomDenyUrlInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	EnableBindingCookie() interface{}
	SetEnableBindingCookie(val interface{})
	EnableBindingCookieInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpOnlyCookieAttribute() interface{}
	SetHttpOnlyCookieAttribute(val interface{})
	HttpOnlyCookieAttributeInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoUrl() *string
	SetLogoUrl(val *string)
	LogoUrlInput() *string
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
	SaasApp() AccessApplicationSaasAppOutputReference
	SaasAppInput() *AccessApplicationSaasApp
	SameSiteCookieAttribute() *string
	SetSameSiteCookieAttribute(val *string)
	SameSiteCookieAttributeInput() *string
	ServiceAuth401Redirect() interface{}
	SetServiceAuth401Redirect(val interface{})
	ServiceAuth401RedirectInput() interface{}
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	SkipInterstitial() interface{}
	SetSkipInterstitial(val interface{})
	SkipInterstitialInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
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
	PutCorsHeaders(value interface{})
	PutSaasApp(value *AccessApplicationSaasApp)
	ResetAccountId()
	ResetAllowedIdps()
	ResetAppLauncherVisible()
	ResetAutoRedirectToIdentity()
	ResetCorsHeaders()
	ResetCustomDenyMessage()
	ResetCustomDenyUrl()
	ResetDomain()
	ResetEnableBindingCookie()
	ResetHttpOnlyCookieAttribute()
	ResetId()
	ResetLogoUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSaasApp()
	ResetSameSiteCookieAttribute()
	ResetServiceAuth401Redirect()
	ResetSessionDuration()
	ResetSkipInterstitial()
	ResetType()
	ResetZoneId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AccessApplication
type jsiiProxy_AccessApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccessApplication) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AllowedIdps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AllowedIdpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AppLauncherVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLauncherVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AppLauncherVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appLauncherVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Aud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AutoRedirectToIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) AutoRedirectToIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRedirectToIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CorsHeaders() AccessApplicationCorsHeadersList {
	var returns AccessApplicationCorsHeadersList
	_jsii_.Get(
		j,
		"corsHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CorsHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CustomDenyMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CustomDenyMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CustomDenyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) CustomDenyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) EnableBindingCookie() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBindingCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) EnableBindingCookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBindingCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) HttpOnlyCookieAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpOnlyCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) HttpOnlyCookieAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpOnlyCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) LogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SaasApp() AccessApplicationSaasAppOutputReference {
	var returns AccessApplicationSaasAppOutputReference
	_jsii_.Get(
		j,
		"saasApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SaasAppInput() *AccessApplicationSaasApp {
	var returns *AccessApplicationSaasApp
	_jsii_.Get(
		j,
		"saasAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SameSiteCookieAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SameSiteCookieAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ServiceAuth401Redirect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAuth401Redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ServiceAuth401RedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAuth401RedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SkipInterstitial() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInterstitial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) SkipInterstitialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipInterstitialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessApplication) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application cloudflare_access_application} Resource.
func NewAccessApplication(scope constructs.Construct, id *string, config *AccessApplicationConfig) AccessApplication {
	_init_.Initialize()

	if err := validateNewAccessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessApplication{}

	_jsii_.Create(
		"cloudflare.accessApplication.AccessApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/access_application cloudflare_access_application} Resource.
func NewAccessApplication_Override(a AccessApplication, scope constructs.Construct, id *string, config *AccessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.accessApplication.AccessApplication",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccessApplication)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetAllowedIdps(val *[]*string) {
	if err := j.validateSetAllowedIdpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIdps",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetAppLauncherVisible(val interface{}) {
	if err := j.validateSetAppLauncherVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherVisible",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetAutoRedirectToIdentity(val interface{}) {
	if err := j.validateSetAutoRedirectToIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRedirectToIdentity",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetCustomDenyMessage(val *string) {
	if err := j.validateSetCustomDenyMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyMessage",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetCustomDenyUrl(val *string) {
	if err := j.validateSetCustomDenyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyUrl",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetEnableBindingCookie(val interface{}) {
	if err := j.validateSetEnableBindingCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBindingCookie",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetHttpOnlyCookieAttribute(val interface{}) {
	if err := j.validateSetHttpOnlyCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpOnlyCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetLogoUrl(val *string) {
	if err := j.validateSetLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUrl",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetSameSiteCookieAttribute(val *string) {
	if err := j.validateSetSameSiteCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sameSiteCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetServiceAuth401Redirect(val interface{}) {
	if err := j.validateSetServiceAuth401RedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAuth401Redirect",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetSkipInterstitial(val interface{}) {
	if err := j.validateSetSkipInterstitialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipInterstitial",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AccessApplication)SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
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
func AccessApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessApplication.AccessApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessApplication.AccessApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccessApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cloudflare.accessApplication.AccessApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cloudflare.accessApplication.AccessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccessApplication) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccessApplication) PutCorsHeaders(value interface{}) {
	if err := a.validatePutCorsHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCorsHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplication) PutSaasApp(value *AccessApplicationSaasApp) {
	if err := a.validatePutSaasAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSaasApp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessApplication) ResetAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetAllowedIdps() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedIdps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetAppLauncherVisible() {
	_jsii_.InvokeVoid(
		a,
		"resetAppLauncherVisible",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetAutoRedirectToIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRedirectToIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetCorsHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetCorsHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetCustomDenyMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDenyMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetCustomDenyUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDenyUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetEnableBindingCookie() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBindingCookie",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetHttpOnlyCookieAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpOnlyCookieAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetLogoUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetLogoUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetSaasApp() {
	_jsii_.InvokeVoid(
		a,
		"resetSaasApp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetSameSiteCookieAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetSameSiteCookieAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetServiceAuth401Redirect() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAuth401Redirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetSkipInterstitial() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipInterstitial",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) ResetZoneId() {
	_jsii_.InvokeVoid(
		a,
		"resetZoneId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

