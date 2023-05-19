package googleapigeekeystoresaliasespkcs12

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeekeystoresaliasespkcs12/internal"
)

type GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList interface {
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
	Get(index *float64) GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList
type jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList {
	_init_.Initialize()

	if err := validateNewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList{}

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesPkcs12.GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList_Override(g GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleApigeeKeystoresAliasesPkcs12.GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Get(index *float64) GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

