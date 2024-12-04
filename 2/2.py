import os

with open("input.txt", 'r') as f:
    safe = 0
    for line in f:
        cur_safe = True
        delta = 0
        parts = line.split(" ")
        last = int(parts[0])
        for part in parts[1:]:
            number = int(part)
            new_delta = number - last
            last = number
            if delta > 0 and new_delta < 0:
                cur_safe = False
                break
            elif delta < 0 and new_delta > 0:
                cur_safe = False
                break
            delta = new_delta
            if abs(delta) == 0 or abs(delta) > 3:
                cur_safe = False
                break
        if cur_safe:
            safe = safe + 1
    print(safe)
