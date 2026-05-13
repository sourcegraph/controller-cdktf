package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/datastreamstream/internal"
)

type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference interface {
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
	InternalValue() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier
	SetInternalValue(val *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier)
	MongodbIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference
	MongodbIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier
	MysqlIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference
	MysqlIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier
	OracleIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference
	OracleIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier
	PostgresqlIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference
	PostgresqlIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier
	SalesforceIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference
	SalesforceIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier
	SpannerIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference
	SpannerIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier
	SqlServerIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference
	SqlServerIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier
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
	PutMongodbIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier)
	PutMysqlIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier)
	PutOracleIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier)
	PutPostgresqlIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier)
	PutSalesforceIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier)
	PutSpannerIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier)
	PutSqlServerIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier)
	ResetMongodbIdentifier()
	ResetMysqlIdentifier()
	ResetOracleIdentifier()
	ResetPostgresqlIdentifier()
	ResetSalesforceIdentifier()
	ResetSpannerIdentifier()
	ResetSqlServerIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference
type jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InternalValue() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MongodbIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference
	_jsii_.Get(
		j,
		"mongodbIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MongodbIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier
	_jsii_.Get(
		j,
		"mongodbIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MysqlIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference
	_jsii_.Get(
		j,
		"mysqlIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MysqlIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier
	_jsii_.Get(
		j,
		"mysqlIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) OracleIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference
	_jsii_.Get(
		j,
		"oracleIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) OracleIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier
	_jsii_.Get(
		j,
		"oracleIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PostgresqlIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference
	_jsii_.Get(
		j,
		"postgresqlIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PostgresqlIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier
	_jsii_.Get(
		j,
		"postgresqlIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SalesforceIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference
	_jsii_.Get(
		j,
		"salesforceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SalesforceIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier
	_jsii_.Get(
		j,
		"salesforceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SpannerIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference
	_jsii_.Get(
		j,
		"spannerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SpannerIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier
	_jsii_.Get(
		j,
		"spannerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SqlServerIdentifier() DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference {
	var returns DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference
	_jsii_.Get(
		j,
		"sqlServerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SqlServerIdentifierInput() *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier {
	var returns *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier
	_jsii_.Get(
		j,
		"sqlServerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference_Override(d DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetInternalValue(val *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutMongodbIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier) {
	if err := d.validatePutMongodbIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongodbIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutMysqlIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier) {
	if err := d.validatePutMysqlIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutOracleIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier) {
	if err := d.validatePutOracleIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutPostgresqlIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier) {
	if err := d.validatePutPostgresqlIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresqlIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSalesforceIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier) {
	if err := d.validatePutSalesforceIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSalesforceIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSpannerIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier) {
	if err := d.validatePutSpannerIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSpannerIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSqlServerIdentifier(value *DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier) {
	if err := d.validatePutSqlServerIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlServerIdentifier",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetMongodbIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetMysqlIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetOracleIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetPostgresqlIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresqlIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSalesforceIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetSalesforceIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSpannerIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetSpannerIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSqlServerIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlServerIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

