import os

def isSafe(parts: list[int]) -> bool:
    last = parts[0]
    delta = 0
    for part in parts[1:]:
        number = int(part)
        new_delta = number - last
        last = number
        if delta > 0 and new_delta < 0:
            return False
        elif delta < 0 and new_delta > 0:
            return False
        delta = new_delta
        if abs(delta) == 0 or abs(delta) > 3:
            return False
    return True

with open("input.txt", 'r') as f:
    safe = 0
    for line in f:
        line_safe = False
        # Forgive me
        parts = line.split(" ")
        parts = [int(x) for x in parts]
        
        permutations = [parts]

        for idx, ignore in enumerate(parts):
            # remove index
            mini_parts = [x for i, x in enumerate(parts) if i != idx]
            permutations.append(mini_parts)
        for permutation in permutations:
            if isSafe(permutation):
                line_safe = True
                break
        if line_safe:
            safe = safe + 1
    print(safe)
