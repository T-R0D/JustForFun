package day24

import scala.util.{Try}

import solution.Solution

class Day24Solution extends Solution {
    override def partOne(input: String): Either[String, String] = {
        val results = parseBridgeSegments(input)
        results match
            case Left(errors) => Left(errors.mkString("; ")) 
            case Right(segments) => {
                val s = segments.sortBy(_.start)
                for {
                    (a, i) <- s.zipWithIndex
                    (b, j) <- s.zipWithIndex
                    if i != j && a.end == b.start
                } do {
                    println(s"$i[${a.start}/${a.end}] --> $j[${b.start}/${b.end}]")
                }

                Left("not done")
            }
    }

    override def partTwo(input: String): Either[String, String] = ???

    case class Segment(start: Int, end: Int)

    def parseBridgeSegments(input: String): Either[Seq[String], Seq[Segment]] = {
        val results = (
            for {
                (line, i) <- input.split("\\n").zipWithIndex
            } yield {
                Try {
                    val parts = line.split("/").map(_.toInt)
                    Segment(parts(0), parts(1))
                }.toEither.left.map(t => s"line $i: $t")
            }
        )

        val failures = results.collect({
            case Left(value) => value
        })
        if !failures.isEmpty then {
            Left(failures)
        } else {
            Right(results.collect({
                case Right(value) => value
            }))
        }
    }

    def findStrongestPossibleBridgeStrength(segments: Seq[Segment]): Int = {
        val (lookup, graph) = constructGraph(segments)

        val potentialFirstSegments = lookup.filter(x => x._2.start == 0 || x._2.end == 0)

        val bestStrength = potentialFirstSegments
            .map { (labelAndSegment) =>
                val segmentsUsed = findStrongestPossibleBridgeStartingWith(lookup, graph, labelAndSegment._1)
                segmentsUsed.foldLeft(0) { (acc, segmentLabel) =>
                    val segment = lookup.get(segmentLabel)
                    segment match
                        case Some(s) => acc + s.start + s.end
                    
                }
            }
            .max

    }

    def constructGraph(segments: Seq[Segment]): (Map[Int, Segment], Map[Int, Int]) = {
        val lookup = segments.zipWithIndex.map(_.swap).toMap

        val graph = (
            for {
                (i, a) <- lookup
                (j, b) <- lookup
                if a.end == b.start || a.start == b.end
            } yield {
                i -> j
            }
        )

        (lookup, graph)
    }

    def findStrongestPossibleBridgeStartingWith(
        lookup: Map[Int, Segment],
        graph: Map[Int, Int],
        startSegment: Int
    ): Seq[Int] = {???}
}

        