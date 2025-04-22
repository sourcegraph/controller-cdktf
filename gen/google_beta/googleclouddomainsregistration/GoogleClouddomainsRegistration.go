package googleclouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddomainsregistration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddomains_registration google_clouddomains_registration}.
type GoogleClouddomainsRegistration interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContactNotices() *[]*string
	SetContactNotices(val *[]*string)
	ContactNoticesInput() *[]*string
	ContactSettings() GoogleClouddomainsRegistrationContactSettingsOutputReference
	ContactSettingsInput() *GoogleClouddomainsRegistrationContactSettings
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsSettings() GoogleClouddomainsRegistrationDnsSettingsOutputReference
	DnsSettingsInput() *GoogleClouddomainsRegistrationDnsSettings
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainNotices() *[]*string
	SetDomainNotices(val *[]*string)
	DomainNoticesInput() *[]*string
	EffectiveLabels() cdktf.StringMap
	ExpireTime() *string
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
	Issues() *[]*string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementSettings() GoogleClouddomainsRegistrationManagementSettingsOutputReference
	ManagementSettingsInput() *GoogleClouddomainsRegistrationManagementSettings
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
	RegisterFailureReason() *string
	State() *string
	SupportedPrivacy() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleClouddomainsRegistrationTimeoutsOutputReference
	TimeoutsInput() interface{}
	YearlyPrice() GoogleClouddomainsRegistrationYearlyPriceOutputReference
	YearlyPriceInput() *GoogleClouddomainsRegistrationYearlyPrice
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
	PutContactSettings(value *GoogleClouddomainsRegistrationContactSettings)
	PutDnsSettings(value *GoogleClouddomainsRegistrationDnsSettings)
	PutManagementSettings(value *GoogleClouddomainsRegistrationManagementSettings)
	PutTimeouts(value *GoogleClouddomainsRegistrationTimeouts)
	PutYearlyPrice(value *GoogleClouddomainsRegistrationYearlyPrice)
	ResetContactNotices()
	ResetDnsSettings()
	ResetDomainNotices()
	ResetId()
	ResetLabels()
	ResetManagementSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for GoogleClouddomainsRegistration
type jsiiProxy_GoogleClouddomainsRegistration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ContactNotices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contactNotices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ContactNoticesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contactNoticesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ContactSettings() GoogleClouddomainsRegistrationContactSettingsOutputReference {
	var returns GoogleClouddomainsRegistrationContactSettingsOutputReference
	_jsii_.Get(
		j,
		"contactSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ContactSettingsInput() *GoogleClouddomainsRegistrationContactSettings {
	var returns *GoogleClouddomainsRegistrationContactSettings
	_jsii_.Get(
		j,
		"contactSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DnsSettings() GoogleClouddomainsRegistrationDnsSettingsOutputReference {
	var returns GoogleClouddomainsRegistrationDnsSettingsOutputReference
	_jsii_.Get(
		j,
		"dnsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DnsSettingsInput() *GoogleClouddomainsRegistrationDnsSettings {
	var returns *GoogleClouddomainsRegistrationDnsSettings
	_jsii_.Get(
		j,
		"dnsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DomainNotices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNotices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) DomainNoticesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainNoticesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ExpireTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Issues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ManagementSettings() GoogleClouddomainsRegistrationManagementSettingsOutputReference {
	var returns GoogleClouddomainsRegistrationManagementSettingsOutputReference
	_jsii_.Get(
		j,
		"managementSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ManagementSettingsInput() *GoogleClouddomainsRegistrationManagementSettings {
	var returns *GoogleClouddomainsRegistrationManagementSettings
	_jsii_.Get(
		j,
		"managementSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) RegisterFailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registerFailureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) SupportedPrivacy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedPrivacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) Timeouts() GoogleClouddomainsRegistrationTimeoutsOutputReference {
	var returns GoogleClouddomainsRegistrationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) YearlyPrice() GoogleClouddomainsRegistrationYearlyPriceOutputReference {
	var returns GoogleClouddomainsRegistrationYearlyPriceOutputReference
	_jsii_.Get(
		j,
		"yearlyPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistration) YearlyPriceInput() *GoogleClouddomainsRegistrationYearlyPrice {
	var returns *GoogleClouddomainsRegistrationYearlyPrice
	_jsii_.Get(
		j,
		"yearlyPriceInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddomains_registration google_clouddomains_registration} Resource.
func NewGoogleClouddomainsRegistration(scope constructs.Construct, id *string, config *GoogleClouddomainsRegistrationConfig) GoogleClouddomainsRegistration {
	_init_.Initialize()

	if err := validateNewGoogleClouddomainsRegistrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddomainsRegistration{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_clouddomains_registration google_clouddomains_registration} Resource.
func NewGoogleClouddomainsRegistration_Override(g GoogleClouddomainsRegistration, scope constructs.Construct, id *string, config *GoogleClouddomainsRegistrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetContactNotices(val *[]*string) {
	if err := j.validateSetContactNoticesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactNotices",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetDomainNotices(val *[]*string) {
	if err := j.validateSetDomainNoticesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNotices",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleClouddomainsRegistration resource upon running "cdktf plan <stack-name>".
func GoogleClouddomainsRegistration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleClouddomainsRegistration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
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
func GoogleClouddomainsRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddomainsRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleClouddomainsRegistration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddomainsRegistration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleClouddomainsRegistration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleClouddomainsRegistration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleClouddomainsRegistration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddomainsRegistration) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) PutContactSettings(value *GoogleClouddomainsRegistrationContactSettings) {
	if err := g.validatePutContactSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContactSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) PutDnsSettings(value *GoogleClouddomainsRegistrationDnsSettings) {
	if err := g.validatePutDnsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDnsSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) PutManagementSettings(value *GoogleClouddomainsRegistrationManagementSettings) {
	if err := g.validatePutManagementSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManagementSettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) PutTimeouts(value *GoogleClouddomainsRegistrationTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) PutYearlyPrice(value *GoogleClouddomainsRegistrationYearlyPrice) {
	if err := g.validatePutYearlyPriceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYearlyPrice",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetContactNotices() {
	_jsii_.InvokeVoid(
		g,
		"resetContactNotices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetDnsSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetDomainNotices() {
	_jsii_.InvokeVoid(
		g,
		"resetDomainNotices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetManagementSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetManagementSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

