// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AssociationSetProvider

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AssociationSetProviderMetaData contains all meta data concerning the AssociationSetProvider contract.
var AssociationSetProviderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyPredicate\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"subset\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"characteristicFunction\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"predicateType\",\"type\":\"uint8\",\"internalType\":\"enumAssociationSetProvider.PredicateType\"}],\"outputs\":[{\"name\":\"elements\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"setCardinality\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCategoryBitmap\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestForScope\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecordHashAndCategoryAt\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecordHashesAndCategories\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"recordHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"categoryBitmaps\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRegistryAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRegistryAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCategoryForRecord\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tryGetCategoryBitmap\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalRecords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InsufficientRole\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LeafAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeafCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeafGreaterThanSnarkScalarField\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecordNotFound\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderIsNotAdmin\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b50336203f480818061003d57604051636116401160e11b8152600060048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff85160217905561006760008261007f565b505050610079816100f060201b60201c565b506102c7565b6000826100dd57600061009a6002546001600160a01b031690565b6001600160a01b0316146100c157604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6100e7838361016e565b90505b92915050565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff166101415760405163125aa2cf60e11b8152336004820152602401610034565b61016b7fbb28eb1a0cfabcecf96003fab466159bc2e051e49d79baf049890044e907244082610218565b50565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16610210576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101c83390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016100ea565b5060006100ea565b8161023657604051631fe1e13d60e11b815260040160405180910390fd5b6102408282610244565b5050565b60008281526020819052604090206001015461025f8161026f565b610269838361007f565b50505050565b61016b81336000828152602081815260408083206001600160a01b038516845290915290205460ff166102405760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610034565b611a84806102d66000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c806391d14854116100f9578063c2b1db0b11610097578063cefc142911610071578063cefc142914610433578063cf6eefb71461043b578063d547741f14610469578063d602b9fd1461047c57600080fd5b8063c2b1db0b146103f0578063c9d9d0ab14610403578063cc8463c81461042b57600080fd5b80639c99e1f1116100d35780639c99e1f11461038b578063a1eda53c146103ac578063a217fddf146103d3578063bf584c4b146103db57600080fd5b806391d14854146103445780639408234b146103575780639bbb195e1461037857600080fd5b8063367d9b8b116101665780637790f90f116101405780637790f90f146102ba57806384ef8ffc146102ed5780638da5cb5b14610312578063904057ec1461031a57600080fd5b8063367d9b8b14610281578063634e93da14610294578063649a5ec7146102a757600080fd5b80630aa6220b116101a25780630aa6220b1461022e578063248a9ca3146102385780632f2ff15d1461025b57806336568abe1461026e57600080fd5b806301ffc9a7146101c9578063022d63fb146101f15780630513a7a61461020d575b600080fd5b6101dc6101d73660046115d6565b610484565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016101e8565b61022061021b366004611600565b6104af565b6040519081526020016101e8565b610236610504565b005b610220610246366004611622565b60009081526020819052604090206001015490565b610236610269366004611657565b61051a565b61023661027c366004611657565b610546565b61023661028f366004611683565b6105ed565b6102366102a23660046116af565b6106e6565b6102366102b53660046116ca565b6106fa565b6102cd6102c8366004611622565b61070e565b6040805194855260208501939093529183015260608201526080016101e8565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016101e8565b6102fa610780565b61032d610328366004611600565b610799565b6040805192151583526020830191909152016101e8565b6101dc610352366004611657565b6107bf565b61036a610365366004611683565b6107e8565b6040516101e892919061172e565b6102366103863660046116af565b61095c565b61039e61039936600461175c565b61099e565b6040516101e89291906117fe565b6103b4610a86565b6040805165ffffffffffff9384168152929091166020830152016101e8565b610220600081565b610220600080516020611a2f83398151915281565b6102366103fe3660046116af565b610ada565b610416610411366004611600565b610b1c565b604080519283526020830191909152016101e8565b6101f6610b36565b610236610b95565b610443610bd5565b604080516001600160a01b03909316835265ffffffffffff9091166020830152016101e8565b610236610477366004611657565b610bf6565b610236610c1e565b60006001600160e01b031982166318a4c3c360e11b14806104a957506104a982610c31565b92915050565b6000828152600360205260408120819081906104cb9085610c66565b91509150816104fc57604051638406f96760e01b815260048101869052602481018590526044015b60405180910390fd5b949350505050565b600061050f81610ca8565b610517610cb2565b50565b8161053857604051631fe1e13d60e11b815260040160405180910390fd5b6105428282610cbf565b5050565b8115801561056157506002546001600160a01b038281169116145b156105e357600080610571610bd5565b90925090506001600160a01b038216151580610593575065ffffffffffff8116155b806105a657504265ffffffffffff821610155b156105ce576040516319ca5ebb60e01b815265ffffffffffff821660048201526024016104f3565b50506001805465ffffffffffff60a01b191690555b6105428282610cea565b600080516020611a2f83398151915261060681336107bf565b61062c576040516309d8965560e31b8152336004820152602481018290526044016104f3565b6000848152600460209081526040808320600181015484526002018252808320548784526003909252909120610663908585610d22565b156106835760008581526004602052604090206106809085610d49565b90505b7f40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e85826106c1600360008a8152602001908152602001600020610f24565b6040805193845260208401929092529082015260600160405180910390a15050505050565b60006106f181610ca8565b61054282610f2f565b600061070581610ca8565b61054282610fa2565b600081815260046020908152604080832060018101548452600201825280832054848452600390925282209091908190819061074990610f24565b905080156107795761077361075f600183611836565b600087815260036020526040902090611012565b90935091505b9193509193565b60006107946002546001600160a01b031690565b905090565b600082815260036020526040812081906107b39084610c66565b915091505b9250929050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b60608082841080156108105750600085815260036020526040902061080c90610f24565b8311155b61084c5760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642072616e676560981b60448201526064016104f3565b60006108588585611836565b90508067ffffffffffffffff81111561087357610873611849565b60405190808252806020026020018201604052801561089c578160200160208202803683370190505b5092508067ffffffffffffffff8111156108b8576108b8611849565b6040519080825280602002602001820160405280156108e1578160200160208202803683370190505b50915060005b81811015610952576109106108fc828861185f565b600089815260036020526040902090611012565b85838151811061092257610922611872565b6020026020010185848151811061093b5761093b611872565b6020908102919091010191909152526001016108e7565b5050935093915050565b6109676000336107bf565b6109865760405163125aa2cf60e11b81523360048201526024016104f3565b610517600080516020611a2f8339815191528261051a565b60606000808567ffffffffffffffff8111156109bc576109bc611849565b6040519080825280602002602001820160405280156109e5578160200160208202803683370190505b50905060005b86811015610a76576000888883818110610a0757610a07611872565b905060200201359050600080610a1d8c84610799565b9150915081610a2e57505050610a6e565b610a39888a8361103d565b15610a6a5782858781518110610a5157610a51611872565b602090810291909101015285610a6681611888565b9650505b5050505b6001016109eb565b5081815291509550959350505050565b600254600090600160d01b900465ffffffffffff168015158015610ab257504265ffffffffffff821610155b610abe57600080610ad2565b600254600160a01b900465ffffffffffff16815b915091509091565b610ae56000336107bf565b610b045760405163125aa2cf60e11b81523360048201526024016104f3565b610517600080516020611a2f83398151915282610bf6565b600082815260036020526040812081906107b39084611012565b600254600090600160d01b900465ffffffffffff168015158015610b6157504265ffffffffffff8216105b610b7c57600154600160d01b900465ffffffffffff16610b8f565b600254600160a01b900465ffffffffffff165b91505090565b6000610b9f610bd5565b509050336001600160a01b03821614610bcd57604051636116401160e11b81523360048201526024016104f3565b6105176110a7565b6001546001600160a01b03811691600160a01b90910465ffffffffffff1690565b81610c1457604051631fe1e13d60e11b815260040160405180910390fd5b6105428282611140565b6000610c2981610ca8565b610517611165565b60006001600160e01b03198216637965db0b60e01b14806104a957506301ffc9a760e01b6001600160e01b03198316146104a9565b6000818152600283016020526040812054819080610c9557610c888585611170565b9250600091506107b89050565b6001925090506107b8565b509250929050565b610517813361117c565b610cbd6000806111b5565b565b600082815260208190526040902060010154610cda81610ca8565b610ce48383611275565b50505050565b6001600160a01b0381163314610d135760405163334bd91960e11b815260040160405180910390fd5b610d1d82826112dd565b505050565b60008281526002840160205260408120829055610d3f848461131a565b90505b9392505050565b60007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018210610d8b576040516361c0541760e11b815260040160405180910390fd5b81600003610dac576040516314b48df160e11b815260040160405180910390fd5b600082815260038401602052604090205415610ddb576040516312c50cad60e11b815260040160405180910390fd5b825460018085015490610def90839061185f565b610dfa82600261197d565b1015610e0c57610e0981611888565b90505b600185018190558360005b82811015610ee9578084901c600116600103610ecd57604080518082018252600083815260028a0160209081529083902054825281018490529051632b0aac7f60e11b815273__$a2daaad8940c9006af3f1557205ebe532d$__9163561558fe91610e859190600401611989565b602060405180830381865af4158015610ea2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec691906119ba565b9150610ee1565b600081815260028801602052604090208290555b600101610e17565b50610ef383611888565b8087556000928352600287016020908152604080852084905596845260039097019096529390209390935550919050565b60006104a982611326565b6000610f39610b36565b610f4242611330565b610f4c91906119d3565b9050610f588282611367565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b6000610fad826113e6565b610fb642611330565b610fc091906119d3565b9050610fcc82826111b5565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60008080611020858561142e565b600081815260029690960160205260409095205494959350505050565b600080846002811115611052576110526119f9565b0361106257508082168214610d42565b6001846002811115611076576110766119f9565b0361108657508082161515610d42565b600284600281111561109a5761109a6119f9565b03610d4257501615919050565b6000806110b2610bd5565b915091506110c78165ffffffffffff16151590565b15806110db57504265ffffffffffff821610155b15611103576040516319ca5ebb60e01b815265ffffffffffff821660048201526024016104f3565b61111f600061111a6002546001600160a01b031690565b6112dd565b5061112b600083611275565b5050600180546001600160d01b031916905550565b60008281526020819052604090206001015461115b81610ca8565b610ce483836112dd565b610cbd600080611367565b6000610d42838361143a565b61118682826107bf565b6105425760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016104f3565b600254600160d01b900465ffffffffffff168015611238574265ffffffffffff8216101561120e57600254600180546001600160d01b0316600160a01b90920465ffffffffffff16600160d01b02919091179055611238565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600280546001600160a01b0316600160a01b65ffffffffffff948516026001600160d01b031617600160d01b9290931691909102919091179055565b6000826112d35760006112906002546001600160a01b031690565b6001600160a01b0316146112b757604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610d428383611452565b6000821580156112fa57506002546001600160a01b038381169116145b1561131057600280546001600160a01b03191690555b610d4283836114e4565b6000610d42838361154f565b60006104a9825490565b600065ffffffffffff821115611363576040516306dfcc6560e41b815260306004820152602481018390526044016104f3565b5090565b6000611371610bd5565b6001805465ffffffffffff8616600160a01b026001600160d01b03199091166001600160a01b0388161717905591506113b390508165ffffffffffff16151590565b15610d1d576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a1505050565b6000806113f1610b36565b90508065ffffffffffff168365ffffffffffff1611611419576114148382611a0f565b610d42565b610d4265ffffffffffff841662069780611596565b6000610d4283836115ac565b60008181526001830160205260408120541515610d42565b600061145e83836107bf565b6114dc576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556114943390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016104a9565b5060006104a9565b60006114f083836107bf565b156114dc576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016104a9565b60008181526001830160205260408120546114dc575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104a9565b60008183106115a55781610d42565b5090919050565b60008260000182815481106115c3576115c3611872565b9060005260206000200154905092915050565b6000602082840312156115e857600080fd5b81356001600160e01b031981168114610d4257600080fd5b6000806040838503121561161357600080fd5b50508035926020909101359150565b60006020828403121561163457600080fd5b5035919050565b80356001600160a01b038116811461165257600080fd5b919050565b6000806040838503121561166a57600080fd5b8235915061167a6020840161163b565b90509250929050565b60008060006060848603121561169857600080fd5b505081359360208301359350604090920135919050565b6000602082840312156116c157600080fd5b610d428261163b565b6000602082840312156116dc57600080fd5b813565ffffffffffff81168114610d4257600080fd5b60008151808452602080850194506020840160005b8381101561172357815187529582019590820190600101611707565b509495945050505050565b60408152600061174160408301856116f2565b828103602084015261175381856116f2565b95945050505050565b60008060008060006080868803121561177457600080fd5b85359450602086013567ffffffffffffffff8082111561179357600080fd5b818801915088601f8301126117a757600080fd5b8135818111156117b657600080fd5b8960208260051b85010111156117cb57600080fd5b602083019650809550505050604086013591506060860135600381106117f057600080fd5b809150509295509295909350565b60408152600061181160408301856116f2565b90508260208301529392505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156104a9576104a9611820565b634e487b7160e01b600052604160045260246000fd5b808201808211156104a9576104a9611820565b634e487b7160e01b600052603260045260246000fd5b60006001820161189a5761189a611820565b5060010190565b600181815b80851115610ca05781600019048211156118c2576118c2611820565b808516156118cf57918102915b93841c93908002906118a6565b6000826118eb575060016104a9565b816118f8575060006104a9565b816001811461190e576002811461191857611934565b60019150506104a9565b60ff84111561192957611929611820565b50506001821b6104a9565b5060208310610133831016604e8410600b8410161715611957575081810a6104a9565b61196183836118a1565b806000190482111561197557611975611820565b029392505050565b6000610d4283836118dc565b60408101818360005b60028110156119b1578151835260209283019290910190600101611992565b50505092915050565b6000602082840312156119cc57600080fd5b5051919050565b65ffffffffffff8181168382160190808211156119f2576119f2611820565b5092915050565b634e487b7160e01b600052602160045260246000fd5b65ffffffffffff8281168282160390808211156119f2576119f261182056febb28eb1a0cfabcecf96003fab466159bc2e051e49d79baf049890044e9072440a26469706673582212207a87fea657dc3f7039b489e76ad51b77db05e5a5f837cc063b6cb6638b80a2b464736f6c63430008190033",
}

// AssociationSetProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AssociationSetProviderMetaData.ABI instead.
var AssociationSetProviderABI = AssociationSetProviderMetaData.ABI

// AssociationSetProviderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssociationSetProviderMetaData.Bin instead.
var AssociationSetProviderBin = AssociationSetProviderMetaData.Bin

// DeployAssociationSetProvider deploys a new Ethereum contract, binding an instance of AssociationSetProvider to it.
func DeployAssociationSetProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssociationSetProvider, error) {
	parsed, err := AssociationSetProviderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssociationSetProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssociationSetProvider{AssociationSetProviderCaller: AssociationSetProviderCaller{contract: contract}, AssociationSetProviderTransactor: AssociationSetProviderTransactor{contract: contract}, AssociationSetProviderFilterer: AssociationSetProviderFilterer{contract: contract}}, nil
}

// AssociationSetProvider is an auto generated Go binding around an Ethereum contract.
type AssociationSetProvider struct {
	AssociationSetProviderCaller     // Read-only binding to the contract
	AssociationSetProviderTransactor // Write-only binding to the contract
	AssociationSetProviderFilterer   // Log filterer for contract events
}

// AssociationSetProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssociationSetProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssociationSetProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssociationSetProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssociationSetProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssociationSetProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssociationSetProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssociationSetProviderSession struct {
	Contract     *AssociationSetProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AssociationSetProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssociationSetProviderCallerSession struct {
	Contract *AssociationSetProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AssociationSetProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssociationSetProviderTransactorSession struct {
	Contract     *AssociationSetProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AssociationSetProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssociationSetProviderRaw struct {
	Contract *AssociationSetProvider // Generic contract binding to access the raw methods on
}

// AssociationSetProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssociationSetProviderCallerRaw struct {
	Contract *AssociationSetProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AssociationSetProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssociationSetProviderTransactorRaw struct {
	Contract *AssociationSetProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssociationSetProvider creates a new instance of AssociationSetProvider, bound to a specific deployed contract.
func NewAssociationSetProvider(address common.Address, backend bind.ContractBackend) (*AssociationSetProvider, error) {
	contract, err := bindAssociationSetProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProvider{AssociationSetProviderCaller: AssociationSetProviderCaller{contract: contract}, AssociationSetProviderTransactor: AssociationSetProviderTransactor{contract: contract}, AssociationSetProviderFilterer: AssociationSetProviderFilterer{contract: contract}}, nil
}

// NewAssociationSetProviderCaller creates a new read-only instance of AssociationSetProvider, bound to a specific deployed contract.
func NewAssociationSetProviderCaller(address common.Address, caller bind.ContractCaller) (*AssociationSetProviderCaller, error) {
	contract, err := bindAssociationSetProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderCaller{contract: contract}, nil
}

// NewAssociationSetProviderTransactor creates a new write-only instance of AssociationSetProvider, bound to a specific deployed contract.
func NewAssociationSetProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AssociationSetProviderTransactor, error) {
	contract, err := bindAssociationSetProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderTransactor{contract: contract}, nil
}

// NewAssociationSetProviderFilterer creates a new log filterer instance of AssociationSetProvider, bound to a specific deployed contract.
func NewAssociationSetProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AssociationSetProviderFilterer, error) {
	contract, err := bindAssociationSetProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderFilterer{contract: contract}, nil
}

// bindAssociationSetProvider binds a generic wrapper to an already deployed contract.
func bindAssociationSetProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssociationSetProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssociationSetProvider *AssociationSetProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssociationSetProvider.Contract.AssociationSetProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssociationSetProvider *AssociationSetProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.AssociationSetProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssociationSetProvider *AssociationSetProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.AssociationSetProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssociationSetProvider *AssociationSetProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssociationSetProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssociationSetProvider *AssociationSetProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssociationSetProvider *AssociationSetProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssociationSetProvider.Contract.DEFAULTADMINROLE(&_AssociationSetProvider.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssociationSetProvider.Contract.DEFAULTADMINROLE(&_AssociationSetProvider.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCaller) REGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderSession) REGISTRYADMINROLE() ([32]byte, error) {
	return _AssociationSetProvider.Contract.REGISTRYADMINROLE(&_AssociationSetProvider.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) REGISTRYADMINROLE() ([32]byte, error) {
	return _AssociationSetProvider.Contract.REGISTRYADMINROLE(&_AssociationSetProvider.CallOpts)
}

// ApplyPredicate is a free data retrieval call binding the contract method 0x9c99e1f1.
//
// Solidity: function applyPredicate(uint256 domain, bytes32[] subset, bytes32 characteristicFunction, uint8 predicateType) view returns(bytes32[] elements, uint256 setCardinality)
func (_AssociationSetProvider *AssociationSetProviderCaller) ApplyPredicate(opts *bind.CallOpts, domain *big.Int, subset [][32]byte, characteristicFunction [32]byte, predicateType uint8) (struct {
	Elements       [][32]byte
	SetCardinality *big.Int
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "applyPredicate", domain, subset, characteristicFunction, predicateType)

	outstruct := new(struct {
		Elements       [][32]byte
		SetCardinality *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Elements = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.SetCardinality = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApplyPredicate is a free data retrieval call binding the contract method 0x9c99e1f1.
//
// Solidity: function applyPredicate(uint256 domain, bytes32[] subset, bytes32 characteristicFunction, uint8 predicateType) view returns(bytes32[] elements, uint256 setCardinality)
func (_AssociationSetProvider *AssociationSetProviderSession) ApplyPredicate(domain *big.Int, subset [][32]byte, characteristicFunction [32]byte, predicateType uint8) (struct {
	Elements       [][32]byte
	SetCardinality *big.Int
}, error) {
	return _AssociationSetProvider.Contract.ApplyPredicate(&_AssociationSetProvider.CallOpts, domain, subset, characteristicFunction, predicateType)
}

// ApplyPredicate is a free data retrieval call binding the contract method 0x9c99e1f1.
//
// Solidity: function applyPredicate(uint256 domain, bytes32[] subset, bytes32 characteristicFunction, uint8 predicateType) view returns(bytes32[] elements, uint256 setCardinality)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) ApplyPredicate(domain *big.Int, subset [][32]byte, characteristicFunction [32]byte, predicateType uint8) (struct {
	Elements       [][32]byte
	SetCardinality *big.Int
}, error) {
	return _AssociationSetProvider.Contract.ApplyPredicate(&_AssociationSetProvider.CallOpts, domain, subset, characteristicFunction, predicateType)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderSession) DefaultAdmin() (common.Address, error) {
	return _AssociationSetProvider.Contract.DefaultAdmin(&_AssociationSetProvider.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) DefaultAdmin() (common.Address, error) {
	return _AssociationSetProvider.Contract.DefaultAdmin(&_AssociationSetProvider.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderSession) DefaultAdminDelay() (*big.Int, error) {
	return _AssociationSetProvider.Contract.DefaultAdminDelay(&_AssociationSetProvider.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _AssociationSetProvider.Contract.DefaultAdminDelay(&_AssociationSetProvider.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _AssociationSetProvider.Contract.DefaultAdminDelayIncreaseWait(&_AssociationSetProvider.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _AssociationSetProvider.Contract.DefaultAdminDelayIncreaseWait(&_AssociationSetProvider.CallOpts)
}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCaller) GetCategoryBitmap(opts *bind.CallOpts, scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "getCategoryBitmap", scope, recordHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderSession) GetCategoryBitmap(scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	return _AssociationSetProvider.Contract.GetCategoryBitmap(&_AssociationSetProvider.CallOpts, scope, recordHash)
}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) GetCategoryBitmap(scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	return _AssociationSetProvider.Contract.GetCategoryBitmap(&_AssociationSetProvider.CallOpts, scope, recordHash)
}

// GetLatestForScope is a free data retrieval call binding the contract method 0x7790f90f.
//
// Solidity: function getLatestForScope(uint256 scope) view returns(uint256 root, bytes32 recordHash, bytes32 categoryBitmap, uint256 index)
func (_AssociationSetProvider *AssociationSetProviderCaller) GetLatestForScope(opts *bind.CallOpts, scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "getLatestForScope", scope)

	outstruct := new(struct {
		Root           *big.Int
		RecordHash     [32]byte
		CategoryBitmap [32]byte
		Index          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RecordHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.CategoryBitmap = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Index = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLatestForScope is a free data retrieval call binding the contract method 0x7790f90f.
//
// Solidity: function getLatestForScope(uint256 scope) view returns(uint256 root, bytes32 recordHash, bytes32 categoryBitmap, uint256 index)
func (_AssociationSetProvider *AssociationSetProviderSession) GetLatestForScope(scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	return _AssociationSetProvider.Contract.GetLatestForScope(&_AssociationSetProvider.CallOpts, scope)
}

// GetLatestForScope is a free data retrieval call binding the contract method 0x7790f90f.
//
// Solidity: function getLatestForScope(uint256 scope) view returns(uint256 root, bytes32 recordHash, bytes32 categoryBitmap, uint256 index)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) GetLatestForScope(scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	return _AssociationSetProvider.Contract.GetLatestForScope(&_AssociationSetProvider.CallOpts, scope)
}

// GetRecordHashAndCategoryAt is a free data retrieval call binding the contract method 0xc9d9d0ab.
//
// Solidity: function getRecordHashAndCategoryAt(uint256 scope, uint256 index) view returns(bytes32 recordHash, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCaller) GetRecordHashAndCategoryAt(opts *bind.CallOpts, scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "getRecordHashAndCategoryAt", scope, index)

	outstruct := new(struct {
		RecordHash     [32]byte
		CategoryBitmap [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecordHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.CategoryBitmap = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetRecordHashAndCategoryAt is a free data retrieval call binding the contract method 0xc9d9d0ab.
//
// Solidity: function getRecordHashAndCategoryAt(uint256 scope, uint256 index) view returns(bytes32 recordHash, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderSession) GetRecordHashAndCategoryAt(scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	return _AssociationSetProvider.Contract.GetRecordHashAndCategoryAt(&_AssociationSetProvider.CallOpts, scope, index)
}

// GetRecordHashAndCategoryAt is a free data retrieval call binding the contract method 0xc9d9d0ab.
//
// Solidity: function getRecordHashAndCategoryAt(uint256 scope, uint256 index) view returns(bytes32 recordHash, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) GetRecordHashAndCategoryAt(scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	return _AssociationSetProvider.Contract.GetRecordHashAndCategoryAt(&_AssociationSetProvider.CallOpts, scope, index)
}

// GetRecordHashesAndCategories is a free data retrieval call binding the contract method 0x9408234b.
//
// Solidity: function getRecordHashesAndCategories(uint256 scope, uint256 from, uint256 to) view returns(bytes32[] recordHashes, bytes32[] categoryBitmaps)
func (_AssociationSetProvider *AssociationSetProviderCaller) GetRecordHashesAndCategories(opts *bind.CallOpts, scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "getRecordHashesAndCategories", scope, from, to)

	outstruct := new(struct {
		RecordHashes    [][32]byte
		CategoryBitmaps [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecordHashes = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.CategoryBitmaps = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetRecordHashesAndCategories is a free data retrieval call binding the contract method 0x9408234b.
//
// Solidity: function getRecordHashesAndCategories(uint256 scope, uint256 from, uint256 to) view returns(bytes32[] recordHashes, bytes32[] categoryBitmaps)
func (_AssociationSetProvider *AssociationSetProviderSession) GetRecordHashesAndCategories(scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	return _AssociationSetProvider.Contract.GetRecordHashesAndCategories(&_AssociationSetProvider.CallOpts, scope, from, to)
}

// GetRecordHashesAndCategories is a free data retrieval call binding the contract method 0x9408234b.
//
// Solidity: function getRecordHashesAndCategories(uint256 scope, uint256 from, uint256 to) view returns(bytes32[] recordHashes, bytes32[] categoryBitmaps)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) GetRecordHashesAndCategories(scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	return _AssociationSetProvider.Contract.GetRecordHashesAndCategories(&_AssociationSetProvider.CallOpts, scope, from, to)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssociationSetProvider.Contract.GetRoleAdmin(&_AssociationSetProvider.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssociationSetProvider.Contract.GetRoleAdmin(&_AssociationSetProvider.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssociationSetProvider.Contract.HasRole(&_AssociationSetProvider.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssociationSetProvider.Contract.HasRole(&_AssociationSetProvider.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderSession) Owner() (common.Address, error) {
	return _AssociationSetProvider.Contract.Owner(&_AssociationSetProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) Owner() (common.Address, error) {
	return _AssociationSetProvider.Contract.Owner(&_AssociationSetProvider.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(struct {
		NewAdmin common.Address
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _AssociationSetProvider.Contract.PendingDefaultAdmin(&_AssociationSetProvider.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _AssociationSetProvider.Contract.PendingDefaultAdmin(&_AssociationSetProvider.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(struct {
		NewDelay *big.Int
		Schedule *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _AssociationSetProvider.Contract.PendingDefaultAdminDelay(&_AssociationSetProvider.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _AssociationSetProvider.Contract.PendingDefaultAdminDelay(&_AssociationSetProvider.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssociationSetProvider.Contract.SupportsInterface(&_AssociationSetProvider.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssociationSetProvider.Contract.SupportsInterface(&_AssociationSetProvider.CallOpts, interfaceId)
}

// TryGetCategoryBitmap is a free data retrieval call binding the contract method 0x904057ec.
//
// Solidity: function tryGetCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bool exists, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCaller) TryGetCategoryBitmap(opts *bind.CallOpts, scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	var out []interface{}
	err := _AssociationSetProvider.contract.Call(opts, &out, "tryGetCategoryBitmap", scope, recordHash)

	outstruct := new(struct {
		Exists         bool
		CategoryBitmap [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CategoryBitmap = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TryGetCategoryBitmap is a free data retrieval call binding the contract method 0x904057ec.
//
// Solidity: function tryGetCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bool exists, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderSession) TryGetCategoryBitmap(scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	return _AssociationSetProvider.Contract.TryGetCategoryBitmap(&_AssociationSetProvider.CallOpts, scope, recordHash)
}

// TryGetCategoryBitmap is a free data retrieval call binding the contract method 0x904057ec.
//
// Solidity: function tryGetCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bool exists, bytes32 categoryBitmap)
func (_AssociationSetProvider *AssociationSetProviderCallerSession) TryGetCategoryBitmap(scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	return _AssociationSetProvider.Contract.TryGetCategoryBitmap(&_AssociationSetProvider.CallOpts, scope, recordHash)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.AcceptDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.AcceptDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.BeginDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.BeginDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.CancelDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.CancelDefaultAdminTransfer(&_AssociationSetProvider.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.ChangeDefaultAdminDelay(&_AssociationSetProvider.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.ChangeDefaultAdminDelay(&_AssociationSetProvider.TransactOpts, newDelay)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) GrantRegistryAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "grantRegistryAdminRole", account)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) GrantRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.GrantRegistryAdminRole(&_AssociationSetProvider.TransactOpts, account)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) GrantRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.GrantRegistryAdminRole(&_AssociationSetProvider.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.GrantRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.GrantRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RenounceRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RenounceRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) RevokeRegistryAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "revokeRegistryAdminRole", account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) RevokeRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RevokeRegistryAdminRole(&_AssociationSetProvider.TransactOpts, account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) RevokeRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RevokeRegistryAdminRole(&_AssociationSetProvider.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RevokeRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RevokeRole(&_AssociationSetProvider.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_AssociationSetProvider *AssociationSetProviderSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RollbackDefaultAdminDelay(&_AssociationSetProvider.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.RollbackDefaultAdminDelay(&_AssociationSetProvider.TransactOpts)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactor) SetCategoryForRecord(opts *bind.TransactOpts, scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _AssociationSetProvider.contract.Transact(opts, "setCategoryForRecord", scope, recordHash, categoryBitmap)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_AssociationSetProvider *AssociationSetProviderSession) SetCategoryForRecord(scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.SetCategoryForRecord(&_AssociationSetProvider.TransactOpts, scope, recordHash, categoryBitmap)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_AssociationSetProvider *AssociationSetProviderTransactorSession) SetCategoryForRecord(scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _AssociationSetProvider.Contract.SetCategoryForRecord(&_AssociationSetProvider.TransactOpts, scope, recordHash, categoryBitmap)
}

// AssociationSetProviderDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminDelayChangeCanceledIterator struct {
	Event *AssociationSetProviderDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*AssociationSetProviderDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderDefaultAdminDelayChangeCanceledIterator{contract: _AssociationSetProvider.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderDefaultAdminDelayChangeCanceled)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminDelayChangeCanceled is a log parse operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*AssociationSetProviderDefaultAdminDelayChangeCanceled, error) {
	event := new(AssociationSetProviderDefaultAdminDelayChangeCanceled)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminDelayChangeScheduledIterator struct {
	Event *AssociationSetProviderDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*AssociationSetProviderDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderDefaultAdminDelayChangeScheduledIterator{contract: _AssociationSetProvider.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderDefaultAdminDelayChangeScheduled)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminDelayChangeScheduled is a log parse operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*AssociationSetProviderDefaultAdminDelayChangeScheduled, error) {
	event := new(AssociationSetProviderDefaultAdminDelayChangeScheduled)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminTransferCanceledIterator struct {
	Event *AssociationSetProviderDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*AssociationSetProviderDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderDefaultAdminTransferCanceledIterator{contract: _AssociationSetProvider.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderDefaultAdminTransferCanceled)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminTransferCanceled is a log parse operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*AssociationSetProviderDefaultAdminTransferCanceled, error) {
	event := new(AssociationSetProviderDefaultAdminTransferCanceled)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminTransferScheduledIterator struct {
	Event *AssociationSetProviderDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the AssociationSetProvider contract.
type AssociationSetProviderDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*AssociationSetProviderDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderDefaultAdminTransferScheduledIterator{contract: _AssociationSetProvider.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderDefaultAdminTransferScheduled)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultAdminTransferScheduled is a log parse operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*AssociationSetProviderDefaultAdminTransferScheduled, error) {
	event := new(AssociationSetProviderDefaultAdminTransferScheduled)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleAdminChangedIterator struct {
	Event *AssociationSetProviderRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderRoleAdminChanged represents a RoleAdminChanged event raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssociationSetProviderRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderRoleAdminChangedIterator{contract: _AssociationSetProvider.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderRoleAdminChanged)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseRoleAdminChanged(log types.Log) (*AssociationSetProviderRoleAdminChanged, error) {
	event := new(AssociationSetProviderRoleAdminChanged)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleGrantedIterator struct {
	Event *AssociationSetProviderRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderRoleGranted represents a RoleGranted event raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssociationSetProviderRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderRoleGrantedIterator{contract: _AssociationSetProvider.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderRoleGranted)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseRoleGranted(log types.Log) (*AssociationSetProviderRoleGranted, error) {
	event := new(AssociationSetProviderRoleGranted)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleRevokedIterator struct {
	Event *AssociationSetProviderRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderRoleRevoked represents a RoleRevoked event raised by the AssociationSetProvider contract.
type AssociationSetProviderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssociationSetProviderRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderRoleRevokedIterator{contract: _AssociationSetProvider.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderRoleRevoked)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseRoleRevoked(log types.Log) (*AssociationSetProviderRoleRevoked, error) {
	event := new(AssociationSetProviderRoleRevoked)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssociationSetProviderUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the AssociationSetProvider contract.
type AssociationSetProviderUpdateIterator struct {
	Event *AssociationSetProviderUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AssociationSetProviderUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssociationSetProviderUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AssociationSetProviderUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AssociationSetProviderUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssociationSetProviderUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssociationSetProviderUpdate represents a Update event raised by the AssociationSetProvider contract.
type AssociationSetProviderUpdate struct {
	Scope        *big.Int
	Root         *big.Int
	TotalRecords *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e.
//
// Solidity: event Update(uint256 scope, uint256 root, uint256 totalRecords)
func (_AssociationSetProvider *AssociationSetProviderFilterer) FilterUpdate(opts *bind.FilterOpts) (*AssociationSetProviderUpdateIterator, error) {

	logs, sub, err := _AssociationSetProvider.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &AssociationSetProviderUpdateIterator{contract: _AssociationSetProvider.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e.
//
// Solidity: event Update(uint256 scope, uint256 root, uint256 totalRecords)
func (_AssociationSetProvider *AssociationSetProviderFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *AssociationSetProviderUpdate) (event.Subscription, error) {

	logs, sub, err := _AssociationSetProvider.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssociationSetProviderUpdate)
				if err := _AssociationSetProvider.contract.UnpackLog(event, "Update", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdate is a log parse operation binding the contract event 0x40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e.
//
// Solidity: event Update(uint256 scope, uint256 root, uint256 totalRecords)
func (_AssociationSetProvider *AssociationSetProviderFilterer) ParseUpdate(log types.Log) (*AssociationSetProviderUpdate, error) {
	event := new(AssociationSetProviderUpdate)
	if err := _AssociationSetProvider.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
