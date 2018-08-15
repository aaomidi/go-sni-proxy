package request

import (
	"fmt"
	"io"
	"net"
)

type AddrSpec struct {
	Host string
	IP   net.IP
	Port uint8
}

func (a *AddrSpec) String() string {
	if a.Host != "" {
		return fmt.Sprintf("%s (%s):%d", a.Host, a.IP, a.Port)
	}
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}

func GetAddrSpec(reader *io.Reader) (AddrSpec, error) {

}

func ConvertAddrToAddrSpec(p *net.Addr) *AddrSpec {
	addr := *p
	spec := AddrSpec{}

	switch addr := addr.(type) {
	case *net.UDPAddr:
		spec.IP = addr.IP
		spec.Port = uint8(addr.Port)
	case *net.TCPAddr:
		spec.IP = addr.IP
		spec.Port = uint8(addr.Port)
	}
}
