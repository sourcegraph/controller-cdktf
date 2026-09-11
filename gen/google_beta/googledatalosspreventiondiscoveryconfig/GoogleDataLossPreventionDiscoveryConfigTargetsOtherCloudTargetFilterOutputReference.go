package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference interface {
	cdktn.ComplexObject
	Collection() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionOutputReference
	CollectionInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter)
	Others() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthersOutputReference
	OthersInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers
	SingleResource() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceOutputReference
	SingleResourceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCollection(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection)
	PutOthers(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers)
	PutSingleResource(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource)
	ResetCollection()
	ResetOthers()
	ResetSingleResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) Collection() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollectionOutputReference
	_jsii_.Get(
		j,
		"collection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) CollectionInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection
	_jsii_.Get(
		j,
		"collectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) Others() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthersOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthersOutputReference
	_jsii_.Get(
		j,
		"others",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) OthersInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers
	_jsii_.Get(
		j,
		"othersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) SingleResource() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResourceOutputReference
	_jsii_.Get(
		j,
		"singleResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) SingleResourceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource
	_jsii_.Get(
		j,
		"singleResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference)SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) PutCollection(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection) {
	if err := g.validatePutCollectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCollection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) PutOthers(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers) {
	if err := g.validatePutOthersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOthers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) PutSingleResource(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource) {
	if err := g.validatePutSingleResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSingleResource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ResetCollection() {
	_jsii_.InvokeVoid(
		g,
		"resetCollection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ResetOthers() {
	_jsii_.InvokeVoid(
		g,
		"resetOthers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ResetSingleResource() {
	_jsii_.InvokeVoid(
		g,
		"resetSingleResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

