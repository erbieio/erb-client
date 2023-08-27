package test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/rossoman/erb-client/client"
	"github.com/ethereum/go-ethereum/common"
)

const (
	endpoint         = "http://192.168.4.240:8560"
	priAddress       = "0x8724fd5d3e4a63e0017b8a2a4fC775B91166eD8d"
	priKey           = "50fd980dab6b010c001fcab754421792b451c48706d5bb69ac0ad93ab8dd7aa1"
	buyerPriKey      = "057b05b9cff85c963c3ab90d26503700646781f938054171461b17ad5f7082db"
	buyerAddress     = "0xED8cfA91F533C47863E520DE828E2f970c8f52DB"
	sellerPriKey     = "132a8ed2918b923a91c324d0e22a358ea6a82330a1faf956042f64bce8bf8e46"
	sellerAddress    = "0xD9DC702C0d3518aa27F82fd75f1a544233a7150f"
	exchangerPriKey  = "0bbbb60fa9ff05081a3b63aa8b043d1281cd860dc06d92800aee5b1fdf5bc8d7"
	exchangeAddress  = "0xaECE03150f0A6565e8308E872Bf3Ec143A0b4879"
	exchangerPriKey1 = "1ba6579ab53049bd9730d644815a49fd065d6626e74862873402a99227564c3d"
	exchangeAddress1 = "0xC83279D0fEdd3A814aBe250007c61E324Ba467E6"
	tempAddress      = "0x1f4c00c477651531f5e9F145B317E16c02FE3b1D"
	tempPriKey       = "474a14156d0a4fec502cd67b62cbc72f14358e84225afb11cc15c8127a43ac40"
)

func TestNewClient(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	_ = worm
}

// Recharge
func TestRecharge(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.NormalTransaction("0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", 1000, "")
	fmt.Println(rs)
}

// Mint
// NFT mint 0
func TestMint(t *testing.T) {
	worm := client.NewClient(sellerPriKey, endpoint)
	//rs, _ := worm.Mint(10, "/ipfs/ddfd90be9408b4", exchangeAddress)
	rs, _ := worm.Mint(10, "/ipfs/Qmf3xw9rEmsjJdQTV3ZcyF4KfYGtxMkXdNQ8YkVqNmLHY8", "")
	fmt.Println(rs)
}

//0x8fa2d4b70013407012d002fa395939cb0d322553e4848aaae78d4fad638bef55

// Transfer
// NFT transfer 1
func TestTransfer(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.Transfer("0x0000000000000000000000000000000000000001", sellerAddress)
	fmt.Println(rs)
}

//0x5e8dd659b0ceb95ab53ce32d37daa8688accab601ce58c75e706f08bb47617f4

// Author Single
// NFT authorization 2
func TestAuthor(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.Author("0x0000000000000000000000000000000000000002", exchangeAddress)
	fmt.Println(rs)
}

//0x2657d46f0c4ef16cadbc6842896c1b50f41333d6a247ee43e5085da5d7e3feff

// AuthorRevoke
// Cancel a single authorization 3
func TestAuthorRevoke(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AuthorRevoke("0x0000000000000000000000000000000000000002", exchangeAddress)
	fmt.Println(rs)
}

//0xe043dc7d8505d01f6cd949b7a7cc4ed685a9e1b640195801c3c6265b7d11efee

// AccountAuthor
// All NFTs under the authorized account 4
func TestAccountAuthor(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AccountAuthor(exchangeAddress)
	fmt.Println(rs)
}

//0x6b42237b9dad13211d89f1e6c66cf947bb371f407a4621ffcf7fd73e385f6fea

// AccountAuthorRevoke
// Cancel all NFTs under the authorized account 5
func TestAccountAuthorRevoke(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AccountAuthorRevoke(exchangeAddress)
	fmt.Println(rs)
}

//0x1dee05dff7ea39874ed8401c91288ae627b56ae1df6dc4c26a856fafab0447c5

// SNFTToERB
// Fragment NFT exchange 6
func TestSNFTToERB(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.SNFTToERB("0x8000000000000000000000000000000000000004")
	fmt.Println(rs)
}

//0x77ff920a3a649378e4c7a58644bece643e379113b2bc99257b894a29e220e157

// TokenPledge
// ERB pledge 9
func TestTokenPledge(t *testing.T) {
	worm := client.NewClient(exchangerPriKey1, endpoint)
	toaddr := common.HexToAddress(exchangeAddress1)
	rs, _ := worm.TokenPledge(toaddr, "", "exchange", "www.exchange.com", 700, 100)
	fmt.Println(rs)
}

