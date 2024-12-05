package day15

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day15.Day15Solution

class Day15SolutionTests extends AnyFunSuite with Matchers {
    test("partOne finds the number of (lower 16 bit) matches in 40 million iterations") {
        val input = """|Generator A starts with 65
                       |Generator B starts with 8921""".stripMargin
        val expected = "588"
        val solution = Day15Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
}
        