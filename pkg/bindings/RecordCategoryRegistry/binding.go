// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RecordCategoryRegistry

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

// RecordCategoryRegistryMetaData contains all meta data concerning the RecordCategoryRegistry contract.
var RecordCategoryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"registryAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCategoryBitmap\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestForScope\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecordHashAndCategoryAt\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecordHashesAndCategories\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"recordHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"categoryBitmaps\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRegistryAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRegistryAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCategoryForRecord\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tryGetCategoryBitmap\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"categoryBitmap\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalRecords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InsufficientRole\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LeafAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeafCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeafGreaterThanSnarkScalarField\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecordNotFound\",\"inputs\":[{\"name\":\"scope\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recordHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderIsNotAdmin\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611b3c380380611b3c83398101604081905261002f916102e4565b6203f480338061005a57604051636116401160e11b8152600060048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff85160217905561008460008261009c565b5050506100968161010d60201b60201c565b50610314565b6000826100fa5760006100b76002546001600160a01b031690565b6001600160a01b0316146100de57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610104838361018b565b90505b92915050565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1661015e5760405163125aa2cf60e11b8152336004820152602401610051565b6101887fbb28eb1a0cfabcecf96003fab466159bc2e051e49d79baf049890044e907244082610235565b50565b6000828152602081815260408083206001600160a01b038516845290915281205460ff1661022d576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101e53390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610107565b506000610107565b8161025357604051631fe1e13d60e11b815260040160405180910390fd5b61025d8282610261565b5050565b60008281526020819052604090206001015461027c8161028c565b610286838361009c565b50505050565b61018881336000828152602081815260408083206001600160a01b038516845290915290205460ff1661025d5760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610051565b6000602082840312156102f657600080fd5b81516001600160a01b038116811461030d57600080fd5b9392505050565b611819806103236000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c8063904057ec116100f9578063c2b1db0b11610097578063cefc142911610071578063cefc1429146103f7578063cf6eefb7146103ff578063d547741f1461042d578063d602b9fd1461044057600080fd5b8063c2b1db0b146103b4578063c9d9d0ab146103c7578063cc8463c8146103ef57600080fd5b80639bbb195e116100d35780639bbb195e1461035d578063a1eda53c14610370578063a217fddf14610397578063bf584c4b1461039f57600080fd5b8063904057ec146102ff57806391d14854146103295780639408234b1461033c57600080fd5b806336568abe11610166578063649a5ec711610140578063649a5ec71461028c5780637790f90f1461029f57806384ef8ffc146102d25780638da5cb5b146102f757600080fd5b806336568abe14610253578063367d9b8b14610266578063634e93da1461027957600080fd5b806301ffc9a7146101ae578063022d63fb146101d65780630513a7a6146101f25780630aa6220b14610213578063248a9ca31461021d5780632f2ff15d14610240575b600080fd5b6101c16101bc366004611445565b610448565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016101cd565b61020561020036600461146f565b610473565b6040519081526020016101cd565b61021b6104c8565b005b61020561022b366004611491565b60009081526020819052604090206001015490565b61021b61024e3660046114c6565b6104de565b61021b6102613660046114c6565b61050a565b61021b6102743660046114f2565b6105b1565b61021b61028736600461151e565b6106aa565b61021b61029a366004611539565b6106be565b6102b26102ad366004611491565b6106d2565b6040805194855260208501939093529183015260608201526080016101cd565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016101cd565b6102df610744565b61031261030d36600461146f565b61075d565b6040805192151583526020830191909152016101cd565b6101c16103373660046114c6565b610783565b61034f61034a3660046114f2565b6107ac565b6040516101cd92919061159d565b61021b61036b36600461151e565b610920565b610378610962565b6040805165ffffffffffff9384168152929091166020830152016101cd565b610205600081565b6102056000805160206117c483398151915281565b61021b6103c236600461151e565b6109b6565b6103da6103d536600461146f565b6109f8565b604080519283526020830191909152016101cd565b6101db610a12565b61021b610a71565b610407610ab1565b604080516001600160a01b03909316835265ffffffffffff9091166020830152016101cd565b61021b61043b3660046114c6565b610ad2565b61021b610afa565b60006001600160e01b031982166318a4c3c360e11b148061046d575061046d82610b0d565b92915050565b60008281526003602052604081208190819061048f9085610b42565b91509150816104c057604051638406f96760e01b815260048101869052602481018590526044015b60405180910390fd5b949350505050565b60006104d381610b84565b6104db610b8e565b50565b816104fc57604051631fe1e13d60e11b815260040160405180910390fd5b6105068282610b9b565b5050565b8115801561052557506002546001600160a01b038281169116145b156105a757600080610535610ab1565b90925090506001600160a01b038216151580610557575065ffffffffffff8116155b8061056a57504265ffffffffffff821610155b15610592576040516319ca5ebb60e01b815265ffffffffffff821660048201526024016104b7565b50506001805465ffffffffffff60a01b191690555b6105068282610bc6565b6000805160206117c48339815191526105ca8133610783565b6105f0576040516309d8965560e31b8152336004820152602481018290526044016104b7565b6000848152600460209081526040808320600181015484526002018252808320548784526003909252909120610627908585610bfe565b156106475760008581526004602052604090206106449085610c1b565b90505b7f40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e8582610685600360008a8152602001908152602001600020610df6565b6040805193845260208401929092529082015260600160405180910390a15050505050565b60006106b581610b84565b61050682610e01565b60006106c981610b84565b61050682610e74565b600081815260046020908152604080832060018101548452600201825280832054848452600390925282209091908190819061070d90610df6565b9050801561073d576107376107236001836115e1565b600087815260036020526040902090610ee4565b90935091505b9193509193565b60006107586002546001600160a01b031690565b905090565b600082815260036020526040812081906107779084610b42565b915091505b9250929050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b60608082841080156107d4575060008581526003602052604090206107d090610df6565b8311155b6108105760405162461bcd60e51b815260206004820152600d60248201526c496e76616c69642072616e676560981b60448201526064016104b7565b600061081c85856115e1565b90508067ffffffffffffffff811115610837576108376115f4565b604051908082528060200260200182016040528015610860578160200160208202803683370190505b5092508067ffffffffffffffff81111561087c5761087c6115f4565b6040519080825280602002602001820160405280156108a5578160200160208202803683370190505b50915060005b81811015610916576108d46108c0828861160a565b600089815260036020526040902090610ee4565b8583815181106108e6576108e661161d565b602002602001018584815181106108ff576108ff61161d565b6020908102919091010191909152526001016108ab565b5050935093915050565b61092b600033610783565b61094a5760405163125aa2cf60e11b81523360048201526024016104b7565b6104db6000805160206117c4833981519152826104de565b600254600090600160d01b900465ffffffffffff16801515801561098e57504265ffffffffffff821610155b61099a576000806109ae565b600254600160a01b900465ffffffffffff16815b915091509091565b6109c1600033610783565b6109e05760405163125aa2cf60e11b81523360048201526024016104b7565b6104db6000805160206117c483398151915282610ad2565b600082815260036020526040812081906107779084610ee4565b600254600090600160d01b900465ffffffffffff168015158015610a3d57504265ffffffffffff8216105b610a5857600154600160d01b900465ffffffffffff16610a6b565b600254600160a01b900465ffffffffffff165b91505090565b6000610a7b610ab1565b509050336001600160a01b03821614610aa957604051636116401160e11b81523360048201526024016104b7565b6104db610f0f565b6001546001600160a01b03811691600160a01b90910465ffffffffffff1690565b81610af057604051631fe1e13d60e11b815260040160405180910390fd5b6105068282610fa8565b6000610b0581610b84565b6104db610fcd565b60006001600160e01b03198216637965db0b60e01b148061046d57506301ffc9a760e01b6001600160e01b031983161461046d565b6000818152600283016020526040812054819080610b7157610b648585610fd8565b92506000915061077c9050565b60019250905061077c565b509250929050565b6104db8133610feb565b610b99600080611024565b565b600082815260208190526040902060010154610bb681610b84565b610bc083836110e4565b50505050565b6001600160a01b0381163314610bef5760405163334bd91960e11b815260040160405180910390fd5b610bf9828261114c565b505050565b600082815260028401602052604081208290556104c08484611189565b60007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018210610c5d576040516361c0541760e11b815260040160405180910390fd5b81600003610c7e576040516314b48df160e11b815260040160405180910390fd5b600082815260038401602052604090205415610cad576040516312c50cad60e11b815260040160405180910390fd5b825460018085015490610cc190839061160a565b610ccc82600261170f565b1015610cde57610cdb8161171b565b90505b600185018190558360005b82811015610dbb578084901c600116600103610d9f57604080518082018252600083815260028a0160209081529083902054825281018490529051632b0aac7f60e11b815273__$a2daaad8940c9006af3f1557205ebe532d$__9163561558fe91610d579190600401611734565b602060405180830381865af4158015610d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d989190611765565b9150610db3565b600081815260028801602052604090208290555b600101610ce9565b50610dc58361171b565b8087556000928352600287016020908152604080852084905596845260039097019096529390209390935550919050565b600061046d82611195565b6000610e0b610a12565b610e144261119f565b610e1e919061177e565b9050610e2a82826111d6565b60405165ffffffffffff821681526001600160a01b038316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b6000610e7f82611255565b610e884261119f565b610e92919061177e565b9050610e9e8282611024565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b60008080610ef2858561129d565b600081815260029690960160205260409095205494959350505050565b600080610f1a610ab1565b91509150610f2f8165ffffffffffff16151590565b1580610f4357504265ffffffffffff821610155b15610f6b576040516319ca5ebb60e01b815265ffffffffffff821660048201526024016104b7565b610f876000610f826002546001600160a01b031690565b61114c565b50610f936000836110e4565b5050600180546001600160d01b031916905550565b600082815260208190526040902060010154610fc381610b84565b610bc0838361114c565b610b996000806111d6565b6000610fe483836112a9565b9392505050565b610ff58282610783565b6105065760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016104b7565b600254600160d01b900465ffffffffffff1680156110a7574265ffffffffffff8216101561107d57600254600180546001600160d01b0316600160a01b90920465ffffffffffff16600160d01b029190911790556110a7565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec590600090a15b50600280546001600160a01b0316600160a01b65ffffffffffff948516026001600160d01b031617600160d01b9290931691909102919091179055565b6000826111425760006110ff6002546001600160a01b031690565b6001600160a01b03161461112657604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b610fe483836112c1565b60008215801561116957506002546001600160a01b038381169116145b1561117f57600280546001600160a01b03191690555b610fe48383611353565b6000610fe483836113be565b600061046d825490565b600065ffffffffffff8211156111d2576040516306dfcc6560e41b815260306004820152602481018390526044016104b7565b5090565b60006111e0610ab1565b6001805465ffffffffffff8616600160a01b026001600160d01b03199091166001600160a01b03881617179055915061122290508165ffffffffffff16151590565b15610bf9576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a960510990600090a1505050565b600080611260610a12565b90508065ffffffffffff168365ffffffffffff16116112885761128383826117a4565b610fe4565b610fe465ffffffffffff841662069780611405565b6000610fe4838361141b565b60008181526001830160205260408120541515610fe4565b60006112cd8383610783565b61134b576000838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556113033390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161046d565b50600061046d565b600061135f8383610783565b1561134b576000838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161046d565b600081815260018301602052604081205461134b5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561046d565b60008183106114145781610fe4565b5090919050565b60008260000182815481106114325761143261161d565b9060005260206000200154905092915050565b60006020828403121561145757600080fd5b81356001600160e01b031981168114610fe457600080fd5b6000806040838503121561148257600080fd5b50508035926020909101359150565b6000602082840312156114a357600080fd5b5035919050565b80356001600160a01b03811681146114c157600080fd5b919050565b600080604083850312156114d957600080fd5b823591506114e9602084016114aa565b90509250929050565b60008060006060848603121561150757600080fd5b505081359360208301359350604090920135919050565b60006020828403121561153057600080fd5b610fe4826114aa565b60006020828403121561154b57600080fd5b813565ffffffffffff81168114610fe457600080fd5b60008151808452602080850194506020840160005b8381101561159257815187529582019590820190600101611576565b509495945050505050565b6040815260006115b06040830185611561565b82810360208401526115c28185611561565b95945050505050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561046d5761046d6115cb565b634e487b7160e01b600052604160045260246000fd5b8082018082111561046d5761046d6115cb565b634e487b7160e01b600052603260045260246000fd5b600181815b80851115610b7c578160001904821115611654576116546115cb565b8085161561166157918102915b93841c9390800290611638565b60008261167d5750600161046d565b8161168a5750600061046d565b81600181146116a057600281146116aa576116c6565b600191505061046d565b60ff8411156116bb576116bb6115cb565b50506001821b61046d565b5060208310610133831016604e8410600b84101617156116e9575081810a61046d565b6116f38383611633565b8060001904821115611707576117076115cb565b029392505050565b6000610fe4838361166e565b60006001820161172d5761172d6115cb565b5060010190565b60408101818360005b600281101561175c57815183526020928301929091019060010161173d565b50505092915050565b60006020828403121561177757600080fd5b5051919050565b65ffffffffffff81811683821601908082111561179d5761179d6115cb565b5092915050565b65ffffffffffff82811682821603908082111561179d5761179d6115cb56febb28eb1a0cfabcecf96003fab466159bc2e051e49d79baf049890044e9072440a26469706673582212208a5dc91c9c94f1268bff4ee2bdbd612ce745b34e4d0c031de19f8a4b26e88ab464736f6c63430008190033",
}

// RecordCategoryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RecordCategoryRegistryMetaData.ABI instead.
var RecordCategoryRegistryABI = RecordCategoryRegistryMetaData.ABI

// RecordCategoryRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecordCategoryRegistryMetaData.Bin instead.
var RecordCategoryRegistryBin = RecordCategoryRegistryMetaData.Bin

// DeployRecordCategoryRegistry deploys a new Ethereum contract, binding an instance of RecordCategoryRegistry to it.
func DeployRecordCategoryRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, registryAdmin common.Address) (common.Address, *types.Transaction, *RecordCategoryRegistry, error) {
	parsed, err := RecordCategoryRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecordCategoryRegistryBin), backend, registryAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecordCategoryRegistry{RecordCategoryRegistryCaller: RecordCategoryRegistryCaller{contract: contract}, RecordCategoryRegistryTransactor: RecordCategoryRegistryTransactor{contract: contract}, RecordCategoryRegistryFilterer: RecordCategoryRegistryFilterer{contract: contract}}, nil
}

// RecordCategoryRegistry is an auto generated Go binding around an Ethereum contract.
type RecordCategoryRegistry struct {
	RecordCategoryRegistryCaller     // Read-only binding to the contract
	RecordCategoryRegistryTransactor // Write-only binding to the contract
	RecordCategoryRegistryFilterer   // Log filterer for contract events
}

// RecordCategoryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecordCategoryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordCategoryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecordCategoryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordCategoryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecordCategoryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordCategoryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecordCategoryRegistrySession struct {
	Contract     *RecordCategoryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RecordCategoryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecordCategoryRegistryCallerSession struct {
	Contract *RecordCategoryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RecordCategoryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecordCategoryRegistryTransactorSession struct {
	Contract     *RecordCategoryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RecordCategoryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecordCategoryRegistryRaw struct {
	Contract *RecordCategoryRegistry // Generic contract binding to access the raw methods on
}

// RecordCategoryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecordCategoryRegistryCallerRaw struct {
	Contract *RecordCategoryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RecordCategoryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecordCategoryRegistryTransactorRaw struct {
	Contract *RecordCategoryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecordCategoryRegistry creates a new instance of RecordCategoryRegistry, bound to a specific deployed contract.
func NewRecordCategoryRegistry(address common.Address, backend bind.ContractBackend) (*RecordCategoryRegistry, error) {
	contract, err := bindRecordCategoryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistry{RecordCategoryRegistryCaller: RecordCategoryRegistryCaller{contract: contract}, RecordCategoryRegistryTransactor: RecordCategoryRegistryTransactor{contract: contract}, RecordCategoryRegistryFilterer: RecordCategoryRegistryFilterer{contract: contract}}, nil
}

// NewRecordCategoryRegistryCaller creates a new read-only instance of RecordCategoryRegistry, bound to a specific deployed contract.
func NewRecordCategoryRegistryCaller(address common.Address, caller bind.ContractCaller) (*RecordCategoryRegistryCaller, error) {
	contract, err := bindRecordCategoryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryCaller{contract: contract}, nil
}

// NewRecordCategoryRegistryTransactor creates a new write-only instance of RecordCategoryRegistry, bound to a specific deployed contract.
func NewRecordCategoryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RecordCategoryRegistryTransactor, error) {
	contract, err := bindRecordCategoryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryTransactor{contract: contract}, nil
}

// NewRecordCategoryRegistryFilterer creates a new log filterer instance of RecordCategoryRegistry, bound to a specific deployed contract.
func NewRecordCategoryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RecordCategoryRegistryFilterer, error) {
	contract, err := bindRecordCategoryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryFilterer{contract: contract}, nil
}

// bindRecordCategoryRegistry binds a generic wrapper to an already deployed contract.
func bindRecordCategoryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecordCategoryRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecordCategoryRegistry *RecordCategoryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecordCategoryRegistry.Contract.RecordCategoryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecordCategoryRegistry *RecordCategoryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RecordCategoryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecordCategoryRegistry *RecordCategoryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RecordCategoryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecordCategoryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.DEFAULTADMINROLE(&_RecordCategoryRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.DEFAULTADMINROLE(&_RecordCategoryRegistry.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) REGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) REGISTRYADMINROLE() ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.REGISTRYADMINROLE(&_RecordCategoryRegistry.CallOpts)
}

