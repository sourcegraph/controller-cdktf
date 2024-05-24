package googleprivilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleprivilegedaccessmanagerentitlement/internal"
)

type GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference interface {
	cdktf.ComplexObject
	ApprovalsNeeded() *float64
	SetApprovalsNeeded(val *float64)
	ApprovalsNeededInput() *float64
	ApproverEmailRecipients() *[]*string
	SetApproverEmailRecipients(val *[]*string)
	ApproverEmailRecipientsInput() *[]*string
	Approvers() GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference
	ApproversInput() *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers
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
	PutApprovers(value *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers)
	ResetApprovalsNeeded()
	ResetApproverEmailRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference
type jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApprovalsNeeded() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsNeeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApprovalsNeededInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsNeededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproverEmailRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approverEmailRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproverEmailRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approverEmailRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Approvers() GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference {
	var returns GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference
	_jsii_.Get(
		j,
		"approvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproversInput() *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers {
	var returns *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers
	_jsii_.Get(
		j,
		"approversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference_Override(g GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivilegedAccessManagerEntitlement.GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetApprovalsNeeded(val *float64) {
	if err := j.validateSetApprovalsNeededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalsNeeded",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetApproverEmailRecipients(val *[]*string) {
	if err := j.validateSetApproverEmailRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approverEmailRecipients",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) PutApprovers(value *GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers) {
	if err := g.validatePutApproversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApprovers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ResetApprovalsNeeded() {
	_jsii_.InvokeVoid(
		g,
		"resetApprovalsNeeded",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ResetApproverEmailRecipients() {
	_jsii_.InvokeVoid(
		g,
		"resetApproverEmailRecipients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

