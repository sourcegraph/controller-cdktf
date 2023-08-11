package googlebigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryjob/internal"
)

type GoogleBigqueryJobQueryOutputReference interface {
	cdktf.ComplexObject
	AllowLargeResults() interface{}
	SetAllowLargeResults(val interface{})
	AllowLargeResultsInput() interface{}
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
	CreateDisposition() *string
	SetCreateDisposition(val *string)
	CreateDispositionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDataset() GoogleBigqueryJobQueryDefaultDatasetOutputReference
	DefaultDatasetInput() *GoogleBigqueryJobQueryDefaultDataset
	DestinationEncryptionConfiguration() GoogleBigqueryJobQueryDestinationEncryptionConfigurationOutputReference
	DestinationEncryptionConfigurationInput() *GoogleBigqueryJobQueryDestinationEncryptionConfiguration
	DestinationTable() GoogleBigqueryJobQueryDestinationTableOutputReference
	DestinationTableInput() *GoogleBigqueryJobQueryDestinationTable
	FlattenResults() interface{}
	SetFlattenResults(val interface{})
	FlattenResultsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryJobQuery
	SetInternalValue(val *GoogleBigqueryJobQuery)
	MaximumBillingTier() *float64
	SetMaximumBillingTier(val *float64)
	MaximumBillingTierInput() *float64
	MaximumBytesBilled() *string
	SetMaximumBytesBilled(val *string)
	MaximumBytesBilledInput() *string
	ParameterMode() *string
	SetParameterMode(val *string)
	ParameterModeInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	SchemaUpdateOptions() *[]*string
	SetSchemaUpdateOptions(val *[]*string)
	SchemaUpdateOptionsInput() *[]*string
	ScriptOptions() GoogleBigqueryJobQueryScriptOptionsOutputReference
	ScriptOptionsInput() *GoogleBigqueryJobQueryScriptOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseLegacySql() interface{}
	SetUseLegacySql(val interface{})
	UseLegacySqlInput() interface{}
	UseQueryCache() interface{}
	SetUseQueryCache(val interface{})
	UseQueryCacheInput() interface{}
	UserDefinedFunctionResources() GoogleBigqueryJobQueryUserDefinedFunctionResourcesList
	UserDefinedFunctionResourcesInput() interface{}
	WriteDisposition() *string
	SetWriteDisposition(val *string)
	WriteDispositionInput() *string
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
	PutDefaultDataset(value *GoogleBigqueryJobQueryDefaultDataset)
	PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobQueryDestinationEncryptionConfiguration)
	PutDestinationTable(value *GoogleBigqueryJobQueryDestinationTable)
	PutScriptOptions(value *GoogleBigqueryJobQueryScriptOptions)
	PutUserDefinedFunctionResources(value interface{})
	ResetAllowLargeResults()
	ResetCreateDisposition()
	ResetDefaultDataset()
	ResetDestinationEncryptionConfiguration()
	ResetDestinationTable()
	ResetFlattenResults()
	ResetMaximumBillingTier()
	ResetMaximumBytesBilled()
	ResetParameterMode()
	ResetPriority()
	ResetSchemaUpdateOptions()
	ResetScriptOptions()
	ResetUseLegacySql()
	ResetUseQueryCache()
	ResetUserDefinedFunctionResources()
	ResetWriteDisposition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryJobQueryOutputReference
