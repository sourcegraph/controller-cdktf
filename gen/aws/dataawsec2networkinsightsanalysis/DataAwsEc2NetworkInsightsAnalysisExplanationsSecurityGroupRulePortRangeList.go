package dataawsec2networkinsightsanalysis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/aws/dataawsec2networkinsightsanalysis/internal"
)

type DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList interface {
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
	Get(index *float64) DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList
type jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList {
	_init_.Initialize()

	if err := validateNewDataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList{}

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList_Override(d DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-aws.dataAwsEc2NetworkInsightsAnalysis.DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) Get(index *float64) DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSecurityGroupRulePortRangeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

