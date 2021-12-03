package day3

import (
	"advent-of-code-2021/utils"
	"container/list"
	"strconv"
)

func Execute(level int) {
	input := utils.LoadData(3, level)
	intInput := utils.ConverStringArrayToBinIntArray(input)

	bits := len(input[0])
	enabledBits := 0
	disabledBits := 0

	output := make([]string, 1)

	if level == 1 {
		// Power consumption
		bit := int64(1)
		gamma := ""
		epsilon := ""

		for i := 0; i < bits; i++ {
			for j := 0; j < len(intInput); j++ {
				value := intInput[j]
				if value&bit > 0 {
					enabledBits++
				} else {
					disabledBits++
				}
			}

			if enabledBits > disabledBits {
				gamma = "1" + gamma
				epsilon = "0" + epsilon
			} else {
				gamma = "0" + gamma
				epsilon = "1" + epsilon
			}

			enabledBits = 0
			disabledBits = 0
			bit <<= 1
		}

		gammaNumber, _ := strconv.ParseInt(gamma, 2, 64)
		epsilonNumber, _ := strconv.ParseInt(epsilon, 2, 64)
		output[0] = strconv.FormatInt(gammaNumber*epsilonNumber, 10)
	} else {
		// Life support rating
		oxygenNumbers := utils.IntArrayToList(intInput)
		co2Numbers := oxygenNumbers

		highNumbers := list.New()
		lowNumbers := list.New()

		bit := int64(1) << int64(bits-1)

		for i := bits - 1; i >= 0; i-- {
			if oxygenNumbers.Len() > 1 {
				for e := oxygenNumbers.Front(); e != nil; e = e.Next() {
					value := e.Value.(int64)
					if value&bit > 0 {
						enabledBits++
						highNumbers.PushBack(value)
					} else {
						disabledBits++
						lowNumbers.PushBack(value)
					}
				}

				if enabledBits >= disabledBits {
					oxygenNumbers = highNumbers
				} else {
					oxygenNumbers = lowNumbers
				}

				enabledBits = 0
				disabledBits = 0
				highNumbers = list.New()
				lowNumbers = list.New()
			}

			if co2Numbers.Len() > 1 {
				for e := co2Numbers.Front(); e != nil; e = e.Next() {
					value := e.Value.(int64)
					if value&bit > 0 {
						enabledBits++
						highNumbers.PushBack(value)
					} else {
						disabledBits++
						lowNumbers.PushBack(value)
					}
				}

				if disabledBits > enabledBits {
					co2Numbers = highNumbers
				} else {
					co2Numbers = lowNumbers
				}

				enabledBits = 0
				disabledBits = 0
				highNumbers = list.New()
				lowNumbers = list.New()
			}

			bit >>= 1
		}

		oxygen := oxygenNumbers.Front().Value.(int64)
		co2 := co2Numbers.Front().Value.(int64)

		output[0] = strconv.FormatInt(oxygen*co2, 10)
	}

	utils.WriteData(3, level, output)
}
