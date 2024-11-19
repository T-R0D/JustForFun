package args

import scopt.OParser

val builder = OParser.builder[Config]

val parser = {
  import builder._

  OParser.sequence(
    programName("aoc2017"),
    head("aoc2017", "0.0.1"),
    cmd("generate")
      .text("Generate files for the project.")
      .action((_, config) => config.copy(mode = Option(Mode.Generate)))
      .children(
        opt[String]("target-dir")
          .text("Path to the directory the project should be generated in.")
          .required()
          .action((targetDir, config) => config.copy(genTargetDir = targetDir)),
        opt[Unit]("delete")
          .text("Delete the files that were generated.")
          .optional()
          .action((_, config) => config.copy(delete = true))
      ),
    cmd("solve")
      .text("Solve an Advent of Code problem.")
      .action((_, config) => config.copy(mode = Option(Mode.Solve)))
      .children(
        opt[String]("input")
          .text("Path to the input file to use.")
          .required()
          .action((input, config) => config.copy(input = input)),
        opt[Int]("day")
          .text("The problem day to solve.")
          .required()
          .validate(x =>
            if (1 <= x && x <= 25) success
            else failure("`day` must be in the range 1-25 (inclusive)")
          )
          .action((day, config) => config.copy(day = day)),
        opt[Int]("part")
          .text("The problem part to solve.")
          .required()
          .validate(x =>
            if (Seq(1, 2).contains(x)) success
            else failure("part must be either 1 or 2")
          )
          .action((part, config) => config.copy(part = part))
      )
  )
}
