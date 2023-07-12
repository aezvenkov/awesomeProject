package main

type TaggantSettings struct {
	TaggantName        string   `json:"taggantName"`
	SignalFirstRange   *IntPair `json:"signalFirstRange"`
	SignalSecondRange  *IntPair `json:"signalSecondRange"`
	SignalThirdRange   *IntPair `json:"signalThirdRange"`
	SignalFourthRange  *IntPair `json:"signalFourthRange"`
	SignalFifthRange   *IntPair `json:"signalFifthRange"`
	SignalSixthRange   *IntPair `json:"signalSixthRange"`
	SignalSeventhRange *IntPair `json:"signalSeventhRange"`
	SignalEightRange   *IntPair `json:"signalEightRange"`
}

type IntPair struct {
	First  int `json:"first"`
	Second int `json:"second"`
}
