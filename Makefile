deploy_ganache:
	@echo "deploy block chain......"
	ganache-cli -m "you are really the very best person in the world"

abi_gen_go:
	abigen --bin ./abi/PumpkinToken.bin --abi ./abi/PumpkinToken.abi --pkg contract --type PumpkinToken --out contract/pumpkin_token.go

