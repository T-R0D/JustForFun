package day19

import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.should.*

import day19.Day19Solution

class Day19SolutionTests extends AnyFunSuite with Matchers {
    test("partOne collects the letters on the path in the correct order") {
        val input = """#    |          
                       #    |  +--+    
                       #    A  |  C     
                       #F---|----E|--+
                       #    |  |  |  D 
                       #    +B-+  +--+ 
                       #""".stripMargin('#')
        val expected = "ABCDEF"
        val solution = Day19Solution()

        val result = solution.partOne(input)

        result should equal(Right(expected))
    }

    test("partOne counts the steps to walk the path") {
        val input = """#    |          
                       #    |  +--+    
                       #    A  |  C     
                       #F---|----E|--+
                       #    |  |  |  D 
                       #    +B-+  +--+ 
                       #""".stripMargin('#')
        val expected = "38"
        val solution = Day19Solution()

        val result = solution.partTwo(input)

        result should equal(Right(expected))
    }
}
        