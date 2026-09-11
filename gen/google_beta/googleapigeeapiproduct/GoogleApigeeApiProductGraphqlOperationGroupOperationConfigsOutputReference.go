package googleapigeeapiproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeeapiproduct/internal"
)

type GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference interface {
	cdktn.ComplexObject
	ApiSource() *string
	SetApiSource(val *string)
	ApiSourceInput() *string
	Attributes() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsAttributesList
	AttributesInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Operations() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOperationsList
	OperationsInput() interface{}
	Quota() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuotaOutputReference
	QuotaInput() *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota
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
	PutAttributes(value interface{})
	PutOperations(value interface{})
	PutQuota(value *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota)
	ResetApiSource()
	ResetAttributes()
	ResetOperations()
	ResetQuota()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference
type jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ApiSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ApiSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) Attributes() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsAttributesList {
	var returns GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) Operations() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOperationsList {
	var returns GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOperationsList
	_jsii_.Get(
		j,
		"operations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) OperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) Quota() GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuotaOutputReference {
	var returns GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuotaOutputReference
	_jsii_.Get(
		j,
		"quota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) QuotaInput() *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota {
	var returns *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota
	_jsii_.Get(
		j,
		"quotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference_Override(g GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleApigeeApiProduct.GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetApiSource(val *string) {
	if err := j.validateSetApiSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiSource",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) PutAttributes(value interface{}) {
	if err := g.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAttributes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) PutOperations(value interface{}) {
	if err := g.validatePutOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOperations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) PutQuota(value *GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsQuota) {
	if err := g.validatePutQuotaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQuota",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ResetApiSource() {
	_jsii_.InvokeVoid(
		g,
		"resetApiSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		g,
		"resetAttributes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ResetOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ResetQuota() {
	_jsii_.InvokeVoid(
		g,
		"resetQuota",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeApiProductGraphqlOperationGroupOperationConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

