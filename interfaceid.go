// https://eips.ethereum.org/EIPS/eip-165 interface IDs
// Used for ERC721 supportsInterface call.
//
// https://github.com/metachris/erc721-go-bindings
package erc721

var (
	INTERFACEID_ERC165            = [4]byte{1, 255, 201, 167}  // 0x01ffc9a7
	INTERFACEID_ERC721            = [4]byte{128, 172, 88, 205} // 0x80ac58cd
	INTERFACEID_ERC721_METADATA   = [4]byte{91, 94, 19, 159}   // 0x5b5e139f
	INTERFACEID_ERC721_ENUMERABLE = [4]byte{120, 14, 157, 99}  // 0x780e9d63
	INTERFACEID_ERC1155           = [4]byte{217, 182, 122, 38} // 0xd9b67a26
)
