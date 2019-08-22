package services

import "fmt"

type GalaxyNumberNode struct {
	galaxyNumber   *GalaxyNumber
	next           *GalaxyNumberNode
	tail           *GalaxyNumberNode
	subtractedWith string
	subtracted     bool
	value          float64
	previous       *GalaxyNumberNode
}

func (gn *GalaxyNumberNode) AddNextNode(next *GalaxyNumberNode, nextAppearance int) error {
	var nextValue float64 = next.value
	var nextName string = next.galaxyNumber.Name
	var nextRepeatable bool = next.galaxyNumber.Repeatable

	var tailValue float64 = gn.tail.value
	var tailSubtractable bool = gn.tail.galaxyNumber.Subtractable
	var tailSubtracted bool = gn.tail.subtracted
	var tailName string = gn.tail.galaxyNumber.Name
	var tailSubtractedWith string = gn.tail.subtractedWith
	var tailPrevious = gn.tail.previous

	var previousValue float64

	if tailPrevious != nil {
		previousValue = tailPrevious.value
	}

	var subtrahedLimit float64 = tailValue * 10

	if tailSubtracted {
		switch true {
		case tailName == nextName:
			return fmt.Errorf("Error : '%v' cannot appear after '%v-%v'", nextName, tailName, tailSubtractedWith)
		case nextValue > tailValue:
			return fmt.Errorf("Error : '%v' after '%v', there cannot not be a bigger or equal number after subtraction", nextName, tailSubtractedWith)
		}
	}
	if nextAppearance >= 1 {
		if nextRepeatable == false {
			return fmt.Errorf("Error : '%v' only can appear once", nextName)
		}

	}
	if nextValue > tailValue {
		switch true {
		case tailSubtractable == false:
			return fmt.Errorf("Error : '%v' cannot be subtracted", tailName)
		case nextValue > subtrahedLimit:
			return fmt.Errorf("Error : '%v' cannot be subtracted with '%v'", tailName, nextName)
		case previousValue != 0 && nextValue > previousValue:
			return fmt.Errorf("Error : subtraction cannot appear after equal or lower number")
		default:
			gn.tail.value = nextValue - tailValue
			gn.tail.subtractedWith = next.galaxyNumber.Name
			gn.tail.subtracted = true
			return nil
		}
	}

	if tailName == nextName {
		switch true {
		case nextAppearance >= 3:
			if tailSubtracted == false {
				return fmt.Errorf("Error :  '%v' appear more than three times in a succession", tailName)
			}

		}
	}
	next.previous = gn.tail
	gn.tail.next = next
	gn.tail = next

	return nil
}

func (gn *GalaxyNumberNode) GetArabicNumber() float64 {
	var result float64
	if gn.next == nil {
		result = gn.value
		return result
	}
	for gn := gn; gn != nil; gn = gn.next {
		result += gn.value
	}
	return result
}
