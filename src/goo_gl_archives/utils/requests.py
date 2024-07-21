import random
import string

import time

import requests

from bs4 import BeautifulSoup
from typing import List, Optional, Dict, Any

from goo_gl_archives.utils.logger import setup_logger


logger = setup_logger()


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
            logger.error(f"Error generating random string: {e}")
    return random_strings


def get_redirect_info(url: str) -> Optional[Dict[str, Any]]:
    """
    Retrieve the redirect URL, domain name, site title, and HTTP status code for a given URL.

    Args:
        url (str): The URL to fetch redirect information for.

    Returns:
        Optional[Dict[str, Any]]: Dictionary containing redirect information if available, otherwise None.
    """
    try:
        response = requests.get(url, allow_redirects=True)
        redirect_url = response.url
        domain_name = requests.utils.urlparse(redirect_url).netloc
        soup = BeautifulSoup(requests.get(redirect_url).text, "html.parser")
        site_title = soup.title.string if soup.title else None
        http_status = response.status_code

        # Delay to avoid overwhelming the server
        time.sleep(0.3)

        if url == redirect_url:
            return None

        return {
            "redirect_url": redirect_url,
            "domain_name": domain_name,
            "site_title": site_title,
            "http_status": http_status,
        }
    except requests.RequestException as e:
        logger.error(f"Request failed for URL {url}: {e}")
    except Exception as e:
        logger.error(f"Error processing URL {url}: {e}")
        return None
