package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference interface {
	cdktf.ComplexObject
	CertSign() cdktf.IResolvable
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
	ContentCommitment() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrlSign() cdktf.IResolvable
	DataEncipherment() cdktf.IResolvable
	DecipherOnly() cdktf.IResolvable
	DigitalSignature() cdktf.IResolvable
	EncipherOnly() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage)
	KeyAgreement() cdktf.IResolvable
	KeyEncipherment() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) CertSign() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"certSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) ContentCommitment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"contentCommitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) CrlSign() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) DataEncipherment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) DecipherOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) DigitalSignature() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) EncipherOnly() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage {
	var returns *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) KeyAgreement() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) KeyEncipherment() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference_Override(g GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageBaseKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

