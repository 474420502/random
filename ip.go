package random

import (
	"encoding/binary"
	"net"
)

func (ext *Extend) IPv4() net.IP {
	var ipbytes net.IP = make([]byte, 4)
	binary.LittleEndian.PutUint32(ipbytes, ext.Uint32())
	return ipbytes
	// return fmt.Sprintf("%d.%d.%d.%d", bs[0], bs[1], bs[2], bs[3])
}

func (ext *Extend) IPv6() net.IP {
	var ipbytes net.IP = make([]byte, 16)
	binary.LittleEndian.PutUint64(ipbytes, ext.Uint64())
	binary.LittleEndian.PutUint64(ipbytes[8:], ext.Uint64())
	return ipbytes
}
