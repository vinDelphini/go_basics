"""
Create a function in Python that accepts a single word and returns the number of vowels in that word.
In this function, only a, e, i, o, and u will be counted as vowels — not y.
"""

def vowels_in_string(word: str) -> int:
    vowels = ("a", "e", "i", "o", "u")
    return len([1 for letter in word if letter in vowels])


if __name__=="__main__":
    print(vowels_in_string("venkatesh"))
