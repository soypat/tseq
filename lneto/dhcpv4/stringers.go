// Code generated by "stringer -type=OptNum,Op,MessageType,ClientState -linecomment -output stringers.go"; DO NOT EDIT.

package dhcpv4

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OptWordAligned-0]
	_ = x[OptSubnetMask-1]
	_ = x[OptTimeOffset-2]
	_ = x[OptRouter-3]
	_ = x[OptTimeServers-4]
	_ = x[OptNameServers-5]
	_ = x[OptDNSServers-6]
	_ = x[OptLogServers-7]
	_ = x[OptCookieServers-8]
	_ = x[OptLPRServers-9]
	_ = x[OptImpressServers-10]
	_ = x[OptRLPServers-11]
	_ = x[OptHostName-12]
	_ = x[OptBootFileSize-13]
	_ = x[OptMeritDumpFile-14]
	_ = x[OptDomainName-15]
	_ = x[OptSwapServer-16]
	_ = x[OptRootPath-17]
	_ = x[OptExtensionFile-18]
	_ = x[OptIPLayerForwarding-19]
	_ = x[OptSrcrouteenabler-20]
	_ = x[OptPolicyFilter-21]
	_ = x[OptMaximumDGReassemblySize-22]
	_ = x[OptDefaultIPTTL-23]
	_ = x[OptPathMTUAgingTimeout-24]
	_ = x[OptMTUPlateau-25]
	_ = x[OptInterfaceMTUSize-26]
	_ = x[OptAllSubnetsAreLocal-27]
	_ = x[OptBroadcastAddress-28]
	_ = x[OptPerformMaskDiscovery-29]
	_ = x[OptProvideMasktoOthers-30]
	_ = x[OptPerformRouterDiscovery-31]
	_ = x[OptRouterSolicitationAddress-32]
	_ = x[OptStaticRoutingTable-33]
	_ = x[OptTrailerEncapsulation-34]
	_ = x[OptARPCacheTimeout-35]
	_ = x[OptEthernetEncapsulation-36]
	_ = x[OptDefaultTCPTimetoLive-37]
	_ = x[OptTCPKeepaliveInterval-38]
	_ = x[OptTCPKeepaliveGarbage-39]
	_ = x[OptNISDomainName-40]
	_ = x[OptNISServerAddresses-41]
	_ = x[OptNTPServersAddresses-42]
	_ = x[OptVendorSpecificInformation-43]
	_ = x[OptNetBIOSNameServer-44]
	_ = x[OptNetBIOSDatagramDistribution-45]
	_ = x[OptNetBIOSNodeType-46]
	_ = x[OptNetBIOSScope-47]
	_ = x[OptXWindowFontServer-48]
	_ = x[OptXWindowDisplayManager-49]
	_ = x[OptRequestedIPaddress-50]
	_ = x[OptIPAddressLeaseTime-51]
	_ = x[OptOptionOverload-52]
	_ = x[OptMessageType-53]
	_ = x[OptServerIdentification-54]
	_ = x[OptParameterRequestList-55]
	_ = x[OptMessage-56]
	_ = x[OptMaximumMessageSize-57]
	_ = x[OptRenewTimeValue-58]
	_ = x[OptRebindingTimeValue-59]
	_ = x[OptClientIdentifier-60]
	_ = x[OptClientIdentifier1-61]
}

const _OptNum_name = "word-alignedsubnet maskTime offset in seconds from UTCN/4 router addressesN/4 time server addressesN/4 IEN-116 server addressesN/4 DNS server addressesN/4 logging server addressesN/4 quote server addressesN/4 printer server addressesN/4 impress server addressesN/4 RLP server addressesHostname stringSize of boot file in 512 byte chunksClient to dump and name of file to dump toThe DNS domain name of the clientSwap server addressesPath name for root diskPatch name for more BOOTP infoEnable or disable IP forwardingEnable or disable source routingRouting policy filtersMaximum datagram reassembly sizeDefault IP time-to-livePath MTU aging timeoutPath MTU plateau tableInterface MTU sizeAll subnets are localBroadcast addressPerform mask discoveryProvide mask to othersPerform router discoveryRouter solicitation addressStatic routing tableTrailer encapsulationARP cache timeoutEthernet encapsulationDefault TCP time to liveTCP keepalive intervalTCP keepalive garbageNIS domain nameNIS server addressesNTP servers addressesVendor specific informationNetBIOS name serverNetBIOS datagram distributionNetBIOS node typeNetBIOS scopeX window font serverX window display managerRequested IP addressIP address lease timeOverload “sname” or “file”DHCP message type.DHCP server identificationParameter request listDHCP error messageDHCP maximum message sizeDHCP renewal (T1) timeDHCP rebinding (T2) timeClient identifierClient identifier(1)"

var _OptNum_index = [...]uint16{0, 12, 23, 54, 74, 99, 127, 151, 179, 205, 233, 261, 285, 300, 336, 378, 411, 432, 455, 485, 516, 548, 570, 602, 625, 647, 669, 687, 708, 725, 747, 769, 793, 820, 840, 861, 878, 900, 924, 946, 967, 982, 1002, 1023, 1050, 1069, 1098, 1115, 1128, 1148, 1172, 1192, 1213, 1247, 1265, 1291, 1313, 1331, 1356, 1378, 1402, 1419, 1439}

func (i OptNum) String() string {
	if i >= OptNum(len(_OptNum_index)-1) {
		return "OptNum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OptNum_name[_OptNum_index[i]:_OptNum_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUndefined-0]
	_ = x[OpRequest-1]
	_ = x[OpReply-2]
}

const _Op_name = "undefinedrequestreply"

var _Op_index = [...]uint8{0, 9, 16, 21}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[msg-0]
	_ = x[MsgDiscover-1]
	_ = x[MsgOffer-2]
	_ = x[MsgRequest-3]
	_ = x[MsgDecline-4]
	_ = x[MsgAck-5]
	_ = x[MsgNack-6]
	_ = x[MsgRelease-7]
	_ = x[MsgInform-8]
}

const _MessageType_name = "undefineddiscoverofferrequestdeclineacknakreleaseinform"

var _MessageType_index = [...]uint8{0, 9, 17, 22, 29, 36, 39, 42, 49, 55}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StateInit-1]
	_ = x[StateSelecting-2]
	_ = x[StateRequesting-3]
	_ = x[StateBound-4]
	_ = x[StateRenewing-5]
	_ = x[StateRebinding-6]
	_ = x[StateInitReboot-7]
	_ = x[StateRebooting-8]
}

const _ClientState_name = "initselectingrequestingboundrenewingrebindinginit-rebootrebooting"

var _ClientState_index = [...]uint8{0, 4, 13, 23, 28, 36, 45, 56, 65}

func (i ClientState) String() string {
	i -= 1
	if i >= ClientState(len(_ClientState_index)-1) {
		return "ClientState(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ClientState_name[_ClientState_index[i]:_ClientState_index[i+1]]
}