package googlebigqueryanalyticshublisting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryanalyticshublisting/internal"
)

type GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference interface {
	cdktn.ComplexObject
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
	DataAffinityRegions() *[]*string
	SetDataAffinityRegions(val *[]*string)
	DataAffinityRegionsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryAnalyticsHubListingPubsubTopic
	SetInternalValue(val *GoogleBigqueryAnalyticsHubListingPubsubTopic)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetDataAffinityRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference
type jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) DataAffinityRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataAffinityRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) DataAffinityRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dataAffinityRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) InternalValue() *GoogleBigqueryAnalyticsHubListingPubsubTopic {
	var returns *GoogleBigqueryAnalyticsHubListingPubsubTopic
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryAnalyticsHubListingPubsubTopicOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference_Override(g GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleBigqueryAnalyticsHubListing.GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetDataAffinityRegions(val *[]*string) {
	if err := j.validateSetDataAffinityRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAffinityRegions",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetInternalValue(val *GoogleBigqueryAnalyticsHubListingPubsubTopic) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) ResetDataAffinityRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetDataAffinityRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryAnalyticsHubListingPubsubTopicOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

