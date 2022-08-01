// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/Test.sol";
import "forge-std/console2.sol"; //console2.log() forge test -vvv で確認できる
import "../src/NFT.sol";

contract NFTTest is Test {
	using stdStorage for StdStorage;

	NFT private nft;

	/**
  setUpというメソッドを定義することでテスト前に実行してくれる
  beforeEach的な挙動.各テストメソッドの実行前に実行.
   */
	function setUp() public {
		//デプロイ
		nft = new NFT("NFT_tutorial", "TUT", "baseUri");
	}

	/**
  testFaildのprefixでエラー検証
   */
	function testFailNoMintPricePaid() public {
		nft.mintTo(address(1));
	}

	/**
  testのprefixで正常系
   */
	function testMintPricePaid() public {
		nft.mintTo{ value: 0.08 ether }(address(1));
	}

	/**
  テストコントラクトのストレージにアクセス.
   */
	function testFailMaxSupplyReached() public {
		//たぶん、コントラクトの関数を見つけてきてモックに差し替えている
		uint256 slot = stdstore
			.target(address(nft))
			.sig("currentTokenId()")
			.find();
		bytes32 loc = bytes32(slot);
		bytes32 mockedCurrentTokenId = bytes32(abi.encode(10000));
		vm.store(address(nft), loc, mockedCurrentTokenId);

		nft.mintTo{ value: 0.08 ether }(address(1));
	}

	function testFailMintToZeroAddress() public {
		nft.mintTo{ value: 0.08 ether }(address(0));
	}

	function testNewMintOwnerRegistered() public {
		nft.mintTo{ value: 0.08 ether }(address(1));
		uint256 slotOfNewOwner = stdstore
			.target(address(nft))
			.sig(nft.ownerOf.selector)
			.with_key(1)
			.find();

		uint160 ownerOfTokenIdOne = uint160(
			uint256(
				(vm.load(address(nft), bytes32(abi.encode(slotOfNewOwner))))
			)
		);

		assertEq(address(ownerOfTokenIdOne), address(1));
	}

	function testBalanceIncremented() public {
		nft.mintTo{ value: 0.08 ether }(address(1));
		uint256 slotBalance = stdstore
			.target(address(nft))
			.sig(nft.balanceOf.selector)
			.with_key(address(1))
			.find();

		uint256 balanceFirstMint = uint256(
			vm.load(address(nft), bytes32(slotBalance))
		);
		assertEq(balanceFirstMint, 1);

		nft.mintTo{ value: 0.08 ether }(address(1));
		uint256 balanceSecondMint = uint256(
			vm.load(address(nft), bytes32(slotBalance))
		);
		assertEq(balanceSecondMint, 2);
	}

	function testSafeContractReceive() public {
		Reciver reciver = new Reciver();
		nft.mintTo{ value: 0.08 ether }(address(reciver));
		uint256 slotBalance = stdstore
			.target(address(nft))
			.sig(nft.balanceOf.selector)
			.with_key(address(reciver))
			.find();

		uint256 balance = uint256(vm.load(address(nft), bytes32(slotBalance)));
		assertEq(balance, 1);
	}

	function testFailUnSafeContractReciver() public {
		vm.etch(address(1), bytes("mock code"));
		nft.mintTo{ value: 0.08 ether }(address(1));
	}

	function testWithdrawWorksAsOwner() public {
		Reciver reciver = new Reciver();
		address payable payee = payable(address(0x1337));
		uint256 priorPayeeBalance = payee.balance;
		uint256 mintPrice = nft.MINT_PRICE();

		nft.mintTo{ value: mintPrice }(address(reciver));

		uint256 afterMintNFTContractBalance = address(nft).balance;
		assertEq(afterMintNFTContractBalance, mintPrice);

		nft.withdrawPayments(payee);
		assertEq(
			payee.balance,
			priorPayeeBalance + afterMintNFTContractBalance
		);
	}

	function testWithdrawFailsAsNotOwner() public {
		Reciver reciver = new Reciver();
		uint256 mintPrice = nft.MINT_PRICE();
		nft.mintTo{ value: mintPrice }(address(reciver));
		assertEq(address(nft).balance, mintPrice);

		vm.expectRevert("Ownable: caller is not the owner");
		vm.startPrank(address(0xd3ad));
		nft.withdrawPayments(payable(address(0x3ad)));
		vm.stopPrank();
	}
}

contract Reciver is ERC721TokenReceiver {
	function onERC721Received(
		address operator,
		address from,
		uint256 id,
		bytes calldata data
	) external override returns (bytes4) {
		return this.onERC721Received.selector;
	}
}