//0x6ceb02802455ab959964866410f37a2f0fcd78e7e64e87d6c9d8102de7f9974b

// TokenRevokesPledge
// ERB revokes pledge 10
func TestTokenRevokesPledge(t *testing.T) {
	worm := client.NewClient(tempPriKey, endpoint)
	toaddr := common.HexToAddress(tempAddress)

	rs, _ := worm.TokenRevokesPledge(toaddr, 1)
	fmt.Println(rs)
}

//0xcbd19c386d8b5944d4a88017680239651edefc527e4ba2c8762ab0df2333a7ca

// Open
// Open an exchange 11
//func TestOpen(t *testing.T) {
//	worm := client.NewClient(exchangerPriKey, endpoint)
//	rs, _ := worm.Open(10, "wormholes", "www.kang123456.com")
//	fmt.Println(rs)
//}

//0xcccd6c9499224e7d216f3bd230447900550b07345841eebd2e62b613f7bd924f

// Close
// close a exchange 12
//func TestClose(t *testing.T) {
//	worm := client.NewClient(exchangerPriKey, endpoint)
//	rs, _ := worm.Close()
//	fmt.Println(rs)
//}

//0x61cd018d6e70af47c6204fea18db5b33fdecc92162cca66b0089783733809e84

// TransactionNFT 14
func TestTransactionNFT(t *testing.T) {
	worm := client.NewClient(buyerPriKey, endpoint)
	number, _ := worm.BlockNumber(context.Background())
	blockNumber := fmt.Sprintf("0x%x", number+10)
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000002", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", blockNumber, "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(buyer))

	worm1 := client.NewClient(sellerPriKey, endpoint)
	rs, _ := worm1.TransactionNFT(buyer, buyerAddress)
	fmt.Println(rs)
}

//0xc9c4e6652ba411a0435d2e3187f019329b084734f19ae6699ee7f1fa9a92123b

