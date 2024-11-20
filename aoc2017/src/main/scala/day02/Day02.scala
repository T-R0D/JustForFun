package day02

import solution.Solution

class Day02Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val spreadsheet = parseSpreadsheet(input)
        
        val checksum = spreadsheet
            .map(row => row.max - row.min)
            .fold(0)((a, b) => a + b)
            .toString

        Right(checksum)

    // 616 too high
    override def partTwo(input: String): Either[String, String] =
        val spreadsheet = parseSpreadsheet(input)

        val quotients =
            for
                row <- spreadsheet
                (a, i) <- row.zipWithIndex
                (b, j) <- row.zipWithIndex
                    if i != j && a % b == 0
            yield
                a / b
        val quotiensSum = quotients.fold(0)((a, b) => a + b)

        Right(quotiensSum.toString)
        
    def parseSpreadsheet(input: String): Seq[Seq[Int]] =
        input
            .split("\n")
            .map { line => 
                line
                    .split("\\s+")
                    .map(s => s.toInt)
                    .toSeq
            }
            .toSeq