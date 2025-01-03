package day17

import solution.Solution

class Day17Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val stepSize = parseSpinlockStepSize(input)
        val nIterations = 2017

        val ringList = simulateSpinlockPath(nIterations, stepSize)

        Right(ringList.nextPosition(nIterations).toString)
    }

    override def partTwo(input: String): Either[String, String] = {
        val stepSize = parseSpinlockStepSize(input)
        val nIterations = 50_000_000

        val finalValue = simulateSpinlockEnteringSinglePosition(1, stepSize, nIterations)

        Right(finalValue.toString)
    }

    def parseSpinlockStepSize(input: String): Int = {
        input.toInt
    }

    def simulateSpinlockPath(nIterations: Int, stepSize: Int): RingList = {     
        val resultRingList = (1 to nIterations).foldLeft(RingList.withCapacity(nIterations + 1)) { (acc, nextValue) =>
            acc.insertNext(steps = stepSize, nextValue)
        }
        
        resultRingList
    }

    class RingList private (data: Array[Int], currentPosition: Int) {
        def insertNext(steps: Int, value: Int): RingList = {
            val newData = data.clone()
            val positionAfterStepping = (0 until steps).foldLeft(currentPosition) { (acc, _) =>
                data(acc)
            }
            val tempNext = data(positionAfterStepping)
            newData(value) = tempNext
            newData(positionAfterStepping) = value
            RingList(newData, value)
        }

        def nextPosition(i: Int): Int = {
            data(i)
        }
    }

    object RingList {
        def withCapacity(capacity: Int): RingList = {
            val data = Array.fill(capacity)(-1)
            data(0) = 0
            RingList(data, 0)
        }
    }

    def simulateSpinlockEnteringSinglePosition(targetPosition: Int, stepSize: Int, nIterations: Int): Int = {
        val (targetValue, _, _) = (1 to nIterations).foldLeft((-1, 0, 0)) { (acc, i) =>
            val (targetValue, offset, currentPosition) = acc
            
            val newPosition = ((currentPosition + stepSize) % i) + 1
            val newOffset = {
                if newPosition == 0 then {
                    offset + 1
                } else {
                    offset
                }
            }
            val newTargetValue = {
                if newPosition == targetPosition + offset then {
                    i
                } else {
                    targetValue
                }
            }
            (newTargetValue, newOffset, newPosition)
        }

        targetValue
    }
}
