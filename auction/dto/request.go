package dto

import "time"

type CreateAuctionRequest struct {
	Name    string                   `json:"name" validate:"required,min=1"`
	EndDate CreateAuctionRequestTime `json:"endDate" validate:"required,dateFormatCheck"`
}

type CreateAuctionRequestTime time.Time

func (ct *CreateAuctionRequestTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// 따옴표 제거
	s = s[1 : len(s)-1]
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*ct = CreateAuctionRequestTime(t)
	return nil
}

type CreateBiddingRequest struct {
	UserID       int `json:"userId"`
	RequestPrice int `json:"requestPrice"`
}
