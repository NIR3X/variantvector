# Variant Vector Serialization/Deserialization (Go)

This Go implementation provides a mechanism for serializing and deserializing a variant vector (`variantvector.Type`). The variant vector can hold elements of different types, including `uint64`, `string`, and `[]uint8`.

## Serialization (Packing)

### `varsizedIntPack` Function
Encodes a `uint64` value using a variable-sized integer encoding and appends the encoded bytes to a slice.

### `Pack` Function
Packs a `variantvector.Type` into a slice of bytes.
- The size of the variant vector is packed first using `varsizedIntPack`.
- For each variant in the slice:
  - The variant index is packed.
  - Depending on the variant index, the corresponding value is packed (either `uint64`, `string`, or `[]uint8`).

## Deserialization (Unpacking)

### `varsizedIntUnpack` Function
Decodes a variable-sized integer from a byte slice and updates the offset.
Returns the decoded value and the updated offset.

### `Unpack` Function
Unpacks a byte slice into a `variantvector.Type`.
- Reads the size of the variant vector using `varsizedIntUnpack`.
- Iterates over the elements in the slice:
  - Reads the variant index using `varsizedIntUnpack`.
  - Based on the variant index, reads and adds the corresponding value to the variant vector.

## Usage

```go
package main

import (
	"fmt"
	"github.com/NIR3X/variantvector"
)

func main() {
	// Creating a variant vector
	v := variantvector.Type{
		uint64(1),
		"Hello",
		[]uint8{0x01, 0x02, 0x03},
	}

	// Packing the variant vector
	packed, err := variantvector.Pack(v)
	if err != nil {
		fmt.Println("Error packing:", err)
		return
	}

	// Displaying the packed bytes
	fmt.Print("Packed Bytes: ")
	for _, b := range packed {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	// Unpacking the bytes
	unpacked, err := variantvector.Unpack(packed)
	if err != nil {
		fmt.Println("Error unpacking:", err)
		return
	}

	// Displaying the unpacked values
	fmt.Println("Unpacked Values:")
	for _, variant := range unpacked {
		switch v := variant.(type) {
		case uint64:
			fmt.Println("uint64:", v)
		case string:
			fmt.Println("string:", v)
		case []uint8:
			fmt.Print("[]uint8: ")
			for _, b := range v {
				fmt.Printf("%02x ", b)
			}
			fmt.Println()
		}
	}
}
```

## Dependencies

* This code requires the `varsizedint` package for encoding and decoding variable-sized integers.

## Notes

* The `variantvector.Type` is used to represent a variant vector.
* Error handling is done using the `error` type.
* Ensure that dependencies are correctly imported for the code to work as intended.

## License

[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
