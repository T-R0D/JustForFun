package day06

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day06.Day06Solution

class Day06SolutionTests extends AnyFunSuite with Matchers:
    test("partOne computes the iterations until repeat") {
        val input = "0\t2\t7\t0"
        val expected = "5"
        val solution = Day06Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
    
    test("partTwo computes the cycle length") {
        val input = "0\t2\t7\t0"
        val expected = "4"
        val solution = Day06Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
        