"""
Create a Python function that accepts a string.
This function should count the number of Xs and the number of Os in the string.
It should then return a boolean value of either True or False.

If the count of Xs and Os are equal, then the function should return True.
If the count isn’t the same, it should return False.
If there are no Xs or Os in the string, it should also return True because 0 equals 0.
The string can contain any type and number of characters.
"""

def count_x_o(word: str) -> bool:
    x_counter = 0
    y_counter = 0
    for letter in word:
        if letter == "X":
            x_counter += 1
        if letter == "O":
            y_counter += 1
    return x_counter == y_counter


if __name__ == "__main__":
    print(count_x_o("XOXOXOX"))
