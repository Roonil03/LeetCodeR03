class MyCalendarThree:
    def __init__(self):
        self.cal = []

    def book(self, startTime: int, endTime: int) -> int:
        self.cal.append((startTime, 1))
        self.cal.append((endTime, -1))
        self.cal.sort()
        book = 0
        maxKBook = 0
        for time, delta in self.cal:
            book += delta
            maxKBook = max(maxKBook, book)
        return maxKBook


# Your MyCalendarThree object will be instantiated and called as such:
# obj = MyCalendarThree()
# param_1 = obj.book(startTime,endTime)