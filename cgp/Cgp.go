package cgp

import "math/rand"

const numberOfOperations = 15
const maxParameterValue = 256

type Cgp struct {
	chromosomes         [][]int
	chromosomeMaxValues []int
	outputs             int
	mutationProbability float32
}

func NewCgp(inputs int, size int, outputs int, numberOfChromosomes int, mutationProbability float32) *Cgp {
	var cgp Cgp

	cgp.chromosomes = make([][]int, numberOfChromosomes)
	for i := 0; i < numberOfChromosomes; i++ {
		cgp.chromosomes[i] = make([]int, 4*size+outputs)
	}

	for i := 0; i < size; i++ {
		cgp.chromosomeMaxValues[i*4] = inputs + i
		cgp.chromosomeMaxValues[i*4+1] = inputs + i
		cgp.chromosomeMaxValues[i*4+2] = numberOfOperations
		cgp.chromosomeMaxValues[i*4+3] = maxParameterValue
	}

	outputsOffset := 4 * size
	for i := 0; i < size; i++ {
		cgp.chromosomeMaxValues[outputsOffset+i] = size
	}

	cgp.outputs = outputs
	cgp.mutationProbability = mutationProbability

	return &cgp
}

func (cgp Cgp) Randomize() {
	numberOfChromosomes := len(cgp.chromosomes)
	size := len(cgp.chromosomeMaxValues)
	for i := 0; i < numberOfChromosomes; i++ {
		for j := 0; j < size; j++ {
			cgp.chromosomes[i][j] = rand.Intn(cgp.chromosomeMaxValues[j])
		}
	}
}

func (cgp Cgp) GetChromosome(number int) *[]int {
	size := len(cgp.chromosomeMaxValues) - cgp.outputs
	chromosome := make([]int, size)
	for i := 0; i < size; i++ {
		chromosome[i] = cgp.chromosomes[number][i]
	}

	return &chromosome
}

func (cgp Cgp) SetParent(parent []int) {
	size := len(cgp.chromosomeMaxValues) - cgp.outputs
	for i := 0; i < size; i++ {
		cgp.chromosomes[0][i] = parent[i]
	}
}

func (cgp Cgp) Mutate() {
	numberOfChromosomes := len(cgp.chromosomes)
	size := len(cgp.chromosomeMaxValues) - cgp.outputs
	for i := 1; i < numberOfChromosomes; i++ {
		for j := 0; j < size; j++ {
			if rand.Float32() < cgp.mutationProbability {
				cgp.chromosomes[i][j] = rand.Intn(cgp.chromosomeMaxValues[j])
			} else {
				cgp.chromosomes[i][j] = cgp.chromosomes[0][j]
			}
		}

	}
}

func (cgp Cgp) evaluate(inputs []int, chromosomeNumber int) *[]int {
	size := len(cgp.chromosomeMaxValues) - cgp.outputs
	numberOfInputs := len(inputs)

	values := make([]int, size)
	for i := 0; i < size; i++ {
		leftOperandAddress := cgp.chromosomes[chromosomeNumber][4*i]
		rightOperandAddress := cgp.chromosomes[chromosomeNumber][4*i+1]
		functionNumber := cgp.chromosomes[chromosomeNumber][4*i+2]
		parameter := cgp.chromosomes[chromosomeNumber][4*i+3]

		var leftOperand int
		var rightOperand int
		if leftOperandAddress < numberOfInputs {
			leftOperand = inputs[leftOperandAddress]
		} else {
			leftOperand = values[leftOperandAddress-numberOfInputs]
		}
		if rightOperandAddress < numberOfInputs {
			rightOperand = inputs[rightOperandAddress]
		} else {
			rightOperand = values[rightOperandAddress-numberOfInputs]
		}

		values[i] = evaluateFunction(leftOperand, rightOperand, functionNumber, parameter)
	}

	outputs := make([]int, cgp.outputs)
	for i := 0; i < cgp.outputs; i++ {
		outputs[i] = values[cgp.chromosomes[chromosomeNumber][4*size+i]]
	}

	return &outputs
}
