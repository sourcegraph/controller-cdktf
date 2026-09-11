package privatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/privatecacertificate/internal"
)

type PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList
type jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList{}

	_jsii_.Create(
		"@cdktn/provider-google.privatecaCertificate.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList_Override(p PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.privatecaCertificate.PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := p.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		p,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) Get(index *float64) PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansOutputReference {
	if err := p.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansOutputReference

	_jsii_.Invoke(
		p,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateCertificateDescriptionSubjectDescriptionSubjectAltNameCustomSansList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

