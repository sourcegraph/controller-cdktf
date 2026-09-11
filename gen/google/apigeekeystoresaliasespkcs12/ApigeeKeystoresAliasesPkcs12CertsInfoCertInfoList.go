package apigeekeystoresaliasespkcs12

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/apigeekeystoresaliasespkcs12/internal"
)

type ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList interface {
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
	Get(index *float64) ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList
type jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList {
	_init_.Initialize()

	if err := validateNewApigeeKeystoresAliasesPkcs12CertsInfoCertInfoListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList{}

	_jsii_.Create(
		"@cdktn/provider-google.apigeeKeystoresAliasesPkcs12.ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList_Override(a ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.apigeeKeystoresAliasesPkcs12.ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Get(index *float64) ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeKeystoresAliasesPkcs12CertsInfoCertInfoList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

