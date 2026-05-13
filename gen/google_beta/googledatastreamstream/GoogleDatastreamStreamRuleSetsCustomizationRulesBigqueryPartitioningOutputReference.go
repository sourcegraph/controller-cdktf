package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference interface {
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
	// Experimental.
	Fqn() *string
	IngestionTimePartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference
	IngestionTimePartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	IntegerRangePartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference
	IntegerRangePartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition
	InternalValue() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning
	SetInternalValue(val *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning)
	RequirePartitionFilter() interface{}
	SetRequirePartitionFilter(val interface{})
	RequirePartitionFilterInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeUnitPartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference
	TimeUnitPartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
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
	PutIngestionTimePartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition)
	PutIntegerRangePartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition)
	PutTimeUnitPartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition)
	ResetIngestionTimePartition()
	ResetIntegerRangePartition()
	ResetRequirePartitionFilter()
	ResetTimeUnitPartition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference
type jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IngestionTimePartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference {
	var returns GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartitionOutputReference
	_jsii_.Get(
		j,
		"ingestionTimePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IngestionTimePartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition {
	var returns *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition
	_jsii_.Get(
		j,
		"ingestionTimePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IntegerRangePartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference {
	var returns GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartitionOutputReference
	_jsii_.Get(
		j,
		"integerRangePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) IntegerRangePartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition {
	var returns *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition
	_jsii_.Get(
		j,
		"integerRangePartitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InternalValue() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning {
	var returns *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) RequirePartitionFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) RequirePartitionFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requirePartitionFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TimeUnitPartition() GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference {
	var returns GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference
	_jsii_.Get(
		j,
		"timeUnitPartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) TimeUnitPartitionInput() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition {
	var returns *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
	_jsii_.Get(
		j,
		"timeUnitPartitionInput",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference_Override(g GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetInternalValue(val *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioning) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetRequirePartitionFilter(val interface{}) {
	if err := j.validateSetRequirePartitionFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirePartitionFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutIngestionTimePartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIngestionTimePartition) {
	if err := g.validatePutIngestionTimePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngestionTimePartition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutIntegerRangePartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningIntegerRangePartition) {
	if err := g.validatePutIntegerRangePartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIntegerRangePartition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) PutTimeUnitPartition(value *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition) {
	if err := g.validatePutTimeUnitPartitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeUnitPartition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetIngestionTimePartition() {
	_jsii_.InvokeVoid(
		g,
		"resetIngestionTimePartition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetIntegerRangePartition() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegerRangePartition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetRequirePartitionFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetRequirePartitionFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ResetTimeUnitPartition() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeUnitPartition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

