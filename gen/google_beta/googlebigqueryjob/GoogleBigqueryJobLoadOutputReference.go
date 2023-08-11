package googlebigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigqueryjob/internal"
)

type GoogleBigqueryJobLoadOutputReference interface {
	cdktf.ComplexObject
	AllowJaggedRows() interface{}
	SetAllowJaggedRows(val interface{})
	AllowJaggedRowsInput() interface{}
	AllowQuotedNewlines() interface{}
	SetAllowQuotedNewlines(val interface{})
	AllowQuotedNewlinesInput() interface{}
	Autodetect() interface{}
	SetAutodetect(val interface{})
	AutodetectInput() interface{}
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
	DestinationEncryptionConfiguration() GoogleBigqueryJobLoadDestinationEncryptionConfigurationOutputReference
	DestinationEncryptionConfigurationInput() *GoogleBigqueryJobLoadDestinationEncryptionConfiguration
	DestinationTable() GoogleBigqueryJobLoadDestinationTableOutputReference
	DestinationTableInput() *GoogleBigqueryJobLoadDestinationTable
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	// Experimental.
	Fqn() *string
	IgnoreUnknownValues() interface{}
	SetIgnoreUnknownValues(val interface{})
	IgnoreUnknownValuesInput() interface{}
	InternalValue() *GoogleBigqueryJobLoad
	SetInternalValue(val *GoogleBigqueryJobLoad)
	JsonExtension() *string
	SetJsonExtension(val *string)
	JsonExtensionInput() *string
	MaxBadRecords() *float64
	SetMaxBadRecords(val *float64)
	MaxBadRecordsInput() *float64
	NullMarker() *string
	SetNullMarker(val *string)
	NullMarkerInput() *string
	ParquetOptions() GoogleBigqueryJobLoadParquetOptionsOutputReference
	ParquetOptionsInput() *GoogleBigqueryJobLoadParquetOptions
	ProjectionFields() *[]*string
	SetProjectionFields(val *[]*string)
	ProjectionFieldsInput() *[]*string
	Quote() *string
	SetQuote(val *string)
	QuoteInput() *string
	SchemaUpdateOptions() *[]*string
	SetSchemaUpdateOptions(val *[]*string)
	SchemaUpdateOptionsInput() *[]*string
	SkipLeadingRows() *float64
	SetSkipLeadingRows(val *float64)
	SkipLeadingRowsInput() *float64
	SourceFormat() *string
	SetSourceFormat(val *string)
	SourceFormatInput() *string
	SourceUris() *[]*string
	SetSourceUris(val *[]*string)
	SourceUrisInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimePartitioning() GoogleBigqueryJobLoadTimePartitioningOutputReference
	TimePartitioningInput() *GoogleBigqueryJobLoadTimePartitioning
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
	PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobLoadDestinationEncryptionConfiguration)
	PutDestinationTable(value *GoogleBigqueryJobLoadDestinationTable)
	PutParquetOptions(value *GoogleBigqueryJobLoadParquetOptions)
	PutTimePartitioning(value *GoogleBigqueryJobLoadTimePartitioning)
	ResetAllowJaggedRows()
	ResetAllowQuotedNewlines()
	ResetAutodetect()
	ResetCreateDisposition()
	ResetDestinationEncryptionConfiguration()
	ResetEncoding()
	ResetFieldDelimiter()
	ResetIgnoreUnknownValues()
	ResetJsonExtension()
	ResetMaxBadRecords()
	ResetNullMarker()
	ResetParquetOptions()
	ResetProjectionFields()
	ResetQuote()
	ResetSchemaUpdateOptions()
	ResetSkipLeadingRows()
	ResetSourceFormat()
	ResetTimePartitioning()
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

