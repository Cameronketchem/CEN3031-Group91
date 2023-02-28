// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.0;

import "./ERC20.sol";
import "./ERC721Sale.sol";
import "./IERC721Receiver.sol";

contract ERC20Executor is ERC20, IERC721Receiver {
  mapping(address => uint256) private _contributions;
  mapping(uint => address) private _holders;
  uint private _totalHolders;
  bool private _soldImmediate;
  uint256 private _pooledAmt;

  // The contract of the NFT sale
  ERC721Sale _sale;
  address _sale_address;

  constructor(string memory name_, string memory symbol_, address nft_)
    ERC20(name_, symbol_){

    _totalHolders = 0;
    _sale = ERC721Sale(nft_);
    _sale_address = nft_;
    _soldImmediate = false;
    _pooledAmt = 0;
  }

  /**
   * @dev Return the price of the NFT being sold
   */
  function price() public view returns (uint256) {
    return _sale.price();
  }

  /**
   * @dev Return whether the NFT sale is still active
   */
  function saleActive() public view returns (bool) {
    return _sale.saleActive();
  }

  /**
   * @dev Buy an NFT immediately, without waiting for other contributors.
   * receipient is the address that receives the NFT. If receipient is a
   * contract, then it must implement IERC20Receiver interface.
   */
  function buyImmediate(address receipient) public payable {
    require(saleActive(), "Sale is not active");
    require(msg.value >= price(), "Msg value must be greater than or equal to price");

    if (!_sale.buy{ value: msg.value }(receipient)) {
      payable(msg.sender).transfer(msg.value);
      return;
    }

    _soldImmediate = true;
  }

  function buyImmediate() public payable {
    return buyImmediate(msg.sender);
  }

  /**
   * @dev Make a contribution to the crowd sale, the contributed amount
   * will be on behalf of the receipient address. If the sum of all
   * contributions matches the price of the NFT, then it will be bought
   * and a share token will be distributed to all contributors.
   */
  function contribute(address receipient) public payable {
    require(saleActive(), "Sale is not active");
    require(msg.value > 0, "Contribution must be greater than 0");

    // Check if contributor is registered
    bool found = false;
    for (uint i = 0; i < _totalHolders; i++) {
      if (_holders[i] == receipient) {
        found = true;
        break;
      }
    }

    // Add address to holders and contributor maps
    if (!found) {
      _totalHolders++;
      _holders[_totalHolders] = receipient;
    }

    // Apply contribution
    _contributions[receipient] += msg.value;
    _pooledAmt += msg.value;

    // Check if sale is complete
    if (_pooledAmt < price())
      return;

    // Purchase NFT
    if (!_sale.buy{ value: _pooledAmt }(address(this)))
      return;

    // Distribute share token if sale succesful
    for (uint i = 0; i < _totalHolders; i++)
      super._mint(_holders[i], _contributions[_holders[i]]);
  }

  function contribute() public payable {
    return contribute(msg.sender);
  }

  /**
   * @dev Returns the contribution amount of address 'addr'
   */
  function contributionAmt(address addr) public view returns (uint256) {
    return _contributions[addr];
  }

  /**
   * @dev Refund a contribution, provided that a crowd sale was not
   * finalized. The sender will receive 'amt', provided that amt does
   * not exceed the sum of his contributions.
   */
  function refund(uint256 amt) public {
    require(saleActive() || (!saleActive() && _soldImmediate), "Cannot refund finalized sale");
    require(amt > 0, "Refund amount must be more than zero");
    require(amt <= _contributions[msg.sender], "Refund cannot exceed contributed amount");

    payable(msg.sender).transfer(amt);
    _contributions[msg.sender] -= amt;
    _pooledAmt -= amt;
  }

  /**
   * @dev Implements IERC20Receiver
   */
  function onERC721Received
    (address operator, address from, uint256 tokenId, bytes calldata data)
    public view returns (bytes4) {
    require(operator == _sale_address, "Cannot receive ERC721 from unregistered address");
    return IERC721Receiver.onERC721Received.selector;
  }

  /**
   * @dev Returns the contract's eth balance.
   */
  function getBalance() external view returns (uint256) {
    return address(this).balance;
  }
}
