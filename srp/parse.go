package srp

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func parse(payload []byte) (SteamRemotePlay, error) {
	logrus.Debugf("Payload: %x", payload)
	srp := SteamRemotePlay{}
	offset := 0

	version, err := getVersion(&offset, payload)
	if err != nil {
		return srp, newParseError(offset, err.Error())
	}
	logrus.Debugf("Version: %d", version)
	srp.Version = version
	name, err := getName(&offset, payload)
	if err != nil {
		return srp, newParseError(offset, err.Error())
	}
	logrus.Debugf("Name: %s", name)
	srp.Name = name

	os, err := getOs(&offset, payload)
	if err != nil {
		return srp, newParseError(offset, err.Error())
	}
	logrus.Debugf("OS: %s", os)
	srp.OS = os

	macs, err := getMacs(&offset, payload)
	if err != nil {
		return srp, newParseError(offset, err.Error())
	}
	logrus.Debugf("MACs %#v", macs)
	srp.MACs = macs

	addrs, err := getAddrs(&offset, payload)
	if err != nil {
		return srp, newParseError(offset, err.Error())
	}
	logrus.Debugf("Addrs: %v", addrs)
	srp.Addrs = addrs
	amp := getAmplification(offset)
	logrus.Debugf("Apmlification: %f", amp)
	srp.Amplification = amp

	return srp, nil
}

func getVersion(offset *int, payload []byte) (int, error) {
	return int(payload[8]), nil
}

func getName(offset *int, payload []byte) (string, error) {
	var name string
	off := *offset
	off += helloAndUnknownBytes

	nullSeen := false
	//name := []byte{}
	var currentByte byte
	// read until we're out of the variable field. I'm missing an obvious length or encoding:(
	for (currentByte != stopAfterNull) || !nullSeen {
		currentByte = payload[off]
		if currentByte == null {
			nullSeen = true
		}
		off++
	}

	length := int(payload[off])
	off++
	name = string(payload[off : off+int(length)])
	off += length
	*offset = off
	return name, nil
}

func safeAdvance(slice []byte, current *int, next int) error {
	maxIndex := len(slice) - 1
	if (*current + next) < maxIndex {
		cast := *current
		*current = cast + next
		return nil
	}
	return errors.New(fmt.Sprintf("Advancing from %d to %d would exceed index of %d", *current, *current+next, maxIndex))
}

func getOs(offset *int, payload []byte) (string, error) {
	var os string

	off := *offset
	osBranch := payload[off+osOffset]

	if osBranch == winBranch {
		*offset = off + winFillerBytes
		return osWin, nil
	}

	if osBranch == linuxBranch {
		*offset = off + linuxFillerBytes
		return osLinux, nil
	}

	if osBranch == macBranch {
		*offset = off + macFillerBytes
		return osMac, nil

	}

	if osBranch == unknownBranch {
		next := payload[off+osOffset+1]
		if next == unknownBranchNext {
			*offset = off + unknownFillerBytes
			return osUnknown, nil
		} else if next == badBranch {
			return os, errors.New(fmt.Sprintf("Known bad byte: %x", next))
		}

	}

	return os, errors.New(fmt.Sprintf("Unknown OS byte: %x", osBranch))
}

func getMacs(offset *int, payload []byte) ([]string, error) {
	off := *offset
	length := int(payload[off])
	currentByte := payload[off]
	//move to the first char
	off++
	macs := []string{}
	for (currentByte != addrDelim) && (currentByte != addrLast) {
		if off+length >= len(payload) {
			return macs, errors.New("Bad parsing causing index overflow")
		}
		newMac := string(payload[off : off+length])
		macs = append(macs, newMac)
		off += length
		if off >= len(payload) {
			return macs, errors.New("Bad parsing causing index overflow")
		}
		currentByte = payload[off]
		if currentByte == macDelim {
			off++
			if off >= len(payload) {
				return macs, errors.New("Bad parsing causing index overflow")
			}
			length = int(payload[off])
			off++
			if off >= len(payload) {
				return macs, errors.New("Bad parsing causing index overflow")
			}
			currentByte = payload[off]
		}
		if off >= len(payload) {
			return macs, errors.New("Bad parsing causing index overflow")
		}

	}
	*offset = off
	return macs, nil
}

func getAddrs(offset *int, payload []byte) ([]string, error) {
	off := *offset
	done := false
	addrs := []string{}
	for !done {
		//jump 2, delimiter is a2/aa 01
		off += addrOffset
		if off >= len(payload) {
			return addrs, errors.New("Bad parsing causing index overflow")
		}
		length := int(payload[off])
		off++
		if off+length >= len(payload) {
			return addrs, errors.New("Bad parsing causing index overflow")
		}
		newAddr := string(payload[off : off+length])
		addrs = append(addrs, newAddr)

		off = off + length
		if off >= len(payload) {
			return addrs, errors.New("Bad parsing causing index overflow")
		}
		currentByte := payload[off]
		if currentByte == addrLast {
			off += addrOffset
			length = int(payload[off])
			off++
			addrs = append(addrs, string(payload[off:off+length]))
			off += length
			done = true
		}
	}
	*offset = off

	return addrs, nil
}
func getAmplification(end int) float64 {
	return float64(end*3) / float64(len(clientHello))
}

type ParseError struct {
	LastOffset int
	Message    string
}

func (p *ParseError) Error() string {
	return p.Message
}

func newParseError(offset int, msg string) *ParseError {
	return &ParseError{Message: msg, LastOffset: offset}
}
