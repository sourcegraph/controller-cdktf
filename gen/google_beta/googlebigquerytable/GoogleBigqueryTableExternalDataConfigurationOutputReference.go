package googlebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerytable/internal"
)

type GoogleBigqueryTableExternalDataConfigurationOutputReference interface {
	cdktf.ComplexObject
	Autodetect() interface{}
	SetAutodetect(val interface{})
	AutodetectInput() interface{}
	AvroOptions() GoogleBigqueryTableExternalDataConfigurationAvroOptionsOutputReference
	AvroOptionsInput() *GoogleBigqueryTableExternalDataConfigurationAvroOptions
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsvOptions() GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference
	CsvOptionsInput() *GoogleBigqueryTableExternalDataConfigurationCsvOptions
	// Experimental.
	Fqn() *string
	GoogleSheetsOptions() GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsOutputReference
	GoogleSheetsOptionsInput() *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions
	HivePartitioningOptions() GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsOutputReference
	HivePartitioningOptionsInput() *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions
	IgnoreUnknownValues() interface{}
	SetIgnoreUnknownValues(val interface{})
	IgnoreUnknownValuesInput() interface{}
	InternalValue() *GoogleBigqueryTableExternalDataConfiguration
	SetInternalValue(val *GoogleBigqueryTableExternalDataConfiguration)
	JsonOptions() GoogleBigqueryTableExternalDataConfigurationJsonOptionsOutputReference
	JsonOptionsInput() *GoogleBigqueryTableExternalDataConfigurationJsonOptions
	MaxBadRecords() *float64
	SetMaxBadRecords(val *float64)
	MaxBadRecordsInput() *float64
	MetadataCacheMode() *string
	SetMetadataCacheMode(val *string)
	MetadataCacheModeInput() *string
	ObjectMetadata() *string
	SetObjectMetadata(val *string)
	ObjectMetadataInput() *string
	ParquetOptions() GoogleBigqueryTableExternalDataConfigurationParquetOptionsOutputReference
	ParquetOptionsInput() *GoogleBigqueryTableExternalDataConfigurationParquetOptions
	ReferenceFileSchemaUri() *string
	SetReferenceFileSchemaUri(val *string)
	ReferenceFileSchemaUriInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	PutAvroOptions(value *GoogleBigqueryTableExternalDataConfigurationAvroOptions)
	PutCsvOptions(value *GoogleBigqueryTableExternalDataConfigurationCsvOptions)
	PutGoogleSheetsOptions(value *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions)
	PutHivePartitioningOptions(value *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions)
	PutJsonOptions(value *GoogleBigqueryTableExternalDataConfigurationJsonOptions)
	PutParquetOptions(value *GoogleBigqueryTableExternalDataConfigurationParquetOptions)
	ResetAvroOptions()
	ResetCompression()
	ResetConnectionId()
	ResetCsvOptions()
	ResetGoogleSheetsOptions()
	ResetHivePartitioningOptions()
	ResetIgnoreUnknownValues()
	ResetJsonOptions()
	ResetMaxBadRecords()
	ResetMetadataCacheMode()
	ResetObjectMetadata()
	ResetParquetOptions()
	ResetReferenceFileSchemaUri()
	ResetSchema()
	ResetSourceFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryTableExternalDataConfigurationOutputReference
type jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) Autodetect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) AutodetectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autodetectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) AvroOptions() GoogleBigqueryTableExternalDataConfigurationAvroOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationAvroOptionsOutputReference
	_jsii_.Get(
		j,
		"avroOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) AvroOptionsInput() *GoogleBigqueryTableExternalDataConfigurationAvroOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationAvroOptions
	_jsii_.Get(
		j,
		"avroOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) CsvOptions() GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference
	_jsii_.Get(
		j,
		"csvOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) CsvOptionsInput() *GoogleBigqueryTableExternalDataConfigurationCsvOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationCsvOptions
	_jsii_.Get(
		j,
		"csvOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GoogleSheetsOptions() GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsOutputReference
	_jsii_.Get(
		j,
		"googleSheetsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GoogleSheetsOptionsInput() *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions
	_jsii_.Get(
		j,
		"googleSheetsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) HivePartitioningOptions() GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsOutputReference
	_jsii_.Get(
		j,
		"hivePartitioningOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) HivePartitioningOptionsInput() *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions
	_jsii_.Get(
		j,
		"hivePartitioningOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) IgnoreUnknownValues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) IgnoreUnknownValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreUnknownValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) InternalValue() *GoogleBigqueryTableExternalDataConfiguration {
	var returns *GoogleBigqueryTableExternalDataConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) JsonOptions() GoogleBigqueryTableExternalDataConfigurationJsonOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationJsonOptionsOutputReference
	_jsii_.Get(
		j,
		"jsonOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) JsonOptionsInput() *GoogleBigqueryTableExternalDataConfigurationJsonOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationJsonOptions
	_jsii_.Get(
		j,
		"jsonOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) MaxBadRecords() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) MaxBadRecordsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) MetadataCacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataCacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) MetadataCacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataCacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ObjectMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ObjectMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ParquetOptions() GoogleBigqueryTableExternalDataConfigurationParquetOptionsOutputReference {
	var returns GoogleBigqueryTableExternalDataConfigurationParquetOptionsOutputReference
	_jsii_.Get(
		j,
		"parquetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ParquetOptionsInput() *GoogleBigqueryTableExternalDataConfigurationParquetOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationParquetOptions
	_jsii_.Get(
		j,
		"parquetOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ReferenceFileSchemaUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceFileSchemaUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ReferenceFileSchemaUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceFileSchemaUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) SourceUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) SourceUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryTableExternalDataConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryTableExternalDataConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryTableExternalDataConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference{}

	_jsii_.Create(
		"google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryTableExternalDataConfigurationOutputReference_Override(g GoogleBigqueryTableExternalDataConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetAutodetect(val interface{}) {
	if err := j.validateSetAutodetectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodetect",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetIgnoreUnknownValues(val interface{}) {
	if err := j.validateSetIgnoreUnknownValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUnknownValues",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetInternalValue(val *GoogleBigqueryTableExternalDataConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetMaxBadRecords(val *float64) {
	if err := j.validateSetMaxBadRecordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBadRecords",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetMetadataCacheMode(val *string) {
	if err := j.validateSetMetadataCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataCacheMode",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetObjectMetadata(val *string) {
	if err := j.validateSetObjectMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectMetadata",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetReferenceFileSchemaUri(val *string) {
	if err := j.validateSetReferenceFileSchemaUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceFileSchemaUri",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetSourceFormat(val *string) {
	if err := j.validateSetSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetSourceUris(val *[]*string) {
	if err := j.validateSetSourceUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceUris",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutAvroOptions(value *GoogleBigqueryTableExternalDataConfigurationAvroOptions) {
	if err := g.validatePutAvroOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAvroOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutCsvOptions(value *GoogleBigqueryTableExternalDataConfigurationCsvOptions) {
	if err := g.validatePutCsvOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCsvOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutGoogleSheetsOptions(value *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions) {
	if err := g.validatePutGoogleSheetsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoogleSheetsOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutHivePartitioningOptions(value *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions) {
	if err := g.validatePutHivePartitioningOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHivePartitioningOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutJsonOptions(value *GoogleBigqueryTableExternalDataConfigurationJsonOptions) {
	if err := g.validatePutJsonOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJsonOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) PutParquetOptions(value *GoogleBigqueryTableExternalDataConfigurationParquetOptions) {
	if err := g.validatePutParquetOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putParquetOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetAvroOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetAvroOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		g,
		"resetCompression",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetConnectionId() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetCsvOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetCsvOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetGoogleSheetsOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetGoogleSheetsOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetHivePartitioningOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetHivePartitioningOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetIgnoreUnknownValues() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreUnknownValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetJsonOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetMaxBadRecords() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxBadRecords",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetMetadataCacheMode() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataCacheMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetObjectMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetObjectMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetParquetOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetParquetOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetReferenceFileSchemaUri() {
	_jsii_.InvokeVoid(
		g,
		"resetReferenceFileSchemaUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

