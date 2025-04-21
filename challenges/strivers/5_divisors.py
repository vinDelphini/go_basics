# Given an integer N, return all divisors of N.
from typing import List

def divisors(n: int) -> List[int]:
    return [num for num in range(1, n + 1) if n % num == 0]


if __name__ == "__main__":
    print(divisors(36))
