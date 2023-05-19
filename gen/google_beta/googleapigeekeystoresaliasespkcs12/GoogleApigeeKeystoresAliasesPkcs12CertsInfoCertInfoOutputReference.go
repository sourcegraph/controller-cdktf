package googleapigeekeystoresaliasespkcs12

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeekeystoresaliasespkcs12/internal"
)

type GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference interface {
	cdktf.ComplexObject
	BasicConstraints() *string
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
	ExpiryDate() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfo
	SetInternalValue(val *GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfo)
	Issuer() *string
	IsValid() *string
	PublicKey() *string
	SerialNumber() *string
	SigAlgName() *string
	Subject() *string
	SubjectAlternativeNames() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidFrom() *string
	Version() *float64
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

// The jsii proxy struct for GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference
type jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) BasicConstraints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ExpiryDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) InternalValue() *GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfo {
	var returns *GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) IsValid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"isValid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) PublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) SigAlgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigAlgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference{}

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesPkcs12.GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference_Override(g GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesPkcs12.GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference)SetInternalValue(val *GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

