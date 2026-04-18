from typing import Callable
from asyncio import Queue


class FooBar:
    def __init__(self, n):
        self.n = n
        self.fooQ = Queue()
        self.barQ = Queue()
        self.barQ.put_nowait(0)

    async def foo(self, printFoo: Callable[[], None]) -> None:
        for i in range(self.n):
            await self.barQ.get()
            # printFoo() outputs "foo". Do not change or remove this line.
            printFoo()
            await self.fooQ.put(i)
        self.fooQ.task_done()

    async def bar(self, printBar: Callable[[], None]) -> None:
        for i in range(self.n):
            await self.fooQ.get()
            # printBar() outputs "bar". Do not change or remove this line.
            printBar()
            await self.barQ.put(i)
        self.barQ.task_done()
