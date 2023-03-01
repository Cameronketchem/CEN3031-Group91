// SPDX-License-Identifier: AGPL-3.0-or-later

pragma solidity ^0.8.0;

import "./ERC721.sol";

contract ERC721Sale is ERC721 {
  address private _admin;
  bool private _sale_active;
  bool private _paused;

  // Address that receives sale proceeds
  address private _sale_receipient;

  // Price in wei
  uint256 private _price;

  constructor (string memory name_, string memory symbol_, uint256 price_, address sale_receipient_)
    ERC721(name_, symbol_) {
    _price = price_;
    _admin = msg.sender;
    _sale_active = true;
    _sale_receipient = sale_receipient_;
  }

  /**
   * @dev Returns the sale price of the NFT
   */
  function price() public view returns (uint256) {
    return _price;
  }

  /**
   * @dev Returns whether the sale is active
   */
  function saleActive() public view returns (bool) {
    return (_sale_active && !_paused);
  }

  /**
   * @dev Change the price of the NFT while sale is still active.
   */
  function changePrice(uint256 price_) public {
    require(saleActive(), "Sale is not active");
    require((msg.sender == _admin) || (msg.sender == _sale_receipient),
            "Only admin can change price");
    _price = price_;
  }

  /**
   * @dev Pause the sale of the NFT, provided that it hasn't been sold.
   */
  function pauseSale() public {
    require(saleActive(), "Cannot pause an inactive sale");
    require((msg.sender == _admin) || (msg.sender == _sale_receipient),
            "Only admin can pause sale");

    _paused = true;
  }

  /**
   * @dev Return whether the sale is paused.
   */
  function paused() public view returns (bool) {
    return _paused;
  }

  /**
   * @dev Unpause the sale of the NFT, provided that it hasn't been sold.
   */
  function resumeSale() public {
    if (!paused()) { return; }

    require(saleActive(), "Cannot unpause an inactive sale");
    require((msg.sender == _admin) || (msg.sender == _sale_receipient),
            "Only admin can unpause sale");

    _paused = false;
  }

  /**
   * @dev Purchase the NFT. The NFT is automatically minted and given to the
   * buyer address. If buyer is a smart contract, then it must implement
   * {IERC721Receiver-onERC721Received}
   */
  function buy(address buyer) public payable returns (bool) {
    require(saleActive(), "Sale is not active");
    require(msg.value >= price(), "Sent value must equal or exceed price");

    // Transfer funds to sale receipient
    payable(_sale_receipient).transfer(msg.value);

    // Mint NFT for buyer address
    super._safeMint(buyer, 0);

    _sale_active = false;
    return true;
  }
}
