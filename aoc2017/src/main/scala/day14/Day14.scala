package day14

import solution.Solution
import scala.collection.BitSet

class Day14Solution extends Solution:
    override def partOne(input: String): Either[String, String] = {
        val hashKey = parseHashKey(input)

        val memoryGrid = computeMemoryGridState(hashKey)
        val usedBlocks = memoryGrid.map(_.sum).sum

        Right(usedBlocks.toString)
    }

    override def partTwo(input: String): Either[String, String] = ???

    def parseHashKey(input: String): String = {
        input
    }

    def computeMemoryGridState(hashKey: String): Seq[Seq[Int]] = {
        val rowHashes = {
            for {
                i <- 0 until 128
            } yield {
                val rowIdentifier = s"$hashKey-$i"
                knotHash(rowIdentifier)
            }
        }
        val memoryGrid = {
            for {
                hash <- rowHashes
            } yield {
                hash.foldLeft(Seq.empty: Seq[Int]) { (acc, hexDigit) =>
                    acc ++ hexDigitToBinaryRepresentation(hexDigit)    
                }
            }
        }

        memoryGrid
    }

    def hexDigitToBinaryRepresentation(digit: Char): Seq[Int] = {
        digit match
            case '0' => Seq(0, 0, 0, 0)
            case '1' => Seq(0, 0, 0, 1)
            case '2' => Seq(0, 0, 1, 0)
            case '3' => Seq(0, 0, 1, 1)
            case '4' => Seq(0, 1, 0, 0)
            case '5' => Seq(0, 1, 0, 1)
            case '6' => Seq(0, 1, 1, 0)
            case '7' => Seq(0, 1, 1, 1)
            case '8' => Seq(1, 0, 0, 0)
            case '9' => Seq(1, 0, 0, 1)
            case 'a' => Seq(1, 0, 1, 0)
            case 'b' => Seq(1, 0, 1, 1)
            case 'c' => Seq(1, 1, 0, 0)
            case 'd' => Seq(1, 1, 0, 1)
            case 'e' => Seq(1, 1, 1, 0)
            case 'f' => Seq(1, 1, 1, 1)
            case _ => Seq(99999)
    }

    def countContiguousMemoryGroups(memoryMap: Seq[Seq[Int]]): Seq[Set[(Int, Int)]] = {
        for {
            (row, i) <- memoryMap.zipWithIndex
            (_, j) <- row.zipWithIndex
        } yield {
            Set.empty
        }
    }

    def knotHash(input: String): String = {
        val inputAsciiValues = input.map(_.toInt)
        val saltedInput = inputAsciiValues ++ Seq(17, 31, 73, 47, 23)

        val initialRopeState = RopeState((0 until 256).map(identity).toSeq, 0, 0)
        val finalRopeState = 
            (0 until 64).foldLeft(initialRopeState) { (ropeState, _) =>
                saltedInput.foldLeft(ropeState) { (intermediateRopeState, byte) =>
                    intermediateRopeState.fold(byte)
                }
            }

        val denseHash = densifyHash(finalRopeState.positions)

        hashBytesToHexString(denseHash)
    }

    case class RopeState(positions: Seq[Int], currentPosition: Int, skip: Int) {
        def fold(length: Int): RopeState = {
            val newPositions =
                for
                    i <- 0 until positions.size
                yield
                    if isInFoldRange(length, i) then
                        positions(reversedIndex(length, i))
                    else
                        positions(i)
            RopeState(
                positions = newPositions,
                currentPosition = (currentPosition + length + skip) % positions.size,
                skip = skip + 1,
            )
        }

        def isInFoldRange(length: Int, i: Int): Boolean = {
            val inNormalSpan = (currentPosition <= i && i < currentPosition + length)
            val lengthWraps = (currentPosition + length) >= positions.size
            val inWrappedSpan = lengthWraps && i < (currentPosition + length) % positions.size
            inNormalSpan || inWrappedSpan
        }

        def reversedIndex(length: Int, i: Int): Int = {
            if currentPosition <= i then
                (currentPosition + (length - (i - currentPosition + 1))) % positions.size
            else
                (currentPosition + (length - (i + positions.size - currentPosition + 1))) % positions.size
        }
    }

    def densifyHash(rope: Seq[Int]): Seq[Int] = {
        rope.grouped(16).map(_.reduce((a, b) => a ^ b)).toSeq
    }

    def hashBytesToHexString(denseHash: Seq[Int]): String = {
        denseHash.map(x => String.format("%02x", x.toByte)).mkString
    }
