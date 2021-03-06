// generated by stringer -type=OptionCode; DO NOT EDIT

package dhcp4

import "fmt"

const (
	_OptionCode_name_0 = "PadOptionSubnetMaskOptionTimeOffsetOptionRouterOptionTimeServerOptionNameServerOptionDomainNameServerOptionLogServerOptionCookieServerOptionLPRServerOptionImpressServerOptionResourceLocationServerOptionHostNameOptionBootFileSizeOptionMeritDumpFileOptionDomainNameOptionSwapServerOptionRootPathOptionExtensionsPathOptionIPForwardingEnableDisableOptionNonLocalSourceRoutingEnableDisableOptionPolicyFilterOptionMaximumDatagramReassemblySizeOptionDefaultIPTimeToLiveOptionPathMTUAgingTimeoutOptionPathMTUPlateauTableOptionInterfaceMTUOptionAllSubnetsAreLocalOptionBroadcastAddressOptionPerformMaskDiscoveryOptionMaskSupplierOptionPerformRouterDiscoveryOptionRouterSolicitationAddressOptionStaticRouteOptionTrailerEncapsulationOptionARPCacheTimeoutOptionEthernetEncapsulationOptionTCPDefaultTTLOptionTCPKeepaliveIntervalOptionTCPKeepaliveGarbageOptionNetworkInformationServiceDomainOptionNetworkInformationServersOptionNetworkTimeProtocolServersOptionVendorSpecificInformationOptionNetBIOSOverTCPIPNameServerOptionNetBIOSOverTCPIPDatagramDistributionServerOptionNetBIOSOverTCPIPNodeTypeOptionNetBIOSOverTCPIPScopeOptionXWindowSystemFontServerOptionXWindowSystemDisplayManagerOptionRequestedIPAddressOptionIPAddressLeaseTimeOptionOverloadOptionDHCPMessageTypeOptionServerIdentifierOptionParameterRequestListOptionMessageOptionMaximumDHCPMessageSizeOptionRenewalTimeValueOptionRebindingTimeValueOptionVendorClassIdentifierOptionClientIdentifier"
	_OptionCode_name_1 = "OptionNetworkInformationServicePlusDomainOptionNetworkInformationServicePlusServersOptionTFTPServerNameOptionBootFileNameOptionMobileIPHomeAgentOptionSimpleMailTransportProtocolOptionPostOfficeProtocolServerOptionNetworkNewsTransportProtocolOptionDefaultWorldWideWebServerOptionDefaultFingerServerOptionDefaultInternetRelayChatServerOptionStreetTalkServerOptionStreetTalkDirectoryAssistanceOptionUserClass"
	_OptionCode_name_2 = "OptionRelayAgentInformation"
	_OptionCode_name_3 = "OptionClientArchitecture"
	_OptionCode_name_4 = "OptionTZPOSIXStringOptionTZDatabaseString"
	_OptionCode_name_5 = "OptionClasslessRouteFormat"
	_OptionCode_name_6 = "End"
)

var (
	_OptionCode_index_0 = [...]uint16{0, 3, 19, 35, 47, 63, 79, 101, 116, 134, 149, 168, 196, 210, 228, 247, 263, 279, 293, 313, 344, 384, 402, 437, 462, 487, 512, 530, 554, 576, 602, 620, 648, 679, 696, 722, 743, 770, 789, 815, 840, 877, 908, 940, 971, 1003, 1051, 1081, 1108, 1137, 1170, 1194, 1218, 1232, 1253, 1275, 1301, 1314, 1342, 1364, 1388, 1415, 1437}
	_OptionCode_index_1 = [...]uint16{0, 41, 83, 103, 121, 144, 177, 207, 241, 272, 297, 333, 355, 390, 405}
	_OptionCode_index_2 = [...]uint8{0, 27}
	_OptionCode_index_3 = [...]uint8{0, 24}
	_OptionCode_index_4 = [...]uint8{0, 19, 41}
	_OptionCode_index_5 = [...]uint8{0, 26}
	_OptionCode_index_6 = [...]uint8{0, 3}
)

func (i OptionCode) String() string {
	switch {
	case 0 <= i && i <= 61:
		return _OptionCode_name_0[_OptionCode_index_0[i]:_OptionCode_index_0[i+1]]
	case 64 <= i && i <= 77:
		i -= 64
		return _OptionCode_name_1[_OptionCode_index_1[i]:_OptionCode_index_1[i+1]]
	case i == 82:
		return _OptionCode_name_2
	case i == 93:
		return _OptionCode_name_3
	case 100 <= i && i <= 101:
		i -= 100
		return _OptionCode_name_4[_OptionCode_index_4[i]:_OptionCode_index_4[i+1]]
	case i == 121:
		return _OptionCode_name_5
	case i == 255:
		return _OptionCode_name_6
	default:
		return fmt.Sprintf("OptionCode(%d)", i)
	}
}