type jsiiProxy_GoogleBigqueryJobQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) AllowLargeResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLargeResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) AllowLargeResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLargeResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) CreateDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) CreateDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DefaultDataset() GoogleBigqueryJobQueryDefaultDatasetOutputReference {
	var returns GoogleBigqueryJobQueryDefaultDatasetOutputReference
	_jsii_.Get(
		j,
		"defaultDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DefaultDatasetInput() *GoogleBigqueryJobQueryDefaultDataset {
	var returns *GoogleBigqueryJobQueryDefaultDataset
	_jsii_.Get(
		j,
		"defaultDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DestinationEncryptionConfiguration() GoogleBigqueryJobQueryDestinationEncryptionConfigurationOutputReference {
	var returns GoogleBigqueryJobQueryDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DestinationEncryptionConfigurationInput() *GoogleBigqueryJobQueryDestinationEncryptionConfiguration {
	var returns *GoogleBigqueryJobQueryDestinationEncryptionConfiguration
	_jsii_.Get(
		j,
		"destinationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DestinationTable() GoogleBigqueryJobQueryDestinationTableOutputReference {
	var returns GoogleBigqueryJobQueryDestinationTableOutputReference
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) DestinationTableInput() *GoogleBigqueryJobQueryDestinationTable {
	var returns *GoogleBigqueryJobQueryDestinationTable
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) FlattenResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) FlattenResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) InternalValue() *GoogleBigqueryJobQuery {
	var returns *GoogleBigqueryJobQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) MaximumBillingTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBillingTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) MaximumBillingTierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBillingTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) MaximumBytesBilled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBytesBilled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) MaximumBytesBilledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBytesBilledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ParameterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ParameterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) SchemaUpdateOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) SchemaUpdateOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ScriptOptions() GoogleBigqueryJobQueryScriptOptionsOutputReference {
	var returns GoogleBigqueryJobQueryScriptOptionsOutputReference
	_jsii_.Get(
		j,
		"scriptOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ScriptOptionsInput() *GoogleBigqueryJobQueryScriptOptions {
	var returns *GoogleBigqueryJobQueryScriptOptions
	_jsii_.Get(
		j,
		"scriptOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UseLegacySql() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacySql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UseLegacySqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacySqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UseQueryCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useQueryCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UseQueryCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useQueryCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UserDefinedFunctionResources() GoogleBigqueryJobQueryUserDefinedFunctionResourcesList {
	var returns GoogleBigqueryJobQueryUserDefinedFunctionResourcesList
	_jsii_.Get(
		j,
		"userDefinedFunctionResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) UserDefinedFunctionResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFunctionResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) WriteDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference) WriteDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDispositionInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryJobQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryJobQueryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryJobQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryJobQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryJobQueryOutputReference_Override(g GoogleBigqueryJobQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetAllowLargeResults(val interface{}) {
	if err := j.validateSetAllowLargeResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLargeResults",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetCreateDisposition(val *string) {
	if err := j.validateSetCreateDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDisposition",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetFlattenResults(val interface{}) {
	if err := j.validateSetFlattenResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenResults",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetInternalValue(val *GoogleBigqueryJobQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetMaximumBillingTier(val *float64) {
	if err := j.validateSetMaximumBillingTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBillingTier",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetMaximumBytesBilled(val *string) {
	if err := j.validateSetMaximumBytesBilledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesBilled",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetParameterMode(val *string) {
	if err := j.validateSetParameterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetSchemaUpdateOptions(val *[]*string) {
	if err := j.validateSetSchemaUpdateOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaUpdateOptions",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetUseLegacySql(val interface{}) {
	if err := j.validateSetUseLegacySqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLegacySql",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetUseQueryCache(val interface{}) {
	if err := j.validateSetUseQueryCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useQueryCache",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobQueryOutputReference)SetWriteDisposition(val *string) {
	if err := j.validateSetWriteDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDisposition",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PutDefaultDataset(value *GoogleBigqueryJobQueryDefaultDataset) {
	if err := g.validatePutDefaultDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDefaultDataset",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobQueryDestinationEncryptionConfiguration) {
	if err := g.validatePutDestinationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PutDestinationTable(value *GoogleBigqueryJobQueryDestinationTable) {
	if err := g.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PutScriptOptions(value *GoogleBigqueryJobQueryScriptOptions) {
	if err := g.validatePutScriptOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putScriptOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) PutUserDefinedFunctionResources(value interface{}) {
	if err := g.validatePutUserDefinedFunctionResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUserDefinedFunctionResources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetAllowLargeResults() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowLargeResults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetCreateDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetDefaultDataset() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultDataset",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetDestinationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetDestinationTable() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetFlattenResults() {
	_jsii_.InvokeVoid(
		g,
		"resetFlattenResults",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetMaximumBillingTier() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximumBillingTier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetMaximumBytesBilled() {
	_jsii_.InvokeVoid(
		g,
		"resetMaximumBytesBilled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetParameterMode() {
	_jsii_.InvokeVoid(
		g,
		"resetParameterMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		g,
		"resetPriority",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetSchemaUpdateOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaUpdateOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetScriptOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetUseLegacySql() {
	_jsii_.InvokeVoid(
		g,
		"resetUseLegacySql",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetUseQueryCache() {
	_jsii_.InvokeVoid(
		g,
		"resetUseQueryCache",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetUserDefinedFunctionResources() {
	_jsii_.InvokeVoid(
		g,
		"resetUserDefinedFunctionResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ResetWriteDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

