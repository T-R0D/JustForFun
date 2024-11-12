import itertools
import re


def part_one(puzzle_input):
    immune_army, infection_army = parse_armies(puzzle_input)

    while immune_army and infection_army:
        immune_army, infection_army, _ = simulate_round_of_battle(
            immune_army, infection_army
        )

    remaining_units = 0
    if immune_army:
        remaining_units = sum([group.n_units for group in immune_army])
    else:
        remaining_units = sum([group.n_units for group in infection_army])

    return str(remaining_units)


def part_two(puzzle_input):
    # Possible improvement: Get this perfected with a binary search.
    # For now, it's good enough: it's correct and runs in just over 1 second.
    immune_army, infection_army = [], []
    immune_boost = 0
    try_again = True
    while not immune_army or try_again:
        try_again = False

        immune_army, infection_army = parse_armies(puzzle_input)
        for group in immune_army:
            group.attack_per_unit += immune_boost

        combat_proceeds_for_immune_army = True
        while immune_army and infection_army and combat_proceeds_for_immune_army:
            immune_army, infection_army, combat_proceeds_for_immune_army = (
                simulate_round_of_battle(immune_army, infection_army)
            )

        if not combat_proceeds_for_immune_army:
            try_again = True

        immune_boost += 1

    return str(sum([group.n_units for group in immune_army]))


def parse_armies(puzzle_input):
    army_blocks = puzzle_input.split("\n\n")
    group_id = 1

    armies = []
    for army_name, block in zip(("immune", "infection"), army_blocks):
        groups = []
        for line in block.split("\n")[1:]:
            groups.append(Group.from_line(group_id, army_name, line))
            group_id += 1

        armies.append(groups)

    immune_army, infection_army = armies

    return immune_army, infection_army


class Group:
    def __init__(
        self,
        group_id,
        army_name,
        n_units,
        hp,
        attack,
        attack_type,
        initiative,
        immunities,
        weaknesses,
    ):
        self.group_id = group_id
        self.army_name = army_name
        self.n_units = n_units
        self.hp_per_unit = hp
        self.attack_per_unit = attack
        self.attack_type = attack_type
        self.initiative = initiative
        self.immunities = immunities
        self.weaknesses = weaknesses

    @classmethod
    def from_line(cls, i, army_name, line):
        slim_str = (
            line.replace("units each with", "")
            .replace("hit points", "")
            .replace("with an attack that does", "")
            .replace("damage at initiative", "")
        )

        weaknesses = set()
        immunities = set()
        weaknesses_and_immunities_re = re.compile(r"\((.+)\)")
        match = weaknesses_and_immunities_re.search(slim_str)
        if match:
            weaknesses_and_immunities_str = match.group(1)
            slim_str = slim_str.replace(f"({weaknesses_and_immunities_str})", "")

            for lst in weaknesses_and_immunities_str.split("; "):
                if lst.startswith("immune"):
                    immunities = set(lst.replace("immune to ", "").split(", "))
                else:
                    weaknesses = set(lst.replace("weak to ", "").split(", "))

        n_units, hp_per_unit, attack_per_unit, attack_type, initiative = (
            slim_str.split()
        )

        return cls(
            i,
            army_name,
            int(n_units),
            int(hp_per_unit),
            int(attack_per_unit),
            attack_type,
            int(initiative),
            immunities,
            weaknesses,
        )

    def is_defeated(self):
        return self.n_units <= 0

    def effective_power(self):
        return self.n_units * self.attack_per_unit

    def compute_potential_damage(self, base_damage, attack_type):
        if attack_type in self.immunities:
            return 0

        multiplier = 1
        if attack_type in self.weaknesses:
            multiplier = 2

        return multiplier * base_damage

    def take_damage(self, base_damage, attack_type):
        units_defeated = (
            self.compute_potential_damage(base_damage, attack_type) // self.hp_per_unit
        )
        self.n_units -= units_defeated
        return self.is_defeated()


def simulate_round_of_battle(immune_army, infection_army):
    immune_attack_assignments = target_selection(immune_army, infection_army)
    infection_attack_assignments = target_selection(infection_army, immune_army)

    new_immune_army, new_infection_army, combat_proceeded_for_immune_army = (
        attack_phase(
            immune_army,
            immune_attack_assignments,
            infection_army,
            infection_attack_assignments,
        )
    )

    return new_immune_army, new_infection_army, combat_proceeded_for_immune_army


def target_selection(attacking_army, defending_army):
    attacking_army = sorted(
        attacking_army, key=lambda g: (-g.effective_power(), -g.initiative)
    )

    attack_assignments = {}
    selected_candidates = set()
    for attacking_group in attacking_army:
        best_target_id = None
        best_target_key = (0, 0, 0)

        for candidate_group in defending_army:
            if candidate_group.group_id in selected_candidates:
                continue

            potential_damage = candidate_group.compute_potential_damage(
                attacking_group.effective_power(),
                attacking_group.attack_type,
            )

            candidate_key = (
                potential_damage,
                candidate_group.effective_power(),
                candidate_group.initiative,
            )

            if potential_damage > 0 and best_target_key < candidate_key:
                best_target_id = candidate_group.group_id
                best_target_key = candidate_key

        attack_assignments[attacking_group.group_id] = best_target_id
        selected_candidates.add(best_target_id)

    return attack_assignments


def attack_phase(
    immune_army, immune_attack_assignments, infection_army, infection_attack_assignments
):
    immune_lookup = {group.group_id: group for group in immune_army}
    infection_lookup = {group.group_id: group for group in infection_army}

    attack_order = sorted(
        itertools.chain(immune_army, infection_army), key=lambda g: -g.initiative
    )

    combat_proceeded = False

    for attacking_group in attack_order:
        if attacking_group.is_defeated():
            continue

        attacker_id = attacking_group.group_id
        attacker_army_name = attacking_group.army_name

        attack_assignments = immune_attack_assignments
        defender_lookup = infection_lookup
        if attacker_army_name == "infection":
            attack_assignments = infection_attack_assignments
            defender_lookup = immune_lookup

        if not attack_assignments[attacker_id]:
            continue

        defending_group = defender_lookup[attack_assignments[attacker_id]]
        if defending_group.is_defeated():
            continue

        units_before = defending_group.n_units
        defending_group.take_damage(
            attacking_group.effective_power(), attacking_group.attack_type
        )
        if units_before > defending_group.n_units and attacker_army_name == "immune":
            combat_proceeded = True

    new_immune_army = [group for group in immune_army if not group.is_defeated()]
    new_infection_army = [group for group in infection_army if not group.is_defeated()]

    return new_immune_army, new_infection_army, combat_proceeded
