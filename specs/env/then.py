import bs4



from . import env

def none(selector):
    def _(s: bs4.BeautifulSoup):
        matches = s.select(selector)
        if len(matches) > 0:
            return f"Expected no matches for '{selector}' got {len(matches)}"

    return env.Then(_)


def count(selector, c: int):
    def _(s: bs4.BeautifulSoup):
        matches = s.select(selector)
        val = len(matches)
        if val != c:
            return f"Expected {c} match(es) for '{selector}' got {val}"

    return env.Then(_)


def text(selector, expected: str, hint=None):



    def _(s: bs4.BeautifulSoup):
        result = s.select_one(selector)

        if hint:
            pretty = f" ({hint})"



        if not result:
            msg = f"Expected result with text '{expected}' for '{selector}'{pretty}"


            return msg
        if result.text.strip() != expected:
            return f"Text for '{selector}'{pretty} should be '{expected}' not '{result.text}'"

    return env.Then(_)


def link(selector, href=str, text=str):
    def _(s: bs4.BeautifulSoup):
        result = s.select_one(selector)
        if not result:
            return f"Expected result for '{selector}'"

        if href and result['href'] != href:
            return f"Link '{selector}' should be '{href}' not '{result['href']}'"

        if text and result.text != text:
            return f"Link '{selector}' should have text '{text}' not '{result.text}'"

    return env.Then(_)


def exists(selector):
    return count(selector, 1)
