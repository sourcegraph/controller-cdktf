package dataoktaappsaml

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/dataoktaappsaml/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/d/app_saml okta_app_saml}.
type DataOktaAppSaml interface {
	cdktf.TerraformDataSource
	AccessibilityErrorRedirectUrl() *string
	AccessibilityLoginRedirectUrl() *string
	AccessibilitySelfService() cdktf.IResolvable
	AcsEndpoints() *[]*string
	ActiveOnly() interface{}
	SetActiveOnly(val interface{})
	ActiveOnlyInput() interface{}
	AppSettingsJson() *string
	AssertionSigned() cdktf.IResolvable
	AttributeStatements() DataOktaAppSamlAttributeStatementsList
	Audience() *string
	AuthnContextClassRef() *string
	AutoSubmitToolbar() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultRelayState() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destination() *string
	DigestAlgorithm() *string
	Features() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groups() *[]*string
	HideIos() cdktf.IResolvable
	HideWeb() cdktf.IResolvable
	HonorForceAuthn() cdktf.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	IdpIssuer() *string
	InlineHookId() *string
	KeyId() *string
	Label() *string
	SetLabel(val *string)
	LabelInput() *string
	LabelPrefix() *string
	SetLabelPrefix(val *string)
	LabelPrefixInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Links() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Recipient() *string
	RequestCompressed() interface{}
	SetRequestCompressed(val interface{})
	RequestCompressedInput() interface{}
	ResponseSigned() cdktf.IResolvable
	SignatureAlgorithm() *string
	SingleLogoutCertificate() *string
	SingleLogoutIssuer() *string
	SingleLogoutUrl() *string
	SkipGroups() interface{}
	SetSkipGroups(val interface{})
	SkipGroupsInput() interface{}
	SkipUsers() interface{}
	SetSkipUsers(val interface{})
	SkipUsersInput() interface{}
	SpIssuer() *string
	SsoUrl() *string
	Status() *string
	SubjectNameIdFormat() *string
	SubjectNameIdTemplate() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserNameTemplate() *string
	UserNameTemplatePushStatus() *string
	UserNameTemplateSuffix() *string
	UserNameTemplateType() *string
	Users() *[]*string
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
	ResetActiveOnly()
	ResetId()
	ResetLabel()
	ResetLabelPrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestCompressed()
	ResetSkipGroups()
	ResetSkipUsers()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOktaAppSaml
type jsiiProxy_DataOktaAppSaml struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOktaAppSaml) AccessibilityErrorRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityErrorRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AccessibilityLoginRedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessibilityLoginRedirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AccessibilitySelfService() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"accessibilitySelfService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AcsEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acsEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) ActiveOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) ActiveOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AppSettingsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSettingsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AssertionSigned() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"assertionSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AttributeStatements() DataOktaAppSamlAttributeStatementsList {
	var returns DataOktaAppSamlAttributeStatementsList
	_jsii_.Get(
		j,
		"attributeStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AuthnContextClassRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authnContextClassRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) AutoSubmitToolbar() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoSubmitToolbar",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) DefaultRelayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRelayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) DigestAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"digestAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Features() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) HideIos() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hideIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) HideWeb() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hideWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) HonorForceAuthn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"honorForceAuthn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) IdpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) InlineHookId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineHookId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) KeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) LabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) LabelPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) LabelPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Links() *string {
	var returns *string
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Recipient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) RequestCompressed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) RequestCompressedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCompressedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) ResponseSigned() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"responseSigned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SingleLogoutCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SingleLogoutIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SingleLogoutUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singleLogoutUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SkipUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SkipUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SpIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SsoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SubjectNameIdFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) SubjectNameIdTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectNameIdTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) UserNameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) UserNameTemplatePushStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplatePushStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) UserNameTemplateSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) UserNameTemplateType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameTemplateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaAppSaml) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/d/app_saml okta_app_saml} Data Source.
func NewDataOktaAppSaml(scope constructs.Construct, id *string, config *DataOktaAppSamlConfig) DataOktaAppSaml {
	_init_.Initialize()

	if err := validateNewDataOktaAppSamlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaAppSaml{}

	_jsii_.Create(
		"okta.dataOktaAppSaml.DataOktaAppSaml",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/d/app_saml okta_app_saml} Data Source.
func NewDataOktaAppSaml_Override(d DataOktaAppSaml, scope constructs.Construct, id *string, config *DataOktaAppSamlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.dataOktaAppSaml.DataOktaAppSaml",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetActiveOnly(val interface{}) {
	if err := j.validateSetActiveOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeOnly",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetLabel(val *string) {
	if err := j.validateSetLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"label",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetLabelPrefix(val *string) {
	if err := j.validateSetLabelPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelPrefix",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetRequestCompressed(val interface{}) {
	if err := j.validateSetRequestCompressedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCompressed",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_DataOktaAppSaml)SetSkipUsers(val interface{}) {
	if err := j.validateSetSkipUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipUsers",
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
func DataOktaAppSaml_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOktaAppSaml_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.dataOktaAppSaml.DataOktaAppSaml",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOktaAppSaml_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.dataOktaAppSaml.DataOktaAppSaml",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOktaAppSaml) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetActiveOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetActiveOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetLabel() {
	_jsii_.InvokeVoid(
		d,
		"resetLabel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetLabelPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetLabelPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetRequestCompressed() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestCompressed",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) ResetSkipUsers() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipUsers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaAppSaml) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaAppSaml) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