// REGISTRYADMINROLE is a free data retrieval call binding the contract method 0xbf584c4b.
//
// Solidity: function REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) REGISTRYADMINROLE() ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.REGISTRYADMINROLE(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) DefaultAdmin() (common.Address, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdmin(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdmin is a free data retrieval call binding the contract method 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) DefaultAdmin() (common.Address, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdmin(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) DefaultAdminDelay() (*big.Int, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdminDelay(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdminDelay is a free data retrieval call binding the contract method 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdminDelay(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdminDelayIncreaseWait(&_RecordCategoryRegistry.CallOpts)
}

// DefaultAdminDelayIncreaseWait is a free data retrieval call binding the contract method 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _RecordCategoryRegistry.Contract.DefaultAdminDelayIncreaseWait(&_RecordCategoryRegistry.CallOpts)
}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) GetCategoryBitmap(opts *bind.CallOpts, scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "getCategoryBitmap", scope, recordHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GetCategoryBitmap(scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.GetCategoryBitmap(&_RecordCategoryRegistry.CallOpts, scope, recordHash)
}

// GetCategoryBitmap is a free data retrieval call binding the contract method 0x0513a7a6.
//
// Solidity: function getCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) GetCategoryBitmap(scope *big.Int, recordHash [32]byte) ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.GetCategoryBitmap(&_RecordCategoryRegistry.CallOpts, scope, recordHash)
}

