package bigqueryanalyticshublisting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigqueryanalyticshublisting/internal"
)

type BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList interface {
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
	Get(index *float64) BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList
type jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList {
	_init_.Initialize()

	if err := validateNewBigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList{}

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryAnalyticsHubListing.BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList_Override(b BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryAnalyticsHubListing.BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := b.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		b,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) Get(index *float64) BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingBigqueryDatasetEffectiveReplicasList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

