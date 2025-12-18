#include <linux/module.h>
#include <linux/timer.h>
#include <linux/jiffies.h>

static struct timer_list my_timer;
static unsigned long count = 0;

// Using a kernel timer.
static void timer_callback(struct timer_list *t) {
    count++;
    printk(KERN_INFO "Hello World from timer! Count: %lu\n", count);
    
    // Reschedule for 1 second later, forming a loop.
    mod_timer(&my_timer, jiffies + HZ);
}

static int __init hello_init(void) {
    printk(KERN_INFO "Starting Hello World timer module\n");
    
    // Create our timer, but don't set an expiration time yet.
    timer_setup(&my_timer, timer_callback, 0);
    // Reschedule for 1 second later by modifying the timer's
    // expiration time to now + 1 second (HZ).
    // - HZ = number of jiffies in a second.
    // - jiffies is a global kernel variable representing the number of
    //   timer interrupts (ticks) since the system booted.
    mod_timer(&my_timer, jiffies + HZ);
    
    return 0;
}

static void __exit hello_exit(void) {
    del_timer_sync(&my_timer);
    printk(KERN_INFO "Stopped timer. Total callbacks: %lu\n", count);
}

module_init(hello_init);
module_exit(hello_exit);
MODULE_LICENSE("GPL");
MODULE_AUTHOR("Zach Croft");
MODULE_DESCRIPTION("Simple timer interrupt demo");
