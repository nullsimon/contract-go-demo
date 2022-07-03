solc --optimize --abi ./contracts/MySmartContract.sol -o build
solc --optimize --bin ./contracts/MySmartContract.sol -o build
abigen --bin=./build/MySmartContract.bin --abi=./build/MySmartContract.abi --pkg=api --out=./api/MySmartContract.go