import random
import string

from typing import List


def generate_random_strings(
    count: int, min_length: int = 5, max_length: int = 8
) -> List[str]:
    """
    Generate random and unique strings of random lengths for the given count.

    Args:
        count (int): Number of unique strings to generate.
        min_length (int): Minimum length of the generated strings.
        max_length (int): Maximum length of the generated strings.

    Returns:
        List[str]: List of generated unique strings.
    """
    random_strings = []
    for _ in range(count):
        try:
            length = random.randint(min_length, max_length)
            random_string = "".join(
                random.choices(string.ascii_letters + string.digits, k=length)
            )
            random_strings.append(random_string)
        except Exception as e:
            print(f"Error generating random string: {e}")
    return random_strings
