#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/interrupt.h>
#include <linux/time.h>

#define TIMER_IRQ 10  // arch_timer from /proc/interrupts

static unsigned long last_jiffies = 0;

static irqreturn_t timer_interrupt_handler(int irq, void *dev_id)
{
    unsigned long now = jiffies;
    
    if (time_after(now, last_jiffies + HZ)) {
        printk(KERN_INFO "Hello World from timer interrupt!\n");
        last_jiffies = now;
    }
    
    return IRQ_HANDLED;
}

static int __init timer_demo_init(void)
{
    int result;
    
    printk(KERN_INFO "Timer Demo: Requesting IRQ %d\n", TIMER_IRQ);
    
    result = request_irq(TIMER_IRQ, 
                        timer_interrupt_handler,
                        IRQF_SHARED,
                        "timer_demo",
                        &last_jiffies);
    
    if (result) {
        printk(KERN_ERR "Timer Demo: Cannot register IRQ %d, error %d\n", 
               TIMER_IRQ, result);
        return result;
    }
    
    printk(KERN_INFO "Timer Demo: Successfully registered\n");
    last_jiffies = jiffies;
    return 0;
}

static void __exit timer_demo_exit(void)
{
    free_irq(TIMER_IRQ, &last_jiffies);
    printk(KERN_INFO "Timer Demo: Unregistered\n");
}

module_init(timer_demo_init);
module_exit(timer_demo_exit);

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Simple timer interrupt demo");
