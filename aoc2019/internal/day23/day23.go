package day23

import (
	"aoc2019/internal/intcode"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	result := simulateNetworkInteractions(input, 50)
	return result, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	result := simulateNetworkInteractions2(input, 50)
	return result, nil
}

type networkComputer struct {
	*intcode.Computer
	inputQueue []packet
}

func newNetworkComputer(prog string, networkAddress int) *networkComputer {
	c := &networkComputer{
		Computer:   intcode.NewComputer(),
		inputQueue: make([]packet, 0),
	}
	c.InputProgram(prog)
	c.SetInterruptibleMode()
	c.RunProgram()
	c.Input(networkAddress)
	c.RunProgram()
	if c.GetState() == intcode.STATE_AWAITING_INPUT {
		c.Input(-1)
		c.RunProgram()
	}

	return c
}

func (n *networkComputer) enqueuePacket(p packet) {
	n.inputQueue = append(n.inputQueue, p)
}

func (n *networkComputer) consumeInput() {
	if len(n.inputQueue) > 0 {
		packet := n.inputQueue[0]
		n.inputQueue = n.inputQueue[1:]

		for i := 0; i < 2; i += 1 {
			if n.GetState() != intcode.STATE_AWAITING_INPUT {
				panic("Computer needs to be awaiting input.")
			}
			switch i {
			case 1:
				n.Input(packet.X)
			case 2:
				n.Input(packet.Y)
			}
			n.RunProgram()
		}
	} else {
		if n.GetState() != intcode.STATE_AWAITING_INPUT {
			panic("Computer needs to be awaiting input.")
		}
		n.Input(-1)
		n.RunProgram()
	}
}

func (n *networkComputer) outputPacket() packet {
	packet := packet{}
	for i := 0; i < 3; i++ {
		if n.GetState() != intcode.STATE_AWAITING_OUTPUT {
			panic("Computer needs to have output to give.")
		}
		switch i {
		case 0:
			packet.Address = n.CollectOutput()
		case 1:
			packet.X = n.CollectOutput()
		case 2:
			packet.Y = n.CollectOutput()
		}
		n.RunProgram()
	}
	return packet
}

type packet struct {
	Address, X, Y int
}

func simulateNetworkInteractions(prog string, nComputers int) int {
	computers := make([]*networkComputer, nComputers)
	for i := range computers {
		computers[i] = newNetworkComputer(prog, i)
	}

	packetQueue := make([]packet, 0)
	for {
		if len(packetQueue) > 0 {
			packet := packetQueue[0]
			packetQueue = packetQueue[1:]

			if packet.Address == 255 {
				return packet.Y
			}

			n := computers[packet.Address]
			n.Input(packet.X)
			n.RunProgram()
			n.Input(packet.Y)
			n.RunProgram()
		}

		for _, n := range computers {
			switch n.GetState() {
			case intcode.STATE_AWAITING_OUTPUT:
				packetQueue = append(packetQueue, n.outputPacket())
			case intcode.STATE_AWAITING_INPUT:
				n.Input(-1)
				n.RunProgram()
			}
		}
	}

	// How I think the problem should have gone:
	// stillRunning := true
	// for stillRunning {
	// 	stillRunning = false
	// 	for _, n := range computers {
	// 		switch n.GetState() {
	// 		case intcode.STATE_AWAITING_INPUT:
	// 			stillRunning = true
	// 			for len(n.inputQueue) > 0 {
	// 				n.consumeInput()
	// 			}
	// 		case intcode.STATE_AWAITING_OUTPUT:
	// 			stillRunning = true
	// 			for n.GetState() == intcode.STATE_AWAITING_OUTPUT {
	// 				packet := n.outputPacket()
	// 				if packet.Address == 255 {
	// 					return packet.Y
	// 				} else {
	// 					computers[packet.Address].enqueuePacket(packet)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return 0
}

func simulateNetworkInteractions2(prog string, nComputers int) int {
	computers := make([]*networkComputer, nComputers)
	for i := range computers {
		computers[i] = newNetworkComputer(prog, i)
	}

	natPacket := packet{X: -1, Y: -1}
	lastSentNatPacket := packet{X: 0, Y: 0}
	packetQueue := make([]packet, 0)
	for {
		for len(packetQueue) > 0 {
			packet := packetQueue[0]
			packetQueue = packetQueue[1:]

			if packet.Address == 255 {
				natPacket = packet
			} else {
				n := computers[packet.Address]
				n.Input(packet.X)
				n.RunProgram()
				n.Input(packet.Y)
				n.RunProgram()
			}
		}

		idle := true
		for _, n := range computers {
			switch n.GetState() {
			case intcode.STATE_AWAITING_OUTPUT:
				packet := n.outputPacket()
				packetQueue = append(packetQueue, packet)
				idle = false
			case intcode.STATE_AWAITING_INPUT:
				n.Input(-1)
				n.RunProgram()
			}
		}

		if idle {
			if natPacket.Y == lastSentNatPacket.Y {
				return natPacket.Y
			}
			natPacket.Address = 0
			packetQueue = append(packetQueue, natPacket)
			lastSentNatPacket = natPacket
		}
	}

	return 0
}
