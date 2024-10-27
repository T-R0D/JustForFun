import copy
import heapq
import itertools


def part_one(puzzle_input):
    grid, elves, goblins = parse_map_and_unit_locations(puzzle_input)

    battle = Battle(grid, elves, goblins)
    result = battle.simulate_battle()

    return str(result)


def part_two(puzzle_input):
    grid, elves, goblins = parse_map_and_unit_locations(puzzle_input)

    lossless_battle_achieved = False
    attack_boost = 0
    battle_outcome = 0
    while not lossless_battle_achieved:
        battle = Battle(grid, elves, goblins)
        battle.boost_side("E", attack_boost)
        battle_outcome, lossless_battle_achieved = battle.simulate_lossless_battle("E")
        attack_boost += 1

    return str(battle_outcome)


def parse_map_and_unit_locations(puzzle_input):
    grid = []
    elves = {}
    goblins = {}

    next_id = 0
    for i, line in enumerate(puzzle_input.split("\n")):
        row = []
        for j, rune in enumerate(line):
            if rune == "E":
                elves[(i, j)] = Unit(next_id, "E")
                next_id += 1
                row.append(".")
            elif rune == "G":
                goblins[(i, j)] = Unit(next_id, "G")
                next_id += 1
                row.append(".")
            else:
                row.append(rune)
        grid.append(row)

    return grid, elves, goblins


STEPS_IN_READING_ORDER = (
    (-1, 0),
    (0, -1),
    (0, 1),
    (1, 0),
)


class Unit:
    def __init__(self, unit_id, kind):
        self.unit_id = unit_id
        self.kind = kind
        self.hp = 200
        self.attack = 3


class Battle:
    def __init__(self, grid, elves, goblins):
        self.grid = copy.deepcopy(grid)
        self.elves = copy.deepcopy(elves)
        self.goblins = copy.deepcopy(goblins)

    def boost_side(self, kind, attack_boost):
        boosted_side = self.elves if kind == "E" else self.goblins
        for unit in boosted_side.values():
            unit.attack += attack_boost

    def simulate_battle(self):
        full_rounds_completed = 0
        while self.elves and self.goblins:
            full_round_completed, _, _ = self.simulate_round()
            if full_round_completed:
                full_rounds_completed += 1

        winners = self.elves if self.elves else self.goblins
        return full_rounds_completed * sum(winner.hp for winner in winners.values())

    def simulate_lossless_battle(self, winning_kind):
        full_rounds_completed = 0
        while self.elves and self.goblins:
            full_round_completed, elf_losses, goblin_losses = self.simulate_round()

            if winning_kind == "E" and elf_losses > 0:
                return 0, False
            elif winning_kind == "G" and goblin_losses > 0:
                return 0, False

            if full_round_completed:
                full_rounds_completed += 1

        winners = self.elves if self.elves else self.goblins
        return (
            full_rounds_completed * sum(winner.hp for winner in winners.values()),
            True,
        )

    def simulate_round(self):
        turn_order = sorted(itertools.chain(self.elves.keys(), self.goblins.keys()))
        already_evaluated_unit_ids = set()

        elf_losses = 0
        goblin_losses = 0

        for unit_location in turn_order:
            if not self.elves or not self.goblins:
                return False, elf_losses, goblin_losses

            active_unit = None
            if unit_location in self.elves:
                active_unit = self.elves[unit_location]
            elif unit_location in self.goblins:
                active_unit = self.goblins[unit_location]

            if not active_unit or active_unit.unit_id in already_evaluated_unit_ids:
                continue

            lost_unit_kind = self.simulate_turn(unit_location)
            if lost_unit_kind == "E":
                elf_losses += 1
            elif lost_unit_kind == "G":
                goblin_losses += 1

            already_evaluated_unit_ids.add(active_unit.unit_id)

        return True, elf_losses, goblin_losses

    def simulate_turn(self, unit_location):
        attacking_group = None
        targets = None
        if unit_location in self.elves:
            attacking_group = self.elves
            targets = self.goblins
        else:
            attacking_group = self.goblins
            targets = self.elves
        unit = attacking_group.pop(unit_location)

        # Move Phase.
        closest_move_target = self.search_closest_move_to_target(unit_location, targets)
        if not closest_move_target:
            closest_move_target = unit_location

        attacking_group[closest_move_target] = unit

        # Attack Phase.
        target_enemy_locations = []
        for step in STEPS_IN_READING_ORDER:
            candidate_location = (
                closest_move_target[0] + step[0],
                closest_move_target[1] + step[1],
            )
            if candidate_location in targets:
                target_enemy_locations.append(candidate_location)

        if target_enemy_locations:
            enemy_units_with_location = [
                (targets[location], location) for location in target_enemy_locations
            ]
            targeted_unit, targeted_enemy_location = sorted(
                enemy_units_with_location, key=lambda x: (x[0].hp, x[1])
            )[0]

            targeted_unit.hp -= unit.attack
            if targeted_unit.hp <= 0:
                targets.pop(targeted_enemy_location)
                return targeted_unit.kind

        return None

    def search_closest_move_to_target(self, unit_location, targets):
        seen = set()
        frontier = [(0, unit_location, None)]

        while frontier:
            distance, location, first_step = heapq.heappop(frontier)

            if location in seen:
                continue

            if (
                location in self.elves
                or location in self.goblins
                or self.grid[location[0]][location[1]] == "#"
            ):
                seen.add(location)
                continue

            i, j = location

            for attack_step in STEPS_IN_READING_ORDER:
                maybe_enemy_location = (
                    i + attack_step[0],
                    j + attack_step[1],
                )
                if maybe_enemy_location in targets:
                    if distance == 0:
                        return None
                    elif first_step:
                        return first_step
                    else:
                        return location

            for step in STEPS_IN_READING_ORDER:
                next_step = (i + step[0], j + step[1])

                heapq.heappush(
                    frontier,
                    (
                        distance + 1,
                        next_step,
                        first_step if first_step else next_step,
                    ),
                )

            seen.add(location)

        return None

    def layout_str(self):
        layout = []
        for i in range(len(self.grid)):
            row = []
            for j in range(len(self.grid[0])):
                if (i, j) in self.elves:
                    row.append(f" ^{self.elves[(i, j)].hp:3} ")
                elif (i, j) in self.goblins:
                    row.append(f" ${self.goblins[(i,j)].hp:3} ")
                elif self.grid[i][j] == "#":
                    row.append("######")
                else:
                    row.append("  ..  ")
            layout.append("".join(row))
        return "\n".join(layout)

    def debug(self, rounds_completed, pause=False):
        print(f"{rounds_completed}:")
        print(self.layout_str())
        print(f"ELVES: {sum(elf.hp for elf in self.elves.values())} HP remaining")
        print(f"GOBLINS: {sum(gob.hp for gob in self.goblins.values())} HP remaining")
        if pause:
            input()
