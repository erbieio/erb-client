package client

type APIs interface {
	NormalTransaction(to string, value int64, data string) (string, error)
	Mint(royalty uint32, metaURL string, exchanger string) (string, error)
	Transfer(nftAddress, to string) (string, error)
	Author(nftAddress, to string) (string, error)
	AuthorRevoke(nftAddress, to string) (string, error)
	AccountAuthor(to string) (string, error)
	AccountAuthorRevoke(to string) (string, error)
	SNFTToERB(nftAddress string) (string, error)
	TokenPledge(proxySign []byte, proxyAddress string) (string, error)
	TokenRevokesPledge() (string, error)
	Open(feeRate uint32, name, url string) (string, error)
	Close() (string, error)
	InsertNFTBlock(dir, startIndex string, number uint64, royalty uint32, creator string) (string, error)
	TransactionNFT(buyer []byte, to string) (string, error)
	BuyerInitiatingTransaction(seller1 []byte) (string, error)
	FoundryTradeBuyer(seller2 []byte) (string, error)
	FoundryExchange(buyer, seller2 []byte, to string) (string, error)
	NftExchangeMatch(buyer, seller, exchangerAuth []byte, to string) (string, error)
	FoundryExchangeInitiated(buyer, seller2, exchangerAuthor []byte, to string) (string, error)
	NFTDoesNotAuthorizeExchanges(buyer, seller1 []byte, to string) (string, error)
	AdditionalPledgeAmount(value int64) (string, error)
	RevokesPledgeAmount(value int64) (string, error)
	VoteOfficialNFT(dir, startIndex string, number uint64, royalty uint32, creator string) (string, error)                                          //23
	VoteOfficialNFTByApprovedExchanger(dir, startIndex string, number uint64, royalty uint32, creator string, exchangerAuth []byte) (string, error) //24
	UnforzenAccount() (string, error)                                                                                                               //25
	RecoverCoefficient() (string, error)
	AccountDelegate(proxySign []byte, proxyAddress string) (string, error) //31
}
