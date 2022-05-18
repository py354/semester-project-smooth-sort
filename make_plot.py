import matplotlib.pyplot as plt

heap_random = [[], []]
heap_sorted = [[], []]
smooth_random = [[], []]
smooth_sorted = [[], []]

for data in ((heap_random, 'heap_random'), (heap_sorted, 'heap_sorted'), (smooth_random, 'smooth_random'), (smooth_sorted, 'smooth_sorted')):
    for line in open(f'result/{data[1]}.csv').readlines():
        if line == '':
            continue

        count, dur = line.split(',')
        count = int(count)
        dur = int(dur)
        data[0][0].append(count)
        data[0][1].append(dur)

print(heap_random)
print(heap_sorted)
print(smooth_random)
print(smooth_sorted)

plt.grid()
plt.ylabel('Время в нс')
plt.xlabel('Количество элементов')

plt.title('Время сортировки массива')

plt.plot(heap_random[0], heap_random[1], label='heap_random')
plt.plot(heap_sorted[0], heap_sorted[1], label='heap_sorted')
plt.plot(smooth_random[0], smooth_random[1], label='smooth_random')
plt.plot(smooth_sorted[0], smooth_sorted[1], label='smooth_sorted')

plt.legend()
plt.savefig('result/plot.png')