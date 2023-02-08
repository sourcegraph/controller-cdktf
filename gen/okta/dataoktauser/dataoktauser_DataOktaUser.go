package dataoktauser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/dataoktauser/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/okta/d/user okta_user}.
type DataOktaUser interface {
	cdktf.TerraformDataSource
	AdminRoles() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	CompoundSearchOperator() *string
	SetCompoundSearchOperator(val *string)
	CompoundSearchOperatorInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostCenter() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CountryCode() *string
	CustomProfileAttributes() *string
	DelayReadSeconds() *string
	SetDelayReadSeconds(val *string)
	DelayReadSecondsInput() *string
	Department() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	Division() *string
	Email() *string
	EmployeeNumber() *string
	FirstName() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GroupMemberships() *[]*string
	HonorificPrefix() *string
	HonorificSuffix() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastName() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	Login() *string
	Manager() *string
	ManagerId() *string
	MiddleName() *string
	MobilePhone() *string
	NickName() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	PostalAddress() *string
	PreferredLanguage() *string
	PrimaryPhone() *string
	ProfileUrl() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Search() DataOktaUserSearchList
	SearchInput() interface{}
	SecondEmail() *string
	SkipGroups() interface{}
	SetSkipGroups(val interface{})
	SkipGroupsInput() interface{}
	SkipRoles() interface{}
	SetSkipRoles(val interface{})
	SkipRolesInput() interface{}
	State() *string
	Status() *string
	StreetAddress() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	Title() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	UserType() *string
	ZipCode() *string
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
	PutSearch(value interface{})
	ResetCompoundSearchOperator()
	ResetDelayReadSeconds()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSearch()
	ResetSkipGroups()
	ResetSkipRoles()
	ResetUserId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataOktaUser
type jsiiProxy_DataOktaUser struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataOktaUser) AdminRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CompoundSearchOperator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compoundSearchOperator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CompoundSearchOperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compoundSearchOperatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) CustomProfileAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProfileAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) DelayReadSeconds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delayReadSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) DelayReadSecondsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delayReadSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) EmployeeNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) GroupMemberships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMemberships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) HonorificPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) HonorificSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Manager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) ManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) NickName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nickName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) PostalAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) PrimaryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) ProfileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Search() DataOktaUserSearchList {
	var returns DataOktaUserSearchList
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SecondEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SkipGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SkipGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SkipRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) SkipRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUser) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/okta/d/user okta_user} Data Source.
func NewDataOktaUser(scope constructs.Construct, id *string, config *DataOktaUserConfig) DataOktaUser {
	_init_.Initialize()

	if err := validateNewDataOktaUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaUser{}

	_jsii_.Create(
		"okta.dataOktaUser.DataOktaUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/okta/d/user okta_user} Data Source.
func NewDataOktaUser_Override(d DataOktaUser, scope constructs.Construct, id *string, config *DataOktaUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.dataOktaUser.DataOktaUser",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataOktaUser)SetCompoundSearchOperator(val *string) {
	if err := j.validateSetCompoundSearchOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compoundSearchOperator",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetDelayReadSeconds(val *string) {
	if err := j.validateSetDelayReadSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delayReadSeconds",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetSkipGroups(val interface{}) {
	if err := j.validateSetSkipGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipGroups",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetSkipRoles(val interface{}) {
	if err := j.validateSetSkipRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRoles",
		val,
	)
}

func (j *jsiiProxy_DataOktaUser)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
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
func DataOktaUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataOktaUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"okta.dataOktaUser.DataOktaUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataOktaUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"okta.dataOktaUser.DataOktaUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataOktaUser) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataOktaUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOktaUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOktaUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOktaUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOktaUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOktaUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOktaUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOktaUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOktaUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaUser) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataOktaUser) PutSearch(value interface{}) {
	if err := d.validatePutSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSearch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataOktaUser) ResetCompoundSearchOperator() {
	_jsii_.InvokeVoid(
		d,
		"resetCompoundSearchOperator",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetDelayReadSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetDelayReadSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetSearch() {
	_jsii_.InvokeVoid(
		d,
		"resetSearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetSkipGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetSkipRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) ResetUserId() {
	_jsii_.InvokeVoid(
		d,
		"resetUserId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataOktaUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

