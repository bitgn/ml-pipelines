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



def bad_name(name:str):
    def _(response: Any):
        if not isinstance(response, client.BadName):
            return "Expected BadName"

        ex: client.BadName = response
        if ex.subject_name != name:
            return f'Expected subject {name} got {ex.subject_name}'
    return env.Then(None, _)



def invalid_argument(subject_uid=None, subject_name=None):
    def _(response: Any):
        if not isinstance(response, client.InvalidArgument):
            return "Expected InvalidArgument"

        ex: client.InvalidArgument = response
        if subject_uid:
            if ex.subject_uid != subject_uid:
                return f'Expected subject {subject_uid} got {ex.subject_uid}'


        if subject_name:
            if ex.subject_name != subject_name:
                return f'Expected subject {subject_name} got {ex.subject_name}'
    return env.Then(None, _)

def name_taken(subject_uid, subject_name):
    def _(response: Any):
        if not isinstance(response, client.AlreadyExists):
            return f"Expected {client.AlreadyExists.__name__}, got {type(response).__name__}: {response}"

        ex: client.AlreadyExists = response
        if ex.subject_uid != subject_uid:
            return f'Expected subject {subject_uid} got {ex.subject_uid}'

        if ex.subject_name != subject_name:
            return f'Expected subject {subject_name} got {ex.subject_name}'



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