// GetLatestForScope is a free data retrieval call binding the contract method 0x7790f90f.
//
// Solidity: function getLatestForScope(uint256 scope) view returns(uint256 root, bytes32 recordHash, bytes32 categoryBitmap, uint256 index)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) GetLatestForScope(opts *bind.CallOpts, scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "getLatestForScope", scope)

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GetLatestForScope(scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.GetLatestForScope(&_RecordCategoryRegistry.CallOpts, scope)
}

// GetLatestForScope is a free data retrieval call binding the contract method 0x7790f90f.
//
// Solidity: function getLatestForScope(uint256 scope) view returns(uint256 root, bytes32 recordHash, bytes32 categoryBitmap, uint256 index)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) GetLatestForScope(scope *big.Int) (struct {
	Root           *big.Int
	RecordHash     [32]byte
	CategoryBitmap [32]byte
	Index          *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.GetLatestForScope(&_RecordCategoryRegistry.CallOpts, scope)
}

// GetRecordHashAndCategoryAt is a free data retrieval call binding the contract method 0xc9d9d0ab.
//
// Solidity: function getRecordHashAndCategoryAt(uint256 scope, uint256 index) view returns(bytes32 recordHash, bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) GetRecordHashAndCategoryAt(opts *bind.CallOpts, scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "getRecordHashAndCategoryAt", scope, index)

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GetRecordHashAndCategoryAt(scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.GetRecordHashAndCategoryAt(&_RecordCategoryRegistry.CallOpts, scope, index)
}

// GetRecordHashAndCategoryAt is a free data retrieval call binding the contract method 0xc9d9d0ab.
//
// Solidity: function getRecordHashAndCategoryAt(uint256 scope, uint256 index) view returns(bytes32 recordHash, bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) GetRecordHashAndCategoryAt(scope *big.Int, index *big.Int) (struct {
	RecordHash     [32]byte
	CategoryBitmap [32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.GetRecordHashAndCategoryAt(&_RecordCategoryRegistry.CallOpts, scope, index)
}

// GetRecordHashesAndCategories is a free data retrieval call binding the contract method 0x9408234b.
//
// Solidity: function getRecordHashesAndCategories(uint256 scope, uint256 from, uint256 to) view returns(bytes32[] recordHashes, bytes32[] categoryBitmaps)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) GetRecordHashesAndCategories(opts *bind.CallOpts, scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "getRecordHashesAndCategories", scope, from, to)

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GetRecordHashesAndCategories(scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.GetRecordHashesAndCategories(&_RecordCategoryRegistry.CallOpts, scope, from, to)
}

