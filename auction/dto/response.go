package dto

type CreateAuctionResponse struct {
	ID int `json:"id"`
}

type CreateBiddingResponse struct {
	IsSuccess bool `json:"isSuccess"`
}
