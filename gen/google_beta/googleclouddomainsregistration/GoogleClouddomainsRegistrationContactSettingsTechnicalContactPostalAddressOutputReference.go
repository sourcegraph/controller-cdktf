package googleclouddomainsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddomainsregistration/internal"
)

type GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference interface {
	cdktf.ComplexObject
	AddressLines() *[]*string
	SetAddressLines(val *[]*string)
	AddressLinesInput() *[]*string
	AdministrativeArea() *string
	SetAdministrativeArea(val *string)
	AdministrativeAreaInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress
	SetInternalValue(val *GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	Recipients() *[]*string
	SetRecipients(val *[]*string)
	RecipientsInput() *[]*string
	RegionCode() *string
	SetRegionCode(val *string)
	RegionCodeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAddressLines()
	ResetAdministrativeArea()
	ResetLocality()
	ResetOrganization()
	ResetPostalCode()
	ResetRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference
type jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) AddressLines() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) AddressLinesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) AdministrativeArea() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeArea",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) AdministrativeAreaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administrativeAreaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) InternalValue() *GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress {
	var returns *GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) Recipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) RecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) RegionCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) RegionCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference_Override(g GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddomainsRegistration.GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetAddressLines(val *[]*string) {
	if err := j.validateSetAddressLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLines",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetAdministrativeArea(val *string) {
	if err := j.validateSetAdministrativeAreaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administrativeArea",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetInternalValue(val *GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetOrganization(val *string) {
	if err := j.validateSetOrganizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetRecipients(val *[]*string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetRegionCode(val *string) {
	if err := j.validateSetRegionCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionCode",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetAddressLines() {
	_jsii_.InvokeVoid(
		g,
		"resetAddressLines",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetAdministrativeArea() {
	_jsii_.InvokeVoid(
		g,
		"resetAdministrativeArea",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		g,
		"resetLocality",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetOrganization() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		g,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ResetRecipients() {
	_jsii_.InvokeVoid(
		g,
		"resetRecipients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddomainsRegistrationContactSettingsTechnicalContactPostalAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

