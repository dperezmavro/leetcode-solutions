class Node:
    def __init__(self):
        self.isWord = False
        self.letters = dict()

    def insert(self, w: str):
        self.val = w[0]
        if len(w) == 1:
            self.isWord = True
            return

        if len(w) > 1:
            c = w[1]
            if c not in self.letters:
                n = Node()
                self.letters[c] = n

            self.letters[c].insert(w[1:])

    def search(self, w: str, strict: bool) -> bool:
        c = w[0]
        if c != self.val:
            return False

        if len(w) == 1:
            if self.isWord or not strict:
                return True
        elif len(w) > 1:
            c = w[1]
            if c not in self.letters:
                return False
            return self.letters[c].search(w[1:], strict)

        return False


class Trie:
    def __init__(self):
        self.letters = dict()

    def insert(self, word: str) -> None:
        c = word[0]
        n: Node = None
        if c in self.letters:
            n = self.letters[c]
        else:
            n = Node()
            self.letters[c] = n

        n.insert(word)

    def search(self, word: str) -> bool:
        c = word[0]
        if c not in self.letters:
            return False
        return self.letters[c].search(word, True)

    def startsWith(self, prefix: str) -> bool:
        c = prefix[0]
        if c not in self.letters:
            return False
        return self.letters[c].search(prefix, False)
