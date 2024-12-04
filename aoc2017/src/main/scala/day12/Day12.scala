package day12

import solution.Solution
import scala.annotation.targetName

class Day12Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val adjacencyList = parseAdjacencyList(input)

        val reachable = findAllReachable(adjacencyList, "0")

        Right(reachable.size.toString)

    override def partTwo(input: String): Either[String, String] =
        val adjacencyList = parseAdjacencyList(input)

        val groups = findAllDistinctGroups(adjacencyList)

        Right(groups.size.toString)
        
    def parseAdjacencyList(input: String): Map[String, Seq[String]] =
        (
            for 
                association <- input.split("\\n")
            yield
                val srcDests = association.split("\\s+<->\\s+")
                val src = srcDests(0)
                val dests = srcDests(1).split(",\\s+").toSeq
                src -> dests
        ).toMap

    def findAllReachable(adjacencyList: Map[String, Seq[String]], src: String): Set[String] =
        findAllReachableInternal(adjacencyList, src, Set.empty)

    def findAllReachableInternal(adjacencyList: Map[String, Seq[String]], src: String, seen: Set[String]): Set[String] =
        if seen.contains(src) then
            seen
        else
            val updatedSeen = seen + src
            adjacencyList.getOrElse(src, Seq.empty).foldLeft(updatedSeen) { (acc, neighbor) =>
                acc | findAllReachableInternal(adjacencyList, neighbor, acc)
            }

    def findAllDistinctGroups(adjacencyList: Map[String, Seq[String]]): Seq[Set[String]] =
        adjacencyList.keys.foldLeft(Seq.empty: Seq[Set[String]]) { (acc, src) =>
            if acc.exists(_.contains(src)) then
                acc
            else
                acc :+ findAllReachable(adjacencyList, src)
        }
