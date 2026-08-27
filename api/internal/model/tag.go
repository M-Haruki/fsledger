package model

type TagType int

const (
	StockTag TagType = iota
	TransactionTag
	FlowTag
)
