package googlefirestoreindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefirestoreindex/internal"
)

type GoogleFirestoreIndexFieldsSearchConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GeoSpec() GoogleFirestoreIndexFieldsSearchConfigGeoSpecOutputReference
	GeoSpecInput() *GoogleFirestoreIndexFieldsSearchConfigGeoSpec
	InternalValue() *GoogleFirestoreIndexFieldsSearchConfig
	SetInternalValue(val *GoogleFirestoreIndexFieldsSearchConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextSpec() GoogleFirestoreIndexFieldsSearchConfigTextSpecOutputReference
	TextSpecInput() *GoogleFirestoreIndexFieldsSearchConfigTextSpec
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
	PutGeoSpec(value *GoogleFirestoreIndexFieldsSearchConfigGeoSpec)
	PutTextSpec(value *GoogleFirestoreIndexFieldsSearchConfigTextSpec)
	ResetGeoSpec()
	ResetTextSpec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFirestoreIndexFieldsSearchConfigOutputReference
type jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GeoSpec() GoogleFirestoreIndexFieldsSearchConfigGeoSpecOutputReference {
	var returns GoogleFirestoreIndexFieldsSearchConfigGeoSpecOutputReference
	_jsii_.Get(
		j,
		"geoSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GeoSpecInput() *GoogleFirestoreIndexFieldsSearchConfigGeoSpec {
	var returns *GoogleFirestoreIndexFieldsSearchConfigGeoSpec
	_jsii_.Get(
		j,
		"geoSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) InternalValue() *GoogleFirestoreIndexFieldsSearchConfig {
	var returns *GoogleFirestoreIndexFieldsSearchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) TextSpec() GoogleFirestoreIndexFieldsSearchConfigTextSpecOutputReference {
	var returns GoogleFirestoreIndexFieldsSearchConfigTextSpecOutputReference
	_jsii_.Get(
		j,
		"textSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) TextSpecInput() *GoogleFirestoreIndexFieldsSearchConfigTextSpec {
	var returns *GoogleFirestoreIndexFieldsSearchConfigTextSpec
	_jsii_.Get(
		j,
		"textSpecInput",
		&returns,
	)
	return returns
}


func NewGoogleFirestoreIndexFieldsSearchConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleFirestoreIndexFieldsSearchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFirestoreIndexFieldsSearchConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirestoreIndex.GoogleFirestoreIndexFieldsSearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleFirestoreIndexFieldsSearchConfigOutputReference_Override(g GoogleFirestoreIndexFieldsSearchConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleFirestoreIndex.GoogleFirestoreIndexFieldsSearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference)SetInternalValue(val *GoogleFirestoreIndexFieldsSearchConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) PutGeoSpec(value *GoogleFirestoreIndexFieldsSearchConfigGeoSpec) {
	if err := g.validatePutGeoSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGeoSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) PutTextSpec(value *GoogleFirestoreIndexFieldsSearchConfigTextSpec) {
	if err := g.validatePutTextSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTextSpec",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ResetGeoSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetGeoSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ResetTextSpec() {
	_jsii_.InvokeVoid(
		g,
		"resetTextSpec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFirestoreIndexFieldsSearchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

