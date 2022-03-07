/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vpcv1

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	ibmVPC "github.com/IBM/vpc-go-sdk/vpcv1"

	"github.com/crossplane-contrib/provider-ibm-cloud/apis/vpcv1/v1alpha1"
	ibmc "github.com/crossplane-contrib/provider-ibm-cloud/pkg/clients"

	"github.com/crossplane/crossplane-runtime/pkg/reference"
)

var (
	// Values below are the ones that will be used if we decide that some parameters has to be non-nil
	addressPrefixVal                   = ibmc.RandomString(false)
	classicAccessVal                   = ibmc.RandomInt(2) == 0
	randomName                         = ibmc.RandomString(true)
	randomResourceGroupID              = ibmc.RandomString(false)
	crnVal                             = ibmc.RandomString(true)
	hrefVal                            = ibmc.RandomString(false)
	idVal                              = ibmc.RandomString(false)
	statusVal                          = ibmc.RandomString(false)
	createdAtVal                       = ibmc.ADateTimeInAYear(2012)
	cseSourceIpsLen                    = ibmc.RandomInt(3)
	cseSourceIpsIPAddress              = ibmc.RandomString(false)
	cseSourceZoneHref                  = ibmc.RandomString(false)
	cseSourceZoneName                  = ibmc.RandomString(false)
	defaultNetworkACLCRN               = ibmc.RandomString(false)
	defaultNetworkACLHref              = ibmc.RandomString(false)
	defaultNetworkACLID                = ibmc.RandomString(false)
	defaultNetworkACLName              = ibmc.RandomString(false)
	defaultNetworkACLDeletedMoreInfo   = ibmc.RandomString(false)
	defaultRoutingTableDeletedMoreInfo = ibmc.RandomString(false)
	defaultRoutingTableHref            = ibmc.RandomString(false)
	defaultRoutingTableID              = ibmc.RandomString(false)
	defaultRoutingTableName            = ibmc.RandomString(false)
	defaultRoutingTableResourceType    = ibmc.RandomString(false)
	defaultSecurityGroupCRN            = ibmc.RandomString(false)
	defaultSecurityGroupHref           = ibmc.RandomString(false)
	defaultSecurityGroupID             = ibmc.RandomString(false)
	defaultSecurityGroupName           = ibmc.RandomString(false)
	resourceGroupName                  = ibmc.RandomString(false)
	resourceGroupHref                  = ibmc.RandomString(false)

	headersMapVal = map[string]string{"a": "b", "c": "d"} // maps cannot be constants hence var. Do not modify.
)

// GenerateSomeCombinations returns "some"  combinations (of booleans) for a given number of elements. Eg if numElems == 1,
// 4 (ie 2^2) combibations will be returned...
//
// Params
// 	  numElems - the number of elements of each array.. in the return array
//    returnSize - size of each return array (may require padding. Which can be with random vars or not...)
//    randAll - randomize the elements that "fill" the random array (we pad at the beginning). false = no randomization => deterministic values
//
// Returns
//    an array of boolean arrays (each of the given size), containing the combinations
func GenerateSomeCombinations(numElems int, returnSize int, randAll bool) [][]bool {
	result := make([][]bool, 0)

	if !randAll {
		rand.Seed(int64(numElems))
	} else {
		rand.Seed(time.Now().UnixNano())
	}

	for _, booleanCombSubset := range generateCombinations(numElems) {
		// Add as many variables as there are parameters
		booleanComb := make([]bool, returnSize)
		copy(booleanComb[returnSize-len(booleanCombSubset):], booleanCombSubset)

		for j := 0; j < returnSize-len(booleanCombSubset); j++ {
			booleanComb[j] = rand.Intn(1) == 1 // nolint - this is ok as we are not doing critical stuff here...)
		}

		result = append(result, booleanComb)
	}

	return result
}

// GetBinaryRep generates a binary representation in string format
//
// Params
//      i  - an integer >= 0
//      size  >= 2^i
//
// Returns
//      a string with binary representation of the integer, of length == size
func GetBinaryRep(i int, size int) string {
	result := strconv.FormatInt(int64(i), 2)

	for len(result) < size {
		result = "0" + result
	}

	return result
}

