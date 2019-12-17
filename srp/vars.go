package srp

const (
	PluginPort   = 27036
	PluginProto  = "udp4"
	PluginString = "steam_remote_play"
	osWin        = "Windows"
	osMac        = "MacOS"
	osLinux      = "Linux"

	macDelim             = byte(0x7a)
	addrDelim            = byte(0xa2)
	addrLast             = byte(0xaa)
	linuxBranch          = byte(0xc6)
	winBranch            = byte(0x10)
	macBranch            = byte(0xac)
	null                 = byte(0x00)
	stopAfterNull        = byte(0x22)
	helloAndUnknownBytes = 26
	winFillerBytes       = 33
	linuxFillerBytes     = 42
	macFillerBytes       = 42
	osOffset             = 3
	addrOffset           = 2
	step                 = 1
)

var clientHello = []byte{0xff, 0xff, 0xff, 0xff, 0x21, 0x4c, 0x5f, 0xa0, 0x05,
	0x00, 0x00, 0x00, 0x08, 0xd2, 0x09, 0x10, 0x00}

var serverHello = []byte{0xff, 0xff, 0xff, 0xff, 0x21, 0x4c, 0x5f, 0xa0, 0x16,
	0x00, 0x00, 0x00, 0x08}
