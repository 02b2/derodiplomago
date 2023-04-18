package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	dero "github.com/deroproject/derohe/rpc"
)

var (
	daemonAddress string
	walletAddress string
	rpcLogin      string
)

func main() {
	// Parse command line arguments
	flag.StringVar(&daemonAddress, "daemon", "127.0.0.1:10102", "Address of the Dero daemon")
	flag.StringVar(&walletAddress, "wallet", "127.0.0.1:10103", "Address of the Dero wallet")
	flag.StringVar(&rpcLogin, "rpclogin", "", "RPC login credentials (username:password)")
	flag.Parse()

	action := flag.Arg(0)
	switch action {
	case "getAddress":
		err := getAddress(walletAddress, rpcLogin)
		if err != nil {
			log.Fatalf("Error getting wallet address: %v", err)
		}
	case "getBalance":
		err := getBalance(walletAddress, rpcLogin)
		if err != nil {
			log.Fatalf("Error getting wallet balance: %v", err)
		}
	case "getHeight":
		err := getHeight(daemonAddress)
		if err != nil {
			log.Fatalf("Error getting daemon height: %v", err)
		}
	default:
		fmt.Println("Unknown command. Supported commands are: getAddress, getBalance, getHeight")
	}
}

func getAddress(walletAddress, rpcLogin string) error {
	rpcClientW, ctx, cancel := setWalletClient(walletAddress, rpcLogin)
	defer cancel()

	var result *dero.GetAddress_Result
	err := rpcClientW.CallFor(ctx, &result, "GetAddress")

	if err != nil {
		return err
	}

	fmt.Printf("Dero Address: %s\n", result.Address)
	return nil
}

func getBalance(walletAddress, rpcLogin string) error {
	rpcClientW, ctx, cancel := setWalletClient(walletAddress, rpcLogin)
	defer cancel()

	var result *dero.GetBalance_Result
	err := rpcClientW.CallFor(ctx, &result, "GetBalance")

	if err != nil {
		return err
	}

	unlocked := float64(result.Unlocked_Balance) / 100000
	locked := float64(result.Balance-result.Unlocked_Balance) / 100000

	fmt.Printf("Unlocked balance: %.5f\nLocked balance: %.5f\n", unlocked, locked)
	return nil
}

func getHeight(daemonAddress string) error {
	rpcClientD, ctx, cancel := setDaemonClient(daemonAddress)
	defer cancel()

	var result *dero.GetHeight_Result
	err := rpcClientD.CallFor(ctx, &result, "GetHeight")

	if err != nil {
		return err
	}

	fmt.Printf("Daemon height: %d\n", result.Height)
	return nil
}

func setWalletClient(walletAddress, rpcLogin string) (*dero.RPCClient, context.Context, context.CancelFunc) {
	rpcEndpoint := fmt.Sprintf("http://%s", walletAddress)
	if rpcLogin != "" {
		rpcEndpoint = fmt.Sprintf("%s:%s@%s", rpcEndpoint, rpcLogin)
	}
	client := dero.NewRPCClient(rpcEndpoint)
	ctx, cancel := context.WithCancel(context.Background())
	return client, ctx, cancel
}

func setDaemonClient(daemonAddress string) (*dero.RPCClient, context.Context, context.CancelFunc) {
	rpcEndpoint := fmt.Sprintf("http://%s", daemonAddress)
	client := dero.NewRPCClient(rpcEndpoint)
	ctx, cancel := context.WithCancel(context.Background())
	return client, ctx, cancel
}
