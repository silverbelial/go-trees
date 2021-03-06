package trie

import (
	"net"
	//"fmt"
	"fmt"
	"bytes"
)

// Bit Trie router
type BitTrail interface{
	Pop() bool
	Empty() bool
	//Remain() string
	Trail() string
}


// pre defined Byte BitTrail
// Big-Endian order
type btByte struct {
	value	uint8
	bit	int
}

// Pop the Left most available bit
func(b *btByte) Pop() bool{
	if b.Empty() {
		return false
	}
	// bit and operation
	r := (uint8(1)<<uint(b.bit)) & b.value
	b.bit -= 1
	// if has value in the certain bit, return true
	return r > 0
}

// Determine if the Byte is "empty"
func(b *btByte) Empty() bool {
	return b.bit < 0
}

func(b *btByte) Trail() string{
	return fmt.Sprintf("%b/%d",b.value, b.bit)
}

// v, byte value
// maxBit, the used bits count (from little-endian)
func NewByte(v uint8, maxBit int) *btByte {
	if maxBit >= 8 {
		maxBit = 8
	}
	return &btByte{
		value:  v,
		bit: maxBit-1,
	}
}

// pre defined IpNet BitTrail
type btIpNet struct {
	net	*net.IPNet
	bit	int
}

func(b *btIpNet) Pop() bool{
	if b.Empty() {
		return false
	}
	step := b.bit/8
	bits := b.bit%8
	if step > 3 {
		return false
	}
	r := (uint8(1)<<uint(7-bits)) & b.net.IP[step]
	b.bit += 1
	return r>0
}

func(b *btIpNet) Empty() bool {
	ones, _ := b.net.Mask.Size()
	return b.bit >= ones
}

func(b *btIpNet) Trail() string {
	var buffer bytes.Buffer
	for i := 0; i< 24; i++ {
		if i >= b.bit {
			//buffer.WriteString("x")
		} else {
			step := i/8
			bits := i%8
			r := (uint8(1)<<uint(7-bits)) & b.net.IP[step]
			if r > 0 {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
	}
	return buffer.String()
}

func NewTrieIpNet(n *net.IPNet) *btIpNet {
	return &btIpNet{
		net: n,
		bit: 0,
	}
}

// pre defined Ip BitTrail
type TrieIpV4 struct {
	ip net.IP
	bit int
}

func(ip *TrieIpV4) Empty() bool {
	return ip.bit >= 32
}

func(ip *TrieIpV4) Pop() bool {
	if ip.Empty() {
		return false
	}
	step := ip.bit/8
	bits := ip.bit%8
	if step > 3 {
		return false
	}
	//fmt.Println("ip pop:",(uint8(1)<<uint(7-bits)), ip.ip.To4()[step] )
	r := (uint8(1)<<uint(7-bits)) & ip.ip.To4()[step]
	ip.bit += 1
	//fmt.Println("ip r:",r, r> 0)
	return r>0
}

func(ip *TrieIpV4) Trail() string {
	var buffer bytes.Buffer
	for i := 0; i< 24; i++ {
		if i >= ip.bit {
			//buffer.WriteString("x")
		} else {
			step := i/8
			bits := i%8
			r := (uint8(1)<<uint(7-bits)) & ip.ip[step]
			if r > 0 {
				buffer.WriteString("1")
			} else {
				buffer.WriteString("0")
			}
		}
	}
	return buffer.String()
}

func NewTrieIpV4(ip net.IP) *TrieIpV4 {
	return &TrieIpV4{
		ip: ip,
		bit: 0,
	}
}