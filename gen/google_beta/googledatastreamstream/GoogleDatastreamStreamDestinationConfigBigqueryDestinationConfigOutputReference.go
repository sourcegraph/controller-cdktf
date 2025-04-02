package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	AppendOnly() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference
	AppendOnlyInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly
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
	DataFreshness() *string
	SetDataFreshness(val *string)
	DataFreshnessInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig
	SetInternalValue(val *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig)
	Merge() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference
	MergeInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge
	SingleTargetDataset() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference
	SingleTargetDatasetInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset
	SourceHierarchyDatasets() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference
	SourceHierarchyDatasetsInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets
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
	PutAppendOnly(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly)
	PutMerge(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge)
	PutSingleTargetDataset(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset)
	PutSourceHierarchyDatasets(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets)
	ResetAppendOnly()
	ResetDataFreshness()
	ResetMerge()
	ResetSingleTargetDataset()
	ResetSourceHierarchyDatasets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference
type jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) AppendOnly() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference
	_jsii_.Get(
		j,
		"appendOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) AppendOnlyInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly
	_jsii_.Get(
		j,
		"appendOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) DataFreshness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFreshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) DataFreshnessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFreshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InternalValue() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Merge() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference
	_jsii_.Get(
		j,
		"merge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) MergeInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge
	_jsii_.Get(
		j,
		"mergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SingleTargetDataset() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference
	_jsii_.Get(
		j,
		"singleTargetDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SingleTargetDatasetInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset
	_jsii_.Get(
		j,
		"singleTargetDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SourceHierarchyDatasets() GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference {
	var returns GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference
	_jsii_.Get(
		j,
		"sourceHierarchyDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SourceHierarchyDatasetsInput() *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets {
	var returns *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets
	_jsii_.Get(
		j,
		"sourceHierarchyDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference_Override(g GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetDataFreshness(val *string) {
	if err := j.validateSetDataFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshness",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetInternalValue(val *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutAppendOnly(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly) {
	if err := g.validatePutAppendOnlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAppendOnly",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutMerge(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigMerge) {
	if err := g.validatePutMergeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMerge",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutSingleTargetDataset(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset) {
	if err := g.validatePutSingleTargetDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSingleTargetDataset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutSourceHierarchyDatasets(value *GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets) {
	if err := g.validatePutSourceHierarchyDatasetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceHierarchyDatasets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetAppendOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetAppendOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetDataFreshness() {
	_jsii_.InvokeVoid(
		g,
		"resetDataFreshness",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetMerge() {
	_jsii_.InvokeVoid(
		g,
		"resetMerge",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetSingleTargetDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetSingleTargetDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetSourceHierarchyDatasets() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceHierarchyDatasets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

