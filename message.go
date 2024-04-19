package main

type Message struct {
	Port     string `json:"port"`
	ClientID string `json:"clientID"`
	Text     string `json:"text"`
}
