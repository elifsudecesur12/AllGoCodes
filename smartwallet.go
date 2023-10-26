package main

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "math/big"
    "log"
)

type WalletContract struct {
    Client       *rpc.Client
    Address      common.Address
    PrivateKey   *ecdsa.PrivateKey
    Nonce        uint64
    Auth         *bind.TransactOpts
}

func NewWalletContract(rpcURL string, privateKey string) (*WalletContract, error) {
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
    auth.GasLimit = uint64(21000) 

    return &WalletContract{
        Client:     client,
        Address:    address,
        PrivateKey: privateKey,
        Nonce:      nonce,
        Auth:       auth,
    }, nil
}

func (wc *WalletContract) SendETH(to common.Address, amount *big.Int) (*common.Hash, error) {
    tx := types.NewTransaction(wc.Nonce, to, amount, wc.Auth.GasLimit, wc.Auth.GasPrice, nil)
    wc.Nonce++
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), wc.PrivateKey)
    if err != nil {
        return nil, err
    }

    err = wc.Client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return nil, err
    }

    return &signedTx.Hash(), nil
}

func main() {
    rpcURL := "https://mainnet.infura.io/v3/your-infura-project-id"
    privateKey := "0xYourPrivateKey"
    wallet, err := NewWalletContract(rpcURL, privateKey)
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("0xReceiverAddress")
    amount := big.NewInt(1000000000000000000) // 1 ETH
    txHash, err := wallet.SendETH(toAddress, amount)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Transaction Hash: %s\n", txHash.Hex())
}