// GetRecordHashesAndCategories is a free data retrieval call binding the contract method 0x9408234b.
//
// Solidity: function getRecordHashesAndCategories(uint256 scope, uint256 from, uint256 to) view returns(bytes32[] recordHashes, bytes32[] categoryBitmaps)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) GetRecordHashesAndCategories(scope *big.Int, from *big.Int, to *big.Int) (struct {
	RecordHashes    [][32]byte
	CategoryBitmaps [][32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.GetRecordHashesAndCategories(&_RecordCategoryRegistry.CallOpts, scope, from, to)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.GetRoleAdmin(&_RecordCategoryRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RecordCategoryRegistry.Contract.GetRoleAdmin(&_RecordCategoryRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RecordCategoryRegistry.Contract.HasRole(&_RecordCategoryRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RecordCategoryRegistry.Contract.HasRole(&_RecordCategoryRegistry.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) Owner() (common.Address, error) {
	return _RecordCategoryRegistry.Contract.Owner(&_RecordCategoryRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) Owner() (common.Address, error) {
	return _RecordCategoryRegistry.Contract.Owner(&_RecordCategoryRegistry.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) PendingDefaultAdmin(opts *bind.CallOpts) (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "pendingDefaultAdmin")

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.PendingDefaultAdmin(&_RecordCategoryRegistry.CallOpts)
}

// PendingDefaultAdmin is a free data retrieval call binding the contract method 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) PendingDefaultAdmin() (struct {
	NewAdmin common.Address
	Schedule *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.PendingDefaultAdmin(&_RecordCategoryRegistry.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "pendingDefaultAdminDelay")

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.PendingDefaultAdminDelay(&_RecordCategoryRegistry.CallOpts)
}

// PendingDefaultAdminDelay is a free data retrieval call binding the contract method 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) PendingDefaultAdminDelay() (struct {
	NewDelay *big.Int
	Schedule *big.Int
}, error) {
	return _RecordCategoryRegistry.Contract.PendingDefaultAdminDelay(&_RecordCategoryRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RecordCategoryRegistry.Contract.SupportsInterface(&_RecordCategoryRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RecordCategoryRegistry.Contract.SupportsInterface(&_RecordCategoryRegistry.CallOpts, interfaceId)
}

// TryGetCategoryBitmap is a free data retrieval call binding the contract method 0x904057ec.
//
// Solidity: function tryGetCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bool exists, bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCaller) TryGetCategoryBitmap(opts *bind.CallOpts, scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	var out []interface{}
	err := _RecordCategoryRegistry.contract.Call(opts, &out, "tryGetCategoryBitmap", scope, recordHash)

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
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) TryGetCategoryBitmap(scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.TryGetCategoryBitmap(&_RecordCategoryRegistry.CallOpts, scope, recordHash)
}

// TryGetCategoryBitmap is a free data retrieval call binding the contract method 0x904057ec.
//
// Solidity: function tryGetCategoryBitmap(uint256 scope, bytes32 recordHash) view returns(bool exists, bytes32 categoryBitmap)
func (_RecordCategoryRegistry *RecordCategoryRegistryCallerSession) TryGetCategoryBitmap(scope *big.Int, recordHash [32]byte) (struct {
	Exists         bool
	CategoryBitmap [32]byte
}, error) {
	return _RecordCategoryRegistry.Contract.TryGetCategoryBitmap(&_RecordCategoryRegistry.CallOpts, scope, recordHash)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.AcceptDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts)
}

// AcceptDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.AcceptDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.BeginDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts, newAdmin)
}

// BeginDefaultAdminTransfer is a paid mutator transaction binding the contract method 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.BeginDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts, newAdmin)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.CancelDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts)
}

// CancelDefaultAdminTransfer is a paid mutator transaction binding the contract method 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.CancelDefaultAdminTransfer(&_RecordCategoryRegistry.TransactOpts)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.ChangeDefaultAdminDelay(&_RecordCategoryRegistry.TransactOpts, newDelay)
}

// ChangeDefaultAdminDelay is a paid mutator transaction binding the contract method 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.ChangeDefaultAdminDelay(&_RecordCategoryRegistry.TransactOpts, newDelay)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) GrantRegistryAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "grantRegistryAdminRole", account)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GrantRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.GrantRegistryAdminRole(&_RecordCategoryRegistry.TransactOpts, account)
}

