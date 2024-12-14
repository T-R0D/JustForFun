package day22

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day22.Day22Solution

class Day22SolutionTests extends AnyFunSuite with Matchers {
    test("partOne counts the times the infection is distributed") {
        val input = """|..#
                       |#..
                       |...""".stripMargin
        val expected = "5587"
        val solution = Day22Solution()

        val result = solution.partOne(input)

        result should equal(Right(expected))
    }
    
    test("partTwo counts the times the infection is distributed by the evolved virus") {
        val input = """|..#
                       |#..
                       |...""".stripMargin
        val expected = "2511944"
        val solution = Day22Solution()

        val result = solution.partTwo(input)

        result should equal(Right(expected))
    }
}
        