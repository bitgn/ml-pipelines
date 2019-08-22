from typing import Any

import bs4

from client import ml_pipelines as client
from . import env




def none(selector):
    def _(s: bs4.BeautifulSoup):
        matches = s.select(selector)
        if len(matches) > 0:
            return f"Expected no matches for '{selector}' got {len(matches)}"

    return env.Then(_, None)


def count(selector, c: int):
    def _(s: bs4.BeautifulSoup):
        matches = s.select(selector)
        val = len(matches)
        if val != c:
            return f"Expected {c} match(es) for '{selector}' got {val}"

    return env.Then(_, None)



def invalid_argument(subject_id=None):
    def _(response: Any):
        if not isinstance(response, client.InvalidArgument):
            return "Expected InvalidArgument"

        ex: client.InvalidArgument = response
        if subject_id:
            if ex.subject_id != subject_id:
                return f'Expected subject {subject_id} got {ex.subject_id}'


    return env.Then(None, _)

def already_exists(subject_id=None):
    def _(response: Any):
        if not isinstance(response, client.AlreadyExists):
            return "Expected AlreadyExists"

        ex: client.AlreadyExists = response
        if subject_id:
            if ex.subject_id != subject_id:
                return f'Expected subject {subject_id} got {ex.subject_id}'


    return env.Then(None, _)

def client_ok():
    def _(response: Any):
        if isinstance(response, Exception):
            return f"Expected valid response, got ({type(response)}) {response}"

    return env.Then(None, _)

def text(selector, expected: str, hint=None, title=None):



    def _(s: bs4.BeautifulSoup):
        result = s.select_one(selector)

        if hint:
            pretty = f" ({hint})"
        else:
            pretty = ""



        if not result:
            msg = f"Expected result with text '{expected}' for '{selector}'{pretty}"


            return msg
        if result.text.strip() != expected:
            return f"Text for '{selector}'{pretty} should be '{expected}' not '{result.text}'"
        if title and result['title'] != title:
                return f"Text for '{selector}'.title {pretty} should be '{title}' not '{result['title']}'"

    return env.Then(_, None)


def link(selector, href=str, text=str):
    def _(s: bs4.BeautifulSoup):
        result = s.select_one(selector)
        if not result:
            return f"Expected result for '{selector}'"

        if href and result['href'] != href:
            return f"Link '{selector}' should be '{href}' not '{result['href']}'"

        if text and result.text != text:
            return f"Link '{selector}' should have text '{text}' not '{result.text}'"

    return env.Then(_, None)


def exists(selector):
    return count(selector, 1)
