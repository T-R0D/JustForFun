package args

import java.io.File

case class Config(
    mode: Option[Mode] = None,
    genTargetDir: String = "",
    delete: Boolean = false,
    input: String = "",
    day: Int = 0,
    part: Int = 0,
)

enum Mode:
    case Generate, Solve
