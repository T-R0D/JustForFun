package day02

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day02.Day02Solution

class Day02SolutionTests extends AnyFunSuite with Matchers:
    test("partOne should checksum spreadsheet") {
        val input = """5 1 9 5
                      |7 5 3
                      |2 4 6 8
                      |""".stripMargin
        val expected = "18"
        val solution = Day02Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
    
    test("partTwo should sum the divisible numbers from each row") {
        val input = """5 9 2 8
                      |9 4 7 3
                      |3 8 6 5
                      |""".stripMargin
        val expected = "18"
        val solution = Day02Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
        