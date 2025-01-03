package lneto

import "errors"

var (
	errShortEth   = errors.New("ethernet length exceeds frame")
	errShortVLAN  = errors.New("ethernet length too short for VLAN")
	errShortUDP   = errors.New("UDP length exceeds frame")
	errBadUDPLen  = errors.New("UDP length invalid")
	errShortIPv4  = errors.New("IPv4 total length exceeds frame")
	errBadIPv4TL  = errors.New("IPv4 short total length")
	errBadIPv4IHL = errors.New("IPv4 bad IHL (<5)")
	errShortIPv6  = errors.New("IPv6 payload length exceeds frame")
	errShortARP   = errors.New("bad ARP size")
	errShortTCP   = errors.New("TCP offset exceeds frame")
	errBadTCPOff  = errors.New("TCP offset invalid")
)

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (efrm EthFrame) ValidateSize() error {
	sz := efrm.EtherTypeOrSize()
	if sz.IsSize() && len(efrm.buf) < int(sz) {
		return errShortEth
	} else if sz == EtherTypeVLAN && len(efrm.buf) < 18 {
		return errShortVLAN
	}
	return nil
}

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (afrm ARPFrame) ValidateSize() error {
	_, hlen := afrm.Hardware()
	_, ilen := afrm.Protocol()
	minLen := 8 + 2*(hlen+ilen)
	if len(afrm.buf) < int(minLen) {
		return errShortARP
	}
	return nil
}

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (ufrm UDPFrame) ValidateSize() error {
	ul := ufrm.Length()
	if ul < sizeHeaderUDP {
		return errBadUDPLen
	} else if int(ul) > len(ufrm.RawData()) {
		return errShortUDP
	}
	return nil
}

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (ifrm IPv4Frame) ValidateSize() error {
	ihl := ifrm.ihl()
	tl := ifrm.TotalLength()
	if tl < sizeHeaderIPv4 {
		return errBadIPv4TL
	} else if int(tl) > len(ifrm.RawData()) {
		return errShortIPv4
	} else if ihl < 5 {
		return errBadIPv4IHL
	}
	return nil
}

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (tfrm TCPFrame) ValidateSize() error {
	off := tfrm.HeaderLength()
	if off < sizeHeaderTCP {
		return errBadTCPOff
	} else if off > len(tfrm.RawData()) {
		return errShortTCP
	}
	return nil
}

// ValidateSize checks the frame's size fields and compares with the actual buffer
// the frame. It returns a non-nil error on finding an inconsistency.
func (i6frm IPv6Frame) ValidateSize() error {
	tl := i6frm.PayloadLength()
	if int(tl)+sizeHeaderIPv6 > len(i6frm.RawData()) {
		return errShortIPv6
	}
	return nil
}
