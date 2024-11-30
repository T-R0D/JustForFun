package day08

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day08.Day08Solution

class Day08SolutionTests extends AnyFunSuite with Matchers:
    test("partOne finds the largest value in any register") {
        val input = """|b inc 5 if a > 1
                       |a inc 1 if b < 5
                       |c dec -10 if a >= 1
                       |c inc -20 if c == 10""".stripMargin
        val expected = "1"
        val solution = Day08Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
    
    test("partTwo finds the largest value seen during execution") {
        val input = """|b inc 5 if a > 1
                       |a inc 1 if b < 5
                       |c dec -10 if a >= 1
                       |c inc -20 if c == 10""".stripMargin
        val expected = "10"
        val solution = Day08Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
        