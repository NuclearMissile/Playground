import heapq

graph = {
    's': {
        'h': 2,
        'edges': {
            'x': 1,
            'y': 1,
        }
    },
    'x': {
        'h': 2,
        'edges': {
            'w': 1,
            'z': 2,
        }
    },
    'y': {
        'h': 1,
        'edges': {
            'z': 3,
            't': 8,
        }
    },
    'w': {
        'h': 1,
        'edges': {
            'z': 3
        }
    },
    'z': {
        'h': 1,
        'edges': {
            't': 2
        }
    },
    't': {
        'h': 0,
        'edges': {}
    }
}


def reheapify(heap: heapq, old, new):
    i = heap.index(old)
    heap[i] = new
    if new > old:
        heapq._siftup(heap, i)
    elif new < old:
        heapq._siftdown(heap, 0, i)


def print_path(parent):
    l = ['t']
    curr = 't'
    while True:
        curr = parent[curr]
        l.append(curr)
        if curr == 's':
            l.reverse()
            print('Shortest Path:', '->'.join(l))
            return


def a_star(graph, start, end):
    OPEN, g, f, parent, CLOSED = [], {start: 0}, {start: graph[start]['h']}, {}, set()
    heapq.heappush(OPEN, (f[start], start))
    OPEN_track, CLOSED_track, f_tarck, v_track, g_track = [], [], [], [], []
    while True:
        OPEN_track.append(OPEN.copy())
        CLOSED_track.append(CLOSED.copy())
        f_tarck.append(f.copy())
        g_track.append(g.copy())
        if len(OPEN) == 0:
            return
        v = heapq.heappop(OPEN)[-1]
        v_track.append(v)
        CLOSED.add(v)
        if v == end:
            print_path(parent)
            print('==============')
            for i in range(len(OPEN_track)):
                print(f'round:  {i}\n'
                      f'v:      {v_track[i]}\n'
                      f'OPEN:   {OPEN_track[i]}\n'
                      f'CLOSED: {CLOSED_track[i]}\n'
                      f'f_vals: {f_tarck[i]}\n'
                      f'g_vals: {g_track[i]}\n')
            return
        for u, cost in graph[v]['edges'].items():
            if (f.get(u, -1), u) not in OPEN and u not in CLOSED:
                g[u] = g[v] + cost
                f[u] = g[u] + graph[u]['h']
                parent[u] = v
                heapq.heappush(OPEN, (f[u], u))
            elif (f[u], u) in OPEN:
                if g[v] + cost < g[u]:
                    g[u] = g[v] + cost
                    old = (f[u], u)
                    f[u] = g[u] + graph[u]['h']
                    new = (f[u], u)
                    reheapify(OPEN, old, new)
                    parent[u] = v


if __name__ == '__main__':
    a_star(graph, 's', 't')
