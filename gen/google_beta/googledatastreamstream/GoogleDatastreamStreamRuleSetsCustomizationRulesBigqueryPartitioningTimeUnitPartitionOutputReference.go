package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference interface {
	cdktn.ComplexObject
	Column() *string
	SetColumn(val *string)
	ColumnInput() *string
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
	InternalValue() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
	SetInternalValue(val *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition)
	PartitioningTimeGranularity() *string
	SetPartitioningTimeGranularity(val *string)
	PartitioningTimeGranularityInput() *string
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
	ResetPartitioningTimeGranularity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference
type jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) InternalValue() *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition {
	var returns *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) PartitioningTimeGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitioningTimeGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) PartitioningTimeGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitioningTimeGranularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference_Override(g GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetColumn(val *string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetInternalValue(val *GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetPartitioningTimeGranularity(val *string) {
	if err := j.validateSetPartitioningTimeGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitioningTimeGranularity",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ResetPartitioningTimeGranularity() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitioningTimeGranularity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsCustomizationRulesBigqueryPartitioningTimeUnitPartitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

