import concurrent.futures

from goo_gl_archives.utils.logger import setup_logger
from goo_gl_archives.utils.requests import generate_random_strings, get_redirect_info
from goo_gl_archives.utils.sql import init_sqlalchemy, insert_data


def main() -> int:
    print("Hello from goo-gl-archives!")
    app = GooGlArchives(database="sqlite:///db/archives.db", count=10)
    app.run()

    return 0


logger = setup_logger()


class GooGlArchives:
    def __init__(self, database: str, count: int = 10):
        self.count = count
        self.base_url = "https://goo.gl/"
        self.filename = "output.csv"

        # Init SQLAlchemy
        self.session = init_sqlalchemy(database)

    def run(self) -> None:
        """
        Main function to generate URLs, retrieve their redirect information, and insert the results to Database.
        """
        logger.info("Starting GooGlArchives")

        # Generate unique strings
        unique_strings = generate_random_strings(self.count)
        logger.info(f"Generated unique strings: {unique_strings}")

        results = []

        with concurrent.futures.ThreadPoolExecutor() as executor:
            future_to_uid = {
                executor.submit(get_redirect_info, self.base_url, uid): uid
                for uid in unique_strings
            }
            for future in concurrent.futures.as_completed(future_to_uid):
                uid = future_to_uid[future]
                try:
                    result = future.result()
                    if result is not None:
                        results.append(result)
                        logger.info(result)
                except Exception as e:
                    logger.error(f"Exception occurred for {uid}: {e}")

        # Insert data to DB
        insert_data(self.session, results)
