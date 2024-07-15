package googlevertexaifeatureonlinestorefeatureview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlevertexaifeatureonlinestorefeatureview/internal"
)

type GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference interface {
	cdktf.ComplexObject
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
	FeatureGroups() GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroupsList
	FeatureGroupsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource
	SetInternalValue(val *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource)
	ProjectNumber() *string
	SetProjectNumber(val *string)
	ProjectNumberInput() *string
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
	PutFeatureGroups(value interface{})
	ResetProjectNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference
type jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) FeatureGroups() GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroupsList {
	var returns GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroupsList
	_jsii_.Get(
		j,
		"featureGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) FeatureGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) InternalValue() *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource {
	var returns *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ProjectNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ProjectNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference_Override(g GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiFeatureOnlineStoreFeatureview.GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetInternalValue(val *GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetProjectNumber(val *string) {
	if err := j.validateSetProjectNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectNumber",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) PutFeatureGroups(value interface{}) {
	if err := g.validatePutFeatureGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFeatureGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ResetProjectNumber() {
	_jsii_.InvokeVoid(
		g,
		"resetProjectNumber",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

