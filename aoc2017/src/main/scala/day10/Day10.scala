package day10

import solution.Solution

class Day10Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val lengths = parseLengths(input)

        val finalRopeState = foldRope(256, lengths)

        Right(
            (finalRopeState.positions(0) * finalRopeState.positions(1)).toString
        )

    override def partTwo(input: String): Either[String, String] =
        val inputAsciiValues = parseInputAsciiValues(input)

        val hash = ropeHash(inputAsciiValues)

        Right(hash)
        
    def parseLengths(input: String): Seq[Int] =
        input.split(",").map(_.toInt).toSeq

    def parseInputAsciiValues(input: String): Seq[Int] =
        input.toSeq.map(_.toInt)
    
    def ropeHash(input: Seq[Int]): String =
        val saltedInput = input ++ Seq(17, 31, 73, 47, 23)

        val initialRopeState = RopeState((0 until 256).map(identity).toSeq, 0, 0)
        val finalRopeState = 
            (0 until 64).foldLeft(initialRopeState) { (ropeState, _) =>
                saltedInput.foldLeft(ropeState) { (intermediateRopeState, byte) =>
                    intermediateRopeState.fold(byte)
                }
            }

        val denseHash = densifyHash(finalRopeState.positions)

        hashBytesToHexString(denseHash)
        

    def foldRope(positionsOnRope: Int, lengths: Seq[Int]): RopeState =
        lengths.foldLeft(
            RopeState((0 until positionsOnRope).map(identity).toSeq, 0, 0)
        ) { (ropeState, length) =>
            ropeState.fold(length)
        }

    case class RopeState(positions: Seq[Int], currentPosition: Int, skip: Int):
        def fold(length: Int): RopeState =
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
        
        def isInFoldRange(length: Int, i: Int): Boolean =
            val inNormalSpan = (currentPosition <= i && i < currentPosition + length)
            val lengthWraps = (currentPosition + length) >= positions.size
            val inWrappedSpan = lengthWraps && i < (currentPosition + length) % positions.size
            inNormalSpan || inWrappedSpan

        def reversedIndex(length: Int, i: Int): Int =
            if currentPosition <= i then
                (currentPosition + (length - (i - currentPosition + 1))) % positions.size
            else
                (currentPosition + (length - (i + positions.size - currentPosition + 1))) % positions.size

    def densifyHash(rope: Seq[Int]): Seq[Int] =
        rope.grouped(16).map(_.reduce((a, b) => a ^ b)).toSeq

    def hashBytesToHexString(denseHash: Seq[Int]): String =
        denseHash.map(x => String.format("%02x", x.toByte)).mkString