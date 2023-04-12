package googleorganizationaccessapprovalsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleorganizationaccessapprovalsettings/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_access_approval_settings google_organization_access_approval_settings}.
type GoogleOrganizationAccessApprovalSettings interface {
	cdktf.TerraformResource
	ActiveKeyVersion() *string
	SetActiveKeyVersion(val *string)
	ActiveKeyVersionInput() *string
	AncestorHasActiveKeyVersion() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnrolledAncestor() cdktf.IResolvable
	EnrolledServices() GoogleOrganizationAccessApprovalSettingsEnrolledServicesList
	EnrolledServicesInput() interface{}
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
	InvalidKeyVersion() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	NotificationEmails() *[]*string
	SetNotificationEmails(val *[]*string)
	NotificationEmailsInput() *[]*string
	OrganizationId() *string
	SetOrganizationId(val *string)
	OrganizationIdInput() *string
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
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleOrganizationAccessApprovalSettingsTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutEnrolledServices(value interface{})
	PutTimeouts(value *GoogleOrganizationAccessApprovalSettingsTimeouts)
	ResetActiveKeyVersion()
	ResetId()
	ResetNotificationEmails()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleOrganizationAccessApprovalSettings
type jsiiProxy_GoogleOrganizationAccessApprovalSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ActiveKeyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ActiveKeyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeKeyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) AncestorHasActiveKeyVersion() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ancestorHasActiveKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) EnrolledAncestor() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enrolledAncestor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) EnrolledServices() GoogleOrganizationAccessApprovalSettingsEnrolledServicesList {
	var returns GoogleOrganizationAccessApprovalSettingsEnrolledServicesList
	_jsii_.Get(
		j,
		"enrolledServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) EnrolledServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrolledServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) InvalidKeyVersion() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"invalidKeyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) NotificationEmails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) NotificationEmailsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationEmailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) OrganizationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) Timeouts() GoogleOrganizationAccessApprovalSettingsTimeoutsOutputReference {
	var returns GoogleOrganizationAccessApprovalSettingsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_access_approval_settings google_organization_access_approval_settings} Resource.
func NewGoogleOrganizationAccessApprovalSettings(scope constructs.Construct, id *string, config *GoogleOrganizationAccessApprovalSettingsConfig) GoogleOrganizationAccessApprovalSettings {
	_init_.Initialize()

	if err := validateNewGoogleOrganizationAccessApprovalSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOrganizationAccessApprovalSettings{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_access_approval_settings google_organization_access_approval_settings} Resource.
func NewGoogleOrganizationAccessApprovalSettings_Override(g GoogleOrganizationAccessApprovalSettings, scope constructs.Construct, id *string, config *GoogleOrganizationAccessApprovalSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetActiveKeyVersion(val *string) {
	if err := j.validateSetActiveKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeKeyVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetNotificationEmails(val *[]*string) {
	if err := j.validateSetNotificationEmailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationEmails",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetOrganizationId(val *string) {
	if err := j.validateSetOrganizationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleOrganizationAccessApprovalSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func GoogleOrganizationAccessApprovalSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationAccessApprovalSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOrganizationAccessApprovalSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationAccessApprovalSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleOrganizationAccessApprovalSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleOrganizationAccessApprovalSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleOrganizationAccessApprovalSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-googlebeta.googleOrganizationAccessApprovalSettings.GoogleOrganizationAccessApprovalSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) PutEnrolledServices(value interface{}) {
	if err := g.validatePutEnrolledServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEnrolledServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) PutTimeouts(value *GoogleOrganizationAccessApprovalSettingsTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ResetActiveKeyVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveKeyVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ResetNotificationEmails() {
	_jsii_.InvokeVoid(
		g,
		"resetNotificationEmails",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOrganizationAccessApprovalSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

