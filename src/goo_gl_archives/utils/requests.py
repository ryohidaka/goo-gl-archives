import random
import string

import time
from typing import List

import requests
from sqlalchemy import Tuple


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


def get_redirect_info(url: str) -> Tuple[str, str, str, int]:
    """
    Retrieve the redirect URL, domain name, and HTTP status code for a given URL.
    """
    try:
        response = requests.get(url, allow_redirects=True)
        redirect_url = response.url
        domain_name = requests.utils.urlparse(redirect_url).netloc
        http_status = response.status_code

        # Delay to avoid overwhelming the server
        time.sleep(0.3)

        return url, redirect_url, domain_name, http_status
    except requests.RequestException as e:
        print(f"Request failed for URL {url}: {e}")
        return url, "Failed", "Failed", "Failed", 0
    except Exception as e:
        print(f"Error processing URL {url}: {e}")
        return url, "Failed", "Failed", "Failed", 0
