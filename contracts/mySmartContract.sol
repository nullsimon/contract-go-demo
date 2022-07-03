// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract MySmartContract {
    function Hello() public view returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public view returns (string memory) {
        return str;
    }
}