// The jsii proxy struct for GoogleBigqueryJobLoadOutputReference
type jsiiProxy_GoogleBigqueryJobLoadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) AllowJaggedRows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) AllowJaggedRowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) AllowQuotedNewlines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) AllowQuotedNewlinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) Autodetect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) AutodetectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) CreateDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) CreateDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) DestinationEncryptionConfiguration() GoogleBigqueryJobLoadDestinationEncryptionConfigurationOutputReference {
	var returns GoogleBigqueryJobLoadDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) DestinationEncryptionConfigurationInput() *GoogleBigqueryJobLoadDestinationEncryptionConfiguration {
	var returns *GoogleBigqueryJobLoadDestinationEncryptionConfiguration
	_jsii_.Get(
		j,
		"destinationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) DestinationTable() GoogleBigqueryJobLoadDestinationTableOutputReference {
	var returns GoogleBigqueryJobLoadDestinationTableOutputReference
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) DestinationTableInput() *GoogleBigqueryJobLoadDestinationTable {
	var returns *GoogleBigqueryJobLoadDestinationTable
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) IgnoreUnknownValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) IgnoreUnknownValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) InternalValue() *GoogleBigqueryJobLoad {
	var returns *GoogleBigqueryJobLoad
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) JsonExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) JsonExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) MaxBadRecords() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) MaxBadRecordsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) NullMarker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullMarker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) NullMarkerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullMarkerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ParquetOptions() GoogleBigqueryJobLoadParquetOptionsOutputReference {
	var returns GoogleBigqueryJobLoadParquetOptionsOutputReference
	_jsii_.Get(
		j,
		"parquetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ParquetOptionsInput() *GoogleBigqueryJobLoadParquetOptions {
	var returns *GoogleBigqueryJobLoadParquetOptions
	_jsii_.Get(
		j,
		"parquetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ProjectionFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectionFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ProjectionFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectionFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) Quote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) QuoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SchemaUpdateOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SchemaUpdateOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SkipLeadingRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SkipLeadingRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SourceUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) SourceUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) TimePartitioning() GoogleBigqueryJobLoadTimePartitioningOutputReference {
	var returns GoogleBigqueryJobLoadTimePartitioningOutputReference
	_jsii_.Get(
		j,
		"timePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) TimePartitioningInput() *GoogleBigqueryJobLoadTimePartitioning {
	var returns *GoogleBigqueryJobLoadTimePartitioning
	_jsii_.Get(
		j,
		"timePartitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) WriteDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference) WriteDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDispositionInput",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryJobLoadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryJobLoadOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryJobLoadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryJobLoadOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobLoadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryJobLoadOutputReference_Override(g GoogleBigqueryJobLoadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryJob.GoogleBigqueryJobLoadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetAllowJaggedRows(val interface{}) {
	if err := j.validateSetAllowJaggedRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJaggedRows",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetAllowQuotedNewlines(val interface{}) {
	if err := j.validateSetAllowQuotedNewlinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowQuotedNewlines",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetAutodetect(val interface{}) {
	if err := j.validateSetAutodetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodetect",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetCreateDisposition(val *string) {
	if err := j.validateSetCreateDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDisposition",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetIgnoreUnknownValues(val interface{}) {
	if err := j.validateSetIgnoreUnknownValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUnknownValues",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetInternalValue(val *GoogleBigqueryJobLoad) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetJsonExtension(val *string) {
	if err := j.validateSetJsonExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonExtension",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetMaxBadRecords(val *float64) {
	if err := j.validateSetMaxBadRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBadRecords",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetNullMarker(val *string) {
	if err := j.validateSetNullMarkerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullMarker",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetProjectionFields(val *[]*string) {
	if err := j.validateSetProjectionFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectionFields",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetQuote(val *string) {
	if err := j.validateSetQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quote",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetSchemaUpdateOptions(val *[]*string) {
	if err := j.validateSetSchemaUpdateOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaUpdateOptions",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetSkipLeadingRows(val *float64) {
	if err := j.validateSetSkipLeadingRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipLeadingRows",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetSourceFormat(val *string) {
	if err := j.validateSetSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetSourceUris(val *[]*string) {
	if err := j.validateSetSourceUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceUris",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryJobLoadOutputReference)SetWriteDisposition(val *string) {
	if err := j.validateSetWriteDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDisposition",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) PutDestinationEncryptionConfiguration(value *GoogleBigqueryJobLoadDestinationEncryptionConfiguration) {
	if err := g.validatePutDestinationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) PutDestinationTable(value *GoogleBigqueryJobLoadDestinationTable) {
	if err := g.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) PutParquetOptions(value *GoogleBigqueryJobLoadParquetOptions) {
	if err := g.validatePutParquetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParquetOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) PutTimePartitioning(value *GoogleBigqueryJobLoadTimePartitioning) {
	if err := g.validatePutTimePartitioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimePartitioning",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetAllowJaggedRows() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowJaggedRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetAllowQuotedNewlines() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowQuotedNewlines",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetAutodetect() {
	_jsii_.InvokeVoid(
		g,
		"resetAutodetect",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetCreateDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetDestinationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		g,
		"resetEncoding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetIgnoreUnknownValues() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreUnknownValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetJsonExtension() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonExtension",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetMaxBadRecords() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBadRecords",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetNullMarker() {
	_jsii_.InvokeVoid(
		g,
		"resetNullMarker",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetParquetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetParquetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetProjectionFields() {
	_jsii_.InvokeVoid(
		g,
		"resetProjectionFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetQuote() {
	_jsii_.InvokeVoid(
		g,
		"resetQuote",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetSchemaUpdateOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetSchemaUpdateOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetSkipLeadingRows() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipLeadingRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetTimePartitioning() {
	_jsii_.InvokeVoid(
		g,
		"resetTimePartitioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ResetWriteDisposition() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteDisposition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryJobLoadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

