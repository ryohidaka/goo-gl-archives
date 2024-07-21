import csv

from datetime import datetime
from typing import List

from goo_gl_archives.utils.logger import setup_logger

logger = setup_logger()
current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")


def write_to_csv(data: List[List[str]], filename: str) -> None:
    """
    Write the collected data to a CSV file with the specified filename.
    """
    try:
        # Append created_at and updated_at to each row
        data = [row + [current_time, current_time] for row in data]

        with open(filename, mode="w", newline="", encoding="utf-8") as file:
            writer = csv.writer(file)
            writer.writerow(
                [
                    "uid",
                    "original_url",
                    "redirect_url",
                    "domain_name",
                    "site_title",
                    "http_status",
                    "created_at",
                    "updated_at",
                ]
            )
            writer.writerows(data)
        logger.info(f"Results written to {filename}")
    except Exception as e:
        logger.error(f"Failed to write to CSV file {filename}: {e}")
