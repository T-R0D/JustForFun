package day25

import solution.Solution

class Day25Solution extends Solution:
  override def partOne(input: String): Either[String, String] =
    val blueprint = parseInput(input)
    Right(runDiagnostic(blueprint).toString)

  case class Blueprint(
      startState: Char,
      diagnosticRunLength: Int,
      instructions: Map[Char, Array[InstructionSet]]
  )

  case class InstructionSet(writeValue: Int, increment: Int, nextState: Char)

  case class MachineState(
      iteration: Int,
      tape: Map[Int, Int],
      currentState: Char,
      location: Int
  )

  def parseInput(input: String): Blueprint =
    val sections = input.split("\n\n")
    val (startState, diagnosticRunLength) = parseConfigurationSection(
      sections(0)
    )

    val instructions = sections.drop(1).map(parseInstructionSet(_)).toMap

    Blueprint(startState, diagnosticRunLength, instructions)

  def parseConfigurationSection(section: String): (Char, Int) =
    val lines = section.split("\n")

    val startState = lines(0).replace("Begin in state ", "").replace(".", "")(0)
    val diagnosticRunLength = lines(1)
      .replace("Perform a diagnostic checksum after ", "")
      .replace(" steps.", "")
      .toInt

    (startState, diagnosticRunLength)

  def parseInstructionSet(section: String): (Char, Array[InstructionSet]) =
    val state =
      section.split("\n")(0).replace("In state ", "").replace(":", "")(0)

    val sets = section
      .split("\n")
      .drop(1)
      .mkString("\n")
      .split("\n")
      .grouped(4)
      .map(g => g.mkString("\n"))

    val instructionSet = sets
      .map(s =>
        val lines = s.split("\n").drop(1)

        val writeValue =
          lines(0)
            .replace("- Write the value ", "")
            .replace(".", "")
            .strip
            .toInt

        val incrementText =
          lines(1).replace("- Move one slot to the ", "").replace(".", "").strip
        val increment = if incrementText == "left" then -1 else 1

        val nextState =
          lines(2)
            .replace("- Continue with state ", "")
            .replace(".", "")
            .strip()(0)

        InstructionSet(writeValue, increment, nextState)
      )
      .toArray

    (state, instructionSet)

  def runDiagnostic(blueprint: Blueprint): Int =
    val resultTape =
      doRunDiagnostic(
        blueprint,
        MachineState(0, Map(), blueprint.startState, 0)
      )
    resultTape.values.sum

  @scala.annotation.tailrec
  final def doRunDiagnostic(
      blueprint: Blueprint,
      state: MachineState
  ): Map[Int, Int] =
    if state.iteration == blueprint.diagnosticRunLength then state.tape
    else
      val nextInstructionSet = blueprint.instructions(state.currentState)
      val currentValue = state.tape.getOrElse(state.location, 0)
      val instruction = nextInstructionSet(currentValue)

      val nextTape =
        if instruction.writeValue == 1 then state.tape + (state.location -> 1)
        else state.tape - state.location
      val nextLocation = state.location + instruction.increment

      val nextState = MachineState(
        state.iteration + 1,
        nextTape,
        instruction.nextState,
        nextLocation
      )

      doRunDiagnostic(blueprint, nextState)

  override def partTwo(input: String): Either[String, String] =
    Right("Merry Christmas!")
