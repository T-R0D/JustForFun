package day01

import solution.Solution

class Day01Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val sequence = input.map(c => c.toInt - '0')
        val rotatedSequence = sequence.drop(1) ++ sequence.take(1)
        Right(sequence.zip(rotatedSequence).foldLeft(0) { (acc, pair) =>
            if (pair._1 == pair._2) {
                acc + pair._1
            } else {
                acc
            }
        }.toString)

    override def partTwo(input: String): Either[String, String] =
        val sequence = input.map(c => c.toInt - '0')
        val halfLength = sequence.length / 2
        val rotatedSequence = sequence.drop(halfLength) ++ sequence.take(halfLength)
        Right(sequence.zip(rotatedSequence).foldLeft(0) { (acc, pair) =>
            if (pair._1 == pair._2) {
                acc + pair._1
            } else {
                acc
            }
        }.toString)
        