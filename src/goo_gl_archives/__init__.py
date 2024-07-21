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

        for uid in unique_strings:
            full_url = self.base_url + uid
            try:
                # Retrieve redirect information for the URL
                redirect_info = get_redirect_info(full_url)

                if redirect_info is None:
                    continue

                results.append(
                    {
                        "uid": uid,
                        "redirect_url": redirect_info["redirect_url"],
                        "domain_name": redirect_info["domain_name"],
                        "site_title": redirect_info["site_title"],
                        "http_status": redirect_info["http_status"],
                    }
                )
                logger.info(redirect_info)
            except Exception as e:
                logger.error(f"Failed to get info for {full_url}: {e}")

        # Insert data to DB
        insert_data(self.session, results)
