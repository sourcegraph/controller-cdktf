package datagoogleiamworkloadidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datagoogleiamworkloadidentitypool/internal"
)

type DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList interface {
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
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList
type jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func NewDataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList {
	_init_.Initialize()

	if err := validateNewDataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleIamWorkloadIdentityPool.DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList_Override(d DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleIamWorkloadIdentityPool.DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList",
		[]any{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]any{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) Get(index *float64) DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]any{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
