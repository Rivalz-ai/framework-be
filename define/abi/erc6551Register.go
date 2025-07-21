package abi 
const ERC6551_REGISTRY_ABI = `[{
	"anonymous": false,
	"inputs": [
		{
		"indexed": false,
		"internalType": "address",
		"name": "account",
		"type": "address"
		},
		{
		"indexed": true,
		"internalType": "address",
		"name": "implementation",
		"type": "address"
		},
		{
		"indexed": false,
		"internalType": "bytes32",
		"name": "salt",
		"type": "bytes32"
		},
		{
		"indexed": false,
		"internalType": "uint256",
		"name": "chainId",
		"type": "uint256"
		},
		{
		"indexed": true,
		"internalType": "address",
		"name": "tokenContract",
		"type": "address"
		},
		{
		"indexed": true,
		"internalType": "uint256",
		"name": "tokenId",
		"type": "uint256"
		}
	],
	"name": "ERC6551AccountCreated",
	"type": "event"
	}]`