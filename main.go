package main

import "sync"

const difficulty = 1

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}


var Blockchain []Block

type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

