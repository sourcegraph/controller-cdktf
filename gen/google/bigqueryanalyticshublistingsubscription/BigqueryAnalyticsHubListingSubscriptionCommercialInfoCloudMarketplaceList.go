package bigqueryanalyticshublistingsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/bigqueryanalyticshublistingsubscription/internal"
)

type BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList interface {
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
	Get(index *float64) BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList
type jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList {
	_init_.Initialize()

	if err := validateNewBigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList{}

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryAnalyticsHubListingSubscription.BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList_Override(b BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.bigqueryAnalyticsHubListingSubscription.BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (b *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) Get(index *float64) BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryAnalyticsHubListingSubscriptionCommercialInfoCloudMarketplaceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

