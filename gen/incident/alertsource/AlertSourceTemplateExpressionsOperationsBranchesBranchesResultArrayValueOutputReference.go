package alertsource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/incident/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/incident/alertsource/internal"
)

type AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference interface {
	cdktn.ComplexObject
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
	Literal() *string
	SetLiteral(val *string)
	LiteralInput() *string
	Reference() *string
	SetReference(val *string)
	ReferenceInput() *string
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
	ResetLiteral()
	ResetReference()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference
type jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) Literal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"literal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) LiteralInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"literalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) Reference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference {
	_init_.Initialize()

	if err := validateNewAlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-incident.alertSource.AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference_Override(a AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-incident.alertSource.AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetLiteral(val *string) {
	if err := j.validateSetLiteralParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"literal",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetReference(val *string) {
	if err := j.validateSetReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reference",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ResetLiteral() {
	_jsii_.InvokeVoid(
		a,
		"resetLiteral",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ResetReference() {
	_jsii_.InvokeVoid(
		a,
		"resetReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertSourceTemplateExpressionsOperationsBranchesBranchesResultArrayValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

