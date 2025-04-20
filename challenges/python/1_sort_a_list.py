"""
Create a function in Python that accepts two parameters.
The first will be a list of numbers.
The second parameter will be a string that can be one of the following values: asc, desc, and none.

If the second parameter is “asc,” then the function should return a list with the numbers in ascending order.
If it’s “desc,” then the list should be in descending order, and if it’s “none,”
it should return the original list unaltered.
"""
from typing import List

def sort_by_input(nums: List[int], order: str) -> List[int]:

    if order == 'asc':
        return sorted(nums)
    elif order == 'desc':
        return sorted(nums, reverse=True)
    else:
        return nums

if __name__=="__main__":
    print(sort_by_input([5, 4, 1, 3, 2], "desc"))
