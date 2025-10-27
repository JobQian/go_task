// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//只有实现了此接口的合约地址才能安全的转移nft，其他的不与转移
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/interfaces/draft-IERC6093.sol";

library ERC721Utils {
    function checkOnERC721Received(
        address operator,
        address from,
        address to,
        uint256 tokenId,
        bytes memory data
    ) internal {
        if (to.code.length > 0) {
            try IERC721Receiver(to).onERC721Received(operator, from, tokenId, data) returns (bytes4 retval) {
                if (retval != IERC721Receiver.onERC721Received.selector) {
                    // Token rejected
                    revert IERC721Errors.ERC721InvalidReceiver(to);
                }
            } catch (bytes memory reason) {
                if (reason.length == 0) {
                    // non-IERC721Receiver implementer
                    revert IERC721Errors.ERC721InvalidReceiver(to);
                } else {
                    assembly ("memory-safe") {
                        revert(add(reason, 0x20), mload(reason))
                    }
                }
            }
        }
    }
}

contract MYERC721 is IERC721Errors{

    address admin;
    string private _name;
    string private _symbol;

    mapping(uint256 tokenId => address owner) private _owners;

    mapping(address owner => uint256 amount) private _balances;

    constructor(string memory name_, string memory symbol_,address owner) {
        _name = name_;
        _symbol = symbol_;
        admin = owner;
    }

    // 只允许管理员调用的修饰符
    modifier onlyAdmin() {
        require(msg.sender == admin, "Caller is not the admin");
        _;
    }

    function name() public view returns (string memory) {
        return _name;
    }

    function symbol() public view returns (string memory) {
        return _symbol;
    }

    function decimals() public pure returns (uint8) {
        return 18;
    }

    //读取了环境状态变量sender，固能必须使用view
    function _msgSender() internal view returns (address) {
        return msg.sender;
    }

    //读取的仅是函数传递的数据不涉及环境状态变量，固能使用pure
    function _msgData() internal pure returns (bytes calldata) {
        return msg.data;
    }

}