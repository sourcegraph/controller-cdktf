package googlebigtableauthorizedview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigtableauthorizedview/internal"
)

type GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference interface {
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
	FamilyName() *string
	SetFamilyName(val *string)
	FamilyNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	QualifierPrefixes() *[]*string
	SetQualifierPrefixes(val *[]*string)
	QualifierPrefixesInput() *[]*string
	Qualifiers() *[]*string
	SetQualifiers(val *[]*string)
	QualifiersInput() *[]*string
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
	ResetQualifierPrefixes()
	ResetQualifiers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference
type jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) FamilyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) FamilyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) QualifierPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"qualifierPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) QualifierPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"qualifierPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) Qualifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"qualifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) QualifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"qualifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigtableAuthorizedView.GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference_Override(g GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigtableAuthorizedView.GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetFamilyName(val *string) {
	if err := j.validateSetFamilyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"familyName",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetQualifierPrefixes(val *[]*string) {
	if err := j.validateSetQualifierPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierPrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetQualifiers(val *[]*string) {
	if err := j.validateSetQualifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifiers",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ResetQualifierPrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetQualifierPrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ResetQualifiers() {
	_jsii_.InvokeVoid(
		g,
		"resetQualifiers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigtableAuthorizedViewSubsetViewFamilySubsetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

