package variantvector

import (
	"errors"
	"fmt"

	"github.com/NIR3X/varsizedint"
)

type Type []any

func varsizedIntPack(packed []uint8, value uint64) []uint8 {
	bytes := [varsizedint.MaxSize]uint8{}
	bytesSize := varsizedint.Encode(bytes[:], value)
	packed = append(packed, bytes[:bytesSize]...)
	return packed
}

func Pack(variantVec Type) ([]uint8, error) {
	packed := []uint8{}
	packed = varsizedIntPack(packed, uint64(len(variantVec)))

	for _, variant := range variantVec {
		switch variant := variant.(type) {
		case uint64:
			packed = varsizedIntPack(packed, uint64(0))
			packed = varsizedIntPack(packed, variant)
		case string:
			packed = varsizedIntPack(packed, uint64(1))
			packed = varsizedIntPack(packed, uint64(len(variant)))
			packed = append(packed, []uint8(variant)...)
		case []uint8:
			packed = varsizedIntPack(packed, uint64(2))
			packed = varsizedIntPack(packed, uint64(len(variant)))
			packed = append(packed, variant...)
		default:
			return nil, fmt.Errorf("invalid variant type: %T", variant)
		}
	}

	return packed, nil
}

func varsizedIntUnpack(data []byte, offset uint64) (uint64, uint64, error) {
	if uint64(len(data)) < offset+1 {
		return 0, 0, errors.New("insufficient data")
	}
	bytesSize := varsizedint.ParseSize(data[offset:])
	if bytesSize < 0 || uint64(len(data)) < offset+uint64(bytesSize) {
		return 0, 0, errors.New("invalid data")
	}
	x := varsizedint.Decode(data[offset:])
	offset += uint64(bytesSize)
	return x, offset, nil
}

func Unpack(data []uint8) (Type, error) {
	offset := uint64(0)
	variantVecSize, offset, err := varsizedIntUnpack(data, offset)
	if err != nil {
		return nil, err
	}

	variantVec := Type{}
	for i := uint64(0); i < variantVecSize; i++ {
		var variantIndex uint64
		variantIndex, offset, err = varsizedIntUnpack(data, offset)
		if err != nil {
			return nil, err
		}

		switch variantIndex {
		case 0:
			var value uint64
			value, offset, err = varsizedIntUnpack(data, offset)
			if err != nil {
				return nil, err
			}
			variantVec = append(variantVec, value)
		case 1:
			var valueSize uint64
			valueSize, offset, err = varsizedIntUnpack(data, offset)
			if err != nil {
				return nil, err
			}
			if uint64(len(data)) < offset+valueSize {
				return nil, errors.New("insufficient data")
			}
			value := string(data[offset : offset+valueSize])
			offset += valueSize
			variantVec = append(variantVec, value)
		case 2:
			var valueSize uint64
			valueSize, offset, err = varsizedIntUnpack(data, offset)
			if err != nil {
				return nil, err
			}
			if uint64(len(data)) < offset+valueSize {
				return nil, errors.New("insufficient data")
			}
			value := data[offset : offset+valueSize]
			offset += valueSize
			variantVec = append(variantVec, value)
		default:
			return nil, fmt.Errorf("invalid variant index: %d", variantIndex)
		}
	}

	return variantVec, nil
}
