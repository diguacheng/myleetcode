import queue


class MaxQueue:

    def __init__(self):
        self.q=queue.Queue()
        self.d=queue.deque()

    def max_value(self):
        if self.q.empty():
            return -1
        return self.d[0]
    def push_back(self, value):
        while self.d and self.d[-1]<value:
            self.d.pop()
        self.q.put(value)
        self.d.append(value)


    def pop_front(self):
        if self.q.empty():
            return -1
        x=self.q.get()
        if x==self.d[0]:
            self.d.popleft()
        return x
