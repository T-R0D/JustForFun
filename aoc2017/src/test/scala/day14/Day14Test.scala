package day14

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day14.Day14Solution

class Day14SolutionTests extends AnyFunSuite with Matchers {
    test("partOne counts used memory blocks") {
        val input = "flqrgnkx"
        val expected = "8108"
        val solution = Day14Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }

    test("partTwo counts contiguous memory groups") {
        val input = "flqrgnkx"
        val expected = "1242"
        val solution = Day14Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
}
        