import io
import sys
import asyncio
from foobar import FooBar
from typing import Callable
import pytest
from foobar import FooBar

# The "Test Table" (Input A, Input B, Expected Result)
test_data = [
    (1, "foobar"),
    (2, "foobarfoobar"),
    (10, "foobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobarfoobar"),
]


def printer(m: str) -> Callable:
    def c():
        print(m, end="")

    return c


async def foo(f: FooBar) -> None:
    await f.foo(printer("foo"))


async def bar(f: FooBar) -> None:
    await f.bar(printer("bar"))


@pytest.mark.asyncio
@pytest.mark.parametrize("n, expected", test_data)
async def test_add(n, expected):
    f = FooBar(n)
    output_buffer = io.StringIO()
    original_stdout = sys.stdout
    sys.stdout = output_buffer

    tasks = [
        asyncio.create_task(foo(f)),
        asyncio.create_task(bar(f)),
    ]

    await asyncio.gather(*tasks, return_exceptions=True)

    try:
        captured_text = output_buffer.getvalue()
    finally:
        sys.stdout = original_stdout

    # print(f"output: {captured_text}")
    assert captured_text == expected
