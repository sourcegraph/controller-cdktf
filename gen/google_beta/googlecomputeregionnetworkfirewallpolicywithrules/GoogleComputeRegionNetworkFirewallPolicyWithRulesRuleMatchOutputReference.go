package googlecomputeregionnetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeregionnetworkfirewallpolicywithrules/internal"
)

type GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference interface {
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
	DestAddressGroups() *[]*string
	SetDestAddressGroups(val *[]*string)
	DestAddressGroupsInput() *[]*string
	DestFqdns() *[]*string
	SetDestFqdns(val *[]*string)
	DestFqdnsInput() *[]*string
	DestIpRanges() *[]*string
	SetDestIpRanges(val *[]*string)
	DestIpRangesInput() *[]*string
	DestNetworkScope() *string
	SetDestNetworkScope(val *string)
	DestNetworkScopeInput() *string
	DestRegionCodes() *[]*string
	SetDestRegionCodes(val *[]*string)
	DestRegionCodesInput() *[]*string
	DestThreatIntelligences() *[]*string
	SetDestThreatIntelligences(val *[]*string)
	DestThreatIntelligencesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch
	SetInternalValue(val *GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch)
	Layer4Config() GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList
	Layer4ConfigInput() interface{}
	SrcAddressGroups() *[]*string
	SetSrcAddressGroups(val *[]*string)
	SrcAddressGroupsInput() *[]*string
	SrcFqdns() *[]*string
	SetSrcFqdns(val *[]*string)
	SrcFqdnsInput() *[]*string
	SrcIpRanges() *[]*string
	SetSrcIpRanges(val *[]*string)
	SrcIpRangesInput() *[]*string
	SrcNetworks() *[]*string
	SetSrcNetworks(val *[]*string)
	SrcNetworkScope() *string
	SetSrcNetworkScope(val *string)
	SrcNetworkScopeInput() *string
	SrcNetworksInput() *[]*string
	SrcRegionCodes() *[]*string
	SetSrcRegionCodes(val *[]*string)
	SrcRegionCodesInput() *[]*string
	SrcSecureTag() GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList
	SrcSecureTagInput() interface{}
	SrcThreatIntelligences() *[]*string
	SetSrcThreatIntelligences(val *[]*string)
	SrcThreatIntelligencesInput() *[]*string
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
	PutLayer4Config(value interface{})
	PutSrcSecureTag(value interface{})
	ResetDestAddressGroups()
	ResetDestFqdns()
	ResetDestIpRanges()
	ResetDestNetworkScope()
	ResetDestRegionCodes()
	ResetDestThreatIntelligences()
	ResetSrcAddressGroups()
	ResetSrcFqdns()
	ResetSrcIpRanges()
	ResetSrcNetworks()
	ResetSrcNetworkScope()
	ResetSrcRegionCodes()
	ResetSrcSecureTag()
	ResetSrcThreatIntelligences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference
type jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestNetworkScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destNetworkScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestNetworkScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destNetworkScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) DestThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InternalValue() *GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch {
	var returns *GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Layer4Config() GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList {
	var returns GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchLayer4ConfigList
	_jsii_.Get(
		j,
		"layer4Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Layer4ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layer4ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcNetworks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcNetworkScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcNetworkScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcNetworkScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcNetworkScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcNetworksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcSecureTag() GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList {
	var returns GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList
	_jsii_.Get(
		j,
		"srcSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcSecureTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcSecureTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) SrcThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionNetworkFirewallPolicyWithRules.GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference_Override(g GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionNetworkFirewallPolicyWithRules.GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestAddressGroups(val *[]*string) {
	if err := j.validateSetDestAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destAddressGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestFqdns(val *[]*string) {
	if err := j.validateSetDestFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destFqdns",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestNetworkScope(val *string) {
	if err := j.validateSetDestNetworkScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destNetworkScope",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestRegionCodes(val *[]*string) {
	if err := j.validateSetDestRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destRegionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetDestThreatIntelligences(val *[]*string) {
	if err := j.validateSetDestThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetInternalValue(val *GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcAddressGroups(val *[]*string) {
	if err := j.validateSetSrcAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAddressGroups",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcFqdns(val *[]*string) {
	if err := j.validateSetSrcFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcFqdns",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcNetworks(val *[]*string) {
	if err := j.validateSetSrcNetworksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcNetworks",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcNetworkScope(val *string) {
	if err := j.validateSetSrcNetworkScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcNetworkScope",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetSrcThreatIntelligences(val *[]*string) {
	if err := j.validateSetSrcThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) PutLayer4Config(value interface{}) {
	if err := g.validatePutLayer4ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLayer4Config",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) PutSrcSecureTag(value interface{}) {
	if err := g.validatePutSrcSecureTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSrcSecureTag",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestAddressGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetDestAddressGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestFqdns() {
	_jsii_.InvokeVoid(
		g,
		"resetDestFqdns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestNetworkScope() {
	_jsii_.InvokeVoid(
		g,
		"resetDestNetworkScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetDestRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetDestThreatIntelligences() {
	_jsii_.InvokeVoid(
		g,
		"resetDestThreatIntelligences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcAddressGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcAddressGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcFqdns() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcFqdns",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcNetworks() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcNetworks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcNetworkScope() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcNetworkScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcSecureTag() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcSecureTag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ResetSrcThreatIntelligences() {
	_jsii_.InvokeVoid(
		g,
		"resetSrcThreatIntelligences",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionNetworkFirewallPolicyWithRulesRuleMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

