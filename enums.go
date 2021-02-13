package main

type EvictionPolicyTypes string

const (
	LRU EvictionPolicyTypes = "LRU"
	LFU EvictionPolicyTypes  = "LFU"
	TBE EvictionPolicyTypes  = "TBE"
)