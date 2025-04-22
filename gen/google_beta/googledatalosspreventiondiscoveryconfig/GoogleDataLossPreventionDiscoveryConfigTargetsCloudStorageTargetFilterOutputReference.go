package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference interface {
	cdktf.ComplexObject
	CloudStorageResourceReference() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReferenceOutputReference
	CloudStorageResourceReferenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference
	Collection() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionOutputReference
	CollectionInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection
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
	InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter
	SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter)
	Others() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthersOutputReference
	OthersInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCloudStorageResourceReference(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference)
	PutCollection(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection)
	PutOthers(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers)
	ResetCloudStorageResourceReference()
	ResetCollection()
	ResetOthers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) CloudStorageResourceReference() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReferenceOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReferenceOutputReference
	_jsii_.Get(
		j,
		"cloudStorageResourceReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) CloudStorageResourceReferenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference
	_jsii_.Get(
		j,
		"cloudStorageResourceReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) Collection() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionOutputReference
	_jsii_.Get(
		j,
		"collection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) CollectionInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection
	_jsii_.Get(
		j,
		"collectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) Others() GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthersOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthersOutputReference
	_jsii_.Get(
		j,
		"others",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) OthersInput() *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers
	_jsii_.Get(
		j,
		"othersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference)SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) PutCloudStorageResourceReference(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference) {
	if err := g.validatePutCloudStorageResourceReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudStorageResourceReference",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) PutCollection(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection) {
	if err := g.validatePutCollectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCollection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) PutOthers(value *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers) {
	if err := g.validatePutOthersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOthers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ResetCloudStorageResourceReference() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudStorageResourceReference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ResetCollection() {
	_jsii_.InvokeVoid(
		g,
		"resetCollection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ResetOthers() {
	_jsii_.InvokeVoid(
		g,
		"resetOthers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