// BuyerInitiatingTransaction 15
func TestBuyerInitiatingTransaction(t *testing.T) {
	worm := client.NewClient(sellerPriKey, "")
	seller1, err := worm.Wallet.SignSeller1("0x38D7EA4C68000", "0x0000000000000000000000000000000000000003", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", "0x677")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(seller1))

	worm1 := client.NewClient(buyerPriKey, endpoint)
	rs, _ := worm1.BuyerInitiatingTransaction(seller1)
	fmt.Println(rs)
}

//0xfb9cf0100340c9bf965fc0f8ef44bb8a75af58175deab0dcff3979a97a8ebefa

// FoundryTradeBuyer 16
func TestFoundryTradeBuyer(t *testing.T) {
	worm := client.NewClient(sellerPriKey, "")
	seller2, err := worm.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", "0x677")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(seller2))

	worm1 := client.NewClient(buyerPriKey, endpoint)
	rs, _ := worm1.FoundryTradeBuyer(seller2)
	fmt.Println(rs)
}

//0x4634d6bbc36b9444914a259c2acf0410af0b99122baef30d7a8701a496bc3b6c

// FoundryExchange 17
func TestFoundryExchange(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller2, err := worm1.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm2.FoundryExchange(buyer, seller2, buyerAddress)
	fmt.Println(rs)
}

//0x70853466fdf5dc4476fab34b79f9be2e66f0448789937094de0b0aa5f3345e8c

// ftExchangeMatch  18
func TestNftExchangeMatch(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000004", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller, err := worm1.Wallet.SignSeller1("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000004", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	rs, _ := worm3.NftExchangeMatch(buyer, seller, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

//0xf11e024297b89e6dfd02bc2da4680cea353ea6956c3ea9084afa40d58477932f

// FoundryExchangeInitiated 19
func TestFoundryExchangeInitiated(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(buyer))

	worm1 := client.NewClient(sellerPriKey, "")
	seller2, err := worm1.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(seller2))

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(exchangeAuth))

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	rs, _ := worm3.FoundryExchangeInitiated(buyer, seller2, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

//0xc9cc570057faf1edd83f48833520f9d546e4972083ee705152b5f35630f1588d

// FtDoesNotAuthorizeExchanges 20
func TestNFTDoesNotAuthorizeExchanges(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller1, err := worm1.Wallet.SignSeller1("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, endpoint)

	rs, _ := worm2.NFTDoesNotAuthorizeExchanges(buyer, seller1, buyerAddress)
	fmt.Println(rs)
}

//0x95615a6c7a164537257492c112a9fcd99907315893706a1b104456d9e3aa8af6

// AdditionalPledgeAmount 21
func TestAdditionalPledgeAmount(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.AdditionalPledgeAmount(100)
	fmt.Println(rs)
}

//0x25f2ed8cf5f1041be9e71d483a32b01fd3f7820ec59e0c060830214c53fea5f9

// AdditionalPledgeAmount 22
func TestRevokesPledgeAmount(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.RevokesPledgeAmount(100)
	fmt.Println(rs)
}

//0xd2c7f943f0f5364b0928c518e7b6de7491c0e8efb6abf912a17e6860f70ebec1

// VoteOfficialNFT
func TestVoteOfficialNFT(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.VoteOfficialNFT("wormholes2", "0x640001", 6553600, 20, "0xab7624f47fd7dadb6b8e255d06a2f10af55990fe")
	fmt.Println(rs)
}

// VoteOfficialNFTByApprovedExchanger
func TestVoteOfficialNFTByApprovedExchanger(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0x0")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(exchangeAuth))
	worm1 := client.NewClient(exchangeAddress1, endpoint)
	rs, _ := worm1.VoteOfficialNFTByApprovedExchanger("wormholes2", "0x640001", 6553600, 20, "0xab7624f47fd7dadb6b8e255d06a2f10af55990fe", exchangeAuth)
	fmt.Println(rs)
}

// ChangeRewardsType
// change revenue model 25
func TestUnforzenAccount(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.UnforzenAccount()
	fmt.Println(rs)
}

// WeightRedemption
// restore the weight 26
func TestWeightRedemption(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.WeightRedemption()
	fmt.Println(rs)
}

// BatchSellTransfer
// Batch buying and selling of minted NFT or S-Nft 27
func TestBatchSellTransfer(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyerauth, err := worm.Wallet.SignBuyerAuth(exchangeAddress, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	sellerauth, err := worm1.Wallet.SignSellerAuth(exchangeAddress, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	buyer, err := worm3.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0x6000", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}
	seller, err := worm3.Wallet.SignSeller1("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}
	rs, _ := worm3.BatchSellTransfer(buyer, seller, buyerauth, sellerauth, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

// ForceBuyingTransfer
// Compulsory purchase of S-Nft 28
func TestForceBuyingTransfer(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyerauth, err := worm.Wallet.SignBuyerAuth(exchangeAddress, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}
	buyer, err := worm.Wallet.SignBuyer("", "0x800000000000000000000000000000000000000", exchangeAddress, "0x6000", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0x6000")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	rs, _ := worm3.ForceBuyingTransfer(buyer, buyerauth, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

// ExtractERB
// Addresses with L3 can initiate this transaction to withdraw ERB 29
func TestExtractERB(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.ExtractERB()
	fmt.Println(rs)
}

// AccountDelegate
// Delegate large accounts to small accounts 31
func TestAccountDelegate(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	proxySign, _ := worm.Wallet.SignDelegate("address", "pledgeAccount")
	rs, _ := worm.AccountDelegate(proxySign, buyerAddress)
	fmt.Println(rs)
}

func TestGetBalance(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	balance, _ := worm.Balance(context.Background(), exchangeAddress)
	fmt.Println(balance)
}

// func TestCheckNFTPool(t *testing.T) {
// 	worm := client.NewClient("c1e74da8e26c5a60870089f59695a1b243887f9d23571d24c7f011b8eb068768", "http://192.168.4.240:8561")

// 	var flag bool
// 	num := int64(22)
// 	for {
// 		if flag {
// 			break
// 		} else {
// 			current, _ := worm.BlockNumber(context.Background())
// 			if uint64(num) > current {
// 				time.Sleep(time.Second * 5)
// 			} else {
// 				fmt.Println("num: ", num)
// 				//res1, _ := worm.NFT.GetBlockBeneficiaryAddressByNumber(context.Background(), num)
// 				//for _, miners := range *res1 {
// 				//	if miners.Address == common.HexToAddress("0xEE3168308949237d395202F134C4243630ebB4A8") {
// 				//		fmt.Println("miner", miners.Address)
// 				//		flag = true
// 				//		//break
// 				//	}
// 				//}

// 				//res1, _ := worm.NFT.GetValidators(context.Background(), num)
// 				//for _, validator := range res1.Validators {
// 				//	fmt.Println("validator", validator)
// 				//	if validator.Addr == common.HexToAddress("0xA7aa3f181aebE59ca697D803B2197cfA50A3913E") {
// 				//		fmt.Println("miner", validator) //0x0F7094Cf6391273AAC99b478b8Eca9D636BBbf0c
// 				//		flag = true
// 				//		break
// 				//	}
// 				//}

// 				res1, _ := worm.GetActiveLivePool(context.Background(), uint64(num))
// 				for _, miners := range res1.ActiveMiners {
// 					fmt.Println("miner", miners)
// 					if miners.Address == common.HexToAddress("0xA7aa3f181aebE59ca697D803B2197cfA50A3913E") {
// 						fmt.Println("miner", miners) //0x0F7094Cf6391273AAC99b478b8Eca9D636BBbf0c
// 						flag = true
// 						break
// 					}
// 				}

// 				//res1, _ := worm.NFT.QueryMinerProxy(context.Background(), num, "0xA7F60Adc80E09F71a7A56044003a2B606Ed1Cac2")
// 				//for _, miners := range res1 {
// 				//	if miners.Address == common.HexToAddress("0x279c59A0DC597276bac3D160Cb1596beFA46bad2") {
// 				//		fmt.Println("miner", miners)
// 				//		flag = true
// 				//		break
// 				//	}
// 				//}
// 				num++
// 			}
// 		}
// 	}
// }

func TestGetSNFT(t *testing.T) {
	exchanger := make(map[string]string)
	exchanger["0x68B14e0F18C3EE322d3e613fF63B87E56D86Df60"] = "d8cf127b1780c0a0e0d2e40519ae2c611d6d7f6b8b706c967ed8183170267d99"
	exchanger["0xeEF79493F62dA884389312d16669455A7E0045c1"] = "9bdbec1e6329a5484105c05aacbbce9ff78a287d20cbd8a8b59c414b5e1edbb6"
	exchanger["0xa5999Cc1DEC36a632dF735064Dc75eF6af0E7389"] = "b6290ad66f10eead80c1371be065af9493ff0cc611fa6d4c207f46e2516f2f38"
	exchanger["0x63d913dfDB75C7B09a1465Fe77B8Ec167793096b"] = "b1c0f70e418cdc851534c6a09c40a50b676466819c3cd65a7aeed9cb581d1643"
	exchanger["0xF50f73B83721c108E8868C5A2706c5b194A0FDB1"] = "f17a19d3d0c4620759e4e365ef79f2553b0639fd1a7bdfbafe570f7e3d59f7aa"
	exchanger["0xB000811Aff6e891f8c0F1aa07f43C1976D4c3076"] = "ec299549a07e9e6202999445dccfe6a1efdc3af75dd942461a403d4a3a03edb3"
	exchanger["0x257F3c6749a0690d39c6FBCd2DceB3fB464f0F94"] = "382b13e70a7e66f7f6d94007b977c1ad6acdc8f454ee77e3e5bb159d0e09f7cb"
	exchanger["0x43ee5cB067F29B920CC44d5d5367BCEb162B4d9E"] = "405321241ccffe1d2bddcac1202209460a5a0caded3a9b203bdbba5c40f45de0"
	exchanger["0x85D3FDA364564c365870233E5aD6B611F2227846"] = "efdb9f92fbae899e8069a41c3ed589f6fdaf9cc0be1da86bb5d0cf77ccf3b5d3"
	exchanger["0xDc807D83d864490C6EEDAC9C9C071E9AAeD8E7d7"] = "ef5664558107effaa7a20d01c328037a15e9a4989a06be79249f517dad7c7eea"
	exchanger["0xFA623BCC71BE5C3aBacfe875E64ef97F91B7b110"] = "f6842d3207b8b81a5ea1e3d08fcb013ec2ef8a320e325252cd2af18c390772fe"
	exchanger["0xb17fAe1710f80Eb9a39732862B0058077F338B21"] = "38f6551752c4c561fe68abe365eae069cc667ae31a92bf3d52df468d918454c6"
	exchanger["0x86FFd3e5a6D310Fcb4668582eA6d0cfC1c35da49"] = "d60c5a8a3fdc26b22533d1c5fffdb11c12b17771cd9f2380e71df30a8970a8b1"
	exchanger["0x6bB0599bC9c5406d405a8a797F8849dB463462D0"] = "04a5ddb33b11fff6923b5eee08f949fead766e9d92a42f4350c726a1b18ffc81"
	exchanger["0x765C83dbA2712582C5461b2145f054d4F85a3080"] = "a1a78a79fb1159a4c871a20a60f1a05ece8189115226fda182565d027b0015da"

	var collects = "0xC65F08C9Dfceb0988631B175E293Af5666535CF0"

	var Empty, _ = new(big.Int).SetString("0x0000000000000000000000000000000000000000", 16)

	worm := client.NewClient("38fc3f36f420ca662e0b423342b61243337a84f992eb60847a67cb8fe90af133", "http://192.168.4.240:8561")
	Nft, _ := new(big.Int).SetString("8000000000000000000000000000000000000000", 16)
	ctx := context.Background()
	for {
		latest, _ := worm.BlockNumber(ctx)
		address := common.BytesToAddress(Nft.Bytes())
		res1, _ := worm.GetAccountInfo(context.Background(), address.String(), int64(latest))

		if (*res1).Nft.Owner.String() == common.BytesToAddress(Empty.Bytes()).String() {
			time.Sleep(time.Second * 5)
		}

		for ex, pri := range exchanger {
			fmt.Println((*res1).Nft.Owner.String())
			fmt.Println(ex)
			if strings.ToLower(ex) == strings.ToLower(res1.Nft.Owner.String()) {
				worms := client.NewClient(pri, "http://192.168.4.240:8561")
				worms.Transfer(common.BytesToAddress(Nft.Bytes()).String(), collects)
				break
			}
		}
		Nft = new(big.Int).Add(Nft, big.NewInt(1))
	}
}

//"number":           (*hexutil.Big)(head.Number),
//"hash":             head.Hash(),
//"parentHash":       head.ParentHash,
//"nonce":            head.Nonce,
//"mixHash":          head.MixDigest,
//"sha3Uncles":       head.UncleHash,
//"logsBloom":        head.Bloom,
//"stateRoot":        head.Root,
//"miner":            miner,
//"difficulty":       (*hexutil.Big)(head.Difficulty),
//"extraData":        hexutil.Bytes(head.Extra),
//"size":             hexutil.Uint64(head.Size()),
//"gasLimit":         hexutil.Uint64(head.GasLimit),
//"gasUsed":          hexutil.Uint64(head.GasUsed),
//"timestamp":        hexutil.Uint64(head.Time),
//"transactionsRoot": head.TxHash,
//"receiptsRoot":     head.ReceiptHash,

type BlockInfo struct {
	Hash    string
	PreHash string
	Number  uint64
	Miner   string
}

func TestAnalysisBlocks(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	blockInfoMap := make(map[uint64]*BlockInfo, 0)
	for {
		time.Sleep(1 * time.Second)
		currentBlockNumber, _ := worm.BlockNumber(context.Background())
		currentBlock, err := worm.GetBlockByNumber(context.Background(), new(big.Int).SetUint64(currentBlockNumber))
		if err != nil {
			continue
		}
		t.Log(currentBlock["hash"])
		t.Log(currentBlock["parentHash"])
		t.Log(currentBlock["miner"])
		hash := currentBlock["hash"].(string)
		prehash := currentBlock["parentHash"].(string)
		miner := currentBlock["miner"].(string)
		t.Log(hash)
		t.Log(prehash)
		t.Log(miner)

		currentBlockInfo := &BlockInfo{
			Hash:    hash,
			PreHash: prehash,
			Number:  currentBlockNumber,
			Miner:   miner,
		}

		v, ok := blockInfoMap[currentBlockNumber]
		if ok {
			if v.Hash != hash {
				t.Error("fork, two blocks have same blocknumber, but not same hash \n",
					"blocknumber ", currentBlockNumber, "\nold hash ", v.Hash, "new hash", hash,
					"\nold miner ", v.Miner, "new miner ", miner)
			}
		} else {
			t.Log("current block number ", currentBlockNumber)
			blockInfoMap[currentBlockNumber] = currentBlockInfo
		}
		if preBlockInfo, ok := blockInfoMap[currentBlockNumber-1]; ok {
			if preBlockInfo.Hash != currentBlockInfo.PreHash {
				t.Error("fork, new block's prehash is not same with the parent block's hash\n",
					"new block number ", currentBlockInfo.Number,
					"\nnew block's prehash ", currentBlockInfo.PreHash,
					"parent block's hash ", preBlockInfo.Hash)
				break
			}
		}

		t.Log("map len ", len(blockInfoMap))
		var startDeleteIndex uint64
		var deleteIndexs []uint64
		if len(blockInfoMap) > 1000 {
			startDeleteIndex = currentBlockNumber - 1000

			for k, _ := range blockInfoMap {
				if k <= startDeleteIndex {
					deleteIndexs = append(deleteIndexs, k)
				}
			}

			for _, v := range deleteIndexs {
				delete(blockInfoMap, v)
			}
			deleteIndexs = deleteIndexs[:0]
		}
	}
}
