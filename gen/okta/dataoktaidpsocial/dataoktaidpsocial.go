package dataoktaidpsocial

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/dataoktaidpsocial/internal"
)

// Represents a {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/data-sources/idp_social okta_idp_social}.
type DataOktaIdpSocial interface {
	cdktf.TerraformDataSource
	AccountLinkAction() *string
	AccountLinkGroupInclude() *[]*string
	AuthorizationBinding() *string
	AuthorizationUrl() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	ClientSecret() *string
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
	DeprovisionedAction() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupsAction() *string
	GroupsAssignment() *[]*string
	GroupsAttribute() *string
	GroupsFilter() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuerMode() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxClockSkew() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProfileMaster() cdktf.IResolvable
	ProtocolType() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisioningAction() *string
	// Experimental.
	RawOverrides() interface{}
	Scopes() *[]*string
	Status() *string
	SubjectMatchAttribute() *string
	SubjectMatchType() *string
	SuspendedAction() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBinding() *string
	TokenUrl() *string
	Type() *string
	UsernameTemplate() *string
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
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataOktaIdpSocial
type jsiiProxy_DataOktaIdpSocial struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOktaIdpSocial) AccountLinkAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountLinkAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) AccountLinkGroupInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountLinkGroupInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) AuthorizationBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) AuthorizationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) DeprovisionedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deprovisionedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) GroupsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) GroupsAssignment() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) GroupsAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) GroupsFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) IssuerMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) MaxClockSkew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClockSkew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ProfileMaster() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"profileMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) ProvisioningAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) SubjectMatchAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) SubjectMatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectMatchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) SuspendedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) TokenBinding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaIdpSocial) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/data-sources/idp_social okta_idp_social} Data Source.
func NewDataOktaIdpSocial(scope constructs.Construct, id *string, config *DataOktaIdpSocialConfig) DataOktaIdpSocial {
	_init_.Initialize()

	if err := validateNewDataOktaIdpSocialParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaIdpSocial{}

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/okta/okta/3.41.0/docs/data-sources/idp_social okta_idp_social} Data Source.
func NewDataOktaIdpSocial_Override(d DataOktaIdpSocial, scope constructs.Construct, id *string, config *DataOktaIdpSocialConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataOktaIdpSocial)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataOktaIdpSocial resource upon running "cdktf plan <stack-name>".
func DataOktaIdpSocial_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataOktaIdpSocial_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
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
func DataOktaIdpSocial_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOktaIdpSocial_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOktaIdpSocial_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOktaIdpSocial_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataOktaIdpSocial_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOktaIdpSocial_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOktaIdpSocial_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-okta.dataOktaIdpSocial.DataOktaIdpSocial",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOktaIdpSocial) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOktaIdpSocial) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOktaIdpSocial) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaIdpSocial) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOktaIdpSocial) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaIdpSocial) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaIdpSocial) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaIdpSocial) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaIdpSocial) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

