package day09

import solution.Solution

class Day09Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val stream = parseStream(input)

        val score = scoreGroups(stream)

        Right(score.toString)

    override def partTwo(input: String): Either[String, String] =
        val stream = parseStream(input)

        val garbageCount = countGarbage(stream)

        Right(garbageCount.toString)

    def parseStream(input: String): Seq[Char] =
        input.toSeq
        
    def scoreGroups(stream: Seq[Char]): Int =
        processStream(
            stream, 
            StreamProcessState(
                scoreSoFar = 0,
                garbageCount = 0,
                openGroups = 0,
                ignored = false,
                inGarbage = false
            ),
        )._1

    def countGarbage(stream: Seq[Char]): Int =
        processStream(
            stream, 
            StreamProcessState(
                scoreSoFar = 0,
                garbageCount = 0,
                openGroups = 0,
                ignored = false,
                inGarbage = false
            ),
        )._2
    
    case class StreamProcessState(
        scoreSoFar: Int,
        garbageCount: Int,
        openGroups: Int,
        ignored: Boolean,
        inGarbage: Boolean,
    )

    @scala.annotation.tailrec
    final def processStream(stream: Seq[Char], state: StreamProcessState): (Int, Int) =
        stream.headOption match
            case None => (state.scoreSoFar, state.garbageCount)
            case Some(c) =>
                val nextState = {
                    if state.ignored then
                        state.copy(ignored = false)
                    else if state.inGarbage && c == '!' then
                        state.copy(ignored = true)
                    else if state.inGarbage && c == '>' then
                        state.copy(inGarbage = false)
                    else if state.inGarbage then
                        state.copy(garbageCount = state.garbageCount + 1)
                    else if c == '<' then
                        state.copy(inGarbage = true)
                    else if c == '{' then
                        state.copy(openGroups = state.openGroups + 1)
                    else if c == '}' then
                        state.copy(
                            openGroups = state.openGroups - 1,
                            scoreSoFar = state.scoreSoFar + state.openGroups
                        )
                    else
                        state
                }
                processStream(stream.drop(1), nextState)
