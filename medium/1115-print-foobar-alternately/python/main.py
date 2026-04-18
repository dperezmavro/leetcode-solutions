import asyncio
from foobar import FooBar


async def foo(f: FooBar):
    while True:

        def pf():
            print("foo")

        await f.foo(pf)


async def bar(f: FooBar):
    while True:

        def pb():
            print("bar")

        await f.foo(pb)


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
