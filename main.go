package main

import (
	"fmt"
	"go-mini-cache/cache"
	"time"
)

func main() {
	c := cache.NewCache()

	// Set karo — alag alag TTL ke saath
	c.Set("name", "Rushi", 5*time.Second)    // 5 sec mein expire
	c.Set("city", "Pune", 10*time.Second)    // 10 sec mein expire
	c.Set("lang", "Go", 2*time.Second)       // 2 sec mein expire

	fmt.Println("=== Cache Demo ===")
	fmt.Printf("Cache size: %d\n\n", c.Size())

	// Turant Get karo — sab milna chahiye
	if val, ok := c.Get("name"); ok {
		fmt.Println("name:", val)
	}
	if val, ok := c.Get("city"); ok {
		fmt.Println("city:", val)
	}
	if val, ok := c.Get("lang"); ok {
		fmt.Println("lang:", val)
	}

	// 3 seconds wait karo
	fmt.Println("\n--- 3 seconds wait kar rahe hain ---")
	time.Sleep(3 * time.Second)

	// Ab "lang" expire ho gayi hogi (TTL 2s tha)
	fmt.Println("\n=== After 3 seconds ===")
	if val, ok := c.Get("lang"); ok {
		fmt.Println("lang:", val)
	} else {
		fmt.Println("lang: EXPIRED ❌")
	}
	if val, ok := c.Get("name"); ok {
		fmt.Println("name:", val, "✅ still alive")
	}

	// Delete test
	c.Delete("city")
	if _, ok := c.Get("city"); !ok {
		fmt.Println("city: DELETED ❌")
	}

	// Cleanup test
	fmt.Println("\n--- 3 aur seconds wait ---")
	time.Sleep(3 * time.Second)

	removed := c.Cleanup()
	fmt.Printf("\nCleanup removed %d expired keys\n", removed)
	fmt.Printf("Cache size after cleanup: %d\n", c.Size())
}