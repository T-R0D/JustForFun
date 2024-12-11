package day20

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day20.Day20Solution

class Day20SolutionTests extends AnyFunSuite with Matchers {
    test("partOne finds the ID of the most homebound particle") {
        val input = """|p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
                       |p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
                       |""".stripMargin
        val expected = "0"
        val solution = Day20Solution()

        val result = solution.partOne(input)

        result should equal(Right(expected))
    }
}
        