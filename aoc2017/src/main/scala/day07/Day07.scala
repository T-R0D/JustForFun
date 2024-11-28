package day07

import solution.Solution

class Day07Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val (_, adjacencyList) = programGraphAndLookup(input)

        val order = topologicalSort(adjacencyList)

        order.headOption match
            case Some(name) => Right(name)
            case None => Left("Topological sort failed")
        

    override def partTwo(input: String): Either[String, String] =
        val (weights, adjacencyList) = programGraphAndLookup(input)

        val order = topologicalSort(adjacencyList)
        val rootName = order.head
        val tree = newWeightTree(weights, adjacencyList, rootName)
        val correctWeight = singleCorrectedWeight(tree)

        Right(correctWeight.toString)

    def programGraphAndLookup(input: String): (Map[String, Int], Map[String, Seq[String]]) =
        val lines = input.split("\\n+").toSeq
        
        val infoPairs = 
            for
                line <- lines
            yield
                val parts = line.split("\\s+->\\s+")
                val (name, weight) = separateNameAndWeight(parts(0))
                if parts.size == 1 then
                    (name -> weight, name -> Seq.empty)
                else
                    (name -> weight, name -> parts(1).split(",\\s+").toSeq)
        val (weights, neighbors) = infoPairs.unzip
        (weights.toMap, neighbors.toMap)

    def separateNameAndWeight(nameAndWeight: String): (String, Int) =
        val parts = nameAndWeight.replace("(", "").replace(")", "").split("\\s+").toSeq
        (parts(0), parts(1).toInt)

    def topologicalSort(programs: Map[String, Seq[String]]): Seq[String] =
        var order: Seq[String] = Seq.empty
        var visited: Set[String] = Set.empty

        def dfs(current: String): Unit =
            if visited.contains(current) then
                ()
            else
                programs.get(current) match
                    case Some(neighbors) =>
                        neighbors.foreach(dfs)
                    case None =>
                        ()
                visited += current
                order = current +: order
                

        for
            program <- programs.view.keys
        do
            dfs(program)

        order

    case class WeightTreeNode(cumulativeWeight: Int, supported: Seq[WeightTreeNode])

    def newWeightTree(
        weights: Map[String, Int],
        adjacencyList: Map[String, Seq[String]],
        root: String
    ): WeightTreeNode =
        def dfs(root: String): WeightTreeNode =
            val weight = weights.getOrElse(root, 0)

            val children =
                for 
                    childName <- adjacencyList.getOrElse(root, Seq.empty)
                yield
                    dfs(childName)

            val childTotalWeight = children.map(_.cumulativeWeight).sum()
            val cumulativeWeight = weight + childTotalWeight

            WeightTreeNode(cumulativeWeight, children)

        dfs(root)

    def singleCorrectedWeight(root: WeightTreeNode): Int =
        def helper(root: WeightTreeNode, targetWeight: Int): Int =
            root.supported match
                case Seq() => targetWeight
                case _ =>
                    val supportedWeights = root.supported.groupBy(_.cumulativeWeight).toSeq.sortBy(_._2.size)
                    supportedWeights.headOption match
                        case Some((supportedWeight, supportedGroup)) =>
                            if supportedGroup.size > 1 then
                                targetWeight - root.supported.map(_.cumulativeWeight).sum
                            else
                                val newTargetWeight = supportedWeights(1)._1
                                helper(supportedGroup(0), newTargetWeight)
                        case None => 0
                            
        helper(root, root.cumulativeWeight)

            