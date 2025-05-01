package googleapihubapihubinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapihubapihubinstance/internal"
)

type GoogleApihubApiHubInstanceConfigAOutputReference interface {
	cdktf.ComplexObject
	CmekKeyName() *string
	SetCmekKeyName(val *string)
	CmekKeyNameInput() *string
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
	DisableSearch() interface{}
	SetDisableSearch(val interface{})
	DisableSearchInput() interface{}
	EncryptionType() *string
	SetEncryptionType(val *string)
	EncryptionTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleApihubApiHubInstanceConfigA
	SetInternalValue(val *GoogleApihubApiHubInstanceConfigA)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VertexLocation() *string
	SetVertexLocation(val *string)
	VertexLocationInput() *string
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
	ResetCmekKeyName()
	ResetDisableSearch()
	ResetEncryptionType()
	ResetVertexLocation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApihubApiHubInstanceConfigAOutputReference
type jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) CmekKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) CmekKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) DisableSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) DisableSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) EncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) InternalValue() *GoogleApihubApiHubInstanceConfigA {
	var returns *GoogleApihubApiHubInstanceConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) VertexLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) VertexLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexLocationInput",
		&returns,
	)
	return returns
}


func NewGoogleApihubApiHubInstanceConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApihubApiHubInstanceConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApihubApiHubInstanceConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApihubApiHubInstance.GoogleApihubApiHubInstanceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApihubApiHubInstanceConfigAOutputReference_Override(g GoogleApihubApiHubInstanceConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApihubApiHubInstance.GoogleApihubApiHubInstanceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetCmekKeyName(val *string) {
	if err := j.validateSetCmekKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmekKeyName",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetDisableSearch(val interface{}) {
	if err := j.validateSetDisableSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableSearch",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetEncryptionType(val *string) {
	if err := j.validateSetEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetInternalValue(val *GoogleApihubApiHubInstanceConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference)SetVertexLocation(val *string) {
	if err := j.validateSetVertexLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vertexLocation",
		val,
	)
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ResetCmekKeyName() {
	_jsii_.InvokeVoid(
		g,
		"resetCmekKeyName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ResetDisableSearch() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableSearch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ResetEncryptionType() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ResetVertexLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetVertexLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApihubApiHubInstanceConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

