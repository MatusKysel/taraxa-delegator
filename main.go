package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	"github.com/MatusKysel/taraxa-delegator/taraxaDposClient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const privateKeyConst = "FILL THE PRIVATE KEY"

var delegateTo = common.HexToAddress("0xe50b5452B2E8435404DBe06E6a05410C47B7583D")

const rpcEndpoint = "https://rpc.mainnet.taraxa.io"

func main() {
	privateKey, account := getKeys()
	fmt.Println("Account address: ", account)

	backend, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	chainID, _ := backend.ChainID(context.Background())
	fmt.Println("Chain ID", chainID)

	auth := prepareTransactOps(backend, account, privateKey, chainID)

	blockNumber, _ := backend.BlockNumber(context.Background())
	fmt.Println("Block number", blockNumber)

	contract, err := taraxaDposClient.NewTaraxaDposClient(common.HexToAddress("0x00000000000000000000000000000000000000FE"), backend)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Get all delegations and claim them
	delegations, err := contract.GetDelegations(nil, account, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Your current delegations are: ")
	for _, delegation := range delegations.Delegations {
		fmt.Println("Validator account:", delegation.Account, "Stake:", ToDecimal(delegation.Delegation.Stake, 18), "Reward:", ToDecimal(delegation.Delegation.Rewards, 18))
		_, err := contract.ClaimRewards(auth, delegation.Account)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	}

	// Get all validators and claim the reward
	validators, err := contract.GetValidatorsFor(nil, account, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Your current validators are: ")
	for _, validator := range validators.Validators {
		fmt.Println("Validator account:", validator.Account, "Stake:", ToDecimal(validator.Info.TotalStake, 18), "Reward:", ToDecimal(validator.Info.CommissionReward, 18))
		_, err := contract.ClaimCommissionRewards(auth, validator.Account)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		auth.Nonce.Add(auth.Nonce, big.NewInt(1))
	}

	// Wait for all transactions to finish
	waitForAllConfirmed(backend, account, auth.Nonce.Uint64())

	// Get the balance
	balance, _ := backend.BalanceAt(context.Background(), account, nil)
	fmt.Println("Current balance:", ToDecimal(balance, 18))

	// Delegate to specified delegator
	if ToRoundedTara(balance) > 1 {
		auth.Value = ToWei(ToRoundedTara(balance), 18)
		_, err = contract.Delegate(auth, delegateTo)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func prepareTransactOps(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainId *big.Int) *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		fmt.Println(err.Error())
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToDecimal
func ToRoundedTara(ivalue interface{}) int64 {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	ethValue := new(big.Int).Quo(value, big.NewInt(int64(math.Pow10(18))))
	return ethValue.Int64()
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func waitForAllConfirmed(c *ethclient.Client, fromAddress common.Address, currentNonce uint64) {
	for {
		nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			fmt.Println(err.Error())
		}
		if nonce == currentNonce {
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func getKeys() (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyConst)
	if err != nil {
		fmt.Println(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA)
}
