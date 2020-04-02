package model

import "time"

type TokenDetails struct {
	AccessToken string
	RefreshToken string
	AccessID string
	RefreshID string
	AtExpires time.Duration
	RtExpires time.Duration
}