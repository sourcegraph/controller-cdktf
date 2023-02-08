package orgconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/orgconfiguration/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/r/org_configuration okta_org_configuration}.
type OrgConfiguration interface {
	cdktf.TerraformResource
	Address1() *string
	SetAddress1(val *string)
	Address1Input() *string
	Address2() *string
	SetAddress2(val *string)
	Address2Input() *string
	BillingContactUser() *string
	SetBillingContactUser(val *string)
	BillingContactUserInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	CompanyName() *string
	SetCompanyName(val *string)
	CompanyNameInput() *string
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
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EndUserSupportHelpUrl() *string
	SetEndUserSupportHelpUrl(val *string)
	EndUserSupportHelpUrlInput() *string
	ExpiresAt() *string
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
	Logo() *string
	SetLogo(val *string)
	LogoInput() *string
	// The tree node.
	Node() constructs.Node
	OptOutCommunicationEmails() interface{}
	SetOptOutCommunicationEmails(val interface{})
	OptOutCommunicationEmailsInput() interface{}
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
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
	State() *string
	SetState(val *string)
	StateInput() *string
	Subdomain() *string
	SupportPhoneNumber() *string
	SetSupportPhoneNumber(val *string)
	SupportPhoneNumberInput() *string
	TechnicalContactUser() *string
	SetTechnicalContactUser(val *string)
	TechnicalContactUserInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Website() *string
	SetWebsite(val *string)
	WebsiteInput() *string
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
	ResetAddress1()
	ResetAddress2()
	ResetBillingContactUser()
	ResetCity()
	ResetCountry()
	ResetEndUserSupportHelpUrl()
	ResetId()
	ResetLogo()
	ResetOptOutCommunicationEmails()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhoneNumber()
	ResetPostalCode()
	ResetState()
	ResetSupportPhoneNumber()
	ResetTechnicalContactUser()
	ResetWebsite()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OrgConfiguration
type jsiiProxy_OrgConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrgConfiguration) Address1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Address1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Address2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Address2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) BillingContactUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingContactUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) BillingContactUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingContactUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) CompanyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) CompanyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) EndUserSupportHelpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endUserSupportHelpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) EndUserSupportHelpUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endUserSupportHelpUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) ExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Logo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) LogoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) OptOutCommunicationEmails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optOutCommunicationEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) OptOutCommunicationEmailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optOutCommunicationEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Subdomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) SupportPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) SupportPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) TechnicalContactUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technicalContactUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) TechnicalContactUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"technicalContactUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) Website() *string {
	var returns *string
	_jsii_.Get(
		j,
		"website",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgConfiguration) WebsiteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/r/org_configuration okta_org_configuration} Resource.
func NewOrgConfiguration(scope constructs.Construct, id *string, config *OrgConfigurationConfig) OrgConfiguration {
	_init_.Initialize()

	if err := validateNewOrgConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgConfiguration{}

	_jsii_.Create(
		"okta.orgConfiguration.OrgConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/r/org_configuration okta_org_configuration} Resource.
func NewOrgConfiguration_Override(o OrgConfiguration, scope constructs.Construct, id *string, config *OrgConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.orgConfiguration.OrgConfiguration",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetAddress1(val *string) {
	if err := j.validateSetAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address1",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetAddress2(val *string) {
	if err := j.validateSetAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address2",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetBillingContactUser(val *string) {
	if err := j.validateSetBillingContactUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingContactUser",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetCompanyName(val *string) {
	if err := j.validateSetCompanyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"companyName",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetEndUserSupportHelpUrl(val *string) {
	if err := j.validateSetEndUserSupportHelpUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endUserSupportHelpUrl",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetLogo(val *string) {
	if err := j.validateSetLogoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logo",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetOptOutCommunicationEmails(val interface{}) {
	if err := j.validateSetOptOutCommunicationEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutCommunicationEmails",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetSupportPhoneNumber(val *string) {
	if err := j.validateSetSupportPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetTechnicalContactUser(val *string) {
	if err := j.validateSetTechnicalContactUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technicalContactUser",
		val,
	)
}

func (j *jsiiProxy_OrgConfiguration)SetWebsite(val *string) {
	if err := j.validateSetWebsiteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"website",
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
func OrgConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrgConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.orgConfiguration.OrgConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrgConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.orgConfiguration.OrgConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrgConfiguration) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OrgConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetAddress1() {
	_jsii_.InvokeVoid(
		o,
		"resetAddress1",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetAddress2() {
	_jsii_.InvokeVoid(
		o,
		"resetAddress2",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetBillingContactUser() {
	_jsii_.InvokeVoid(
		o,
		"resetBillingContactUser",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetCity() {
	_jsii_.InvokeVoid(
		o,
		"resetCity",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetCountry() {
	_jsii_.InvokeVoid(
		o,
		"resetCountry",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetEndUserSupportHelpUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetEndUserSupportHelpUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetLogo() {
	_jsii_.InvokeVoid(
		o,
		"resetLogo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetOptOutCommunicationEmails() {
	_jsii_.InvokeVoid(
		o,
		"resetOptOutCommunicationEmails",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		o,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetPostalCode() {
	_jsii_.InvokeVoid(
		o,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetState() {
	_jsii_.InvokeVoid(
		o,
		"resetState",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetSupportPhoneNumber() {
	_jsii_.InvokeVoid(
		o,
		"resetSupportPhoneNumber",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetTechnicalContactUser() {
	_jsii_.InvokeVoid(
		o,
		"resetTechnicalContactUser",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) ResetWebsite() {
	_jsii_.InvokeVoid(
		o,
		"resetWebsite",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

