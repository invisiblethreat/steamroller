package srp

func parse(payload []byte) SteamRemotePlay {

	macDelim := byte(0x7a)
	addrDelim := byte(0xa2)
	addrLast := byte(0xaa)
	linuxBranch := byte(0xc6)
	winBranch := byte(0x10)
	null := byte(0x00)
	helloAndUnknown := 26 //byte initial offset
	stopAfterNull := byte(0x22)
	winFillerBytes := 33
	linuxFillerBytes := 42

	offset := helloAndUnknown
	//jump to first unspecified field

	nullSeen := false
	//name := []byte{}
	var currentByte byte
	// read until we're out of the variable field. I'm missing an obvious length or encoding:(
	for (currentByte != stopAfterNull) || !nullSeen {
		currentByte = payload[offset]
		if currentByte == null {
			nullSeen = true
		}
		offset++
	}

	length := int(payload[offset])
	offset++
	name := string(payload[offset : offset+int(length)])
	offset += length

	osBranch := payload[offset+3]
	os := ""
	// Users name leaked the OS branch in the payload
	if osBranch == winBranch {
		offset += winFillerBytes
		os = "windows"
	}

	if osBranch == linuxBranch {
		offset += linuxFillerBytes
		os = "linux"
	}

	length = int(payload[offset])
	currentByte = payload[offset]
	//move to the first char
	offset++
	macs := []string{}
	for (currentByte != addrDelim) && (currentByte != addrLast) {
		newMac := string(payload[offset : offset+length])
		macs = append(macs, newMac)
		offset += length
		currentByte = payload[offset]
		if currentByte == macDelim {
			offset++
			length = int(payload[offset])
			offset++
			currentByte = payload[offset]
		}

	}

	done := false
	addrs := []string{}
	for !done {
		//jump 2, delimiter is a2/aa 01
		offset += 2
		length = int(payload[offset])
		offset++
		newAddr := string(payload[offset : offset+length])
		addrs = append(addrs, newAddr)

		offset = offset + length
		currentByte = payload[offset]
		if currentByte == addrLast {
			offset += 2
			length = int(payload[offset])
			offset++
			//newAddr := string(payload[offset : offset+length])
			//fmt.Printf("new addr %s\n", newAddr)
			addrs = append(addrs, string(payload[offset:offset+length]))
			offset += length
			done = true
		}
	}
	//fmt.Printf("%#v\n", addrs)
	amplification := float64(offset*3) / float64(len(steamHello))
	return SteamRemotePlay{Name: name, MACs: macs, Addrs: addrs,
		Amplification: amplification, OS: os}

}
