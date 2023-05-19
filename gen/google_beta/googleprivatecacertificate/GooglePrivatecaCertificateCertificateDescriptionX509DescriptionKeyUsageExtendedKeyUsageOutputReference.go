package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference interface {
	cdktf.ComplexObject
	ClientAuth() cdktf.IResolvable
	CodeSigning() cdktf.IResolvable
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
	EmailProtection() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage
	SetInternalValue(val *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage)
	OcspSigning() cdktf.IResolvable
	ServerAuth() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeStamping() cdktf.IResolvable
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

// The jsii proxy struct for GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference
type jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ClientAuth() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"clientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) CodeSigning() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"codeSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) EmailProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) InternalValue() *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage {
	var returns *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) OcspSigning() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ocspSigning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ServerAuth() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"serverAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) TimeStamping() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"timeStamping",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference{}

	_jsii_.Create(
		"google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference_Override(g GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference)SetInternalValue(val *GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionX509DescriptionKeyUsageExtendedKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

