package day24

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day24.Day24Solution

class Day24SolutionTests extends AnyFunSuite with Matchers {
    test("partOne finds the strength of the strongest possible bridge") {
        val input = """|0/2
                       |2/2
                       |2/3
                       |3/4
                       |3/5
                       |0/1
                       |10/1
                       |9/10""".stripMargin
        val expected = "31"
        val solution = Day24Solution()

        val result = solution.partOne(input)

        result should equal(Right(expected))
    }

    test("partTwo finds the strength of the strongest longest bridge") {
        val input = """|0/2
                       |2/2
                       |2/3
                       |3/4
                       |3/5
                       |0/1
                       |10/1
                       |9/10""".stripMargin
        val expected = "19"
        val solution = Day24Solution()

        val result = solution.partTwo(input)

        result should equal(Right(expected))
    }
}
        