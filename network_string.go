// Code generated by "stringer -trimprefix Network -type Network"; DO NOT EDIT.

package glx

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NetworkTCP-1]
	_ = x[NetworkUDP-2]
}

const _Network_name = "TCPUDP"

var _Network_index = [...]uint8{0, 3, 6}

func (i Network) String() string {
	i -= 1
	if i < 0 || i >= Network(len(_Network_index)-1) {
		return "Network(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Network_name[_Network_index[i]:_Network_index[i+1]]
}
