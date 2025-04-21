"""
Write a function in Python that accepts a list of any length that contains
a mix of non-negative integers and strings.
The function should return a list with only the integers in the original list in the same order.
"""
from typing import List


def justNumbers(arr: List) -> List[int]:
    return [ele for ele in arr if isinstance(ele, int) and not isinstance(ele, bool)]


if __name__ == "__main__":
    print(justNumbers(["one", 2, 3, "four", 5, 0, True]))
