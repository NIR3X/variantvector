package variantvector

import (
	"bytes"
	"testing"
)

func TestVariantVector(t *testing.T) {
	v := Type{
		uint64(1),
		"Hello",
		[]uint8{0x01, 0x02, 0x03},
	}

	packed, err := Pack(v)
	if err != nil {
		t.Error(err)
	}
	for _, b := range packed {
		t.Logf("%02x", b)
	}
	t.Log()

	unpacked, err := Unpack(packed)
	if err != nil {
		t.Error(err)
	}

	if len(unpacked) != len(v) {
		t.Errorf("Assertion failed: len(unpacked) should be equal to len(v)")
	}

	_, ok := unpacked[0].(uint64)
	if !ok {
		t.Errorf("Assertion failed: unpacked[0] should be of type uint64")
	}
	if unpacked[0].(uint64) != v[0].(uint64) {
		t.Errorf("Assertion failed: unpacked[0] should be equal to v[0]")
	}

	_, ok = unpacked[1].(string)
	if !ok {
		t.Errorf("Assertion failed: unpacked[1] should be of type string")
	}
	if unpacked[1].(string) != v[1].(string) {
		t.Errorf("Assertion failed: unpacked[1] should be equal to v[1]")
	}

	_, ok = unpacked[2].([]uint8)
	if !ok {
		t.Errorf("Assertion failed: unpacked[2] should be of type []uint8")
	}
	if !bytes.Equal(unpacked[2].([]uint8), v[2].([]uint8)) {
		t.Errorf("Assertion failed: unpacked[2] should be equal to v[2]")
	}
}
