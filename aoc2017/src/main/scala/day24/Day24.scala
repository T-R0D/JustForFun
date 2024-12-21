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

    def constructGraph(segments: Seq[Segment]): (Map[Int, Segment], Map[Int, Seq[Int]]) = {
        val lookup = segments.zipWithIndex.map(_.swap).toMap

        val graph = (
            for {
                (i, a) <- lookup
                (j, b) <- lookup
                if i != j && (a.end == b.start || a.start == b.end)
            } yield {
                if a.end == b.start then {
                    (i, j)
                } else {
                    (j, i)
                }
            }
        ).toSeq.groupBy(_._1).mapValues(x => x.map(_._2)).toMap

        (lookup, graph)
    }

    case class SegmentUsage(label: Int, segment: Segment, reversed: Boolean) {
        def nextMatch(): Int = {
            if reversed then {
                segment.start
            } else {
                segment.end
            }
        }
    }

    def findStrongestPossibleBridgeStartingWith(
        lookup: Map[Int, Segment],
        graph: Map[Int, Seq[Int]],
        usedSegments: Seq[SegmentUsage],
    ): (Seq[Int], Int) = {
        val recentSegment = usedSegments.last
        val candidateLabels = graph
            .getOrElse(recentSegment.label, Seq.empty)
            .filter { (candidateLabel) =>
                val candidateSegment = lookup.get(candidateLabel)
                candidateSegment match {
                    case None => false
                    case Some(segment) => !usedSegments.contains(segment)
                }
            }

        if candidateLabels.isEmpty then {
            val s = usedSegments.map(_.segment)
            (s, scoreBridge(s))
        } else {
            candidateLabels.foldLeft(0) { (acc, candidateLabel) =>
                
            }
        }
    }

    def scoreBridge(segments: Seq[Segment]): Int = {
        segments.foldLeft(0) { (acc, segment) =>
            acc + (segment.start + segment.end)
        }
    }
}

        