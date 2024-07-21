from goo_gl_archives.utils.requests import generate_random_strings, get_redirect_info


def main() -> int:
    print("Hello from goo-gl-archives!")
    app = GooGlArchives(count=10)
    app.run()

    return 0


class GooGlArchives:
    def __init__(self, count: int = 10):
        self.count = count
        self.base_url = "https://goo.gl/"

    def run(self) -> None:
        """
        Main function to generate URLs, retrieve their redirect information.
        """
        # Generate unique strings
        unique_strings = generate_random_strings(self.count)
        print(unique_strings)

        for uid in unique_strings:
            full_url = self.base_url + uid
            try:
                # Retrieve redirect information for the URL
                redirect_info = get_redirect_info(full_url)
                print(redirect_info)
            except Exception as e:
                print(f"Failed to get info for {full_url}: {e}")