// GrantRegistryAdminRole is a paid mutator transaction binding the contract method 0x9bbb195e.
//
// Solidity: function grantRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) GrantRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.GrantRegistryAdminRole(&_RecordCategoryRegistry.TransactOpts, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.GrantRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.GrantRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RenounceRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RenounceRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) RevokeRegistryAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "revokeRegistryAdminRole", account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) RevokeRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RevokeRegistryAdminRole(&_RecordCategoryRegistry.TransactOpts, account)
}

// RevokeRegistryAdminRole is a paid mutator transaction binding the contract method 0xc2b1db0b.
//
// Solidity: function revokeRegistryAdminRole(address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) RevokeRegistryAdminRole(account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RevokeRegistryAdminRole(&_RecordCategoryRegistry.TransactOpts, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RevokeRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RevokeRole(&_RecordCategoryRegistry.TransactOpts, role, account)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RollbackDefaultAdminDelay(&_RecordCategoryRegistry.TransactOpts)
}

// RollbackDefaultAdminDelay is a paid mutator transaction binding the contract method 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.RollbackDefaultAdminDelay(&_RecordCategoryRegistry.TransactOpts)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactor) SetCategoryForRecord(opts *bind.TransactOpts, scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _RecordCategoryRegistry.contract.Transact(opts, "setCategoryForRecord", scope, recordHash, categoryBitmap)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistrySession) SetCategoryForRecord(scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.SetCategoryForRecord(&_RecordCategoryRegistry.TransactOpts, scope, recordHash, categoryBitmap)
}

// SetCategoryForRecord is a paid mutator transaction binding the contract method 0x367d9b8b.
//
// Solidity: function setCategoryForRecord(uint256 scope, bytes32 recordHash, bytes32 categoryBitmap) returns()
func (_RecordCategoryRegistry *RecordCategoryRegistryTransactorSession) SetCategoryForRecord(scope *big.Int, recordHash [32]byte, categoryBitmap [32]byte) (*types.Transaction, error) {
	return _RecordCategoryRegistry.Contract.SetCategoryForRecord(&_RecordCategoryRegistry.TransactOpts, scope, recordHash, categoryBitmap)
}

// RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator is returned from FilterDefaultAdminDelayChangeCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeCanceled events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator struct {
	Event *RecordCategoryRegistryDefaultAdminDelayChangeCanceled // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryDefaultAdminDelayChangeCanceled)
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
		it.Event = new(RecordCategoryRegistryDefaultAdminDelayChangeCanceled)
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
func (it *RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminDelayChangeCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeCanceled is a free log retrieval operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryDefaultAdminDelayChangeCanceledIterator{contract: _RecordCategoryRegistry.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeCanceled is a free log subscription operation binding the contract event 0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryDefaultAdminDelayChangeCanceled)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*RecordCategoryRegistryDefaultAdminDelayChangeCanceled, error) {
	event := new(RecordCategoryRegistryDefaultAdminDelayChangeCanceled)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator is returned from FilterDefaultAdminDelayChangeScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminDelayChangeScheduled events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator struct {
	Event *RecordCategoryRegistryDefaultAdminDelayChangeScheduled // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryDefaultAdminDelayChangeScheduled)
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
		it.Event = new(RecordCategoryRegistryDefaultAdminDelayChangeScheduled)
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
func (it *RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminDelayChangeScheduled is a free log retrieval operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryDefaultAdminDelayChangeScheduledIterator{contract: _RecordCategoryRegistry.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminDelayChangeScheduled is a free log subscription operation binding the contract event 0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryDefaultAdminDelayChangeScheduled)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*RecordCategoryRegistryDefaultAdminDelayChangeScheduled, error) {
	event := new(RecordCategoryRegistryDefaultAdminDelayChangeScheduled)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryDefaultAdminTransferCanceledIterator is returned from FilterDefaultAdminTransferCanceled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferCanceled events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminTransferCanceledIterator struct {
	Event *RecordCategoryRegistryDefaultAdminTransferCanceled // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryDefaultAdminTransferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryDefaultAdminTransferCanceled)
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
		it.Event = new(RecordCategoryRegistryDefaultAdminTransferCanceled)
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
func (it *RecordCategoryRegistryDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminTransferCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferCanceled is a free log retrieval operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*RecordCategoryRegistryDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryDefaultAdminTransferCanceledIterator{contract: _RecordCategoryRegistry.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferCanceled is a free log subscription operation binding the contract event 0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109.
//
// Solidity: event DefaultAdminTransferCanceled()
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryDefaultAdminTransferCanceled)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*RecordCategoryRegistryDefaultAdminTransferCanceled, error) {
	event := new(RecordCategoryRegistryDefaultAdminTransferCanceled)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryDefaultAdminTransferScheduledIterator is returned from FilterDefaultAdminTransferScheduled and is used to iterate over the raw logs and unpacked data for DefaultAdminTransferScheduled events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminTransferScheduledIterator struct {
	Event *RecordCategoryRegistryDefaultAdminTransferScheduled // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryDefaultAdminTransferScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryDefaultAdminTransferScheduled)
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
		it.Event = new(RecordCategoryRegistryDefaultAdminTransferScheduled)
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
func (it *RecordCategoryRegistryDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultAdminTransferScheduled is a free log retrieval operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*RecordCategoryRegistryDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryDefaultAdminTransferScheduledIterator{contract: _RecordCategoryRegistry.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

// WatchDefaultAdminTransferScheduled is a free log subscription operation binding the contract event 0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryDefaultAdminTransferScheduled)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*RecordCategoryRegistryDefaultAdminTransferScheduled, error) {
	event := new(RecordCategoryRegistryDefaultAdminTransferScheduled)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleAdminChangedIterator struct {
	Event *RecordCategoryRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryRoleAdminChanged)
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
		it.Event = new(RecordCategoryRegistryRoleAdminChanged)
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
func (it *RecordCategoryRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RecordCategoryRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryRoleAdminChangedIterator{contract: _RecordCategoryRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryRoleAdminChanged)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RecordCategoryRegistryRoleAdminChanged, error) {
	event := new(RecordCategoryRegistryRoleAdminChanged)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleGrantedIterator struct {
	Event *RecordCategoryRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryRoleGranted)
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
		it.Event = new(RecordCategoryRegistryRoleGranted)
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
func (it *RecordCategoryRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryRoleGranted represents a RoleGranted event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RecordCategoryRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryRoleGrantedIterator{contract: _RecordCategoryRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryRoleGranted)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseRoleGranted(log types.Log) (*RecordCategoryRegistryRoleGranted, error) {
	event := new(RecordCategoryRegistryRoleGranted)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleRevokedIterator struct {
	Event *RecordCategoryRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryRoleRevoked)
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
		it.Event = new(RecordCategoryRegistryRoleRevoked)
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
func (it *RecordCategoryRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryRoleRevoked represents a RoleRevoked event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RecordCategoryRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryRoleRevokedIterator{contract: _RecordCategoryRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryRoleRevoked)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseRoleRevoked(log types.Log) (*RecordCategoryRegistryRoleRevoked, error) {
	event := new(RecordCategoryRegistryRoleRevoked)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordCategoryRegistryUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryUpdateIterator struct {
	Event *RecordCategoryRegistryUpdate // Event containing the contract specifics and raw log

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
func (it *RecordCategoryRegistryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordCategoryRegistryUpdate)
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
		it.Event = new(RecordCategoryRegistryUpdate)
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
func (it *RecordCategoryRegistryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordCategoryRegistryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordCategoryRegistryUpdate represents a Update event raised by the RecordCategoryRegistry contract.
type RecordCategoryRegistryUpdate struct {
	Scope        *big.Int
	Root         *big.Int
	TotalRecords *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0x40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e.
//
// Solidity: event Update(uint256 scope, uint256 root, uint256 totalRecords)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) FilterUpdate(opts *bind.FilterOpts) (*RecordCategoryRegistryUpdateIterator, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.FilterLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return &RecordCategoryRegistryUpdateIterator{contract: _RecordCategoryRegistry.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0x40c2d32ec54ffac45999120b79eb2516538010ed69ffd59c3f8a184d83b1309e.
//
// Solidity: event Update(uint256 scope, uint256 root, uint256 totalRecords)
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *RecordCategoryRegistryUpdate) (event.Subscription, error) {

	logs, sub, err := _RecordCategoryRegistry.contract.WatchLogs(opts, "Update")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordCategoryRegistryUpdate)
				if err := _RecordCategoryRegistry.contract.UnpackLog(event, "Update", log); err != nil {
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
func (_RecordCategoryRegistry *RecordCategoryRegistryFilterer) ParseUpdate(log types.Log) (*RecordCategoryRegistryUpdate, error) {
	event := new(RecordCategoryRegistryUpdate)
	if err := _RecordCategoryRegistry.contract.UnpackLog(event, "Update", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
