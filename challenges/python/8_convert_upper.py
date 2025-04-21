"""
For this challenge, create a Python function that accepts a string.
The function should return a string, with each lowercase character in the original
string returned as uppercase characters. If you send the function “goodbye” as a parameter,
it should return “GOODBYE”
"""

def convertToUpper(word: str) -> str:
    return word.upper()


if __name__ == "__main__":
    print(convertToUpper("goodbye"))
