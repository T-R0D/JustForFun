package day12

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day12.Day12Solution

class Day12SolutionTests extends AnyFunSuite with Matchers {
    test("partOne determines all IDs in a group") {
        val input = """|0 <-> 2
                       |1 <-> 1
                       |2 <-> 0, 3, 4
                       |3 <-> 2, 4
                       |4 <-> 2, 3, 6
                       |5 <-> 6
                       |6 <-> 4, 5""".stripMargin
        val expected = "6"
        val solution = Day12Solution()

        val result = solution.partOne(input)

        result should equal (Right(expected))
    }
    
    test("partTwo counts all distinct groups") {
        val input = """|0 <-> 2
                       |1 <-> 1
                       |2 <-> 0, 3, 4
                       |3 <-> 2, 4
                       |4 <-> 2, 3, 6
                       |5 <-> 6
                       |6 <-> 4, 5""".stripMargin
        val expected = "2"
        val solution = Day12Solution()

        val result = solution.partTwo(input)

        result should equal (Right(expected))
    }
}
        