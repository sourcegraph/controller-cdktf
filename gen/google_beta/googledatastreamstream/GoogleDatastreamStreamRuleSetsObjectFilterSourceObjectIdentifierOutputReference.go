package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatastreamstream/internal"
)

type GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference interface {
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
	InternalValue() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier
	SetInternalValue(val *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier)
	MongodbIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference
	MongodbIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier
	MysqlIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference
	MysqlIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier
	OracleIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference
	OracleIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier
	PostgresqlIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference
	PostgresqlIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier
	SalesforceIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference
	SalesforceIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier
	SpannerIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference
	SpannerIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier
	SqlServerIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference
	SqlServerIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier
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
	PutMongodbIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier)
	PutMysqlIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier)
	PutOracleIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier)
	PutPostgresqlIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier)
	PutSalesforceIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier)
	PutSpannerIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier)
	PutSqlServerIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier)
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

// The jsii proxy struct for GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference
type jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InternalValue() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MongodbIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifierOutputReference
	_jsii_.Get(
		j,
		"mongodbIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MongodbIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier
	_jsii_.Get(
		j,
		"mongodbIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MysqlIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifierOutputReference
	_jsii_.Get(
		j,
		"mysqlIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) MysqlIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier
	_jsii_.Get(
		j,
		"mysqlIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) OracleIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifierOutputReference
	_jsii_.Get(
		j,
		"oracleIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) OracleIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier
	_jsii_.Get(
		j,
		"oracleIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PostgresqlIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifierOutputReference
	_jsii_.Get(
		j,
		"postgresqlIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PostgresqlIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier
	_jsii_.Get(
		j,
		"postgresqlIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SalesforceIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifierOutputReference
	_jsii_.Get(
		j,
		"salesforceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SalesforceIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier
	_jsii_.Get(
		j,
		"salesforceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SpannerIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifierOutputReference
	_jsii_.Get(
		j,
		"spannerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SpannerIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier
	_jsii_.Get(
		j,
		"spannerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SqlServerIdentifier() GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference {
	var returns GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifierOutputReference
	_jsii_.Get(
		j,
		"sqlServerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) SqlServerIdentifierInput() *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier {
	var returns *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier
	_jsii_.Get(
		j,
		"sqlServerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference_Override(g GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetInternalValue(val *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifier) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutMongodbIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier) {
	if err := g.validatePutMongodbIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMongodbIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutMysqlIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMysqlIdentifier) {
	if err := g.validatePutMysqlIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysqlIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutOracleIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOracleIdentifier) {
	if err := g.validatePutOracleIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutPostgresqlIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier) {
	if err := g.validatePutPostgresqlIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSalesforceIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSalesforceIdentifier) {
	if err := g.validatePutSalesforceIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSalesforceIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSpannerIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSpannerIdentifier) {
	if err := g.validatePutSpannerIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSpannerIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) PutSqlServerIdentifier(value *GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierSqlServerIdentifier) {
	if err := g.validatePutSqlServerIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlServerIdentifier",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetMongodbIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetMongodbIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetMysqlIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetMysqlIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetOracleIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetPostgresqlIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSalesforceIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetSalesforceIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSpannerIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetSpannerIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ResetSqlServerIdentifier() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlServerIdentifier",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

