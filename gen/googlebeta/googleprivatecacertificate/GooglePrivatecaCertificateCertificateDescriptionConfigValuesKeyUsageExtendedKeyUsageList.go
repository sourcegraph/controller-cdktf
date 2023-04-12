package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googleprivatecacertificate/internal"
)

type GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList
type jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList_Override(g GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googlePrivatecaCertificate.GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) Get(index *float64) GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateCertificateDescriptionConfigValuesKeyUsageExtendedKeyUsageList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

