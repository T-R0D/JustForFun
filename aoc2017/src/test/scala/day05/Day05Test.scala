package day05

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day05.Day05Solution

class Day05SolutionTests extends AnyFunSuite with Matchers:
    test("partOne executes the program in the correct number of steps") {
        val input = """|0
                       |3
                       |0
                       |1
                       |-3
                       |""".stripMargin
        val expected = "5"
        val solution = Day05Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
    
    test("partTwo executes the program in the correct number of steps") {
        val input = """|0
                       |3
                       |0
                       |1
                       |-3
                       |""".stripMargin
        val expected = "10"
        val solution = Day05Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
        