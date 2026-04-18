import asyncio
from foobar import FooBar
from typing import Callable


def printer(m: str) -> Callable:
    def c() -> None:
        print(m, end="")

    return c


async def foo(f: FooBar) -> None:
    await f.foo(printer("foo"))


async def bar(f: FooBar) -> None:
    await f.bar(printer("bar"))


async def main():
    n = 10
    f = FooBar(n)

    tasks = [
        asyncio.create_task(foo(f)),
        asyncio.create_task(bar(f)),
    ]

    await asyncio.gather(*tasks, return_exceptions=True)


if __name__ == "__main__":
    asyncio.run(main())
