package cgp

const numberOfOperations = 15
const maxParameterValue = 256

type Cgp struct {
	chromosomeMaxValues []int
	outputs             int
	mutationProbability float32
}

func NewCgp(inputs int, size int, outputs int, mutationProbability float32) *Cgp {
	var cgp Cgp
	cgp.chromosomeMaxValues = make([]int, 4*size+outputs)

	for i := 0; i < size; i++ {
		cgp.chromosomeMaxValues[i*4] = inputs + i
		cgp.chromosomeMaxValues[i*4+1] = inputs + i
		cgp.chromosomeMaxValues[i*4+2] = numberOfOperations
		cgp.chromosomeMaxValues[i*4+3] = maxParameterValue
	}

	outputsOffset := 4 * size
	for i := 0; i < outputs; i++ {
		cgp.chromosomeMaxValues[outputsOffset+i] = size
	}

	cgp.outputs = outputs
	cgp.mutationProbability = mutationProbability

	return &cgp
}

func (cgp Cgp) GenerateChromosome() *Chromosome {
	return GenerateChromosome(&cgp.chromosomeMaxValues)
}

func (cgp Cgp) Evaluate(inputs []int, chromosome *Chromosome) *[]int {
	size := len(cgp.chromosomeMaxValues) - cgp.outputs
	numberOfInputs := len(inputs)

	values := make([]int, size)
	for i := 0; i < size; i++ {
		leftOperandAddress := (*chromosome)[4*i]
		rightOperandAddress := (*chromosome)[4*i+1]
		functionNumber := (*chromosome)[4*i+2]
		parameter := (*chromosome)[4*i+3]

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
		outputs[i] = values[(*chromosome)[4*size+i]]
	}

	return &outputs
}
