import signal
import time

count = 0

def handler(sig, frame):
    global count
    count += 1
    print(f"START handler #{count} at {time.time():.1f}")
    time.sleep(3)
    print(f"END handler #{count} at {time.time():.1f}")

signal.signal(signal.SIGINT, handler)

print("Press Ctrl+C multiple times rapidly!")
time.sleep(5)
print(f"Total handler runs: {count}")
