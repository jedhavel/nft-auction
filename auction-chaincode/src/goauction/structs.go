package main

// User - Create Buyer, Seller , Auction House, Authenticator
// Could establish valid UserTypes -
// AH (Auction House)
// TR (Buyer or Seller)
// AP (Appraiser)
// IN (Insurance)
// BK (Bank)
// SH (Shipper)
/////////////////////////////////////////////////////////////

// The json tags just tell the reader how to name the json fields //
type User struct {
	UserID    string `json:"userId,required"`
	docType   string `json:"docType"` // Type = USER
	Name      string `json:"name,required"`
	Password  string `json:"password,required"`
	UserType  string `json:"userType,required"` // Auction House (AH), Bank (BK), Buyer or Seller (TR), Shipper (SH), Appraiser (AP)
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	PaymentID string `json:"paymentID"`
	Timestamp string `json:"timeStamp"`
}
type itemobject struct {
	ItemID            string `json:"itemID,required"`
	docType           string `json:"docType"`
	ItemDesc          string `json:"itemDescription"`
	ItemDetail        string `json:"itemDetail"` // Could included details such as who created the Art work if item is a Painting
	NumberOfCopies    string `json:"numberOfCopies"`
	ItemDate          string `json:"itemDate"`
	ItemType          string `json:"itemType"`
	ItemSubject       string `json:"itemSubject"`
	ItemHash          string `json:"itemHash"`          // Hash of the Item Image
	ItemHashSignature string `json:"itemHashSignature"` // Signed version of ItemHash by the item creator
	ItemMedia         string `json:"itemMedia"`
	ItemSize          string `json:"itemSize"`
	itemImageName     string `json:"itemImageName"` // Item Subject + Extension
	itemImageType     string `json:"itemImageType"` // should be used to regenerate the appropriate image type
	TimeStamp         string `json:"timeStamp"`     // This is the time stamp
	ObjectCat         string `json:"objectCat"`     // ART, MUSIC et cetra
}

// AuctionRequest - Register a request for participating in an auction
// Usually posted by a seller who owns a piece of ITEM
// The Auction house will determine when to open the item for Auction
// The Auction House may conduct an appraisal and genuineness of the item
// ======================================================================================
type AuctionRequest struct {
	AuctionID      string `json:"auctionID"`
	docType        string `json:"docType"` // AUCREQ
	NftId          string `json:"nftId"`
	itemImage      string `json:"itemImage"`
	itemImageName  string `json:"itemImageName"`  // Item Subject + Extension
	AESKey         string `json:"aesKey"`         // AES Key of the Item supplied by Seller. This is used to validate ownership
	AuctionHouseID string `json:"auctionHouseID"` // ID of the Auction House managing the auction
	SellerID       string `json:"sellerID"`       // ID Of Seller - to verified against the Item CurrentOwnerId
	RequestDate    string `json:"requestDate"`    // Date on which Auction Request was filed
	ReservePrice   string `json:"reservePrice"`   // reserver price > previous purchase price
	BuyItNowPrice  string `json:"buyItNowPrice"`  // 0 (Zero) if not applicable else specify price
	Status         string `json:"status"`         // INIT, OPEN, CLOSED (To be Updated by Trgger Auction)
	OpenDate       string `json:"openDate"`       // Date on which auction will occur (To be Updated by Trigger Auction)
	CloseDate      string `json:"closeDate"`      // Date and time when Auction will close (To be Updated by Trigger Auction)
	TimeStamp      string `json:"timeStamp"`      // The transaction Date and Time
}

// OpenAuctionRequest - Open an Auction for Bids
// ======================================================================================
type OpenAuctionRequest struct {
	AuctionRequestID     string `json:"auctionRequestID"`
	docType              string `json:"docType"` // AUCOPEN
	Duration             string `json:"duration"`
	AuctionStartDateTime string `json:"auctionStartDateTime"` // The transaction Date and Time
}

// Bid - Submit bids against Auction Request
type Bid struct {
	AuctionID string `json:"auctionID"`
	DocType   string `json:"docType"` // BID
	BidID     string `json:"bidID"`
	NftId     string `json:"nftId"`
	BuyerID   string `json:"buyerID"`  // ID Of Buyer - to be verified against the Item CurrentOwnerId
	BidPrice  string `json:"bidPrice"` // BidPrice > Previous Bid
	BidTime   string `json:"bidTime"`  // Time the bid was received
}
