import heapq

graph_1 = {
    's': {'y': 6, 'x': 1},
    'y': {'w': 2, 'z': 3},
    'w': {'t': 9},
    'x': {'y': 4, 'z': 9},
    'z': {'t': 5},
    't': {},
}

graph_2 = {
    's': {'a': 2, 'b': 2, 'c': 4},
    'a': {'c': 1},
    'b': {'c': 1},
    'c': {'d': 1, 'e': 1, 't': 3},
    'd': {'t': 1},
    'e': {'t': 1},
    't': {},
}

INF = 1e5


def reheapify(heap: heapq, old, new):
    i = heap.index(old)
    heap[i] = new
    if new > old:
        heapq._siftup(heap, i)
    elif new < old:
        heapq._siftdown(heap, 0, i)


def init(start: str):
    global g, parent, OPEN, CLOSED
    g, parent, OPEN = {start: 0}, {}, [(0, start)]
    heapq.heapify(OPEN)
    CLOSED = set()


def print_path(start: str, end: str):
    l = [end]
    curr = end
    while True:
        curr = parent[curr]
        l.append(curr)
        if curr == start:
            l.reverse()
            print('Shortest Path:', '->'.join(l))
            return


def dijkstra(graph: dict, start: str, end: str):
    init(start)
    OPEN_track, v_track = [], []
    while len(OPEN) != 0:
        OPEN_track.append(OPEN.copy())
        v = heapq.heappop(OPEN)[-1]
        CLOSED.add(v)
        v_track.append(v)
        if v == end:
            print_path(start, end)
            for i in range(len(OPEN_track)):
                print(f'{i}, {OPEN_track[i]}, {v_track[i]}')
            return
        for u, cost in graph[v].items():
            if (g.get(u, -1), u) not in OPEN and u not in CLOSED:
                parent[u] = v
                g[u] = g[v] + cost
                heapq.heappush(OPEN, (g[u], u))
            elif (g[u], u) in OPEN:
                if g[v] + cost < g[u]:
                    old = (g[u], u)
                    parent[u] = v
                    g[u] = g[v] + cost
                    new = (g[u], u)
                    reheapify(OPEN, old, new)


def count_shortest_path(graph: dict, start: str, end: str):
    init(start)
    num = {start: 1}
    while len(OPEN) != 0:
        u = heapq.heappop(OPEN)[-1]
        if u not in CLOSED:
            CLOSED.add(u)
            for v, w in graph[u].items():
                if g.get(v, INF) == g[u] + w:
                    num[v] += num[u]
                elif g.get(v, INF) > g[u] + w:
                    g[v] = g[u] + w
                    heapq.heappush(OPEN, (g[v], v))
                    num[v] = num[u]
    print('Shortest path count:', num[end])


if __name__ == '__main__':
    print('==========Q1===========')
    dijkstra(graph_1, 's', 't')
    print('==========Q2===========')
    count_shortest_path(graph_2, 's', 't')
