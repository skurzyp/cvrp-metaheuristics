package model

type Problem struct {
	Instance       Instance
	ElitismCount   int
	MaxGenerations int
	MutationRate   float32
	PopulationSize int
}
