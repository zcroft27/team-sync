import signal
import time

count = 0

def handler(sig, frame):
    global count
    count += 1
    n = count
    print(f"START handler #{n} at {time.time():.1f}")
    time.sleep(3)
    print(f"END handler #{n} at {time.time():.1f}")

signal.signal(signal.SIGINT, handler)

print("Press Ctrl+C!")
time.sleep(5)
print(f"Total handler runs: {count}")














































# Even though Python is single-threaded, a signal can interrupt a blocking call (time.sleep(3)) inside a signal handler,
# causing another handler invocation to start before the previous one finishes.
#
# Signal handlers allow a handler to be re-invoked while running, unlike a interrupt handler.
