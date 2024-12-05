package day13

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day13.Day13Solution

class Day13SolutionTests extends AnyFunSuite with Matchers {
    test("partOne finds the total severity cost of leaving immediately") {
        val input = """|0: 3
                       |1: 2
                       |4: 4
                       |6: 4""".stripMargin
        val expected = "24"
        val solution = Day13Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }

    test("partTwo finds the minimum start delay to make it through with no severity") {
        val input = """|0: 3
                       |1: 2
                       |4: 4
                       |6: 4""".stripMargin
        val expected = "10"
        val solution = Day13Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
}
        