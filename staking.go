package main

import (
    "context"
    "math/big"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    "log"
)

type StakingContract struct {
    Client       *rpc.Client
    Address      common.Address
    PrivateKey   *ecdsa.PrivateKey
    Nonce        uint64
    Auth         *bind.TransactOpts
}

func NewStakingContract(rpcURL string, privateKey string) (*StakingContract, error) {
    client, err := rpc.Dial(rpcURL)
    if err != nil {
        return nil, err
    }

    privateKey, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return nil, err
    }

    address := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), address)
    if err != nil {
        return nil, err
    }

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(21000) // uygun bir gas sınırı belirleyin

    return &StakingContract{
        Client:     client,
        Address:    address,
        PrivateKey: privateKey,
        Nonce:      nonce,
        Auth:       auth,
    }, nil
}

func (sc *StakingContract) Stake(amount *big.Int) (*common.Hash, error) {

    inputData, err := yourSmartContract.ABI.Pack("stake", amount)
    if err != nil {
        return nil, err
    }

    tx := types.NewTransaction(sc.Nonce, yourSmartContractAddress, big.NewInt(0), sc.Auth.GasLimit, sc.Auth.GasPrice, inputData)
    sc.Nonce++
    
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), sc.PrivateKey)
    if err != nil {
        return nil, err
    }

    err = sc.Client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return nil, err
    }

    return &signedTx.Hash(), nil
}

func (sc *StakingContract) Unstake(amount *big.Int) (*common.Hash, error) {

    inputData, err := yourSmartContract.ABI.Pack("unstake", amount)
    if err != nil {
        return nil, err
    }

    tx := types.NewTransaction(sc.Nonce, yourSmartContractAddress, big.NewInt(0), sc.Auth.GasLimit, sc.Auth.GasPrice, inputData)
    sc.Nonce++
    
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), sc.PrivateKey)
    if err != nil {
        return nil, err
    }

    err = sc.Client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return nil, err
    }

    return &signedTx.Hash(), nil
}

func (sc *StakingContract) ClaimRewards() (*common.Hash, error) {
  
    inputData, err := yourSmartContract.ABI.Pack("claimRewards")
    if err != nil {
        return nil, err
    }

    tx := types.NewTransaction(sc.Nonce, yourSmartContractAddress, big.NewInt(0), sc.Auth.GasLimit, sc.Auth.GasPrice, inputData)
    sc.Nonce++
    
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), sc.PrivateKey)
    if err != nil {
        return nil, err
    }

    err = sc.Client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return nil, err
    }

    return &signedTx.Hash(), nil
}

