package dataoktausers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/okta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/okta/dataoktausers/internal"
)

type DataOktaUsersUsersOutputReference interface {
	cdktf.ComplexObject
	AdminRoles() *[]*string
	City() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CostCenter() *string
	CountryCode() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomProfileAttributes() *string
	Department() *string
	DisplayName() *string
	Division() *string
	Email() *string
	EmployeeNumber() *string
	FirstName() *string
	// Experimental.
	Fqn() *string
	GroupMemberships() *[]*string
	HonorificPrefix() *string
	HonorificSuffix() *string
	Id() *string
	InternalValue() *DataOktaUsersUsers
	SetInternalValue(val *DataOktaUsersUsers)
	LastName() *string
	Locale() *string
	Login() *string
	Manager() *string
	ManagerId() *string
	MiddleName() *string
	MobilePhone() *string
	NickName() *string
	Organization() *string
	PostalAddress() *string
	PreferredLanguage() *string
	PrimaryPhone() *string
	ProfileUrl() *string
	SecondEmail() *string
	State() *string
	Status() *string
	StreetAddress() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	Title() *string
	UserType() *string
	ZipCode() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataOktaUsersUsersOutputReference
type jsiiProxy_DataOktaUsersUsersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) AdminRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) CostCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"costCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) CustomProfileAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProfileAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Department() *string {
	var returns *string
	_jsii_.Get(
		j,
		"department",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Division() *string {
	var returns *string
	_jsii_.Get(
		j,
		"division",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) EmployeeNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"employeeNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) GroupMemberships() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupMemberships",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) HonorificPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) HonorificSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) InternalValue() *DataOktaUsersUsers {
	var returns *DataOktaUsersUsers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Manager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) ManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) MobilePhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) NickName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nickName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) PostalAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) PrimaryPhone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryPhone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) ProfileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) SecondEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}


func NewDataOktaUsersUsersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataOktaUsersUsersOutputReference {
	_init_.Initialize()

	if err := validateNewDataOktaUsersUsersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataOktaUsersUsersOutputReference{}

	_jsii_.Create(
		"okta.dataOktaUsers.DataOktaUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataOktaUsersUsersOutputReference_Override(d DataOktaUsersUsersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"okta.dataOktaUsers.DataOktaUsersUsersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference)SetInternalValue(val *DataOktaUsersUsers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataOktaUsersUsersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataOktaUsersUsersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

