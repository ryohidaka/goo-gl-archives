from goo_gl_archives.utils.requests import generate_random_strings


def main() -> int:
    print("Hello from goo-gl-archives!")
    app = GooGlArchives(count=10)
    app.run()

    return 0


class GooGlArchives:
    def __init__(self, count: int = 10):
        self.count = count

    def run(self) -> None:
        """
        Main function to generate URLs.
        """
        # Generate unique strings
        unique_strings = generate_random_strings(self.count)
        print(unique_strings)
