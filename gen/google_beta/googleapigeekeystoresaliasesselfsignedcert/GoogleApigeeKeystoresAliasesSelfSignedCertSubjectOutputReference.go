package googleapigeekeystoresaliasesselfsignedcert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeekeystoresaliasesselfsignedcert/internal"
)

type GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference interface {
	cdktf.ComplexObject
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleApigeeKeystoresAliasesSelfSignedCertSubject
	SetInternalValue(val *GoogleApigeeKeystoresAliasesSelfSignedCertSubject)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	Org() *string
	SetOrg(val *string)
	OrgInput() *string
	OrgUnit() *string
	SetOrgUnit(val *string)
	OrgUnitInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
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
	ResetCommonName()
	ResetCountryCode()
	ResetEmail()
	ResetLocality()
	ResetOrg()
	ResetOrgUnit()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference
type jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) InternalValue() *GoogleApigeeKeystoresAliasesSelfSignedCertSubject {
	var returns *GoogleApigeeKeystoresAliasesSelfSignedCertSubject
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) Org() *string {
	var returns *string
	_jsii_.Get(
		j,
		"org",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) OrgInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) OrgUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) OrgUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference{}

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference_Override(g GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesSelfSignedCert.GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetInternalValue(val *GoogleApigeeKeystoresAliasesSelfSignedCertSubject) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetLocality(val *string) {
	if err := j.validateSetLocalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetOrg(val *string) {
	if err := j.validateSetOrgParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"org",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetOrgUnit(val *string) {
	if err := j.validateSetOrgUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orgUnit",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		g,
		"resetCommonName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		g,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		g,
		"resetLocality",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetOrg() {
	_jsii_.InvokeVoid(
		g,
		"resetOrg",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetOrgUnit() {
	_jsii_.InvokeVoid(
		g,
		"resetOrgUnit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		g,
		"resetState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesSelfSignedCertSubjectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

