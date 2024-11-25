package day05

import solution.Solution

class Day05Solution extends Solution:
    override def partOne(input: String): Either[String, String] =
        val program = parseProgram(input)

        val stepsToTerminate = stepsToCompleteProgram(
            program,
            stepsToCompleteProgramHelper
        )

        Right(stepsToTerminate.toString)

    override def partTwo(input: String): Either[String, String] =
        val program = parseProgram(input)

        val stepsToTerminate = stepsToCompleteProgram(
            program,
            alternateStepsToCompleteProgramHelper
        )

        Right(stepsToTerminate.toString)

    def parseProgram(input: String): Seq[Int] =
        input.split("\\n").map(_.toInt).toSeq

    def stepsToCompleteProgram(
        program: Seq[Int],
        executionPlan: (Map[Int, Int], Int, Int) => Int,
    ): Int =
        val mappedProgram = program.zipWithIndex.map(p => (p._2, p._1)).toMap
        executionPlan(mappedProgram, 0, 0)

    @scala.annotation.tailrec
    final def stepsToCompleteProgramHelper(
        program: Map[Int, Int],
        programCounter: Int, stepsSoFar: Int
    ): Int =
        if !program.contains(programCounter) then
            stepsSoFar
        else
            stepsToCompleteProgramHelper(
                program.updatedWith(programCounter)(valOpt => valOpt.map(x => x + 1)), 
                programCounter + program.getOrElse(programCounter, 0),
                stepsSoFar + 1,
            )

    @scala.annotation.tailrec
    final def alternateStepsToCompleteProgramHelper(
        program: Map[Int, Int],
        programCounter: Int, stepsSoFar: Int
    ): Int =
        if !program.contains(programCounter) then
            stepsSoFar
        else
            alternateStepsToCompleteProgramHelper(
                program.updatedWith(programCounter) { valOpt =>
                    valOpt.map { x =>
                        if x >= 3 then
                            x - 1
                        else
                            x + 1
                    }
                }, 
                programCounter + program.getOrElse(programCounter, 0),
                stepsSoFar + 1,
            )

