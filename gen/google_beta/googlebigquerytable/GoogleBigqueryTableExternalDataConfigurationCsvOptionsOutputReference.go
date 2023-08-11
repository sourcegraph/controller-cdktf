package googlebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebigquerytable/internal"
)

type GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference interface {
	cdktf.ComplexObject
	AllowJaggedRows() interface{}
	SetAllowJaggedRows(val interface{})
	AllowJaggedRowsInput() interface{}
	AllowQuotedNewlines() interface{}
	SetAllowQuotedNewlines(val interface{})
	AllowQuotedNewlinesInput() interface{}
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
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBigqueryTableExternalDataConfigurationCsvOptions
	SetInternalValue(val *GoogleBigqueryTableExternalDataConfigurationCsvOptions)
	Quote() *string
	SetQuote(val *string)
	QuoteInput() *string
	SkipLeadingRows() *float64
	SetSkipLeadingRows(val *float64)
	SkipLeadingRowsInput() *float64
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
	ResetAllowJaggedRows()
	ResetAllowQuotedNewlines()
	ResetEncoding()
	ResetFieldDelimiter()
	ResetSkipLeadingRows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference
type jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) AllowJaggedRows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) AllowJaggedRowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowJaggedRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) AllowQuotedNewlines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) AllowQuotedNewlinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowQuotedNewlinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) InternalValue() *GoogleBigqueryTableExternalDataConfigurationCsvOptions {
	var returns *GoogleBigqueryTableExternalDataConfigurationCsvOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) Quote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) QuoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) SkipLeadingRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) SkipLeadingRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipLeadingRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference_Override(g GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBigqueryTable.GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetAllowJaggedRows(val interface{}) {
	if err := j.validateSetAllowJaggedRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowJaggedRows",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetAllowQuotedNewlines(val interface{}) {
	if err := j.validateSetAllowQuotedNewlinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowQuotedNewlines",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetInternalValue(val *GoogleBigqueryTableExternalDataConfigurationCsvOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetQuote(val *string) {
	if err := j.validateSetQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quote",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetSkipLeadingRows(val *float64) {
	if err := j.validateSetSkipLeadingRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipLeadingRows",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ResetAllowJaggedRows() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowJaggedRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ResetAllowQuotedNewlines() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowQuotedNewlines",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		g,
		"resetEncoding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ResetSkipLeadingRows() {
	_jsii_.InvokeVoid(
		g,
		"resetSkipLeadingRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBigqueryTableExternalDataConfigurationCsvOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