// Returns all the combinations (of booleans) for a given number of elements
//
// Params
// 	  numElems - the number of elements
//
// Returns
//    an array of boolean arrays
func generateCombinations(numElems int) [][]bool {
	result := make([][]bool, 0)

	for i := 0; i < int(math.Pow(2, float64(numElems))); i++ {
		str := GetBinaryRep(i, numElems)

		boolArray := make([]bool, numElems)
		boolArrayIdx := len(boolArray) - 1
		for j := len(str) - 1; j >= 0; j-- {
			boolArray[boolArrayIdx] = (str[j] == '1')

			boolArrayIdx--
		}

		result = append(result, boolArray)
	}

	return result
}

// GetDummyCloudVPCObservation returns a dummy object, ready to be used in create-VPC-in-the-cloud request. Non-nil values will be the
// ones of the local constants above.
//
// Params
//		each param is used in controlling the value of the similarly-named field in the returned structure
//
// Returns
//	    an object appropriately populated
func GetDummyCloudVPCObservation( // nolint:gocyclo
	createdAtNonNil bool,
	hrefNonNil bool,
	statusNonNil bool,
	cseSourceIpsIPAdressNonNil bool,
	cseSourceIpsZoneNonNil bool,
	cseSourceIpsZoneHrefNonNil bool,
	cseSourceIpsZoneNameNonNil bool,
	defaultNetworkACLNonNil bool,
	defaultNetworkACLCRNNonNil bool,
	defaultNetworkACLDeletedNonNil bool,
	defaultNetworkACLDeletedMoreInfoNonNil bool,
	defaultNetworkACLHrefNonNil bool,
	defaultNetworkACLIDNonNil bool,
	defaultNetworkACLNameNonNil bool,
	defaultRoutingTableNonNil bool,
	defaultRoutingTableDeletedNonNil bool,
	defaultRoutingTableDeletedMoreInfoNonNil,
	defaultRoutingTableHrefNonNil bool,
	defaultRoutingTableIDNonNil bool,
	defaultRoutingTableNameNonNil bool,
	defaultRoutingTableResourceTypeNonNil bool,
	defaultSecurityGroupNonNil bool,
	defaultSecurityGroupCRNNonNil bool,
	defaultSecurityGroupDeletedNonNil bool,
	defaultSecurityGroupDeletedMoreInfoNonNil bool,
	defaultSecurityGroupHrefNonNil bool,
	defaultSecurityGroupIDNonNil bool,
	defaultSecurityGroupNameNonNil bool,
	resourceGroupNonNil bool,
	resourceGroupNameNonNil bool,
	resourceGroupHrefNonNil bool,
	resourceGroupIDNonNil bool) ibmVPC.VPC {

	result := ibmVPC.VPC{
		ClassicAccess: &classicAccessVal,
		CreatedAt:     ibmc.ReturnConditionalDate(createdAtNonNil, createdAtVal),
		CRN:           reference.ToPtrValue(crnVal),
		Href:          ibmc.ReturnConditionalStr(hrefNonNil, hrefVal),
		ID:            reference.ToPtrValue(idVal),
		Name:          reference.ToPtrValue(randomName),
		Status:        ibmc.ReturnConditionalStr(statusNonNil, statusVal),
	}

	if cseSourceIpsLen > 0 {
		result.CseSourceIps = make([]ibmVPC.VpccseSourceIP, cseSourceIpsLen)

		for i := 0; i < cseSourceIpsLen; i++ {
			result.CseSourceIps[i] = ibmVPC.VpccseSourceIP{
				IP: &ibmVPC.IP{
					Address: ibmc.ReturnConditionalStr(cseSourceIpsIPAdressNonNil, cseSourceIpsIPAddress),
				},
			}

			if cseSourceIpsZoneNonNil {
				result.CseSourceIps[i].Zone = &ibmVPC.ZoneReference{
					Href: ibmc.ReturnConditionalStr(cseSourceIpsZoneHrefNonNil, cseSourceZoneHref),
					Name: ibmc.ReturnConditionalStr(cseSourceIpsZoneNameNonNil, cseSourceZoneName),
				}
			}
		}
	}

	if defaultNetworkACLNonNil {
		result.DefaultNetworkACL = &ibmVPC.NetworkACLReference{
			CRN:  ibmc.ReturnConditionalStr(defaultNetworkACLCRNNonNil, defaultNetworkACLCRN),
			Href: ibmc.ReturnConditionalStr(defaultNetworkACLHrefNonNil, defaultNetworkACLHref),
			ID:   ibmc.ReturnConditionalStr(defaultNetworkACLIDNonNil, defaultNetworkACLID),
			Name: ibmc.ReturnConditionalStr(defaultNetworkACLNameNonNil, defaultNetworkACLName),
		}

		if defaultNetworkACLDeletedNonNil {
			result.DefaultNetworkACL.Deleted = &ibmVPC.NetworkACLReferenceDeleted{
				MoreInfo: ibmc.ReturnConditionalStr(defaultNetworkACLDeletedMoreInfoNonNil, defaultNetworkACLDeletedMoreInfo),
			}
		}
	}

	if defaultRoutingTableNonNil {
		result.DefaultRoutingTable = &ibmVPC.RoutingTableReference{
			Href:         ibmc.ReturnConditionalStr(defaultRoutingTableHrefNonNil, defaultRoutingTableHref),
			ID:           ibmc.ReturnConditionalStr(defaultRoutingTableIDNonNil, defaultRoutingTableID),
			Name:         ibmc.ReturnConditionalStr(defaultRoutingTableNameNonNil, defaultRoutingTableName),
			ResourceType: ibmc.ReturnConditionalStr(defaultRoutingTableResourceTypeNonNil, defaultRoutingTableResourceType),
		}

		if defaultRoutingTableDeletedNonNil {
			result.DefaultRoutingTable.Deleted = &ibmVPC.RoutingTableReferenceDeleted{
				MoreInfo: ibmc.ReturnConditionalStr(defaultRoutingTableDeletedMoreInfoNonNil, defaultRoutingTableDeletedMoreInfo),
			}
		}
	}

	if defaultSecurityGroupNonNil {
		result.DefaultSecurityGroup = &ibmVPC.SecurityGroupReference{
			CRN:  ibmc.ReturnConditionalStr(defaultSecurityGroupCRNNonNil, defaultSecurityGroupCRN),
			Href: ibmc.ReturnConditionalStr(defaultSecurityGroupHrefNonNil, defaultSecurityGroupHref),
			ID:   ibmc.ReturnConditionalStr(defaultSecurityGroupIDNonNil, defaultSecurityGroupID),
			Name: ibmc.ReturnConditionalStr(defaultSecurityGroupNameNonNil, defaultSecurityGroupName),
		}

		if defaultSecurityGroupDeletedNonNil {
			result.DefaultSecurityGroup.Deleted = &ibmVPC.SecurityGroupReferenceDeleted{
				MoreInfo: ibmc.ReturnConditionalStr(defaultNetworkACLDeletedMoreInfoNonNil, defaultNetworkACLDeletedMoreInfo),
			}
		}
	}

	if resourceGroupNonNil {
		result.ResourceGroup = &ibmVPC.ResourceGroupReference{
			Href: ibmc.ReturnConditionalStr(resourceGroupHrefNonNil, resourceGroupHref),
			ID:   ibmc.ReturnConditionalStr(resourceGroupIDNonNil, randomResourceGroupID),
			Name: ibmc.ReturnConditionalStr(resourceGroupNameNonNil, resourceGroupName),
		}
	}

	return result
}

// GetDummyCrossplaneVPCParams returns a dummy object, ready to be used in a create-VPC-in-k8s request. Non-nil values will be the
// ones of the local constants above
//
// Params
//		addressNil - whether to set the 'AddressPrefixManagement' member to nil
// 		nameNil - whether to set the 'Name' member to nil
//		resourceGroupIDNil - whether to set the 'resourceGroupIDNil' member to nil
//      noHeaders - whether to include headers
//
// Returns
//	    an object appropriately populated
func GetDummyCrossplaneVPCParams(addressNil bool, nameNil bool, resourceGroupIDNil bool, noHeaders bool) v1alpha1.VPCParameters {
	result := v1alpha1.VPCParameters{}

	if !addressNil {
		result.AddressPrefixManagement = reference.ToPtrValue(addressPrefixVal)
	}

	result.ClassicAccess = classicAccessVal

	if !nameNil {
		result.Name = reference.ToPtrValue(randomName)
	}

	if !resourceGroupIDNil {
		result.ResourceGroup = &v1alpha1.ResourceGroupIdentity{
			ID: randomResourceGroupID,
		}
	}

	if !noHeaders {
		result.Headers = &headersMapVal
	}

	return result
}
